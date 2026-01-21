# AGENT-CHECKLIST

## Sebelum Coding

- Baca urutan referensi wajib (lihat "Aturan Baca & Konflik").
- Pastikan perubahan tidak melanggar "Aturan Keras" root.
- Pahami boundary layer yang disentuh.
- Identifikasi file yang akan diubah, hindari membuat folder baru tanpa alasan.

## Saat Coding

- Ikuti konvensi naming dan response/error.
- Jaga controller tipis, domain murni, infra hanya implementasi.
- Pertahankan struktur folder yang ada.

## Sebelum Selesai

- Pastikan rujukan terbaru sudah dibaca ulang.
- Cek error/response konsisten.
- Pastikan perubahan siap commit dan tidak menyentuh file terlarang.

## Stop Conditions

- Jika melanggar "Aturan Keras" root -> STOP dan laporkan konflik.
- Jika butuh folder baru -> wajib alasan + update Decision Log di ARCHITECTURE.md.
