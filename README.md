# AitiGo

[![CI](https://github.com/afandisini/AitiGo/actions/workflows/ci.yml/badge.svg)](https://github.com/afandisini/AitiGo/actions/workflows/ci.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/afandisini/AitiGo)](https://goreportcard.com/report/github.com/afandisini/AitiGo) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**AitiGo** adalah kerangka backend berbasis **Go** dengan gaya **MVC-like berlapis**, dibuat untuk konsistensi arsitektur, domain yang murni, dan kolaborasi AI yang tidak menyimpang.

> Dibangun untuk jangka panjang.
> Controller tipis. Domain bersih. Infra terkontrol.

---

## Tujuan dan Filosofi

- **Domain murni**: domain tidak tahu HTTP, DB, atau framework.
- **Controller tipis**: hanya parsing, validasi ringan, dan mapping DTO.
- **Infra terpisah**: detail teknis dibatasi di layer infra.
- **Anti magic**: struktur jelas, tidak ada arsitektur improvisasi.

---

## Struktur Folder Inti

```
aitigo/
|-- cmd/                         # Entrypoint & wiring
|-- internal/
|   |-- app/                     # HTTP layer, kernel
|   |-- domain/                  # Entity, service, repo interface
|   `-- infra/                   # DB, cache, external implementation
|-- templates/                   # Project templates (next-ts, nuxt-ts)
|-- ARCHITECTURE.md
|-- CONVENTIONS.md
|-- AGENT.md                     # Aturan keras (absolut)
|-- AGENT-CHECKLIST.md
`-- README.md
```

Detail lengkap lihat `ARCHITECTURE.md` dan `CONVENTIONS.md`.

---

## Build & Run CLI

Jalankan langsung:

```
go run ./cmd/aitigo help
```

Atau build binary:

```
go build -o aitigo ./cmd/aitigo
```

---

## CLI Commands

```
aitigo help
aitigo check
aitigo make:module <module>
aitigo make:controller <Name> --module <module> [--force]
aitigo make:service <Name> --module <module> [--force]
aitigo make:repository <Name> --module <module> [--force]
aitigo make:crud <module> [--force]
```

Contoh:

```
aitigo make:module user
aitigo make:controller UserController --module user
aitigo make:service UserService --module user
aitigo make:repository UserRepository --module user
aitigo make:crud article
aitigo check
```

---

## How to use AitiGo with Next.js (next-ts)

List template:

```
aitigo templates
```

Init project:

```
aitigo init next-ts myapp
aitigo init next-ts .\myapp
```

Run dev server:

```
cd myapp
pnpm i
pnpm dev
```

Run checks:

```
cd ..
aitigo check ./myapp
```

---

## How to use AitiGo with Nuxt 3 (nuxt-ts)

List template:

```
aitigo templates
```

Init project:

```
aitigo init nuxt-ts myapp
aitigo init nuxt-ts .\myapp
```

Run dev server:

```
cd myapp
pnpm i
pnpm dev
```

Run checks:

```
cd ..
aitigo check ./myapp
```

Catatan Windows:
- Gunakan path relatif `.\myapp` jika perlu.
- Pastikan `pnpm` tersedia di PATH.

Catatan template Nuxt:
- `templates/nuxt-ts` menyertakan shim `nuxt/config` dan placeholder `.nuxt/tsconfig.json` agar editor tidak error sebelum install dependency.

---

## Workflow Rekomendasi (Contoh)

1. Buat module baru:

```
aitigo make:crud article
```

2. Tambahkan handler di controller `internal/app/http/controller/article_controller.go`.

3. Implement usecase di `internal/domain/article`.

4. Implement repository di `internal/infra/repository`.

5. Wiring dependensi di `cmd/*`.

---

## Hasil Generate (Gambaran)

```
internal/domain/article/entity.go
internal/domain/article/repository.go
internal/domain/article/service.go
internal/domain/article/article_service.go
internal/domain/article/article_repository.go
internal/infra/repository/article_repository_impl.go
internal/app/http/controller/article_controller.go
```

---

## Boundary Rules (Ringkas)

- Domain **tidak boleh** import HTTP/infra/DB driver.
- Controller **tidak boleh** query DB atau logic bisnis kompleks.
- Infra **tidak boleh** mengubah aturan domain.

Rujuk detail di `ARCHITECTURE.md` dan `CONVENTIONS.md`.

---

## Donasi

### Donasi & Beli Kopi

Kalau AitiGo ngebantu kerjaanmu dan bikin hidup sedikit lebih waras,
boleh traktir kopi biar maintainer kuat begadang.

[![Buy me a coffee](https://img.icons8.com/emoji/96/hot-beverage.png)](https://saweria.co/aitisolutions)

https://saweria.co/aitisolutions
