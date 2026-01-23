# Agent Rules: Middleware Recover

Tujuan:

- Menangkap panic dan mengembalikan 500 dengan aman.

Non-goals:

- Tidak membocorkan detail internal ke response.

Public API:

- `recovermw.Middleware`
- `recovermw.DefaultConfig`

Acceptance Criteria:

- Panic tidak membuat server crash.
- Response 500 sesuai format (json/plain).
- Stack trace dicatat ke logger.

Checklist PR:

- [ ] Test panic recovery ada.
- [ ] Response tidak memuat secret.
- [ ] Docs ter-update.

Contoh penggunaan:

```go
handler := recovermw.Middleware(recovermw.DefaultConfig())(mux)
```
