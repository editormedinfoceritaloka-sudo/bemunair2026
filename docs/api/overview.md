# Overview

Dokumentasi API BEM UNAIR Digital Submission System. Semua endpoint backend melewati reverse proxy nginx dengan prefix `/api`.

## Base URL

| Environment | Base URL |
|---|---|
| Browser via Docker | `http://localhost:8081/api` |
| Internal backend container | `http://bemunair_server:8080/api` |
| WA Engine internal | `http://bemunair_wa_engine:3001` |

## Authentication

Gunakan JWT dari `POST /api/auth/login`.

```http
Authorization: Bearer <jwt_token>
```

Role yang valid hanya:

| Role | Akses |
|---|---|
| `ADMIN` | Kelola user, template, queue PJ, semua submission, update status |
| `MENTRI` | Buat submission dan melihat submission milik sendiri/kementerian |

## Standard Success Envelope

```json
{
  "success": true,
  "message": "Submission berhasil dibuat",
  "data": {
    "id": 1
  }
}
```

## Standard List Envelope

```json
{
  "success": true,
  "message": "Daftar content submission",
  "data": [],
  "meta": {
    "page": 1,
    "per_page": 20,
    "total": 0,
    "total_pages": 0
  }
}
```

## Standard Error Envelope

```json
{
  "success": false,
  "message": "Validasi gagal",
  "error": {
    "code": "VALIDATION_ERROR",
    "details": [
      { "field": "email", "message": "wajib diisi" }
    ]
  }
}
```

## Error Codes

| Code | HTTP Status | Kapan terjadi |
|---|---:|---|
| `VALIDATION_ERROR` | 422 | Field request tidak valid |
| `UNAUTHENTICATED` | 401 | Token tidak ada, token rusak, token expired, atau kredensial salah |
| `FORBIDDEN` | 403 | Role tidak memiliki akses |
| `NOT_FOUND` | 404 | Data atau dokumen tidak ditemukan |
| `CONFLICT` | 409 | Konflik data atau transisi status tidak valid |
| `RATE_LIMITED` | 429 | Request terlalu banyak |
| `INTERNAL_ERROR` | 500 | Kesalahan server |
| `WA_ENGINE_ERROR` | 502 | WA Engine gagal mengirim pesan |

## Common Headers

| Header | Required | Keterangan |
|---|---:|---|
| `Authorization: Bearer <token>` | Authenticated endpoints | JWT |
| `Content-Type: application/json` | JSON body | Request JSON |
| `Content-Type: multipart/form-data` | Upload file | Content submission |
| `X-Request-Id` | No | Bisa dikirim client; jika kosong dibuat server |

## Timestamp Format

Semua field waktu memakai RFC3339 dengan timezone Jakarta.

```json
{
  "deadline": "2026-06-10T12:00:00+07:00",
  "created_at": "2026-06-01T03:00:00+07:00"
}
```
