# Letter Submissions

Pengajuan surat digital. Jika `ministry` bernilai `MEDINFO`, sistem auto-assign PJ dari queue Medinfo dan mengirim notifikasi WA.

## Letter Submission Object

```json
{
  "id": 1,
  "submitter_id": 2,
  "ministry": "MEDINFO",
  "letter_type": "Undangan",
  "subject": "Undangan Rapat",
  "body": "Isi surat",
  "deadline": "2026-06-10T12:00:00+07:00",
  "assigned_pj_id": 3,
  "status": "PENDING",
  "notes": null,
  "created_at": "2026-06-01T03:00:00+07:00",
  "updated_at": "2026-06-01T03:00:00+07:00"
}
```

---

## POST /api/letter-submissions

Membuat pengajuan surat.

**Auth:** Authenticated  
**Content-Type:** `application/json`

### Request Body

| Field | Type | Required | Keterangan |
|---|---|---:|---|
| `ministry` | string | no | Jika kosong memakai ministry dari JWT |
| `letterType` | string | yes | Jenis surat. Field JSON mengikuti handler saat ini |
| `subject` | string | yes | Perihal surat |
| `body` | string | no | Isi surat |
| `deadline` | string | yes | RFC3339 |

```json
{
  "ministry": "MEDINFO",
  "letterType": "Undangan",
  "subject": "Undangan Rapat",
  "body": "Dengan hormat...",
  "deadline": "2026-06-10T12:00:00+07:00"
}
```

### Response 201

```json
{
  "success": true,
  "message": "Submission surat berhasil dibuat",
  "data": {
    "id": 1,
    "submitter_id": 2,
    "ministry": "MEDINFO",
    "letter_type": "Undangan",
    "subject": "Undangan Rapat",
    "body": "Dengan hormat...",
    "deadline": "2026-06-10T12:00:00+07:00",
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
| 422 | `VALIDATION_ERROR` | JSON tidak valid atau deadline bukan RFC3339 |
| 500 | `INTERNAL_ERROR` | Database gagal menyimpan |

### Curl

```bash
curl -X POST http://localhost:8081/api/letter-submissions \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"ministry":"MEDINFO","letterType":"Undangan","subject":"Undangan Rapat","body":"Dengan hormat...","deadline":"2026-06-10T12:00:00+07:00"}'
```

---

## GET /api/letter-submissions

Mengambil daftar pengajuan surat. `ADMIN` melihat semua. `MENTRI` hanya melihat milik sendiri atau ministry-nya.

**Auth:** Authenticated  
**Content-Type:** none

### Response 200

```json
{
  "success": true,
  "message": "Daftar letter submission",
  "data": [
    {
      "id": 1,
      "submitter_id": 2,
      "ministry": "MEDINFO",
      "letter_type": "Undangan",
      "subject": "Undangan Rapat",
      "deadline": "2026-06-10T12:00:00+07:00",
      "assigned_pj_id": 3,
      "status": "PENDING"
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
curl http://localhost:8081/api/letter-submissions \
  -H "Authorization: Bearer $TOKEN"
```

---

## GET /api/letter-submissions/:id

Mengambil detail pengajuan surat.

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
  "message": "Detail letter submission",
  "data": {
    "id": 1,
    "submitter_id": 2,
    "ministry": "MEDINFO",
    "letter_type": "Undangan",
    "subject": "Undangan Rapat",
    "body": "Dengan hormat...",
    "deadline": "2026-06-10T12:00:00+07:00",
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
curl http://localhost:8081/api/letter-submissions/1 \
  -H "Authorization: Bearer $TOKEN"
```

---

## PUT /api/letter-submissions/:id/status

Memperbarui status pengajuan surat.

**Auth:** ADMIN only  
**Content-Type:** `application/json`

### Request Body

| Field | Type | Required | Keterangan |
|---|---|---:|---|
| `status` | string | yes | `IN_REVIEW`, `APPROVED`, atau `REJECTED` |
| `notes` | string/null | no | Catatan admin |

```json
{
  "status": "APPROVED",
  "notes": "Surat disetujui"
}
```

### Response 200

```json
{
  "success": true,
  "message": "Status berhasil diperbarui",
  "data": {
    "id": 1,
    "status": "APPROVED",
    "notes": "Surat disetujui"
  }
}
```

### Error Responses

| Status | Code | Keterangan |
|---:|---|---|
| 403 | `FORBIDDEN` | Role bukan `ADMIN` |
| 422 | `VALIDATION_ERROR` | Status kosong |
| 409 | `CONFLICT` | Transisi status tidak valid |

### Curl

```bash
curl -X PUT http://localhost:8081/api/letter-submissions/1/status \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"status":"APPROVED","notes":"Surat disetujui"}'
```

---

## DELETE /api/letter-submissions/:id

Menghapus pengajuan surat.

**Auth:** ADMIN only  
**Content-Type:** none

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
curl -X DELETE http://localhost:8081/api/letter-submissions/1 \
  -H "Authorization: Bearer $TOKEN"
```
