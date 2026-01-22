### Project Structure

<details>
<summary>Klik untuk lihat detail struktur folder AitiGo</summary>

```text
aitigo/
|-- .github/               # Konfigurasi CI/CD
|   `-- workflows/
|-- cmd/                   # Entry point aplikasi
|   `-- aitigo/            # Main main.go
|-- internal/              # Kode privat (logic utama)
|   |-- app/               # Transport layer (HTTP)
|   |   `-- http/
|   |       `-- controller/
|   |-- domain/            # Business logic & entities
|   |   |-- article/
|   |   |-- book/
|   |   `-- user/
|   |-- infra/             # Implementasi infrastruktur (DB/repo)
|   |   `-- repository/
|   `-- tooling/           # Helper & CLI internal
|       |-- checker/
|       `-- cli/
|-- templates/             # Template project (next-ts, nuxt-ts)
|   `-- nuxt-ts/           # Nuxt 3 + TypeScript starter
|       `-- .nuxt/         # Placeholder tsconfig untuk editor
```

</details>
