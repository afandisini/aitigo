# Agent Rules: Integrations S3

Tujuan:

- Menyediakan wrapper sederhana untuk put/get/presign.

Non-goals:

- Tidak mengelola bucket lifecycle atau ACL kompleks.

Public API:

- `s3.FromEnv`
- `s3.New`
- `(*Client).Put`
- `(*Client).Get`
- `(*Client).PresignGet`

Acceptance Criteria:

- Config dari env.
- Error handling jelas.
- Unit test minimal ada.

Checklist PR:

- [ ] Docs ter-update.
- [ ] Example compile.

Contoh penggunaan:

```go
cfg, _ := s3.FromEnv()
client, _ := s3.New(cfg)
```
