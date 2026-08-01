# Letter Templates

Template surat disimpan sebagai PDF pada ImageKit. Metadata file disimpan pada `media_assets`, sedangkan `letter_templates` menyimpan jenis, nama, dan relasinya.

## Authorization

- `GET /api/v1/letter-templates` dan `GET /api/v1/letter-templates/:id`: `ADMIN` dan `ADMIN_MEDINFO`.
- `POST`, `PUT`, dan `DELETE`: hanya `ADMIN_MEDINFO`.
- Admin Kementerian hanya memakai URL download PDF dari response.

## Template Object

```json
{
  "id": 900001,
  "name": "Undangan Kegiatan Resmi",
  "type": "UNDANGAN",
  "subject": "Undangan Kegiatan BEM Universitas Airlangga",
  "media_asset_id": 910107,
  "file": {
    "id": 910107,
    "url": "https://cdn.example/template.pdf",
    "name": "template-undangan.pdf",
    "mime_type": "application/pdf",
    "size_bytes": 25000
  },
  "download_url": "https://cdn.example/template.pdf",
  "is_active": true,
  "display_order": 0
}
```

## Upload flow

Halaman Admin Medinfo menjalankan tiga langkah:

1. Upload PDF ke `/api/uploads/imagekit` dengan purpose `letter_template`.
2. Simpan metadata hasil upload ke `POST /api/v1/admin/media-assets` dengan `purpose: letter_template`.
3. Simpan template ke `POST /api/v1/letter-templates` dengan `media_asset_id`.

File harus berformat PDF, memiliki signature PDF yang valid, dan berukuran paling besar 10 MB.

## GET /api/v1/letter-templates

Mengambil daftar template untuk download.

```bash
curl http://localhost:8081/api/v1/letter-templates \
  -H "Authorization: Bearer $TOKEN"
```

## POST /api/v1/letter-templates

Membuat metadata template setelah media PDF tersimpan.

**Auth:** `ADMIN_MEDINFO`

```json
{
  "name": "Undangan Kegiatan Resmi",
  "type": "UNDANGAN",
  "subject": "Undangan Kegiatan BEM Universitas Airlangga",
  "body": "",
  "media_asset_id": 910107,
  "is_active": true,
  "display_order": 0
}
```

## PUT /api/v1/letter-templates/:id

Memperbarui metadata atau mengganti relasi PDF. Jika `media_asset_id` diisi, media tersebut harus berupa PDF dengan purpose `letter_template`.

## DELETE /api/v1/letter-templates/:id

Menghapus metadata template. Penghapusan record tidak menghapus file fisik ImageKit secara otomatis sehingga asset perlu ditangani melalui proses cleanup terpisah.
