# Agent Rules: Integrations Redis

Tujuan:

- Menyediakan helper client Redis dan optional rate limit store.

Non-goals:

- Tidak mengatur cache policy atau data modeling.

Public API:

- `redis.FromEnv`
- `redis.NewClient`
- `redis.Ping`
- `redis.NewRateLimitStore`

Acceptance Criteria:

- Env parsing aman.
- Adapter rate limit opsional.
- Unit test minimal ada.

Checklist PR:

- [ ] Docs ter-update.
- [ ] Example compile.

Contoh penggunaan:

```go
cfg, _ := redis.FromEnv()
client := redis.NewClient(cfg)
```
