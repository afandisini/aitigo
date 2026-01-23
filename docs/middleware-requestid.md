# Request ID Middleware

Package: `pkg/middleware/requestid`

Behavior:

- Generates request ID when missing.
- Uses configurable header (default `X-Request-ID`).
- Injects request ID into context and response header.

Usage:

```go
cfg := requestid.DefaultConfig()
cfg.HeaderName = "X-Request-ID"

handler := requestid.Middleware(cfg)(mux)
```

Integration with logging:

```go
logger := logging.With(r.Context())
logger.Info("request started")
```
