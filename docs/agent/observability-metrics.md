# Agent Rules: Observability Metrics

Tujuan:

- Menyediakan Prometheus metrics dasar untuk HTTP.

Non-goals:

- Tidak membuat metric registries global baru secara otomatis.

Public API:

- `metrics.New`
- `(*Metrics).Handler`
- `(*Metrics).Middleware`

Acceptance Criteria:

- Metrics default tersedia.
- Handler `/metrics` dapat dipasang.
- Unit test minimal ada.

Checklist PR:

- [ ] Docs ter-update.
- [ ] Example compile.

Contoh penggunaan:

```go
metricsSvc, _ := metrics.New(metrics.Config{})
mux.Handle("/metrics", metricsSvc.Handler())
```
