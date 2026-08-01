# Letter Templates Admin

Admin Medinfo memiliki UI Template Surat yang konsisten dengan workspace admin saat ini. UI tersebut menyediakan upload PDF, metadata nama dan jenis surat, daftar template aktif, download preview, dan penghapusan metadata.

Admin Medinfo tidak melihat menu Pengajuan Surat di workspace mereka. Admin Kementerian tetap dapat membuka Pengajuan Surat dari workspace mereka dan hanya memiliki aksi download template PDF sebelum mengisi pengajuan.

Alur utama:

1. Admin Medinfo memilih PDF maksimal 10 MB.
2. Client dan server memvalidasi MIME type serta signature PDF.
3. File diunggah ke ImageKit pada folder `letter-templates`.
4. Metadata media disimpan pada `media_assets`.
5. Relasi template disimpan pada `letter_templates`.
6. Admin Kementerian mengambil daftar template melalui API dan mengunduh PDF.
7. Isi pengajuan surat tetap disimpan pada `letter_submissions`; template PDF tidak dicampur dengan isi submission.

UI Admin Medinfo tidak memproses atau mengelola pengajuan surat. Endpoint operasional submission tetap dipertahankan untuk workflow Admin Kementerian dan kompatibilitas backend.
