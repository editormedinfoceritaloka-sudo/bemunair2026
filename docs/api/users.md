# Users

Endpoint untuk mengelola user. Semua endpoint di modul ini hanya untuk `ADMIN`.

## User Object

```json
{
  "id": 2,
  "name": "Mentri Medinfo",
  "email": "mentri.medinfo@bem.unair.ac.id",
  "role": "MENTRI",
  "ministry": "MEDINFO",
  "phone": "6281222222222",
  "created_at": "2026-06-01T03:00:00+07:00",
  "updated_at": "2026-06-01T03:00:00+07:00"
}
```

---

## GET /api/users

Mengambil semua user.

**Auth:** ADMIN only  
**Content-Type:** none

### Request

Tidak ada query parameter.

### Response 200

```json
{
  "success": true,
  "message": "Daftar user",
  "data": [
    {
      "id": 1,
      "name": "Admin BEM UNAIR",
      "email": "admin@bem.unair.ac.id",
      "role": "ADMIN",
      "ministry": null,
      "phone": "6281111111111"
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

### Curl

```bash
curl http://localhost:8081/api/users \
  -H "Authorization: Bearer $TOKEN"
```

---

## GET /api/users/:id

Mengambil detail user.

**Auth:** ADMIN only  
**Content-Type:** none

### Path Parameters

| Field | Type | Required | Keterangan |
|---|---|---:|---|
| `id` | integer | yes | ID user |

### Response 200

```json
{
  "success": true,
  "message": "Detail user",
  "data": {
    "id": 2,
    "name": "Mentri Medinfo",
    "email": "mentri.medinfo@bem.unair.ac.id",
    "role": "MENTRI",
    "ministry": "MEDINFO",
    "phone": "6281222222222"
  }
}
```

### Error Responses

| Status | Code | Keterangan |
|---:|---|---|
| 401 | `UNAUTHENTICATED` | Token tidak valid |
| 403 | `FORBIDDEN` | Role bukan `ADMIN` |
| 404 | `NOT_FOUND` | User tidak ditemukan |

### Curl

```bash
curl http://localhost:8081/api/users/2 \
  -H "Authorization: Bearer $TOKEN"
```

---

## POST /api/users

Membuat user baru.

**Auth:** ADMIN only  
**Content-Type:** `application/json`

### Request Body

| Field | Type | Required | Keterangan |
|---|---|---:|---|
| `name` | string | yes | Nama |
| `email` | string | yes | Email unik |
| `password` | string | yes | Password plaintext |
| `role` | string | yes | `ADMIN` atau `MENTRI` |
| `ministry` | string/null | no | Kementerian untuk `MENTRI` |
| `phone` | string/null | no | Nomor WA |

```json
{
  "name": "Mentri PSDM",
  "email": "mentri.psdm@bem.unair.ac.id",
  "password": "password",
  "role": "MENTRI",
  "ministry": "PSDM",
  "phone": "6281333333333"
}
```

### Response 201

```json
{
  "success": true,
  "message": "User berhasil dibuat",
  "data": {
    "id": 6,
    "name": "Mentri PSDM",
    "email": "mentri.psdm@bem.unair.ac.id",
    "role": "MENTRI",
    "ministry": "PSDM",
    "phone": "6281333333333"
  }
}
```

### Error Responses

| Status | Code | Keterangan |
|---:|---|---|
| 422 | `VALIDATION_ERROR` | Field wajib kosong |
| 409 | `CONFLICT` | Email sudah dipakai |

### Curl

```bash
curl -X POST http://localhost:8081/api/users \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name":"Mentri PSDM","email":"mentri.psdm@bem.unair.ac.id","password":"password","role":"MENTRI","ministry":"PSDM","phone":"6281333333333"}'
```

---

## PUT /api/users/:id

Memperbarui user.

**Auth:** ADMIN only  
**Content-Type:** `application/json`

### Path Parameters

| Field | Type | Required | Keterangan |
|---|---|---:|---|
| `id` | integer | yes | ID user |

### Request Body

| Field | Type | Required | Keterangan |
|---|---|---:|---|
| `name` | string | no | Nama baru |
| `email` | string | no | Email baru |
| `role` | string | no | `ADMIN` atau `MENTRI` |
| `ministry` | string/null | no | Kementerian |
| `phone` | string/null | no | Nomor WA |

```json
{
  "name": "Mentri Medinfo Updated",
  "email": "mentri.medinfo@bem.unair.ac.id",
  "role": "MENTRI",
  "ministry": "MEDINFO",
  "phone": "6281222222222"
}
```

### Response 200

```json
{
  "success": true,
  "message": "User berhasil diperbarui",
  "data": {
    "id": 2,
    "name": "Mentri Medinfo Updated",
    "email": "mentri.medinfo@bem.unair.ac.id",
    "role": "MENTRI",
    "ministry": "MEDINFO",
    "phone": "6281222222222"
  }
}
```

### Error Responses

| Status | Code | Keterangan |
|---:|---|---|
| 404 | `NOT_FOUND` | User tidak ditemukan |
| 422 | `VALIDATION_ERROR` | JSON tidak valid |
| 500 | `INTERNAL_ERROR` | Gagal update |

### Curl

```bash
curl -X PUT http://localhost:8081/api/users/2 \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name":"Mentri Medinfo Updated","role":"MENTRI","ministry":"MEDINFO","phone":"6281222222222"}'
```

---

## DELETE /api/users/:id

Menghapus user.

**Auth:** ADMIN only  
**Content-Type:** none

### Path Parameters

| Field | Type | Required | Keterangan |
|---|---|---:|---|
| `id` | integer | yes | ID user |

### Response 200

```json
{
  "success": true,
  "message": "User berhasil dihapus"
}
```

### Error Responses

| Status | Code | Keterangan |
|---:|---|---|
| 403 | `FORBIDDEN` | Role bukan admin |
| 500 | `INTERNAL_ERROR` | User gagal dihapus, misalnya masih direferensikan FK |

### Curl

```bash
curl -X DELETE http://localhost:8081/api/users/2 \
  -H "Authorization: Bearer $TOKEN"
```
