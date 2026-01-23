# Agent Rules: Observability

Tujuan:

- Memberikan logging, tracing, dan metrics opsional yang mudah dipakai.

Non-goals:

- Tidak memaksa exporter tertentu.

Public API:

- `pkg/observability/logging`
- `pkg/observability/otel`
- `pkg/observability/metrics`

Acceptance Criteria:

- Semua fitur bisa dipakai tanpa konfigurasi berat.
- Integrasi request_id tersedia.
- Unit test minimal ada.

Checklist PR:

- [ ] Tidak menambah dependency wajib untuk core.
- [ ] Docs dan examples ada.
- [ ] Default aman.

Contoh penggunaan:

```go
logger := logging.New(logging.Config{Level: "info"})
logging.SetDefault(logger)
```
