package requestid

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"net/http"
)

const defaultHeader = "X-Request-ID"

type contextKey struct{}

type Config struct {
	HeaderName string
	Generator  func() string
}

func DefaultConfig() Config {
	return Config{
		HeaderName: defaultHeader,
		Generator:  defaultGenerator,
	}
}

func Middleware(cfg Config) func(http.Handler) http.Handler {
	if cfg.HeaderName == "" {
		cfg.HeaderName = defaultHeader
	}
	if cfg.Generator == nil {
		cfg.Generator = defaultGenerator
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			id := r.Header.Get(cfg.HeaderName)
			if id == "" {
				id = cfg.Generator()
			}

			ctx := context.WithValue(r.Context(), contextKey{}, id)
			r = r.WithContext(ctx)
			w.Header().Set(cfg.HeaderName, id)
			next.ServeHTTP(w, r)
		})
	}
}

func FromContext(ctx context.Context) (string, bool) {
	value := ctx.Value(contextKey{})
	id, ok := value.(string)
	return id, ok
}

func WithContext(ctx context.Context, id string) context.Context {
	if id == "" {
		return ctx
	}
	return context.WithValue(ctx, contextKey{}, id)
}

func defaultGenerator() string {
	var buf [16]byte
	if _, err := rand.Read(buf[:]); err != nil {
		return "unknown"
	}
	return hex.EncodeToString(buf[:])
}
