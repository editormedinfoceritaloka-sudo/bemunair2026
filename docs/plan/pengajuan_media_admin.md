# Pengajuan Media dan Role Admin

Status: implemented, 26 Juli 2026.

## Arsitektur

Alur implementasi mengikuti:

```text
MySQL migration + seeder
→ GORM entity/repository
→ domain service dan state transition
→ Gin controller + role middleware
→ SvelteKit server load/action
→ shadcn-svelte UI
```

Role final:

- `ADMIN`: membuat pengajuan media/surat, melihat data milik akun atau kementeriannya, melihat timeline, dan mengirim ulang revisi.
- `ADMIN_MEDINFO`: seluruh kemampuan `ADMIN` ditambah review, approval, scheduling, publish/complete, master data, konfigurasi, antrean PJ, pengguna, serta CMS artikel.

Legacy `MENTRI` hanya dipertahankan sebagai alias JWT sementara dan dimigrasikan menjadi `ADMIN`.

## Data

Migration `008_add_roles_and_submission_workflows.sql` menambahkan:

- master `ministries`;
- snapshot pengaju dan kementerian pada pengajuan;
- request code `MED-YYYY-NNNNNN` dan `SUR-YYYY-NNNNNN`;
- status history untuk media dan surat;
- attachment ImageKit untuk materi final dan brief;
- pengaturan SOP/template/PIC/jadwal;
- blackout publikasi, request sequence, notification outbox;
- hubungan idempotent `articles.source_submission_id`.

Seeder menyediakan data untuk seluruh tabel dan dapat dijalankan ulang tanpa menggandakan data development.

## Workflow

Media:

```text
SUBMITTED
→ PENDING_REVIEW
→ REVISION_REQUIRED
→ REVISION_SUBMITTED
→ PENDING_REVIEW
→ APPROVED
→ SCHEDULED
→ PUBLISHED
```

`REJECTED` dapat menjadi keputusan dari tahap review. Artikel yang sudah `APPROVED` dapat dibuat menjadi draft CMS melalui aksi idempotent “Buat / buka draft artikel”.

Surat:

```text
SUBMITTED
→ PENDING_REVIEW
→ REVISION_REQUIRED
→ REVISION_SUBMITTED
→ PENDING_REVIEW
→ APPROVED
→ COMPLETED
```

## UI pengajuan media

Route utama:

- `/admin/content-submissions/new/select`
- `/admin/content-submissions/new/content`
- `/admin/content-submissions/new/content/form`
- `/admin/content-submissions/new/article`
- `/admin/content-submissions/new/article/form`
- `/admin/content-submissions/success`

Form konten mencakup Feed, Reels, dan Instastory; identitas akun; H-7; waktu 08.00–17.00; caption; Canva; lagu kondisional; upload media; upload brief; catatan; dan empat pernyataan.

Form artikel mencakup identitas akun; H+3; jadwal; judul; Google Drive dokumentasi; Google Docs naskah; informasi wajib; caption; catatan; dan empat pernyataan.

## Upload ImageKit

Browser meminta autentikasi singkat dari:

```text
GET /api/uploads/imagekit/auth?purpose=submission_media
GET /api/uploads/imagekit/auth?purpose=submission_brief
```

Endpoint hanya mengembalikan `token`, `expire`, `signature`, public key, dan folder. Private key tidak pernah dikirim ke browser. Browser kemudian mengunggah langsung ke Upload API V1 ImageKit dan mengirim metadata hasil upload bersama form. Backend menyimpan metadata tersebut pada `content_submission_attachments`.

Environment wajib:

```text
IMAGEKIT_PUBLIC_KEY=
IMAGEKIT_PRIVATE_KEY=
IMAGEKIT_UPLOAD_FOLDER=/bemunair
```

## Master dan konfigurasi

- `/admin/ministries`: master kementerian aktif/nonaktif.
- `/admin/settings/media`: SOP, template, PIC, terms, lead day, rentang jam, interval slot, dan kapasitas harian.

Kedua halaman hanya dapat diakses `ADMIN_MEDINFO`.

## Verifikasi

- `go test ./...`
- `pnpm check`
- `pnpm build`
- migrasi dan seeder MySQL melalui Docker
- smoke test API untuk RBAC, ownership, timeline, metadata attachment, dan article handoff
- smoke test SSR untuk halaman Admin, redirect Medinfo-only, form baru, settings, master kementerian, dan rich editor.
