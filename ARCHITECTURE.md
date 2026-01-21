# ARCHITECTURE

## Tujuan

AITIGO adalah kerangka Go bergaya MVC-like untuk layanan HTTP. Prinsip utamanya: domain murni, controller tipis, dan detail teknis dipisah ke infra.

## Layer dan Boundary

- `internal/app/http/controller`: HTTP adapter, parsing/validasi ringan, mapping DTO.
- `internal/domain`: entity, value object, service, dan repository interface.
- `internal/infra`: implementasi teknis (DB, cache, external API) yang memenuhi interface domain.
- `cmd`: wiring/entrypoint aplikasi.

## Diagram Boundary (Request -> DB)

```mermaid
flowchart LR
  Client --> C[HTTP Controller]
  C --> S[Domain Service/Usecase]
  S --> RI[Repository Interface]
  RI --> R[Infra Repository Impl]
  R --> DB[(Database)]

  subgraph App
    C
  end
  subgraph Domain
    S
    RI
  end
  subgraph Infra
    R
    DB
  end
```

## Diagram Aturan Dependency

```mermaid
flowchart TD
  CMD[cmd/*] --> APP[internal/app]
  CMD --> INFRA[internal/infra]
  APP --> DOMAIN[internal/domain]
  INFRA --> DOMAIN
  DOMAIN --> DOMAIN
```

## Aturan Dependency

- `internal/domain` tidak boleh import `internal/infra`, `internal/app/http`, `net/http` handler, framework router, atau DB driver.
- `internal/app` boleh import `internal/domain` dan contract interface, tidak boleh akses DB langsung.
- `internal/infra` boleh import `internal/domain` untuk implementasi interface.
- Wiring dependensi hanya di `cmd/*` atau layer entrypoint.

## Pola Modul (Contoh: user)

- `internal/domain/user/entity.go` (User entity)
- `internal/domain/user/repository.go` (interface `UserRepository`)
- `internal/domain/user/service.go` (service/usecase)
- `internal/app/http/controller/user_controller.go` memanggil `UserService`
- `internal/infra/repository/user_repo.go` mengimplement `UserRepository`

Alur: Controller -> Service -> Repository Interface -> Infra Repo -> DB.

## Decision Log (Ringkas)

- `internal/`: mencegah package dipakai di luar module, menjaga boundary.
- `cmd/`: entrypoint per service/binary, wiring hanya di sini.
- Repository interface di domain: agar domain tetap murni dan infra dapat diganti tanpa ubah bisnis.
- `internal/tooling/cli`: CLI generator internal untuk scaffolding tanpa dependency eksternal.
