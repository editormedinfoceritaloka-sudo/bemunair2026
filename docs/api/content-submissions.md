# Content Submissions

Pengajuan konten untuk platform `INSTAGRAM` atau `TWITTER`. Saat dibuat, sistem melakukan round-robin assign PJ dari queue Medinfo dan mencoba mengirim notifikasi WhatsApp. Kegagalan WA tidak menggagalkan pembuatan submission.

## Content Submission Object

```json
{
  "id": 1,
  "submitter_id": 2,
  "ministry": "MEDINFO",
  "platform": "INSTAGRAM",
  "submission_type": "Feed",
  "caption": "Caption konten",
  "deadline": "2026-06-10T12:00:00+07:00",
  "brief_file": "brief.pdf",
  "poster_file": "poster.png",
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
| `platform` | string | yes | `INSTAGRAM` atau `TWITTER` |
| `submission_type` | string | yes | Contoh: `Feed`, `Story`, `Thread` |
| `caption` | string | no | Caption konten |
| `deadline` | string | yes | RFC3339, contoh `2026-06-10T12:00:00+07:00` |
| `brief_file` | file | no | File brief |
| `poster_file` | file | no | File poster |

### Response 201

```json
{
  "success": true,
  "message": "Submission berhasil dibuat",
  "data": {
    "id": 1,
    "submitter_id": 2,
    "ministry": "MEDINFO",
    "platform": "INSTAGRAM",
    "submission_type": "Feed",
    "caption": "Caption konten",
    "deadline": "2026-06-10T12:00:00+07:00",
    "brief_file": "brief.pdf",
    "poster_file": "poster.png",
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
| 422 | `VALIDATION_ERROR` | Deadline bukan RFC3339 atau field wajib kosong |
| 500 | `INTERNAL_ERROR` | Database gagal menyimpan |

```json
{
  "success": false,
  "message": "Deadline harus RFC3339",
  "error": { "code": "VALIDATION_ERROR" }
}
```

### Curl

```bash
curl -X POST http://localhost:8081/api/content-submissions \
  -H "Authorization: Bearer $TOKEN" \
  -F ministry=MEDINFO \
  -F platform=INSTAGRAM \
  -F submission_type=Feed \
  -F caption="Caption konten" \
  -F deadline=2026-06-10T12:00:00+07:00 \
  -F brief_file=@brief.pdf \
  -F poster_file=@poster.png
```

---

## GET /api/content-submissions

Mengambil daftar pengajuan konten. `ADMIN` melihat semua. `MENTRI` hanya melihat milik sendiri atau ministry-nya.

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
      "platform": "INSTAGRAM",
      "submission_type": "Feed",
      "caption": "Caption konten",
      "deadline": "2026-06-10T12:00:00+07:00",
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
    "platform": "INSTAGRAM",
    "submission_type": "Feed",
    "caption": "Caption konten",
    "deadline": "2026-06-10T12:00:00+07:00",
    "brief_file": "brief.pdf",
    "poster_file": "poster.png",
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
