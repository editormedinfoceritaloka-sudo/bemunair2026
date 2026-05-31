# Authentication

Endpoint untuk registrasi, login, dan mengambil profil user aktif.

---

## POST /api/auth/register

Membuat akun baru. Endpoint ini public pada implementasi saat ini.

**Auth:** Public  
**Content-Type:** `application/json`

### Request Body

| Field | Type | Required | Keterangan |
|---|---|---:|---|
| `name` | string | yes | Nama lengkap user |
| `email` | string | yes | Email unik |
| `password` | string | yes | Password plaintext, akan di-hash bcrypt |
| `role` | string | yes | `ADMIN` atau `MENTRI` |
| `ministry` | string/null | no | Wajib secara bisnis untuk `MENTRI` |
| `phone` | string/null | no | Nomor WhatsApp |

```json
{
  "name": "Mentri Medinfo",
  "email": "mentri.medinfo@bem.unair.ac.id",
  "password": "password",
  "role": "MENTRI",
  "ministry": "MEDINFO",
  "phone": "6281222222222"
}
```

### Response 201

```json
{
  "success": true,
  "message": "User berhasil dibuat",
  "data": {
    "id": 5,
    "name": "Mentri Medinfo",
    "email": "mentri.medinfo@bem.unair.ac.id",
    "role": "MENTRI",
    "ministry": "MEDINFO",
    "phone": "6281222222222",
    "created_at": "2026-06-01T03:00:00+07:00",
    "updated_at": "2026-06-01T03:00:00+07:00"
  }
}
```

### Error Responses

| Status | Code | Keterangan |
|---:|---|---|
| 422 | `VALIDATION_ERROR` | Field wajib kosong atau role bukan `ADMIN`/`MENTRI` |
| 409 | `CONFLICT` | Email sudah dipakai |

```json
{
  "success": false,
  "message": "invalid role",
  "error": { "code": "VALIDATION_ERROR" }
}
```

### Curl

```bash
curl -X POST http://localhost:8081/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{"name":"Mentri Medinfo","email":"mentri.medinfo@bem.unair.ac.id","password":"password","role":"MENTRI","ministry":"MEDINFO","phone":"6281222222222"}'
```

---

## POST /api/auth/login

Login user dan mengembalikan JWT.

**Auth:** Public  
**Content-Type:** `application/json`

### Request Body

| Field | Type | Required | Keterangan |
|---|---|---:|---|
| `email` | string | yes | Email user |
| `password` | string | yes | Password |

```json
{
  "email": "admin@bem.unair.ac.id",
  "password": "password"
}
```

### Response 200

```json
{
  "success": true,
  "message": "Login berhasil",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIs...",
    "user": {
      "id": 1,
      "name": "Admin BEM UNAIR",
      "email": "admin@bem.unair.ac.id",
      "role": "ADMIN",
      "ministry": null,
      "phone": "6281111111111"
    }
  }
}
```

### Error Responses

| Status | Code | Keterangan |
|---:|---|---|
| 422 | `VALIDATION_ERROR` | Email/password kosong |
| 401 | `UNAUTHENTICATED` | Email tidak ditemukan atau password salah |

```json
{
  "success": false,
  "message": "Kredensial salah",
  "error": { "code": "UNAUTHENTICATED" }
}
```

### Curl

```bash
curl -X POST http://localhost:8081/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@bem.unair.ac.id","password":"password"}'
```

---

## GET /api/auth/me

Mengambil user aktif berdasarkan JWT.

**Auth:** Authenticated  
**Content-Type:** none

### Request Headers

| Header | Required | Keterangan |
|---|---:|---|
| `Authorization` | yes | `Bearer <token>` |

### Response 200

```json
{
  "success": true,
  "message": "User aktif",
  "data": {
    "id": 1,
    "name": "Admin BEM UNAIR",
    "email": "admin@bem.unair.ac.id",
    "role": "ADMIN",
    "ministry": null,
    "phone": "6281111111111"
  }
}
```

### Error Responses

| Status | Code | Keterangan |
|---:|---|---|
| 401 | `UNAUTHENTICATED` | Token kosong, rusak, atau expired |
| 404 | `NOT_FOUND` | User dari token tidak ditemukan |

### Curl

```bash
curl http://localhost:8081/api/auth/me \
  -H "Authorization: Bearer $TOKEN"
```
