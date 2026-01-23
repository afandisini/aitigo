# Agent Rules: Integrations

Tujuan:

- Menyediakan adapter optional untuk layanan pihak ketiga.

Non-goals:

- Tidak membuat framework auth atau storage penuh.

Public API:

- `pkg/integrations/redis`
- `pkg/integrations/s3`
- `pkg/integrations/auth`

Acceptance Criteria:

- Semua adapter optional dan terpisah.
- Config dapat diambil dari env.
- Example tersedia.

Checklist PR:

- [ ] Docs ter-update.
- [ ] Test unit minimal ada.

Contoh penggunaan:

```go
cfg, _ := redis.FromEnv()
client := redis.NewClient(cfg)
```
