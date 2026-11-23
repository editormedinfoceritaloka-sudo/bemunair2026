# Upload media

Endpoint SvelteKit ini menyimpan file di direktori upload server dan mengembalikan URL di bawah `/uploads/`.

## `POST /api/uploads/imagekit`

Memerlukan cookie sesi admin yang valid dan role `ADMIN` atau `ADMIN_MEDINFO`.

Request menggunakan `multipart/form-data`:

| Field | Wajib | Nilai |
| --- | --- | --- |
| `file` | Ya | JPEG, PNG, WebP, GIF, atau AVIF |
| `purpose` | Tidak | `article`, `cover`, `profile`, `documentation`, `submission_media`, `submission_brief`, atau `letter_template`. Default `article`. |

PDF diterima untuk `letter_template`. PDF, DOC, dan DOCX diterima untuk `submission_brief`.

Batas ukuran default adalah 10 MB dan dapat diubah melalui `UPLOAD_MAX_BYTES`.

Contoh:

```bash
curl -X POST http://localhost:3000/api/uploads/imagekit \
  -b "bemunair_admin_session=<session>" \
  -F "file=@cover.webp" \
  -F "purpose=cover"
```

Response berhasil:

```json
{
  "status": true,
  "message": "File berhasil diunggah.",
  "data": {
    "file_id": "f59c7d20-b778-4d34-b564-48732e92ee13",
    "name": "f59c7d20-b778-4d34-b564-48732e92ee13-cover.webp",
    "file_path": "/uploads/article-covers/f59c7d20-b778-4d34-b564-48732e92ee13-cover.webp",
    "url": "/uploads/article-covers/f59c7d20-b778-4d34-b564-48732e92ee13-cover.webp",
    "thumbnail_url": "/uploads/article-covers/f59c7d20-b778-4d34-b564-48732e92ee13-cover.webp",
    "width": null,
    "height": null,
    "size": 245000,
    "file_type": "image/webp"
  }
}
```

## Keamanan

- Endpoint memvalidasi sesi melalui backend dan membatasi akses ke admin.
- MIME type, ukuran, dan signature biner file diperiksa sebelum upload.
- SVG tidak diterima untuk mencegah active-content injection.
- Folder upload dipetakan dari allowlist dan tidak menerima path bebas dari pengguna.
- Nama file dinormalisasi. ID file berupa UUID dan terpisah dari path penyimpanan.
