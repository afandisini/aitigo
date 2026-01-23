package ratelimit

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type Limit struct {
	Rate  int
	Burst int
	Per   time.Duration
}

type Store interface {
	Allow(ctx context.Context, key string, limit Limit) (bool, time.Duration, error)
}

type KeyFunc func(r *http.Request) string

type Config struct {
	Limit   Limit
	Store   Store
	KeyFunc KeyFunc
}

func DefaultConfig() Config {
	return Config{
		Limit: Limit{
			Rate:  10,
			Burst: 20,
			Per:   time.Second,
		},
		Store:   NewInMemoryStore(),
		KeyFunc: func(r *http.Request) string { return "global" },
	}
}

func Middleware(cfg Config) func(http.Handler) http.Handler {
	if cfg.Store == nil {
		cfg.Store = NewInMemoryStore()
	}
	if cfg.KeyFunc == nil {
		cfg.KeyFunc = func(r *http.Request) string { return "global" }
	}
	if cfg.Limit.Rate <= 0 {
		cfg.Limit.Rate = 1
	}
	if cfg.Limit.Burst <= 0 {
		cfg.Limit.Burst = cfg.Limit.Rate
	}
	if cfg.Limit.Per <= 0 {
		cfg.Limit.Per = time.Second
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			key := cfg.KeyFunc(r)
			allowed, retryAfter, err := cfg.Store.Allow(r.Context(), key, cfg.Limit)
			if err != nil {
				writeLimitError(w, http.StatusInternalServerError, "INTERNAL", "rate limiter error")
				return
			}
			if !allowed {
				if retryAfter > 0 {
					w.Header().Set("Retry-After", intToString(int(retryAfter.Seconds())))
				}
				writeLimitError(w, http.StatusTooManyRequests, "RATE_LIMITED", "rate limit exceeded")
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

type errorPayload struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

type errorResponse struct {
	Success bool         `json:"success"`
	Error   errorPayload `json:"error"`
}

func writeLimitError(w http.ResponseWriter, status int, code, message string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(errorResponse{
		Success: false,
		Error: errorPayload{
			Code:    code,
			Message: message,
		},
	})
}
