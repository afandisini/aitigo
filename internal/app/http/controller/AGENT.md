# Agent: HTTP Controller

## Rujukan Wajib

- [ARCHITECTURE](/ARCHITECTURE.md)
- [CONVENTIONS](/CONVENTIONS.md)
- [AGENT-CHECKLIST](/AGENT-CHECKLIST.md)

## Aturan Baca & Konflik

- Urutan baca wajib: root/AGENT.md -> /ARCHITECTURE.md -> /CONVENTIONS.md -> AGENT.md folder target.
- Konflik aturan: AGENT.md terdekat menang, KECUALI "Aturan Keras" di root/AGENT.md yang absolut.

## Tugas

- Bind request (JSON/form)
- Validasi ringan
- Panggil domain service/usecase
- Return JSON/HTML

## Dilarang

- Query DB langsung
- Logic bisnis kompleks
- Mengakses env secara langsung (pakai config provider)
