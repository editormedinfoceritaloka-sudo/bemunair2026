# SOP Form Media dan Assignment PJ

Implementasi ini menggunakan `media_submission_settings` sebagai sumber URL SOP/template. Form Konten dan Artikel langsung membuka dialog SOP wajib sebelum field dapat diakses.

## Referensi aktif

- SOP Konten/Artikel: `https://docs.google.com/document/d/1RgIdOLXwUJz6pRrktK2fLQ6srvCZn0y8/edit`
- Template kementerian/artikel: `https://linktr.ee/templateeditingkementerian`
- Brief Konten: `https://docs.google.com/document/d/1bBIs7dcqL1OMSJWAoXZduy2c6r0MpvLIC84T5t3OZJ8/edit?tab=t.0`
- Caption Artikel: `https://docs.google.com/document/d/1--ki9SXtg1tTQQFFNQElnQ4brK1c8mMA6Ww540XCKzk/edit?tab=t.0`

Konten memakai minimum H-7 dan Artikel H+3. Feed/Instastory mengunggah gambar ke ImageKit, sedangkan video Reels dikirim melalui Google Drive. Brief tetap diunggah ke ImageKit.

## Assignment

Pengajuan baru tidak memiliki PJ otomatis. `ADMIN_MEDINFO` memilih PJ melalui detail request. Backend mengunci row roster dan menolak assignment apabila target sedang menangani task aktif pada media maupun surat. Assignment dan reassignment dicatat pada timeline.

Endpoint:

- `PUT /api/v1/content-submissions/:id/assignee`
- `PUT /api/v1/letter-submissions/:id/assignee`
- `GET /api/v1/medinfo-pj/queue` mengembalikan status availability dan task aktif.
