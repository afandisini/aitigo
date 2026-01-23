# Agent Rules: Middleware Rate Limit

Tujuan:

- Membatasi request agar server tetap stabil.

Non-goals:

- Tidak mengatur auth atau quota billing.

Public API:

- `ratelimit.Middleware`
- `ratelimit.Config`
- `ratelimit.Limit`
- `ratelimit.Store`

Acceptance Criteria:

- Default rate limit aman.
- Response 429 JSON.
- In-memory store tersedia.

Checklist PR:

- [ ] Test limit ada.
- [ ] Redis adapter opsional (di package integrations).
- [ ] Docs ter-update.

Contoh penggunaan:

```go
handler := ratelimit.Middleware(ratelimit.DefaultConfig())(mux)
```
