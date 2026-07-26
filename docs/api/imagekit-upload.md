# Upload gambar ImageKit

Endpoint SvelteKit ini melakukan upload server-side dengan SDK resmi ImageKit. Private API key tidak pernah dikirim ke browser.

## `POST /api/uploads/imagekit`

Memerlukan cookie sesi admin yang valid dan role `ADMIN`.

Request menggunakan `multipart/form-data`:

| Field | Wajib | Nilai |
| --- | --- | --- |
| `file` | Ya | JPEG, PNG, WebP, GIF, atau AVIF |
| `purpose` | Tidak | `article`, `cover`, `profile`, atau `submission`; default `article` |

Batas ukuran default adalah 10 MB dan dapat diubah melalui `IMAGEKIT_MAX_UPLOAD_BYTES`.

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
  "message": "Gambar berhasil diunggah.",
  "data": {
    "file_id": "imagekit-file-id",
    "name": "cover_unique.webp",
    "file_path": "/bemunair/article-covers/cover_unique.webp",
    "url": "https://ik.imagekit.io/example/cover_unique.webp",
    "thumbnail_url": "https://ik.imagekit.io/example/tr:n-media_library_thumbnail/cover_unique.webp",
    "width": 1600,
    "height": 900,
    "size": 245000,
    "file_type": "image"
  }
}
```

## Keamanan

- Private key hanya dibaca dari environment server.
- Endpoint memvalidasi sesi melalui backend dan membatasi akses ke admin.
- MIME type, ukuran, dan signature biner file diperiksa sebelum upload.
- SVG tidak diterima untuk mencegah active-content injection.
- Folder upload dipetakan dari allowlist dan tidak menerima path bebas dari pengguna.
- Nama file dinormalisasi dan ImageKit menggunakan nama unik.
