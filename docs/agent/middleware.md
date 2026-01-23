# Agent Rules: Middleware

Tujuan:

- Menyediakan middleware net/http yang konsisten dan aman secara default.

Non-goals:

- Tidak mengikat ke framework router spesifik.
- Tidak menambahkan logic bisnis.

Public API:

- `Middleware` functions di `pkg/middleware/*`.
- Config structs dan `DefaultConfig`.

Acceptance Criteria:

- Middleware tidak memaksa dependency eksternal.
- Semua middleware punya unit test dan contoh.
- Default behavior aman (deny-all untuk CORS, limit konservatif).

Checklist PR:

- [ ] Tidak melanggar boundary layer.
- [ ] Test unit minimal ada.
- [ ] Example compile.
- [ ] Docs ter-update.

Contoh penggunaan:

```go
handler := requestid.Middleware(requestid.DefaultConfig())(mux)
```
