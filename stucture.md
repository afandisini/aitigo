### 📂 Project Structure

<details>
<summary>Klik untuk lihat detail struktur folder AitiGo</summary>

```text
aitigo/
├── .github/           # Konfigurasi CI/CD
│   └── workflows/
├── cmd/               # Entry point aplikasi
│   └── aitigo/        # Main main.go
├── internal/          # Kode privat (Logic Utama)
│   ├── app/           # Transport layer (HTTP)
│   │   └── http/
│   │       └── controller/
│   ├── domain/        # Business logic & Entities
│   │   ├── article/
│   │   ├── book/
│   │   └── user/
│   ├── infra/         # Implementasi infrastruktur (DB/Repo)
│   │   └── repository/
│   └── tooling/       # Helper & CLI internal
│       ├── checker/
│       └── cli/
</details>


-----

### Kenapa Pakai Cara Ini?

  * **Font Monospace:** Menggunakan block \`\`\`text membuat font jadi sejajar, jadi garis `├──` tidak akan bengkok.
  * **Hemat Ruang:** Dengan `<details>`, halaman utama README kamu tidak akan penuh sesak kalau foldernya makin banyak nanti.
  * **Professional Look:** Ini standar yang sering dipakai di dokumentasi framework Go populer.

**Next step buat framework AitiGo kamu:**
Setelah tampilannya rapi, apakah kamu mau aku bantu buatkan isi file `main.go` di folder `cmd/aitigo/` yang sudah mendukung **Graceful Shutdown** agar framework kamu tidak langsung mati saat ada error?
