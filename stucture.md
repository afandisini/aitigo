# 📂 Project Structure

````text
aitigo/
├── .github/           # Konfigurasi CI/CD GitHub Actions
│   └── workflows/
├── cmd/               # Entry point aplikasi
│   └── aitigo/        # Main main.go aplikasi
├── internal/          # Kode privat (tidak bisa di-import project lain)
│   ├── app/           # Logika aplikasi (HTTP/gRPC transport)
│   │   └── http/      # Transport layer via HTTP
│   │       └── controller/
│   ├── domain/        # Business logic & Entities (Domain Driven Design)
│   │   ├── article/
│   │   ├── book/
│   │   └── user/
│   ├── infra/         # Implementasi infrastruktur (Database, Mailer, dll)
│   │   └── repository/ # Implementasi query database (Postgres)
│   └── tooling/       # Helper internal atau CLI tools
│       ├── checker/
│       └── cli/

---

### Tips Tambahan Biar makin "Bening":

1.  **Gunakan Code Block:** Selalu bungkus struktur folder kamu dengan triple backticks (\`\`\`) dan tambahkan label `text` atau `bash` di sebelahnya agar font-nya menjadi *monospace* (font khusus koding) yang sejajar secara vertikal.
2.  **Tambahkan Keterangan:** Jangan cuma list foldernya saja. Menambahkan deskripsi singkat di sebelah kanan folder (seperti contoh di atas) sangat membantu orang lain (atau kamu di masa depan) paham fungsi tiap bagian.
3.  **Gunakan Emoji:** Menambahkan emoji folder (📂) atau file (📄) di judul memberikan sentuhan modern pada dokumentasi kamu.
4.  **Fitur Collapsible (Opsional):** Kalau strukturnya makin panjang, kamu bisa pakai tag `<details>` biar bisa di-expand/collapse:

```html
<details>
  <summary>Klik untuk lihat detail struktur folder</summary>

  (Masukkan code block struktur di sini)

</details>
````
