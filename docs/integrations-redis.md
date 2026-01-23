# Redis Integration

Package: `pkg/integrations/redis`

Features:

- Client initialization from env
- Healthcheck ping
- Optional rate limit store adapter

Env:

- `REDIS_ADDR` (default `localhost:6379`)
- `REDIS_PASSWORD`
- `REDIS_DB`

Usage:

```go
cfg, _ := redis.FromEnv()
client := redis.NewClient(cfg)
_ = redis.Ping(context.Background(), client)
```
