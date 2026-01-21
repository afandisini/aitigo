# Agent: Infrastructure

## Rujukan Wajib

- [ARCHITECTURE](/ARCHITECTURE.md)
- [CONVENTIONS](/CONVENTIONS.md)
- [AGENT-CHECKLIST](/AGENT-CHECKLIST.md)

## Aturan Baca & Konflik

- Urutan baca wajib: root/AGENT.md -> /ARCHITECTURE.md -> /CONVENTIONS.md -> AGENT.md folder target.
- Konflik aturan: AGENT.md terdekat menang, KECUALI "Aturan Keras" di root/AGENT.md yang absolut.

## Tugas

Implementasi detail teknis:

- koneksi DB
- repo PostgreSQL/MySQL
- cache, storage, external API

## Aturan

- Implement interface domain, jangan bikin bisnis di sini.
- Semua query & mapping data di repo file masing-masing.
