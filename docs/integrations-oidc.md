# OIDC / OAuth2 Integration

Package: `pkg/integrations/auth`

Features:

- OIDC discovery
- JWT verification helper
- OAuth2 client config builder

Usage:

```go
provider, _ := auth.Discover(ctx, "https://issuer.example.com")
verifier, _ := auth.NewVerifier(provider, "client-id")
token, _ := auth.VerifyJWT(ctx, verifier, rawToken)
_ = token
```
