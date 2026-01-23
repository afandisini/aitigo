# Logging (slog)

Package: `pkg/observability/logging`

Default logger uses JSON handler with configurable level.

Usage:

```go
logger := logging.New(logging.Config{Level: "info"})
logging.SetDefault(logger)

handler := requestid.Middleware(requestid.DefaultConfig())(mux)
```

Context enrichment:

```go
ctx := logging.WithUserID(r.Context(), "user-1")
logging.With(ctx).Info("request")
```
