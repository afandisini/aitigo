# Agent Rules: Architecture

Purpose:

- Menjaga arsitektur berlapis tetap konsisten dan mudah diaudit.
- Memastikan domain tetap murni dan wiring tidak bocor ke layer lain.

Non-goals:

- Menentukan framework HTTP tertentu.
- Mengatur detail implementasi storage atau vendor SDK.

Acceptance criteria:

- `internal/domain` tidak mengimpor `internal/app`, `internal/infra`, `net/http`, router framework, atau driver DB.
- `internal/app` (controller) hanya berisi adapter HTTP, validasi ringan, dan mapping DTO.
- `internal/infra` hanya mengimplementasikan interface domain.
- Wiring dependency hanya di `cmd/*`.
- Error domain dipetakan ke HTTP status sesuai `CONVENTIONS.md`.

Checklist PR:

- [ ] Domain tidak mengimpor infra/http/router/DB driver.
- [ ] Controller tidak query DB langsung atau memuat logic bisnis kompleks.
- [ ] Infra tidak mengakses layer HTTP.
- [ ] Wiring dependency hanya di `cmd/*`.
- [ ] Error mapping mengikuti kode `DomainError` di `CONVENTIONS.md`.