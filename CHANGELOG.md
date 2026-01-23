## v0.2.0 (draft)

- Add middleware core: request ID, recover, CORS, rate limit
- Add observability packages: slog logging, OpenTelemetry tracing, Prometheus metrics
- Add migration runner CLI and database config helpers
- Add integrations: Redis, S3/MinIO, OIDC/OAuth2
- Add testing utilities and example tests
- Add documentation index, tutorials, and agent rules
- Build: Go toolchain 1.23.12 (dependency baseline)
- CLI: add `aitigo version` and improve `migrate` help UX
- Docs: add installation instructions and Windows build guidance

## v0.1.1

- Fix Nuxt TypeScript template config by properly importing defineNuxtConfig
- Add local Nuxt config type shim for template editing before dependency install
- Add placeholder .nuxt/tsconfig.json to avoid TS `extends` resolution errors

## v0.1.0

- First public release
- CLI scaffolding
- Boundary check
- GitHub Actions CI
