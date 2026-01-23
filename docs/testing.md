# Testing Utilities

Package: `pkg/testingutil`

Included helpers:

- HTTP request/response helpers
- Minimal assertions
- Mock clock and UUID generator
- JSON fixture loader

Example:

```go
req := testingutil.NewRequest(http.MethodGet, "/health", nil)
rec := testingutil.ExecuteRequest(handler, req)
testingutil.RequireEqual(t, rec.Code, http.StatusOK, "status")
```
