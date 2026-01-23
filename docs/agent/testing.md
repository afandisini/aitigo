# Agent Rules: Testing Utilities

Tujuan:

- Menyediakan helper testing ringan tanpa dependency berat.

Non-goals:

- Tidak menggantikan testing framework eksternal.

Public API:

- `testingutil.NewRequest`
- `testingutil.ExecuteRequest`
- `testingutil.RequireNoError`
- `testingutil.NewMockClock`
- `testingutil.LoadJSONFixture`

Acceptance Criteria:

- Helper mudah dipakai dan deterministic.
- Unit test minimal ada.
- Example test tersedia.

Checklist PR:

- [ ] Docs ter-update.
- [ ] Example test compile.

Contoh penggunaan:

```go
req := testingutil.NewRequest(http.MethodGet, "/health", nil)
```
