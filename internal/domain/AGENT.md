# Agent: Domain Layer

## Rujukan Wajib

- [ARCHITECTURE](/ARCHITECTURE.md)
- [CONVENTIONS](/CONVENTIONS.md)
- [AGENT-CHECKLIST](/AGENT-CHECKLIST.md)

## Aturan Baca & Konflik

- Urutan baca wajib: root/AGENT.md -> /ARCHITECTURE.md -> /CONVENTIONS.md -> AGENT.md folder target.
- Konflik aturan: AGENT.md terdekat menang, KECUALI "Aturan Keras" di root/AGENT.md yang absolut.

## Tugas

Menulis bisnis inti: entity, value object, interface repository, service.

## Boleh

- entity.go, service.go, repository.go
- pure logic, tanpa HTTP/DB/driver

## Dilarang

- import gin/echo/fiber, sqlx/gorm, net/http handler
- akses env/config langsung

## Pola

- Repository berupa interface
- Service menerima interface repository
