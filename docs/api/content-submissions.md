# Content Submissions

Pengajuan konten dengan tiga jenis: `FEEDS_REELS`, `INSTASTORY`, dan `ARTIKEL`. Saat dibuat, sistem melakukan round-robin assign PJ dari queue Medinfo dan mencoba mengirim notifikasi WhatsApp. Kegagalan WA tidak menggagalkan pembuatan submission.

## Jenis Pengajuan & Field

Field bersama untuk semua jenis: `title`, `add_song` (opsional), `caption`, `additional_notes`, `brief_link` (wajib).

`brief_link` berisi URL Google Docs brief yang dibuat pengaju (bukan file upload). Medinfo menyediakan template gdocs brief di luar sistem; pengaju membuat gdocs sendiri berisi brief lalu menempelkan link-nya di field ini.

| Field | `FEEDS_REELS` | `INSTASTORY` | `ARTIKEL` |
|---|:---:|:---:|:---:|
| `publish_date` | wajib | wajib | - |
| `publish_time` | wajib | wajib | - |
| `design_drive_link` | wajib | wajib | - |
| `canva_link` | wajib | wajib | - |
| `article_drive_link` | - | - | wajib |

## Deadline

Tidak ada field input `deadline`. Server menurunkannya otomatis:

- `FEEDS_REELS` / `INSTASTORY`: deadline diisi dari `publish_date` + `publish_time`.
- `ARTIKEL`: tidak punya tanggal publikasi, sehingga `deadline` bernilai `null`.

Konsekuensi: Artikel tetap menerima notifikasi WhatsApp saat pengajuan dibuat, tetapi tidak ikut reminder harian (cron jam 12 hanya memproses submission yang punya deadline). Pada daftar, submission tanpa deadline diurutkan paling bawah.

## Content Submission Object

```json
{
  "id": 1,
  "submitter_id": 2,
  "ministry": "MEDINFO",
  "submission_type": "FEEDS_REELS",
  "title": "Kampanye Hari Bumi",
  "add_song": "Coldplay - Paradise",
  "caption": "Caption konten",
  "additional_notes": "Tolong tag akun rektorat",
  "publish_date": "2026-06-10T00:00:00+07:00",
  "publish_time": "14:30",
  "design_drive_link": "https://drive.google.com/...",
  "canva_link": "https://canva.com/...",
  "article_drive_link": null,
  "deadline": "2026-06-10T14:30:00+07:00",
  "brief_link": "https://docs.google.com/document/d/...",
  "assigned_pj_id": 3,
  "status": "PENDING",
  "notes": null,
  "created_at": "2026-06-01T03:00:00+07:00",
  "updated_at": "2026-06-01T03:00:00+07:00"
}
```

Status valid: `PENDING`, `IN_REVIEW`, `APPROVED`, `REJECTED`.

---

## POST /api/content-submissions

Membuat pengajuan konten baru.

**Auth:** Authenticated  
**Content-Type:** `multipart/form-data`

### Form Data

| Field | Type | Required | Keterangan |
|---|---|---:|---|
| `ministry` | string | no | Jika kosong memakai ministry dari JWT |
| `submission_type` | string | yes | `FEEDS_REELS`, `INSTASTORY`, atau `ARTIKEL` |
| `title` | string | yes | Judul konten / nama kegiatan |
| `add_song` | string | no | Lagu yang ditambahkan (opsional) |
| `caption` | string | yes | Caption Instagram |
| `additional_notes` | string | no | Keterangan tambahan |
| `publish_date` | string | conditional | Format `YYYY-MM-DD`. Wajib untuk `FEEDS_REELS` / `INSTASTORY` |
| `publish_time` | string | conditional | Format `HH:MM` (24 jam). Wajib untuk `FEEDS_REELS` / `INSTASTORY` |
| `design_drive_link` | string | conditional | Link drive desain video/gambar. Wajib untuk `FEEDS_REELS` / `INSTASTORY` |
| `canva_link` | string | conditional | Link Canva template kementerian. Wajib untuk `FEEDS_REELS` / `INSTASTORY` |
| `article_drive_link` | string | conditional | Link GDrive kebutuhan artikel. Wajib untuk `ARTIKEL` |
| `brief_link` | string | yes | Link GDocs brief yang dibuat pengaju. Wajib untuk semua jenis |

### Response 201

```json
{
  "success": true,
  "message": "Submission berhasil dibuat",
  "data": {
    "id": 1,
    "submitter_id": 2,
    "ministry": "MEDINFO",
    "submission_type": "FEEDS_REELS",
    "title": "Kampanye Hari Bumi",
    "caption": "Caption konten",
    "publish_date": "2026-06-10T00:00:00+07:00",
    "publish_time": "14:30",
    "design_drive_link": "https://drive.google.com/...",
    "canva_link": "https://canva.com/...",
    "deadline": "2026-06-10T14:30:00+07:00",
    "brief_link": "https://docs.google.com/document/d/...",
    "assigned_pj_id": 3,
    "status": "PENDING",
    "notes": null
  }
}
```

### Error Responses

| Status | Code | Keterangan |
|---:|---|---|
| 401 | `UNAUTHENTICATED` | Token tidak valid |
| 422 | `VALIDATION_ERROR` | `publish_date` bukan `YYYY-MM-DD`, jenis tidak dikenal, atau field wajib per jenis kosong |
| 500 | `INTERNAL_ERROR` | Database gagal menyimpan |

```json
{
  "success": false,
  "message": "Validasi gagal",
  "error": "article_drive_link wajib diisi untuk Artikel"
}
```

### Curl

Feeds & Reels / Instastory:

```bash
curl -X POST http://localhost:8081/api/content-submissions \
  -H "Authorization: Bearer $TOKEN" \
  -F ministry=MEDINFO \
  -F submission_type=FEEDS_REELS \
  -F title="Kampanye Hari Bumi" \
  -F caption="Caption konten" \
  -F publish_date=2026-06-10 \
  -F publish_time=14:30 \
  -F design_drive_link="https://drive.google.com/..." \
  -F canva_link="https://canva.com/..." \
  -F brief_link="https://docs.google.com/document/d/..."
```

Artikel:

```bash
curl -X POST http://localhost:8081/api/content-submissions \
  -H "Authorization: Bearer $TOKEN" \
  -F ministry=MEDINFO \
  -F submission_type=ARTIKEL \
  -F title="Liputan Seminar Nasional" \
  -F caption="Caption artikel" \
  -F additional_notes="Penyelenggara: BEM UNAIR, tempat: Aula, waktu: 10 Juni" \
  -F article_drive_link="https://drive.google.com/..." \
  -F brief_link="https://docs.google.com/document/d/..."
```

---

## GET /api/content-submissions

Mengambil daftar pengajuan konten. `ADMIN` melihat semua. `MENTRI` hanya melihat milik sendiri atau ministry-nya. Submission tanpa deadline (Artikel) diurutkan paling bawah.

**Auth:** Authenticated  
**Content-Type:** none

### Request

Tidak ada query parameter pada implementasi saat ini.

### Response 200

```json
{
  "success": true,
  "message": "Daftar content submission",
  "data": [
    {
      "id": 1,
      "submitter_id": 2,
      "ministry": "MEDINFO",
      "submission_type": "FEEDS_REELS",
      "title": "Kampanye Hari Bumi",
      "caption": "Caption konten",
      "deadline": "2026-06-10T14:30:00+07:00",
      "assigned_pj_id": 3,
      "status": "PENDING",
      "submitter": { "id": 2, "name": "Mentri Medinfo" },
      "assigned_pj": { "id": 3, "name": "Mentri PSDM" }
    }
  ],
  "meta": { "page": 1, "per_page": 1, "total": 1, "total_pages": 1 }
}
```

### Error Responses

| Status | Code | Keterangan |
|---:|---|---|
| 401 | `UNAUTHENTICATED` | Token tidak valid |
| 500 | `INTERNAL_ERROR` | Query gagal |

### Curl

```bash
curl http://localhost:8081/api/content-submissions \
  -H "Authorization: Bearer $TOKEN"
```

---

## GET /api/content-submissions/:id

Mengambil detail pengajuan konten.

**Auth:** Authenticated  
**Content-Type:** none

### Path Parameters

| Field | Type | Required | Keterangan |
|---|---|---:|---|
| `id` | integer | yes | ID submission |

### Response 200

```json
{
  "success": true,
  "message": "Detail content submission",
  "data": {
    "id": 1,
    "submitter_id": 2,
    "ministry": "MEDINFO",
    "submission_type": "FEEDS_REELS",
    "title": "Kampanye Hari Bumi",
    "caption": "Caption konten",
    "additional_notes": "Tolong tag akun rektorat",
    "publish_date": "2026-06-10T00:00:00+07:00",
    "publish_time": "14:30",
    "design_drive_link": "https://drive.google.com/...",
    "canva_link": "https://canva.com/...",
    "deadline": "2026-06-10T14:30:00+07:00",
    "brief_link": "https://docs.google.com/document/d/...",
    "assigned_pj_id": 3,
    "status": "PENDING",
    "notes": null
  }
}
```

### Error Responses

| Status | Code | Keterangan |
|---:|---|---|
| 401 | `UNAUTHENTICATED` | Token tidak valid |
| 404 | `NOT_FOUND` | Submission tidak ditemukan |

### Curl

```bash
curl http://localhost:8081/api/content-submissions/1 \
  -H "Authorization: Bearer $TOKEN"
```

---

## PUT /api/content-submissions/:id/status

Memperbarui status pengajuan konten.

**Auth:** ADMIN only  
**Content-Type:** `application/json`

### Path Parameters

| Field | Type | Required | Keterangan |
|---|---|---:|---|
| `id` | integer | yes | ID submission |

### Request Body

| Field | Type | Required | Keterangan |
|---|---|---:|---|
| `status` | string | yes | `IN_REVIEW`, `APPROVED`, atau `REJECTED` |
| `notes` | string/null | no | Catatan admin |

Transisi valid: `PENDING -> IN_REVIEW -> APPROVED/REJECTED`.

```json
{
  "status": "IN_REVIEW",
  "notes": "Sedang dicek PJ"
}
```

### Response 200

```json
{
  "success": true,
  "message": "Status berhasil diperbarui",
  "data": {
    "id": 1,
    "status": "IN_REVIEW",
    "notes": "Sedang dicek PJ",
    "updated_at": "2026-06-01T03:10:00+07:00"
  }
}
```

### Error Responses

| Status | Code | Keterangan |
|---:|---|---|
| 403 | `FORBIDDEN` | Role bukan `ADMIN` |
| 422 | `VALIDATION_ERROR` | Status kosong |
| 409 | `CONFLICT` | Transisi status tidak valid |

```json
{
  "success": false,
  "message": "invalid transition",
  "error": { "code": "CONFLICT" }
}
```

### Curl

```bash
curl -X PUT http://localhost:8081/api/content-submissions/1/status \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"status":"IN_REVIEW","notes":"Sedang dicek PJ"}'
```

---

## DELETE /api/content-submissions/:id

Menghapus pengajuan konten.

**Auth:** ADMIN only  
**Content-Type:** none

### Path Parameters

| Field | Type | Required | Keterangan |
|---|---|---:|---|
| `id` | integer | yes | ID submission |

### Response 200

```json
{
  "success": true,
  "message": "Submission berhasil dihapus"
}
```

### Error Responses

| Status | Code | Keterangan |
|---:|---|---|
| 403 | `FORBIDDEN` | Role bukan `ADMIN` |
| 500 | `INTERNAL_ERROR` | Gagal hapus |

### Curl

```bash
curl -X DELETE http://localhost:8081/api/content-submissions/1 \
  -H "Authorization: Bearer $TOKEN"
```
