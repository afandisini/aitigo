# Agent Rules: Migrations

Tujuan:

- Menyediakan runner migrasi SQL yang sederhana.

Non-goals:

- Tidak menggantikan ORM atau migrator eksternal.

Public API:

- `migrate.NewRunner`
- `migrate.Load`
- `migrate.Create`

Acceptance Criteria:

- Format file up/down konsisten.
- CLI `aitigo migrate` bekerja.
- Test parsing migrasi ada.

Checklist PR:

- [ ] Docs ter-update.
- [ ] Example compile.

Contoh penggunaan:

```go
runner := migrate.NewRunner(db, "migrations")
_, _ = runner.Up(ctx)
```
