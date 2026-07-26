# BEM UNAIR 2026 Digital Workspace

Platform terpusat BEM UNAIR 2026 untuk pengajuan media, pengajuan surat, pengelolaan artikel, antrean PJ Medinfo, notifikasi WhatsApp, serta pemantauan progres melalui dashboard admin.

## Daftar Isi

- [Fitur utama](#fitur-utama)
- [Arsitektur](#arsitektur)
- [Role dan hak akses](#role-dan-hak-akses)
- [Menjalankan development](#menjalankan-development)
- [Database dan seeder](#database-dan-seeder)
- [API dan contoh curl](#api-dan-contoh-curl)
- [WhatsApp dan ImageKit](#whatsapp-dan-imagekit)
- [Testing](#testing)
- [Deployment production](#deployment-production)
- [GitHub Actions](#github-actions)
- [Rollback dan backup](#rollback-dan-backup)
- [Checklist keamanan](#checklist-keamanan)

## Fitur Utama

- Login admin dengan role `ADMIN` dan `ADMIN_MEDINFO`.
- Pengajuan Konten Medinfo: Feed Instagram, Reels, Instastory, caption, Canva, brief, gambar, serta link Google Drive untuk video.
- Pengajuan Artikel dengan SOP, template, Google Docs, dokumentasi, caption, dan workflow editorial.
- Pengajuan surat, template surat, timeline status, dan assignment PJ.
- Dashboard admin dengan ringkasan, kalender bulanan, deadline, detail kegiatan, dan timeline request.
- Master kementerian yang tersimpan di database dan digunakan sebagai dropdown konsisten.
- Antrean PJ Medinfo, assignment manual, serta proteksi satu task aktif per PJ.
- Rich text editor artikel bergaya Notion dengan upload gambar melalui ImageKit.
- WhatsApp untuk konfirmasi request, assignment PJ, pembaruan status, dan reminder deadline H-3.
- Dokumentasi API yang dapat dibaca dari aplikasi.

## Arsitektur

```text
Browser
   |
   v
TLS reverse proxy pada host VPS (Caddy/Nginx/Traefik)
   |
   v
Nginx container :8081
   |--------------------------|
   v                          v
SvelteKit client :3000        Go API :8080
   |                          |
   |                          +---- MySQL :3306
   |                          |
   +---- ImageKit             +---- WA Engine :3001 ---- WhatsApp
```

Semua port backend hanya tersedia di jaringan Docker. Compose production secara default hanya mengikat Nginx ke `127.0.0.1:8081`, sehingga TLS harus dihentikan oleh reverse proxy pada host VPS.

### Stack

| Bagian | Teknologi |
| --- | --- |
| Client | SvelteKit 2, Svelte 5, TypeScript, Tailwind CSS, shadcn-svelte, Tiptap |
| Server | Go 1.25, Gin, GORM |
| Database | MySQL 8, migrasi SQL embedded |
| WhatsApp | Node.js 22, TypeScript, Baileys |
| Upload | ImageKit |
| Runtime | Docker Compose dan Nginx |
| CI/CD | GitHub Actions, GHCR, SSH ke VPS |

## Role dan Hak Akses

| Kemampuan | `ADMIN` | `ADMIN_MEDINFO` |
| --- | :---: | :---: |
| Login dashboard | Ya | Ya |
| Membuat request | Ya | Ya |
| Melihat request kementeriannya | Ya | Ya |
| Melihat timeline dan progres | Ya | Ya |
| Mengelola seluruh request | Tidak | Ya |
| Assignment PJ dan status operasional | Tidak | Ya |
| Mengelola user, kementerian, setting, artikel, dan template | Tidak | Ya |

PJ bukan role terpisah. PJ merupakan user `ADMIN_MEDINFO` dari kementerian `MEDINFO` yang dimasukkan ke roster `medinfo_pj_queues`.

## Struktur Repository

```text
.
├── client/                 # SvelteKit dan UI admin
├── server/                 # Go API, migrasi, dan seeder
├── wa-engine/              # Integrasi WhatsApp
├── docs/api/               # Dokumentasi endpoint
├── docs/plan/              # Planning UI dan arsitektur
├── infra/
│   ├── docker-compose.dev.yml
│   ├── docker-compose.yml  # Production
│   ├── deploy.sh
│   └── nginx/
└── .github/
    ├── workflows/deploy.yml
    └── dependabot.yml
```

## Menjalankan Development

### Prasyarat

- Docker Engine dengan Docker Compose v2.
- Port lokal `3000`, `3001`, `3308`, `8080`, dan `8081` tersedia.
- Akun ImageKit bila ingin menguji upload nyata.
- Akun WhatsApp khusus organisasi bila ingin menguji notifikasi nyata.

### Persiapan environment

```bash
cp server/.env.example server/.env
cp client/.env.example client/.env
cp wa-engine/.env.example wa-engine/.env
```

Pastikan `WA_ENGINE_API_KEY` pada `server/.env` sama dengan `wa-engine/.env`. Private key ImageKit hanya boleh diletakkan di `client/.env` dan tidak boleh memakai prefix `PUBLIC_`.

### Start development stack

```bash
docker compose -f infra/docker-compose.dev.yml up --build
```

Jalankan di background:

```bash
docker compose -f infra/docker-compose.dev.yml up -d --build
docker compose -f infra/docker-compose.dev.yml logs -f
```

Stop container tanpa menghapus data:

```bash
docker compose -f infra/docker-compose.dev.yml down
```

Reset seluruh database dan session development:

```bash
docker compose -f infra/docker-compose.dev.yml down -v
```

Perintah terakhir bersifat destruktif dan hanya boleh digunakan untuk data development.

### URL development

| Service | URL |
| --- | --- |
| Frontend langsung | `http://localhost:3000` |
| Nginx gateway | `http://localhost:8081` |
| API langsung | `http://localhost:8080/api/v1` |
| API melalui gateway | `http://localhost:8081/api/v1` |
| Docs UI | `http://localhost:3000/api/docs` |
| WA Engine | `http://localhost:3001` |
| MySQL host | `127.0.0.1:3308` |

### Hot reload Docker

Development stack sudah mendukung update otomatis:

- perubahan Go dipantau Air dan server dibangun ulang;
- perubahan Svelte dipantau Vite;
- perubahan WA Engine dipantau `tsx watch`;
- database tersimpan pada named volume dan tidak hilang saat container restart.

Jika dependency atau Dockerfile berubah, jalankan ulang dengan `--build`.

## Environment

### Server

| Variabel | Keterangan |
| --- | --- |
| `APP_ENV`, `PORT`, `TZ` | Mode, port, dan zona waktu server |
| `DB_HOST`, `DB_PORT`, `DB_NAME`, `DB_USER`, `DB_PASSWORD` | Koneksi MySQL |
| `JWT_SECRET` | Kunci penandatanganan JWT; gunakan random minimal 32 byte |
| `CORS_ALLOWED_ORIGINS` | Daftar origin dipisahkan koma |
| `WA_ENGINE_URL`, `WA_ENGINE_API_KEY` | Koneksi internal WA Engine |
| `WA_GROUP_JID_1`, `WA_GROUP_JID_2` | Grup penerima reminder opsional |
| `ADMIN_DASHBOARD_URL` | URL publik untuk link detail pada pesan WhatsApp |
| `DOCS_DIR` | Lokasi dokumentasi API |

### Client

| Variabel | Keterangan |
| --- | --- |
| `API_INTERNAL_URL` | URL Go API dari proses SSR |
| `WA_INTERNAL_URL` | URL WA Engine dari proses SSR |
| `ORIGIN` | Origin publik SvelteKit di production |
| `IMAGEKIT_PUBLIC_KEY` | Public key ImageKit |
| `IMAGEKIT_PRIVATE_KEY` | Private key server-side, jangan diekspos ke browser |
| `IMAGEKIT_UPLOAD_FOLDER` | Root folder upload |
| `IMAGEKIT_MAX_UPLOAD_BYTES` | Batas upload gambar artikel |
| `IMAGEKIT_MAX_IMAGE_BYTES` | Batas gambar submission |
| `IMAGEKIT_MAX_BRIEF_BYTES` | Batas dokumen brief |

### WA Engine

| Variabel | Keterangan |
| --- | --- |
| `WA_ENGINE_API_KEY` | Bearer key internal, harus sama dengan server |
| `WA_SESSION_DIR` | Volume session WhatsApp |
| `WA_AUTO_CONNECT` | Menghubungkan session otomatis saat start |
| `WA_SENDER_PHONE` | Label nomor sender dalam format `62...` |

Gunakan file `.env.example` pada masing-masing service sebagai sumber lengkap. Jangan commit `.env`, `.env.production`, private key, password, QR, atau folder session.

## Database dan Seeder

Migrasi dan seeder di-embed ke binary Go saat image dibangun.

```bash
cd server
go run ./cmd migrate  # hanya migrasi yang belum dijalankan
go run ./cmd seed     # jalankan seeder idempotent
go run ./cmd setup    # migrate lalu seed
```

Migration manager mencatat migrasi pada tabel `schema_migrations`. File migrasi saat ini berada di `server/database/migrations/`, sedangkan master data development berada di `server/database/seeders/`.

Seeder hanya untuk development atau bootstrap yang disetujui. Deployment production menjalankan `./server migrate` dan tidak pernah menjalankan seeder otomatis.

Akun seed development memakai password contoh `password`. Password tersebut wajib diganti dan tidak boleh digunakan di production.

## API dan Contoh Curl

Base path API adalah `/api/v1`. Dokumentasi rinci tersedia di `docs/api/` dan endpoint `GET /api/v1/docs`.

### Kelompok endpoint

| Modul | Endpoint utama |
| --- | --- |
| Auth | `/api/v1/auth/login`, `/api/v1/auth/me` |
| User | `/api/v1/users` |
| Kementerian | `/api/v1/ministries` |
| Media submission | `/api/v1/content-submissions` |
| Surat | `/api/v1/letter-submissions` |
| PJ Medinfo | `/api/v1/medinfo-pj/queue` |
| Template surat | `/api/v1/letter-templates` |
| Artikel publik | `/api/v1/articles` |
| Artikel admin | `/api/v1/admin/articles` |
| Setting media | `/api/v1/media-submission-settings/:serviceType` |
| Docs | `/api/v1/docs` |

### Login dan membaca user aktif

```bash
TOKEN=$(curl -s http://localhost:8081/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@bem.unair.ac.id","password":"password"}' \
  | jq -r '.data.token')

curl -s http://localhost:8081/api/v1/auth/me \
  -H "Authorization: Bearer $TOKEN" \
  | jq
```

Response sukses memakai envelope berikut:

```json
{
  "status": true,
  "message": "Login berhasil",
  "data": {}
}
```

## WhatsApp dan ImageKit

### WhatsApp

Lihat status koneksi dan QR pada development:

```bash
docker compose -f infra/docker-compose.dev.yml logs -f wa-engine
curl -H "Authorization: Bearer $WA_ENGINE_API_KEY" http://localhost:3001/api/connect
curl -H "Authorization: Bearer $WA_ENGINE_API_KEY" http://localhost:3001/api/qr
```

Notifikasi assignment dikirim ke nomor user yang dipilih sebagai PJ. Reminder deadline berjalan pukul 12.00 WIB untuk task dalam rentang H-3 sampai maksimal 24 jam setelah deadline. Pesan menggunakan format WhatsApp ringkas dengan kode request, layanan, judul, kementerian, deadline, status, dan tindak lanjut.

WA Engine menggunakan Baileys, bukan WhatsApp Business Cloud API resmi. Gunakan akun organisasi khusus, pantau perubahan kompatibilitas, dan siapkan rencana migrasi ke API resmi bila sistem menjadi kritikal.

### ImageKit

Browser meminta signature jangka pendek dari SvelteKit, kemudian mengunggah file ke ImageKit. Private key tetap berada di server SvelteKit dan tidak pernah dikirim ke browser.

- Artikel menggunakan uploader gambar pada rich editor.
- Gambar submission menggunakan upload form.
- Video final diarahkan ke Google Drive agar tidak membebani server.

## Testing

Jalankan seluruh test utama:

```bash
cd server && go test ./...
cd ../client && pnpm install --frozen-lockfile && pnpm check && pnpm build
cd ../wa-engine && pnpm install --frozen-lockfile && pnpm test
```

Validasi infrastruktur production:

```bash
bash -n infra/deploy.sh
docker compose --env-file infra/.env.production.example \
  -f infra/docker-compose.yml config --quiet
docker run --rm \
  --add-host client:127.0.0.1 --add-host server:127.0.0.1 \
  -v "$PWD/infra/nginx/nginx.conf:/etc/nginx/nginx.conf:ro" \
  nginx:alpine nginx -t
```

## Deployment Production

### Model deployment

Production tidak membangun source di VPS. GitHub Actions:

1. menjalankan test Go, SvelteKit, dan WA Engine;
2. membangun tiga image Docker;
3. menambahkan SBOM dan provenance;
4. push image bertag immutable `sha-<commit>` ke GHCR;
5. mengunggah Compose, Nginx, skrip deploy, dan docs ke VPS;
6. membuat backup MySQL;
7. menjalankan migrasi one-shot;
8. mengganti container;
9. memeriksa health API, client, dan WA Engine;
10. mengembalikan image sebelumnya jika deployment gagal.

Database dan session WhatsApp berada di named volume dan tidak diganti saat release baru.

### Persiapan VPS

VPS membutuhkan:

- Linux 64-bit;
- Docker Engine dan Docker Compose v2;
- `curl`, `tar`, dan OpenSSH server;
- user deploy non-root yang memiliki akses terbatas ke Docker;
- firewall yang hanya membuka SSH, HTTP, dan HTTPS;
- reverse proxy TLS pada host;
- ruang disk untuk image, database, session WhatsApp, dan backup.

Siapkan direktori aplikasi sesuai variable `VPS_APP_DIR`:

```bash
sudo install -d -m 750 -o deploy -g deploy /opt/bemunair2026/infra
sudo install -d -m 750 -o deploy -g deploy /opt/bemunair2026/backups
```

Pada bootstrap pertama, workflow akan mengunggah `infra/.env.production.example`, tetapi deploy sengaja berhenti aman selama `.env.production` belum ada. Setelah job pertama mencapai tahap VPS, masuk ke VPS dan buat konfigurasi rahasia:

```bash
cd /opt/bemunair2026/infra
cp .env.production.example .env.production
chmod 600 .env.production
nano .env.production
```

Ganti seluruh placeholder. Jangan menjalankan production sebelum nilai JWT, database, ImageKit, domain, dan WA API key sudah benar. Setelah file tersimpan, jalankan ulang workflow melalui `Actions -> Deploy Production -> Run workflow`; deploy berikutnya tidak menimpa `.env.production`.

### TLS pada host

Compose mengikat Nginx ke `127.0.0.1:8081`. Contoh sederhana Caddy:

```caddyfile
bem.example.org {
  reverse_proxy 127.0.0.1:8081
}
```

Arahkan DNS ke VPS, aktifkan HTTPS, lalu samakan `APP_ORIGIN`, `CORS_ALLOWED_ORIGINS`, dan `ADMIN_DASHBOARD_URL` dengan domain production. Jangan mengekspos port MySQL `3306` atau WA Engine `3001` ke internet.

## GitHub Actions

Workflow berada di `.github/workflows/deploy.yml`. Pull request hanya menjalankan validasi. Push ke `main` atau eksekusi manual akan melakukan build, publish, dan deploy.

Workflow menggunakan GitHub Container Registry dengan permission minimum, image commit SHA, SBOM, provenance, SSH known-host verification, environment `production`, concurrency lock, backup, healthcheck, dan rollback.

### GitHub Environment

Buat environment bernama `production` melalui:

```text
Repository Settings -> Environments -> New environment -> production
```

Disarankan:

- batasi deployment hanya dari branch `main`;
- aktifkan required reviewer bila paket GitHub mendukung;
- cegah self-approval bila tersedia;
- simpan secret VPS sebagai environment secret.

Environment GitHub menahan akses secret sampai protection rule terpenuhi. Dokumentasi resmi: [GitHub deployment environments](https://docs.github.com/en/actions/concepts/workflows-and-actions/deployment-environments).

### Secrets wajib

| Secret | Isi |
| --- | --- |
| `VPS_HOST` | IP atau hostname VPS tanpa protokol |
| `VPS_USER` | User SSH deployment, misalnya `deploy` |
| `VPS_SSH_PRIVATE_KEY` | Private key Ed25519 khusus CI/CD |
| `VPS_KNOWN_HOSTS` | Baris host key VPS yang sudah diverifikasi |

Secret opsional:

| Secret | Isi |
| --- | --- |
| `GHCR_PULL_TOKEN` | Fine-grained/classic token `read:packages` untuk private image; bila kosong workflow memakai `GITHUB_TOKEN` |

Gunakan deploy key khusus, tanpa password interaktif, dan jangan memakai private key personal. Public key-nya ditempatkan pada `$HOME/.ssh/authorized_keys` user deploy.

Buat known-host entry dari jaringan tepercaya, lalu cocokkan fingerprint dengan panel/provider VPS:

```bash
ssh-keyscan -p 22 your-vps.example.org
```

Simpan hasil yang sudah diverifikasi sebagai `VPS_KNOWN_HOSTS`. Workflow tidak menjalankan `ssh-keyscan` otomatis agar tidak mempercayai host yang belum diverifikasi.

### Variables

| Variable | Default | Keterangan |
| --- | --- | --- |
| `VPS_PORT` | `22` | Port SSH |
| `VPS_APP_DIR` | `/opt/bemunair2026` | Direktori deployment |
| `PRODUCTION_URL` | kosong | URL publik untuk healthcheck tambahan, misalnya `https://bem.example.org` |

### Menyesuaikan workflow

- Branch deploy: ubah `on.push.branches` dan kondisi `refs/heads/main`.
- Registry: ubah login registry dan `image_prefix` jika tidak memakai GHCR.
- Direktori VPS: ubah variable `VPS_APP_DIR`, bukan hardcode workflow.
- Domain/port aplikasi: ubah `.env.production` pada VPS.
- Approval: atur protection rule pada environment `production`.
- Image action: action dipin ke commit SHA. Dependabot akan membuat PR saat versi aman baru tersedia.

GitHub merekomendasikan `GITHUB_TOKEN` untuk publish GHCR dan permission `packages: write`. Lihat [publishing Docker images](https://docs.github.com/en/actions/tutorials/publish-packages/publish-docker-images). Docker menjelaskan SBOM/provenance pada [build attestations](https://docs.docker.com/build/ci/github-actions/attestations/).

## Rollback dan Backup

Setiap deploy menghasilkan:

```text
/opt/bemunair2026/backups/pre-deploy-<sha>-<timestamp>.sql
/opt/bemunair2026/infra/.release.env
/opt/bemunair2026/infra/.release.env.previous
```

Rollback manual image aplikasi:

```bash
cd /opt/bemunair2026/infra
cp .release.env.previous .release.env
docker compose --env-file .env.production --env-file .release.env \
  -f docker-compose.yml up -d --remove-orphans
```

Rollback image tidak otomatis membatalkan migrasi database. Semua migrasi production harus backward-compatible. Untuk restore database, hentikan traffic, verifikasi file backup, dan lakukan prosedur restore yang sudah diuji pada staging.

Backup lokal VPS bukan backup final. Salin backup terenkripsi ke penyimpanan off-site dan terapkan retention policy.

## Checklist Keamanan

### Sudah diterapkan di repository

- Tidak ada password database hardcoded pada Compose production.
- Database, Go API, SvelteKit, dan WA Engine tidak mempublikasikan port host.
- Nginx production bind ke loopback secara default.
- Secret production di-ignore Git.
- Image aplikasi bertag commit SHA dan workflow action dipin ke commit SHA.
- Container aplikasi Node dan Go berjalan sebagai non-root.
- Filesystem server/client read-only dengan tmpfs terbatas.
- `no-new-privileges`, log rotation, healthcheck, dan restart policy.
- Migrasi terpisah dan seeder tidak berjalan otomatis di production.
- Backup sebelum migrasi dan rollback image otomatis.
- SSH known-host verification dan deployment environment GitHub.
- SBOM dan provenance image.

### Wajib diselesaikan pada VPS

- Aktifkan HTTPS dan renewal sertifikat.
- Aktifkan firewall serta batasi SSH dengan key-only authentication.
- Gunakan secret random kuat dan lakukan rotasi berkala.
- Lindungi environment `production` dengan branch policy/reviewer.
- Pastikan package GHCR private hanya dapat dibaca akun deploy.
- Simpan backup terenkripsi di lokasi off-site dan uji restore.
- Pantau disk, memory, container health, error rate, dan status WhatsApp.
- Jangan menjalankan seeder development di database production.
- Pin `MYSQL_IMAGE` dan `NGINX_IMAGE` ke digest setelah diuji di staging.

Konfigurasi ini merupakan baseline deployment yang aman, tetapi keamanan akhir tetap bergantung pada hardening VPS, TLS, firewall, pengelolaan secret, monitoring, dan backup eksternal.

## Dokumentasi API

Dokumentasi modul berada di `docs/api/`:

- overview;
- authentication;
- users;
- content submissions;
- letter submissions;
- Medinfo PJ queue;
- letter templates;
- articles;
- WA Engine.

Untuk menambah halaman, buat `docs/api/<slug>.md` lalu tambahkan entri pada `docs/api/index.json`.
