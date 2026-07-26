# Authentication

Endpoint untuk login dan mengambil profil user aktif. Pembuatan akun dilakukan admin melalui endpoint Users.

---

## Public registration disabled

`POST /api/v1/auth/register` tidak diekspos dan mengembalikan `404`. Akun baru dibuat oleh admin melalui `POST /api/v1/users`.

---

## POST /api/v1/auth/login

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
  "status": true,
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
  "status": false,
  "message": "Kredensial salah",
  "error": { "code": "UNAUTHENTICATED" }
}
```

### Curl

```bash
curl -X POST http://localhost:8081/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@bem.unair.ac.id","password":"password"}'
```

---

## GET /api/v1/auth/me

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
  "status": true,
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
curl http://localhost:8081/api/v1/auth/me \
  -H "Authorization: Bearer $TOKEN"
```
