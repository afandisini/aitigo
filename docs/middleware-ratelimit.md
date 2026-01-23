# Rate Limit Middleware

Package: `pkg/middleware/ratelimit`

Behavior:

- In-memory token bucket.
- Configurable rate, burst, and per-window duration.
- Optional Redis store adapter in `pkg/integrations/redis`.

Usage:

```go
handler := ratelimit.Middleware(ratelimit.Config{
	Limit: ratelimit.Limit{Rate: 10, Burst: 20, Per: time.Second},
})(mux)
```
