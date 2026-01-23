# Agent Rules: Middleware CORS

Tujuan:

- Mengelola header CORS dan preflight secara aman.

Non-goals:

- Tidak mengelola auth atau caching.

Public API:

- `cors.Middleware`
- `cors.Config`

Acceptance Criteria:

- Default deny-all.
- Preflight ditangani dengan benar.
- Config dapat mengatur origins/methods/headers.

Checklist PR:

- [ ] Test preflight ada.
- [ ] Behavior default aman.
- [ ] Docs ter-update.

Contoh penggunaan:

```go
handler := cors.Middleware(cors.Config{
	AllowOrigins: []string{"https://example.com"},
})(mux)
```
