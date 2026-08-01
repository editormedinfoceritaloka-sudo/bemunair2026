# AGENTS.md

## Ringkasan proyek

BEM UNAIR 2026 Digital Workspace adalah platform internal untuk pengajuan media, surat, artikel, antrean PJ Medinfo, dashboard admin, dan notifikasi WhatsApp.

- `client/`: SvelteKit 2, Svelte 5, TypeScript, Tailwind, shadcn-svelte, dan Tiptap.
- `server/`: Go 1.25 dengan Gin dan GORM, termasuk migration SQL dan seeder.
- `wa-engine/`: Node.js 22, TypeScript, Express, Baileys, dan Vitest.
- `infra/`: Docker Compose development/production dan konfigurasi Nginx.
- `docs/api/`: dokumentasi endpoint API.
- `docs/plan/`: rancangan UI dan arsitektur.

## Arsitektur development

Alur service development adalah `nginx:8081 -> client:3000`, `nginx:8081 -> server:8080`, `server -> db:3306`, dan `server -> wa-engine:3001`.

Gunakan `infra/docker-compose.dev.yml` sebagai sumber konfigurasi service aplikasi. File `.devcontainer/docker-compose.extend.yml` hanya menambahkan container `workspace` untuk VS Code. Compose production berada di `infra/docker-compose.yml` dan tidak digunakan untuk development.

## Menjalankan project

Siapkan environment lokal dari contoh:

- `cp server/.env.example server/.env`
- `cp client/.env.example client/.env`
- `cp wa-engine/.env.example wa-engine/.env`

Pastikan `WA_ENGINE_API_KEY` pada server dan WA Engine sama. Jalankan stack biasa dengan `docker compose -f infra/docker-compose.dev.yml up --build`. Untuk VS Code gunakan `Dev Containers: Reopen in Container`.

Port development: `3000` client, `3001` WA Engine, `3308` MySQL host, `8080` API, dan `8081` Nginx gateway.

## Aturan perubahan

- Baca pola yang sudah ada sebelum menambah abstraksi atau dependency.
- Buat perubahan sekecil mungkin dan jangan membatalkan perubahan pengguna yang sudah ada.
- Jangan commit `.env`, password, JWT secret, private key ImageKit, QR WhatsApp, atau data `sessions/`.
- Jangan menjalankan `docker compose down -v`, menghapus volume, atau reset database tanpa kebutuhan jelas dan persetujuan eksplisit.
- Jika mengubah endpoint, DTO, response, atau auth, perbarui client dan dokumentasi `docs/api/` yang terkait.
- Jika mengubah skema database, tambahkan migration SQL baru bernomor urut. Jangan mengedit migration yang sudah pernah dijalankan; perbarui entity dan seeder terkait.

## Konvensi service

### Client

- Gunakan Node.js 22 dan `pnpm`.
- Route ada di `client/src/routes/`; secret harus tetap server-side.
- Jangan mengekspos `IMAGEKIT_PRIVATE_KEY` atau secret lain dengan prefix `PUBLIC_`.
- Ikuti komponen dan utilitas yang sudah ada di `client/src/lib/`.

### Server

- Pertahankan pemisahan route, controller, DTO, validation, service, dan repository.
- Jalankan `gofmt` pada file Go yang berubah.
- Pertahankan envelope error/response API dan jangan menambahkan kredensial hard-coded.
- Perubahan role, authorization, status submission, assignment PJ, dan cron reminder harus ditinjau terhadap endpoint terkait.

### WA Engine

- Gunakan TypeScript dan `pnpm`.
- Pertahankan autentikasi internal melalui `WA_ENGINE_API_KEY`.
- Jangan menghapus atau mengubah session WhatsApp secara otomatis; session disimpan pada volume `wa_auth_data`.
- Perubahan event, template pesan, atau health endpoint perlu dicek terhadap pemanggilnya di server.

## Validasi

Perintah test utama adalah `make test`. Validasi per service: `cd client && pnpm check && pnpm lint && pnpm build`, `cd wa-engine && pnpm test`, dan `cd server && go test ./... -v`.

Untuk perubahan Compose atau Dev Container, jalankan `docker compose -f infra/docker-compose.dev.yml config --quiet`, validasi compose gabungan dengan tambahan `-f .devcontainer/docker-compose.extend.yml config --quiet`, serta `bash -n .devcontainer/scripts/post-create.sh`.

Jika validasi gagal karena dependency, Docker daemon, credential, atau service eksternal, laporkan perintah yang gagal dan penyebabnya secara spesifik.
