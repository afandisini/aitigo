package recovermw

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"runtime/debug"
)

const (
	defaultMessage   = "internal server error"
	defaultErrorCode = "INTERNAL"
)

type Config struct {
	Format string
	Logger *slog.Logger
}

func DefaultConfig() Config {
	return Config{
		Format: "json",
		Logger: slog.Default(),
	}
}

func Middleware(cfg Config) func(http.Handler) http.Handler {
	if cfg.Format == "" {
		cfg.Format = "json"
	}
	if cfg.Logger == nil {
		cfg.Logger = slog.Default()
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if rec := recover(); rec != nil {
					cfg.Logger.Error("panic recovered", "error", rec, "stack", string(debug.Stack()))
					writeResponse(w, cfg)
				}
			}()

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

func writeResponse(w http.ResponseWriter, cfg Config) {
	w.WriteHeader(http.StatusInternalServerError)
	if cfg.Format == "plain" {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		_, _ = w.Write([]byte(defaultMessage))
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	payload := errorResponse{
		Success: false,
		Error: errorPayload{
			Code:    defaultErrorCode,
			Message: defaultMessage,
		},
	}
	_ = json.NewEncoder(w).Encode(payload)
}
