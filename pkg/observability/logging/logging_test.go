package logging

import (
	"bytes"
	"context"
	"log/slog"
	"strings"
	"testing"

	"aitigo/pkg/middleware/requestid"
)

func TestWithAddsContextFields(t *testing.T) {
	var buf bytes.Buffer
	logger := slog.New(slog.NewJSONHandler(&buf, &slog.HandlerOptions{Level: slog.LevelInfo}))

	ctx := context.Background()
	ctx = IntoContext(ctx, logger)
	ctx = requestid.WithContext(ctx, "req-123")
	ctx = WithUserID(ctx, "user-9")

	With(ctx).Info("hello")

	out := buf.String()
	if !strings.Contains(out, `"request_id":"req-123"`) {
		t.Fatalf("expected request_id in log, got %s", out)
	}
	if !strings.Contains(out, `"user_id":"user-9"`) {
		t.Fatalf("expected user_id in log, got %s", out)
	}
}
