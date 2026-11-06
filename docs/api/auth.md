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
| 429 | `RATE_LIMITED` | Terlalu banyak percobaan gagal untuk kombinasi IP+email ini; lihat header `Retry-After` (detik) |

```json
{
  "status": false,
  "message": "Kredensial salah",
  "error": { "code": "UNAUTHENTICATED" }
}
```

### Rate limiting

Dua lapisan independen melindungi endpoint ini dari brute force:

- **nginx** (`limit_req zone=login_zone`): dibatasi per alamat IP koneksi nyata (`$remote_addr`, bukan header `X-Forwarded-For` yang bisa dipalsukan klien).
- **Server Go** (`middlewares.LoginThrottle`): maksimum 5 percobaan gagal per 10 menit untuk kombinasi IP asli (TCP peer, independen dari header apa pun) + email yang dicoba. Percobaan yang gagal (401) menambah hitungan; login berhasil (200) mereset hitungan untuk kombinasi tersebut.

Kedua lapisan sengaja tidak memakai `X-Forwarded-For` dari klien sebagai kunci, agar tidak bisa dilewati dengan memalsukan header tersebut.

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
