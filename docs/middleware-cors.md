# CORS Middleware

Package: `pkg/middleware/cors`

Behavior:

- Deny-all by default until configured.
- Supports origins, methods, headers, credentials, max-age.
- Handles preflight requests.

Usage:

```go
handler := cors.Middleware(cors.Config{
	AllowOrigins:     []string{"https://example.com"},
	AllowMethods:     []string{"GET", "POST"},
	AllowHeaders:     []string{"Content-Type"},
	AllowCredentials: true,
	MaxAgeSeconds:    600,
})(mux)
```
