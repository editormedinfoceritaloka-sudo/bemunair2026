# Articles

Publikasi artikel (CMS) oleh medinfo. Artikel yang berstatus `PUBLISHED` dapat dibaca publik tanpa login. Operasi tulis (buat, edit, publish, hapus) hanya untuk medinfo (role `ADMIN`).

Body artikel berupa HTML dari rich text editor dan **disanitasi di backend** (bluemonday) sebelum disimpan — tag berbahaya seperti `<script>`, event handler (`onerror`, dll), dan `javascript:` URI dibuang otomatis.

## Article Object

```json
{
  "id": 1,
  "slug": "seminar-nasional-2026",
  "title": "Seminar Nasional 2026",
  "excerpt": "Ringkasan singkat artikel",
  "body": "<p>Isi artikel dalam HTML.</p>",
  "cover_image": "https://drive.google.com/...",
  "author_id": 2,
  "author": { "id": 2, "name": "Mentri Medinfo" },
  "status": "PUBLISHED",
  "published_at": "2026-07-21T10:00:00+07:00",
  "created_at": "2026-07-20T09:00:00+07:00",
  "updated_at": "2026-07-21T10:00:00+07:00"
}
```

`slug` dibuat otomatis dari `title` (unik; ditambah sufiks angka bila bentrok). Slug tidak berubah saat edit kecuali `title` diubah, agar URL artikel lama tetap stabil.

Status valid: `DRAFT`, `PUBLISHED`.

---

## GET /api/v1/articles

Daftar artikel `PUBLISHED` (publik, tanpa auth). Diurutkan dari `published_at` terbaru.

**Auth:** Tidak perlu

### Query Parameters

| Field | Type | Default | Keterangan |
|---|---|---:|---|
| `page` | integer | 1 | Halaman |
| `per_page` | integer | 10 | Item per halaman (maks 50) |

Response memakai bentuk ringkas (`ArticleListItem`) tanpa `body`.

### Response 200

```json
{
  "success": true,
  "message": "Daftar artikel",
  "data": [
    {
      "id": 1,
      "slug": "seminar-nasional-2026",
      "title": "Seminar Nasional 2026",
      "excerpt": "Ringkasan singkat artikel",
      "cover_image": "https://drive.google.com/...",
      "author_id": 2,
      "author": { "id": 2, "name": "Mentri Medinfo" },
      "status": "PUBLISHED",
      "published_at": "2026-07-21T10:00:00+07:00"
    }
  ],
  "meta": { "page": 1, "per_page": 10, "total": 1, "total_pages": 1 }
}
```

### Curl

```bash
curl "http://localhost:8081/api/v1/articles?page=1&per_page=10"
```

---

## GET /api/v1/articles/:slug

Detail satu artikel `PUBLISHED` (publik, tanpa auth). Artikel `DRAFT` mengembalikan 404.

**Auth:** Tidak perlu

### Path Parameters

| Field | Type | Required | Keterangan |
|---|---|---:|---|
| `slug` | string | yes | Slug artikel |

### Response 200

Mengembalikan Article Object lengkap (termasuk `body`).

### Error Responses

| Status | Code | Keterangan |
|---:|---|---|
| 404 | `NOT_FOUND` | Artikel tidak ditemukan atau belum dipublikasi |

### Curl

```bash
curl http://localhost:8081/api/v1/articles/seminar-nasional-2026
```

---

## GET /api/v1/admin/articles

Daftar SEMUA artikel termasuk `DRAFT` (medinfo). Diurutkan dari `created_at` terbaru.

**Auth:** ADMIN only

### Query Parameters

Sama seperti endpoint publik (`page`, `per_page`).

### Curl

```bash
curl http://localhost:8081/api/v1/admin/articles \
  -H "Authorization: Bearer $TOKEN"
```

---

## GET /api/v1/admin/articles/:id

Detail artikel berdasarkan ID (medinfo), berlaku untuk `DRAFT` maupun `PUBLISHED`.

**Auth:** ADMIN only

### Path Parameters

| Field | Type | Required | Keterangan |
|---|---|---:|---|
| `id` | integer | yes | ID artikel |

### Error Responses

| Status | Code | Keterangan |
|---:|---|---|
| 403 | `FORBIDDEN` | Role bukan `ADMIN` |
| 404 | `NOT_FOUND` | Artikel tidak ditemukan |

---

## POST /api/v1/admin/articles

Membuat artikel baru. Selalu tersimpan sebagai `DRAFT`.

**Auth:** ADMIN only
**Content-Type:** `application/json`

### Request Body

| Field | Type | Required | Keterangan |
|---|---|---:|---|
| `title` | string | yes | Judul artikel |
| `body` | string | yes | Isi HTML (akan disanitasi) |
| `excerpt` | string | no | Ringkasan singkat |
| `cover_image` | string | no | URL gambar sampul |

`author_id` diambil dari JWT, bukan dari body.

```json
{
  "title": "Seminar Nasional 2026",
  "excerpt": "Ringkasan singkat artikel",
  "body": "<p>Isi artikel.</p>",
  "cover_image": "https://drive.google.com/..."
}
```

### Response 201

Mengembalikan Article Object (status `DRAFT`, `published_at` null).

### Error Responses

| Status | Code | Keterangan |
|---:|---|---|
| 401 | `UNAUTHENTICATED` | Token tidak valid |
| 403 | `FORBIDDEN` | Role bukan `ADMIN` |
| 422 | `VALIDATION_ERROR` | `title` atau `body` kosong |

### Curl

```bash
curl -X POST http://localhost:8081/api/v1/admin/articles \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"title":"Seminar Nasional 2026","body":"<p>Isi artikel.</p>"}'
```

---

## PUT /api/v1/admin/articles/:id

Memperbarui artikel. Slug diregenerasi hanya jika `title` berubah.

**Auth:** ADMIN only
**Content-Type:** `application/json`

### Request Body

Sama seperti POST (`title`, `body` wajib; `excerpt`, `cover_image` opsional).

### Error Responses

| Status | Code | Keterangan |
|---:|---|---|
| 403 | `FORBIDDEN` | Role bukan `ADMIN` |
| 404 | `NOT_FOUND` | Artikel tidak ditemukan |
| 422 | `VALIDATION_ERROR` | `title` atau `body` kosong |

### Curl

```bash
curl -X PUT http://localhost:8081/api/v1/admin/articles/1 \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"title":"Judul Baru","body":"<p>Isi diperbarui.</p>"}'
```

---

## PUT /api/v1/admin/articles/:id/publish

Publikasikan atau batalkan publikasi artikel. Saat pertama kali dipublikasi, `published_at` di-set ke waktu sekarang; unpublish mengembalikan status ke `DRAFT` dan `published_at` menjadi null.

**Auth:** ADMIN only
**Content-Type:** `application/json`

### Request Body

| Field | Type | Required | Keterangan |
|---|---|---:|---|
| `published` | boolean | yes | `true` untuk publish, `false` untuk unpublish |

```json
{ "published": true }
```

### Error Responses

| Status | Code | Keterangan |
|---:|---|---|
| 403 | `FORBIDDEN` | Role bukan `ADMIN` |
| 404 | `NOT_FOUND` | Artikel tidak ditemukan |

### Curl

```bash
curl -X PUT http://localhost:8081/api/v1/admin/articles/1/publish \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"published":true}'
```

---

## DELETE /api/v1/admin/articles/:id

Menghapus artikel.

**Auth:** ADMIN only

### Error Responses

| Status | Code | Keterangan |
|---:|---|---|
| 403 | `FORBIDDEN` | Role bukan `ADMIN` |
| 500 | `INTERNAL_ERROR` | Gagal hapus |

### Curl

```bash
curl -X DELETE http://localhost:8081/api/v1/admin/articles/1 \
  -H "Authorization: Bearer $TOKEN"
```
