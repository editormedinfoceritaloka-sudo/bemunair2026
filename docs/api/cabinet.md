# Cabinet Content

Modul ini menyediakan data publik Kabinet Cerita Loka, hierarki Kemenkoan dan Kementerian, pimpinan, program kerja, timeline, dokumentasi, serta endpoint admin untuk mengelola konten tersebut.

## Public API

Semua endpoint publik tidak memerlukan autentikasi dan hanya mengembalikan kabinet, unit, program kerja, serta media yang aktif dan dipublikasikan.

### GET /api/v1/cabinet

Mengambil kabinet aktif beserta Kemenkoan dan Kementerian di bawahnya.

### GET /api/v1/cabinet/:slug

Mengambil kabinet yang sudah dipublikasikan berdasarkan slug.

### GET /api/v1/cabinet/units/:slug

Mengambil detail Kemenkoan atau Kementerian berdasarkan slug. Response unit memuat profil pimpinan yang aktif, logo, cover, dan anak unit bila tersedia.

### GET /api/v1/cabinet/units/:slug/programs

Mengambil program kerja yang dipublikasikan milik unit tersebut.

| Query | Type | Default | Keterangan |
|---|---|---:|---|
| `page` | integer | 1 | Nomor halaman |

Ukuran halaman publik selalu 8 item. Parameter `per_page` tidak digunakan agar pagination konsisten dengan UI.

### GET /api/v1/cabinet/units/:slug/programs/:programSlug

Mengambil detail program kerja yang dipublikasikan, termasuk milestone dan dokumentasi yang aktif.

Program tanpa dokumentasi tetap mengembalikan `documentations: []`; frontend menampilkan empty state “Coming Soon”.

### GET /api/v1/cabinet/programs/:programSlug

Mengambil detail program kerja yang dipublikasikan hanya berdasarkan slug program, tanpa perlu slug unit. Response sama dengan endpoint `GET /api/v1/cabinet/units/:slug/programs/:programSlug` (termasuk milestone dan dokumentasi yang aktif).

Endpoint lama `GET /api/v1/cabinet/units/:slug/programs/:programSlug` tetap dipertahankan untuk kompatibilitas.

## Response Shape

Unit organisasi memakai `unit_type` `KEMENKOAN` atau `KEMENTERIAN`. Kementerian dapat memiliki profil `MINISTER` dan `DIRECTOR_GENERAL`; UI detail Kementerian menampilkan dua posisi tersebut.

Media menyimpan metadata file ImageKit, bukan binary file.

```json
{
  "status": true,
  "message": "Program kerja",
  "data": {
    "id": 910301,
    "ministry_id": 1,
    "name": "Cerita Visual",
    "slug": "cerita-visual",
    "status": "ONGOING",
    "milestones": [],
    "documentations": []
  }
}
```

## Admin API

Endpoint admin memerlukan JWT dan role admin. `ADMIN_MEDINFO` dapat mengelola seluruh unit dan dokumentasi. `ADMIN` hanya dapat mengelola konten unit yang ditugaskan melalui `user_organization_roles` atau unit yang sama dengan `users.ministry_id`.

| Method | Route | Akses | Tujuan |
|---|---|---|---|
| `GET` | `/api/v1/admin/cabinet-terms` | Admin | Daftar kabinet |
| `POST` | `/api/v1/admin/cabinet-terms` | `ADMIN_MEDINFO` | Membuat kabinet |
| `PUT` | `/api/v1/admin/cabinet-terms/:id` | `ADMIN_MEDINFO` | Memperbarui kabinet |
| `GET` | `/api/v1/admin/organizations` | Admin | Daftar unit |
| `POST` | `/api/v1/admin/organizations` | Admin sesuai scope | Membuat unit |
| `PUT` | `/api/v1/admin/organizations/:id` | Admin sesuai scope | Memperbarui unit |
| `GET` | `/api/v1/admin/organizations/:id/members` | Admin | Daftar Menteri/Dirjen |
| `POST` | `/api/v1/admin/organizations/:id/members` | Admin sesuai scope | Membuat profil pimpinan |
| `PUT` | `/api/v1/admin/organizations/members/:id` | Admin sesuai scope | Memperbarui profil pimpinan |
| `GET` | `/api/v1/admin/work-programs?unit_id=:id` | Admin | Daftar program kerja |
| `GET` | `/api/v1/admin/work-programs/:id` | Admin sesuai scope | Detail program kerja |
| `POST` | `/api/v1/admin/work-programs` | Admin sesuai scope | Membuat program kerja |
| `PUT` | `/api/v1/admin/work-programs/:id` | Admin sesuai scope | Memperbarui program kerja |
| `PUT` | `/api/v1/admin/work-programs/:id/publish` | Admin sesuai scope | Publish/unpublish program |
| `POST` | `/api/v1/admin/work-programs/:id/milestones` | Admin sesuai scope | Menambah milestone |
| `POST` | `/api/v1/admin/work-programs/:id/documentations` | `ADMIN` atau `ADMIN_MEDINFO` sesuai scope | Menambah relasi dokumentasi |
| `PUT` | `/api/v1/admin/work-programs/:id/documentations/reorder` | `ADMIN` atau `ADMIN_MEDINFO` sesuai scope | Mengatur urutan dokumentasi |
| `POST` | `/api/v1/admin/media-assets` | Admin sesuai scope | Menyimpan metadata hasil upload ImageKit |

## Documentation Request

Upload file memakai mekanisme ImageKit yang sudah tersedia. Setelah upload berhasil, simpan metadata media lalu hubungkan media tersebut ke program kerja.

```json
{
  "media_asset_id": 910106,
  "title": "Dokumentasi kegiatan",
  "caption": "Pelaksanaan program kerja",
  "display_order": 1,
  "is_cover": true
}
```

`media_asset_id` harus menunjuk media aktif. Satu program dapat memiliki banyak dokumentasi dan hanya relasinya yang diubah saat reorder atau remove.

## Status dan Validation

Status program kerja yang didukung: `DRAFT`, `PLANNED`, `ONGOING`, `COMPLETED`, `POSTPONED`, `CANCELLED`, dan `ARCHIVED`.

- Kemenkoan tidak memiliki parent.
- Kementerian wajib memiliki parent Kemenkoan dari kabinet yang sama.
- Hierarchy cycle ditolak.
- Slug unit unik dalam satu kabinet.
- Slug program unik di seluruh program kerja (global), divalidasi di service. Saat membuat program tanpa slug, slug dihasilkan otomatis dari nama + timestamp, misal `cerita-visual-20260807-153045`. Saat update dengan slug kosong, slug lama dipertahankan.
- Data unpublished tidak muncul di public API.
- Menghapus media yang masih direferensikan perlu dilakukan dengan menghapus relasi terlebih dahulu.

