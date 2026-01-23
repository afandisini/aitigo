# Tutorial: JWT Validation + Request ID + Logging

This tutorial shows how to combine request ID middleware, logging, and OIDC JWT validation.

```go
mux := http.NewServeMux()
mux.HandleFunc("/me", func(w http.ResponseWriter, r *http.Request) {
	logger := logging.With(r.Context())
	token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")

	verifier, _ := auth.NewVerifier(provider, "client-id")
	if _, err := auth.VerifyJWT(r.Context(), verifier, token); err != nil {
		logger.Warn("invalid token", "error", err)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	logger.Info("token verified")
	w.WriteHeader(http.StatusOK)
})

handler := requestid.Middleware(requestid.DefaultConfig())(mux)
```

Notes:

- Avoid logging raw tokens.
- Use `logging.With` to include request IDs.
