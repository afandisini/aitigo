# Recover Middleware

Package: `pkg/middleware/recover`

Behavior:

- Recovers from panics and returns HTTP 500.
- Logs stack trace using slog.
- Response format: JSON by default, plain text optional.

Usage:

```go
cfg := recovermw.DefaultConfig()
cfg.Format = "plain"

handler := recovermw.Middleware(cfg)(mux)
```
