# Agent Rules: Integrations OIDC

Tujuan:

- Menyediakan helper OIDC discovery dan JWT validation.

Non-goals:

- Tidak membuat auth framework.

Public API:

- `auth.Discover`
- `auth.NewVerifier`
- `auth.VerifyJWT`
- `auth.BuildOAuth2Config`

Acceptance Criteria:

- Discovery menggunakan OIDC provider.
- JWT validation terpisah dari HTTP layer.
- Unit test minimal ada.

Checklist PR:

- [ ] Docs ter-update.
- [ ] Example compile.

Contoh penggunaan:

```go
provider, _ := auth.Discover(ctx, issuer)
verifier, _ := auth.NewVerifier(provider, clientID)
```
