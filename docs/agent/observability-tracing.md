# Agent Rules: Observability Tracing

Tujuan:

- Menyediakan setup OpenTelemetry sederhana.

Non-goals:

- Tidak memaksa exporter tertentu.

Public API:

- `otel.Setup`
- `otel.Middleware`

Acceptance Criteria:

- Stdout exporter default.
- OTLP optional via env.
- Middleware menghasilkan span per request.

Checklist PR:

- [ ] Test setup minimal ada.
- [ ] Docs ter-update.

Contoh penggunaan:

```go
shutdown, _ := otel.Setup(ctx, otel.Config{ServiceName: "svc"})
defer shutdown(ctx)
```
