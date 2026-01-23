# Agent Rules: Observability Logging

Tujuan:

- Menyediakan structured logging via slog.

Non-goals:

- Tidak membuat logging framework baru.

Public API:

- `logging.New`
- `logging.SetDefault`
- `logging.With`
- `logging.WithUserID`

Acceptance Criteria:

- JSON handler default.
- Level configurable.
- request_id otomatis terpasang jika ada.

Checklist PR:

- [ ] Test log fields ada.
- [ ] Docs ter-update.

Contoh penggunaan:

```go
logger := logging.New(logging.Config{Level: "info"})
logging.SetDefault(logger)
```
