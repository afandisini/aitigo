# Agent Rules: Middleware Request ID

Tujuan:

- Menjamin setiap request punya ID yang konsisten.

Non-goals:

- Tidak mengatur format log secara global.

Public API:

- `requestid.Middleware`
- `requestid.DefaultConfig`
- `requestid.FromContext`

Acceptance Criteria:

- Header configurable, default `X-Request-ID`.
- ID masuk ke context dan response header.
- Unit test ada untuk generate dan propagate.

Checklist PR:

- [ ] Header default tetap sama.
- [ ] Tidak ada dependency non-standar.
- [ ] Test dan docs ada.

Contoh penggunaan:

```go
handler := requestid.Middleware(requestid.DefaultConfig())(mux)
```
