package logging

import (
	"context"
	"log/slog"
	"os"
	"strings"

	"aitigo/pkg/middleware/requestid"
)

type Config struct {
	Level     string
	AddSource bool
}

type contextKey struct{}
type userIDKey struct{}

func New(config Config) *slog.Logger {
	level := parseLevel(config.Level)
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     level,
		AddSource: config.AddSource,
	})
	return slog.New(handler)
}

func SetDefault(logger *slog.Logger) {
	if logger == nil {
		return
	}
	slog.SetDefault(logger)
}

func With(ctx context.Context) *slog.Logger {
	logger := FromContext(ctx)
	if logger == nil {
		logger = slog.Default()
	}
	if id, ok := requestid.FromContext(ctx); ok {
		logger = logger.With("request_id", id)
	}
	if userID, ok := UserIDFromContext(ctx); ok {
		logger = logger.With("user_id", userID)
	}
	return logger
}

func IntoContext(ctx context.Context, logger *slog.Logger) context.Context {
	if logger == nil {
		return ctx
	}
	return context.WithValue(ctx, contextKey{}, logger)
}

func FromContext(ctx context.Context) *slog.Logger {
	value := ctx.Value(contextKey{})
	logger, _ := value.(*slog.Logger)
	return logger
}

func WithUserID(ctx context.Context, userID string) context.Context {
	if userID == "" {
		return ctx
	}
	return context.WithValue(ctx, userIDKey{}, userID)
}

func UserIDFromContext(ctx context.Context) (string, bool) {
	value := ctx.Value(userIDKey{})
	userID, ok := value.(string)
	return userID, ok
}

func parseLevel(level string) slog.Level {
	switch strings.ToLower(level) {
	case "debug":
		return slog.LevelDebug
	case "warn", "warning":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
