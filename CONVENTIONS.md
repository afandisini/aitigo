# CONVENTIONS

## Naming

- Folder/package: lowercase singkat, satu kata (`user`, `auth`, `health`).
- File: `snake_case.go` (`user_controller.go`).
- Struct: `PascalCase` (`UserProfile`).
- Interface: `PascalCase` + peran (`UserRepository`, `AuthService`).
- Method/func: `PascalCase` untuk export, `camelCase` untuk non-export.

## Controller/Endpoint

- Controller tipis: hanya parsing, validasi ringan, mapping request/response DTO.
- Business logic ada di domain service/usecase.
- Hindari akses DB, config/env, atau SDK langsung di controller.

## Error Standar

DomainError:

- `Code` (string, stabil untuk client)
- `Message` (human-readable)
- `Details` (opsional, untuk debug)

Mapping ke HTTP status:

- `BAD_REQUEST` -> 400
- `UNAUTHORIZED` -> 401
- `FORBIDDEN` -> 403
- `NOT_FOUND` -> 404
- `CONFLICT` -> 409
- `INTERNAL` -> 500

## Response JSON Standar

Sukses:

```json
{ "success": true, "data": {}, "meta": {} }
```

Error:

```json
{
  "success": false,
  "error": { "code": "BAD_REQUEST", "message": "invalid payload", "details": {} }
}
```

## Logging

Minimal fields:

- `request_id` / `trace_id`
- `method`, `path`, `status`
- `latency_ms`
- `error_code` (jika ada)

## Testing

- Domain/service: unit test wajib.
- Repo: integration test opsional (gunakan DB test jika diperlukan).
