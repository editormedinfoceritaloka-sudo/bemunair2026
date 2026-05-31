# Letter Templates

Template surat reusable. Semua endpoint hanya untuk `ADMIN`.

## Template Object

```json
{
  "id": 1,
  "name": "Undangan Rapat",
  "type": "UNDANGAN",
  "subject": "Undangan Rapat",
  "body": "Dengan hormat...",
  "created_at": "2026-06-01T03:00:00+07:00",
  "updated_at": "2026-06-01T03:00:00+07:00"
}
```

---

## POST /api/letter-templates

Membuat template surat.

**Auth:** ADMIN only  
**Content-Type:** `application/json`

### Request Body

| Field | Type | Required | Keterangan |
|---|---|---:|---|
| `name` | string | yes | Nama template |
| `type` | string | yes | Jenis template |
| `subject` | string | yes | Subject default |
| `body` | string | yes | Isi template |

```json
{
  "name": "Undangan Rapat",
  "type": "UNDANGAN",
  "subject": "Undangan Rapat",
  "body": "Dengan hormat..."
}
```

### Response 201

```json
{
  "success": true,
  "message": "Template berhasil dibuat",
  "data": {
    "id": 1,
    "name": "Undangan Rapat",
    "type": "UNDANGAN",
    "subject": "Undangan Rapat",
    "body": "Dengan hormat..."
  }
}
```

### Error Responses

| Status | Code | Keterangan |
|---:|---|---|
| 422 | `VALIDATION_ERROR` | JSON tidak valid |
| 500 | `INTERNAL_ERROR` | Gagal simpan |

### Curl

```bash
curl -X POST http://localhost:8081/api/letter-templates \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name":"Undangan Rapat","type":"UNDANGAN","subject":"Undangan Rapat","body":"Dengan hormat..."}'
```

---

## GET /api/letter-templates

Mengambil semua template.

**Auth:** ADMIN only  
**Content-Type:** none

### Response 200

```json
{
  "success": true,
  "message": "Daftar template",
  "data": [
    {
      "id": 1,
      "name": "Undangan Rapat",
      "type": "UNDANGAN",
      "subject": "Undangan Rapat",
      "body": "Dengan hormat..."
    }
  ],
  "meta": { "page": 1, "per_page": 1, "total": 1, "total_pages": 1 }
}
```

### Error Responses

| Status | Code | Keterangan |
|---:|---|---|
| 403 | `FORBIDDEN` | Role bukan `ADMIN` |
| 500 | `INTERNAL_ERROR` | Query gagal |

### Curl

```bash
curl http://localhost:8081/api/letter-templates \
  -H "Authorization: Bearer $TOKEN"
```

---

## GET /api/letter-templates/:id

Mengambil detail template.

**Auth:** ADMIN only  
**Content-Type:** none

### Response 200

```json
{
  "success": true,
  "message": "Detail template",
  "data": {
    "id": 1,
    "name": "Undangan Rapat",
    "type": "UNDANGAN",
    "subject": "Undangan Rapat",
    "body": "Dengan hormat..."
  }
}
```

### Error Responses

| Status | Code | Keterangan |
|---:|---|---|
| 404 | `NOT_FOUND` | Template tidak ditemukan |

### Curl

```bash
curl http://localhost:8081/api/letter-templates/1 \
  -H "Authorization: Bearer $TOKEN"
```

---

## PUT /api/letter-templates/:id

Memperbarui template.

**Auth:** ADMIN only  
**Content-Type:** `application/json`

### Request Body

| Field | Type | Required | Keterangan |
|---|---|---:|---|
| `name` | string | yes | Nama template |
| `type` | string | yes | Jenis template |
| `subject` | string | yes | Subject |
| `body` | string | yes | Isi |

```json
{
  "name": "Undangan Rapat Bulanan",
  "type": "UNDANGAN",
  "subject": "Undangan Rapat Bulanan",
  "body": "Dengan hormat..."
}
```

### Response 200

```json
{
  "success": true,
  "message": "Template berhasil diperbarui",
  "data": {
    "id": 1,
    "name": "Undangan Rapat Bulanan",
    "type": "UNDANGAN",
    "subject": "Undangan Rapat Bulanan",
    "body": "Dengan hormat..."
  }
}
```

### Error Responses

| Status | Code | Keterangan |
|---:|---|---|
| 404 | `NOT_FOUND` | Template tidak ditemukan |
| 422 | `VALIDATION_ERROR` | JSON tidak valid |
| 500 | `INTERNAL_ERROR` | Gagal update |

### Curl

```bash
curl -X PUT http://localhost:8081/api/letter-templates/1 \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name":"Undangan Rapat Bulanan","type":"UNDANGAN","subject":"Undangan Rapat Bulanan","body":"Dengan hormat..."}'
```

---

## DELETE /api/letter-templates/:id

Menghapus template.

**Auth:** ADMIN only  
**Content-Type:** none

### Response 200

```json
{
  "success": true,
  "message": "Template berhasil dihapus"
}
```

### Error Responses

| Status | Code | Keterangan |
|---:|---|---|
| 403 | `FORBIDDEN` | Role bukan `ADMIN` |
| 500 | `INTERNAL_ERROR` | Gagal hapus |

### Curl

```bash
curl -X DELETE http://localhost:8081/api/letter-templates/1 \
  -H "Authorization: Bearer $TOKEN"
```
