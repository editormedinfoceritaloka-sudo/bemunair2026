# Medinfo PJ Queue

Queue PJ Medinfo dipakai untuk auto-assign submission secara round-robin. Semua endpoint hanya untuk `ADMIN`.

## Queue Object

```json
{
  "id": 1,
  "user_id": 2,
  "position": 1,
  "is_current": true,
  "user": {
    "id": 2,
    "name": "Mentri Medinfo",
    "email": "mentri.medinfo@bem.unair.ac.id"
  },
  "created_at": "2026-06-01T03:00:00+07:00",
  "updated_at": "2026-06-01T03:00:00+07:00"
}
```

---

## GET /api/medinfo-pj/queue

Mengambil queue PJ, urut berdasarkan `position`.

**Auth:** ADMIN only  
**Content-Type:** none

### Response 200

```json
{
  "success": true,
  "message": "Queue PJ Medinfo",
  "data": [
    {
      "id": 1,
      "user_id": 2,
      "position": 1,
      "is_current": true,
      "user": { "id": 2, "name": "Mentri Medinfo" }
    }
  ],
  "meta": { "page": 1, "per_page": 1, "total": 1, "total_pages": 1 }
}
```

### Error Responses

| Status | Code | Keterangan |
|---:|---|---|
| 401 | `UNAUTHENTICATED` | Token tidak valid |
| 403 | `FORBIDDEN` | Role bukan `ADMIN` |
| 500 | `INTERNAL_ERROR` | Query gagal |

### Curl

```bash
curl http://localhost:8081/api/medinfo-pj/queue \
  -H "Authorization: Bearer $TOKEN"
```

---

## POST /api/medinfo-pj/queue

Menambahkan user ke queue PJ.

**Auth:** ADMIN only  
**Content-Type:** `application/json`

### Request Body

| Field | Type | Required | Keterangan |
|---|---|---:|---|
| `user_id` | integer | yes | ID user yang akan menjadi PJ |
| `position` | integer | no | Posisi queue; default `1` jika kosong |

```json
{
  "user_id": 2,
  "position": 1
}
```

### Response 201

```json
{
  "success": true,
  "message": "Queue berhasil dibuat",
  "data": {
    "id": 4,
    "user_id": 2,
    "position": 1,
    "is_current": false
  }
}
```

### Error Responses

| Status | Code | Keterangan |
|---:|---|---|
| 422 | `VALIDATION_ERROR` | `user_id` kosong |
| 409 | `CONFLICT` | User sudah ada di queue |

### Curl

```bash
curl -X POST http://localhost:8081/api/medinfo-pj/queue \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"user_id":2,"position":1}'
```

---

## PUT /api/medinfo-pj/queue/reorder

Mengurutkan ulang queue. Item pertama otomatis menjadi `is_current = true`.

**Auth:** ADMIN only  
**Content-Type:** `application/json`

### Request Body

| Field | Type | Required | Keterangan |
|---|---|---:|---|
| `ids` | integer[] | yes | Urutan ID row queue |

```json
{
  "ids": [1, 2, 3]
}
```

### Response 200

```json
{
  "success": true,
  "message": "Queue berhasil diurutkan"
}
```

### Error Responses

| Status | Code | Keterangan |
|---:|---|---|
| 422 | `VALIDATION_ERROR` | `ids` kosong |
| 500 | `INTERNAL_ERROR` | Update gagal |

### Curl

```bash
curl -X PUT http://localhost:8081/api/medinfo-pj/queue/reorder \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"ids":[1,2,3]}'
```

---

## DELETE /api/medinfo-pj/queue/:id

Menghapus row queue PJ.

**Auth:** ADMIN only  
**Content-Type:** none

### Path Parameters

| Field | Type | Required | Keterangan |
|---|---|---:|---|
| `id` | integer | yes | ID row queue |

### Response 200

```json
{
  "success": true,
  "message": "Queue berhasil dihapus"
}
```

### Error Responses

| Status | Code | Keterangan |
|---:|---|---|
| 403 | `FORBIDDEN` | Role bukan `ADMIN` |
| 500 | `INTERNAL_ERROR` | Gagal hapus |

### Curl

```bash
curl -X DELETE http://localhost:8081/api/medinfo-pj/queue/1 \
  -H "Authorization: Bearer $TOKEN"
```
