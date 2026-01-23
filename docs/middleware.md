# Middleware Overview

Packages live under `pkg/middleware/*` and provide standard net/http middleware.

Included middleware:

- Request ID: `pkg/middleware/requestid`
- Recover: `pkg/middleware/recover`
- CORS: `pkg/middleware/cors`
- Rate limit: `pkg/middleware/ratelimit`

Example:

```go
handler := requestid.Middleware(requestid.DefaultConfig())(mux)
handler = recovermw.Middleware(recovermw.DefaultConfig())(handler)
handler = cors.Middleware(cors.Config{AllowOrigins: []string{"https://example.com"}})(handler)
handler = ratelimit.Middleware(ratelimit.DefaultConfig())(handler)
```
