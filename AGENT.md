# AitiGo Agent (ROOT · ABSOLUT)

Dokumen ini adalah **otoritas tertinggi** dalam repo AitiGo.
Semua AGENT.md lain, AI agent, tool, dan manusia **WAJIB patuh**.

Jika terjadi konflik aturan:
➡️ **Aturan Keras di dokumen ini SELALU MENANG.**

---

## Rujukan Wajib (Wajib Dibaca Berurutan)

1. Dokumen ini (`root/AGENT.md`)
2. [`ARCHITECTURE.md`](./ARCHITECTURE.md)
3. [`CONVENTIONS.md`](./CONVENTIONS.md)
4. `AGENT.md` di folder target (jika ada)

❗ Melanggar urutan baca = desain dianggap tidak sah.

---

## Identitas AitiGo

- Nama: **AitiGo**
- Bahasa: Go
- Gaya: **MVC-like berlapis (Layered Architecture)**
- Prinsip inti:
  - Domain **murni**
  - Controller **tipis**
  - Infra **detail teknis**
  - Tidak ada “magic”, tidak ada arsitektur improvisasi

---

## ATURAN KERAS (ABSOLUT)

### 1. Aturan Boundary Layer (TIDAK BISA DITAWAR)

#### ❌ Domain (`internal/domain`) DILARANG:
- Import:
  - `internal/infra/*`
  - `internal/app/*`
  - `net/http`
  - framework router (gin, echo, fiber, dll)
  - ORM / DB driver (`gorm`, `sql`, `sqlx`, dll)
  - env/config loader
- Akses:
  - HTTP request/response
  - DB connection langsung
  - cache / external SDK

➡️ Domain **HANYA BOLEH**:
- entity
- value object
- service/usecase
- interface repository
- error domain

---

#### ❌ Controller (`internal/app/http/controller`) DILARANG:
- Query DB langsung
- Logic bisnis kompleks
- Implementasi repository
- Import DB driver / ORM
- Import package infra

➡️ Controller **HANYA BOLEH**:
- Parsing request
- Validasi ringan
- Panggil domain service
- Mapping response DTO

---

#### ❌ Infra (`internal/infra`) DILARANG:
- Menyimpan logic bisnis
- Mengubah aturan domain
- Mengakses HTTP layer

➡️ Infra **HANYA BOLEH**:
- Implement interface domain
- Mapping data teknis (DB, cache, external API)

---

### 2. Aturan Dependency (WAJIB)

- `internal/domain` -> **tidak bergantung ke siapa pun**
- `internal/app` -> boleh ke `internal/domain`
- `internal/infra` -> boleh ke `internal/domain`
- Wiring dependency **HANYA** di:
  - `cmd/*`
  - atau entrypoint kernel/app

Jika import graph melanggar -> **STOP. Jangan lanjut coding.**

---

### 3. Aturan Folder & File (ANTI SAMPAH)

❌ DILARANG MEMBUAT:
- Folder/filename acak:
  - `helpers_final`
  - `fix_fix`
  - `temp`
  - `coba`
  - `utils2`
  - `misc`
- Duplikasi layer (misal `service` di infra)

➡️ Folder baru **HANYA BOLEH** jika:
1. Ada alasan arsitektural
2. Dicatat di `ARCHITECTURE.md` (Decision Log)
3. Tidak melanggar boundary

---

### 4. Aturan Error & Response (WAJIB SERAGAM)

- Semua error bisnis = **DomainError**
- Tidak boleh return error mentah ke client
- Mapping error ke HTTP status **WAJIB** mengikuti `CONVENTIONS.md`

Response API **WAJIB** mengikuti format standar:
- `success: true|false`
- Tidak boleh format custom per endpoint

---

### 5. Aturan Konflik Agent

- Jika agent folder **berkonflik** dengan agent lain:
  - Agent terdekat dengan folder target menang
- **KECUALI**:
  - Jika konflik dengan **Aturan Keras** di dokumen ini
    -> **Aturan ini SELALU MENANG**

---

### 6. Aturan Kerja AI (WAJIB DIPATUHI)

Setiap kali AI mengerjakan task:
1. Tentukan layer target (domain/app/infra)
2. Sebutkan file yang akan diubah/dibuat
3. Pastikan tidak melanggar boundary
4. Maksimal perubahan relevan (hindari scope melebar)
5. Tidak “kreatif” di luar arsitektur

Jika instruksi user:
- Bertentangan dengan aturan ini -> **STOP & LAPORKAN**
- Meminta shortcut melanggar arsitektur -> **TOLAK DENGAN ALASAN**

---

## Prinsip Penutup

> AitiGo lebih mementingkan **konsistensi jangka panjang**
> daripada kecepatan sesaat.

Kode boleh lambat ditulis,
tapi **arsitektur tidak boleh rusak**.
