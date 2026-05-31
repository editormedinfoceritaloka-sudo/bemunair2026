# WA Engine

Service internal untuk WhatsApp berbasis Baileys. Semua endpoint `/api/*` memakai API key internal.

WA Engine melakukan auto-connect saat container start jika `WA_AUTO_CONNECT=true`. QR login akan dicetak di Docker logs, jadi operator cukup membuka logs dan scan QR dari terminal. Session tersimpan di volume `/app/sessions`, sehingga setelah scan pertama tidak perlu login ulang selama session tidak logout atau volume tidak dihapus.

Nomor `WA_SENDER_PHONE` hanya label sender di log. WhatsApp tetap mewajibkan autentikasi device via QR; tidak bisa login hanya dengan mengetik nomor.

## Authentication

```http
Authorization: Bearer <WA_ENGINE_API_KEY>
```

## Status Object

```json
{
  "state": "open",
  "qr": "optional_qr_string"
}
```

State valid: `disconnected`, `connecting`, `open`, `reconnecting`, `logged_out`.

---

## GET /health

Health check public.

**Auth:** Public  
**Content-Type:** none

### Response 200

```json
{
  "success": true,
  "message": "ok"
}
```

### Curl

```bash
curl http://localhost:3001/health
```

---

## GET /api/status

Mengambil status socket WhatsApp.

**Auth:** API key  
**Content-Type:** none

### Response 200

```json
{
  "success": true,
  "data": {
    "state": "connecting",
    "qr": "2@abc..."
  }
}
```

### Error Responses

| Status | Keterangan |
|---:|---|
| 401 | API key kosong atau salah |

```json
{
  "success": false,
  "message": "Unauthorized"
}
```

### Curl

```bash
curl http://localhost:3001/api/status \
  -H "Authorization: Bearer $WA_ENGINE_API_KEY"
```

---

## GET /api/qr

Mengambil QR login sebagai base64 image jika state belum `open`.

**Auth:** API key  
**Content-Type:** none

### Response 200

```json
{
  "success": true,
  "data": {
    "image": "data:image/png;base64,iVBORw0KGgo...",
    "qr": "2@abc..."
  }
}
```

### Response 200 Jika Sudah Login

```json
{
  "success": true,
  "data": {
    "state": "open",
    "qr": null
  }
}
```

### Error Responses

| Status | Keterangan |
|---:|---|
| 401 | API key salah |
| 404 | QR belum tersedia |

### Curl

```bash
curl http://localhost:3001/api/qr \
  -H "Authorization: Bearer $WA_ENGINE_API_KEY"
```

---

## POST /api/connect

Memicu koneksi WhatsApp dan pembuatan QR secara manual. Biasanya tidak perlu dipanggil karena auto-connect sudah aktif saat container start.

**Auth:** API key  
**Content-Type:** `application/json`

### Request Body

Tidak ada body.

### Response 200

```json
{
  "success": true,
  "message": "connect triggered",
  "data": {
    "state": "connecting",
    "qr": "2@abc..."
  }
}
```

### Error Responses

| Status | Keterangan |
|---:|---|
| 401 | API key salah |
| 500 | Baileys gagal start |

### Curl

```bash
curl -X POST http://localhost:3001/api/connect \
  -H "Authorization: Bearer $WA_ENGINE_API_KEY"
```

## Login via Docker Logs

```bash
docker compose -f infra/docker-compose.yml logs -f wa-engine
```

Cari log:

```text
[WA] QR code received. Scan this QR from Docker logs:
```

Scan QR tersebut dari WhatsApp mobile melalui menu Linked Devices. Setelah berhasil, log akan berisi:

```text
[WA] open - connected as <whatsapp-user-id>
```

---

## GET /api/groups

Mengambil daftar grup WhatsApp yang diikuti akun sender. Endpoint ini dipakai untuk mendapatkan `groupJid` seperti `120363xxxxxxxx@g.us`.

**Auth:** API key  
**Content-Type:** none

### Response 200

```json
{
  "success": true,
  "message": "groups fetched",
  "data": [
    {
      "jid": "120363xxxxxxxx@g.us",
      "subject": "BEM UNAIR",
      "size": 42
    }
  ]
}
```

### Docker Logs

Saat endpoint dipanggil, WA Engine juga mencetak daftar grup:

```text
[WA] groups found count=2
[WA-GROUP] jid=120363xxxxxxxx@g.us subject="BEM UNAIR" size=42
```

### Error Responses

| Status | Keterangan |
|---:|---|
| 401 | API key salah |
| 500 | Socket belum `open` |

### Curl

```bash
curl http://localhost:3001/api/groups \
  -H "Authorization: Bearer $WA_ENGINE_API_KEY"
```

---

## POST /api/send-message

Mengirim pesan WhatsApp personal.

**Auth:** API key  
**Content-Type:** `application/json`

### Request Body

| Field | Type | Required | Keterangan |
|---|---|---:|---|
| `to` | string | yes | Nomor telepon, contoh `08123456789` |
| `message` | string | yes | Isi pesan |

```json
{
  "to": "08123456789",
  "message": "Halo dari BEM UNAIR"
}
```

### Response 200

```json
{
  "success": true,
  "message": "message sent"
}
```

### Error Responses

| Status | Keterangan |
|---:|---|
| 401 | API key salah |
| 500 | Socket belum `open` atau pengiriman gagal |

```json
{
  "success": false,
  "message": "WA socket is not open"
}
```

### Curl

```bash
curl -X POST http://localhost:3001/api/send-message \
  -H "Authorization: Bearer $WA_ENGINE_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{"to":"08123456789","message":"Halo dari BEM UNAIR"}'
```

---

## POST /api/send-group-message

Mengirim pesan ke grup WhatsApp.

**Auth:** API key  
**Content-Type:** `application/json`

### Request Body

| Field | Type | Required | Keterangan |
|---|---|---:|---|
| `groupJid` | string | yes | Format `120363xxx@g.us` |
| `message` | string | yes | Isi pesan |

```json
{
  "groupJid": "120363xxxxxxxx@g.us",
  "message": "Reminder harian BEM UNAIR"
}
```

### Response 200

```json
{
  "success": true,
  "message": "group message sent"
}
```

### Error Responses

| Status | Keterangan |
|---:|---|
| 401 | API key salah |
| 500 | Socket belum `open` atau pengiriman gagal |

### Curl

```bash
curl -X POST http://localhost:3001/api/send-group-message \
  -H "Authorization: Bearer $WA_ENGINE_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{"groupJid":"120363xxxxxxxx@g.us","message":"Reminder harian BEM UNAIR"}'
```
