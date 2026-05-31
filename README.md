# BEM UNAIR Digital Submission System

Sistem pengajuan konten dan surat digital untuk BEM UNAIR.

## Stack

Backend Go + Gin + GORM MySQL, frontend SvelteKit, MySQL 8, WA Engine TypeScript + Baileys, Docker Compose, dan Nginx reverse proxy.

## Menjalankan

```bash
cp server/.env.example server/.env
cp client/.env.example client/.env
cp wa-engine/.env.example wa-engine/.env
cd infra
docker compose up --build
```

Akses:

| Service | URL |
|---|---|
| Frontend | http://localhost:8081 |
| API | http://localhost:8081/api |
| Docs viewer | http://localhost:8081/api/docs |
| WA Engine internal | http://bemunair_wa_engine:3001 |
| MySQL | localhost:3308 |

## Database

Schema produksi berasal dari SQL canonical di `server/database/migrations/`:

1. `001_create_users_table.sql`
2. `002_create_content_submissions_table.sql`
3. `003_create_letter_submissions_table.sql`
4. `004_create_medinfo_pj_queues_table.sql`
5. `005_create_letter_templates_table.sql`

Seeder eksplisit ada di `server/database/seeders/001_seed_users_and_queue.sql`. Password seed memakai bcrypt untuk nilai contoh `password`.

Untuk verifikasi dev saja:

```bash
cd server
go run ./cmd/migrate
```

## Endpoint Utama

Auth: `POST /api/auth/register`, `POST /api/auth/login`, `GET /api/auth/me`.

Users: `GET|POST /api/users`, `GET|PUT|DELETE /api/users/:id`.

Content submissions: `POST|GET /api/content-submissions`, `GET /api/content-submissions/:id`, `PUT /api/content-submissions/:id/status`, `DELETE /api/content-submissions/:id`.

Letter submissions: `POST|GET /api/letter-submissions`, `GET /api/letter-submissions/:id`, `PUT /api/letter-submissions/:id/status`, `DELETE /api/letter-submissions/:id`.

Queue PJ: `GET|POST /api/medinfo-pj/queue`, `PUT /api/medinfo-pj/queue/reorder`, `DELETE /api/medinfo-pj/queue/:id`.

Templates: `POST|GET /api/letter-templates`, `GET|PUT|DELETE /api/letter-templates/:id`.

Docs API: `GET /api/docs`, `GET /api/docs/:slug`.

## Contoh Curl

```bash
TOKEN=$(curl -s -X POST http://localhost:8081/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@bem.unair.ac.id","password":"password"}')

curl -X POST http://localhost:8081/api/content-submissions \
  -H "Authorization: Bearer $JWT" \
  -F platform=INSTAGRAM \
  -F submission_type=Feed \
  -F deadline=2026-06-10T12:00:00+07:00 \
  -F caption="Caption konten"
```

Success envelope:

```json
{ "success": true, "message": "Submission berhasil dibuat", "data": {} }
```

Error envelope:

```json
{ "success": false, "message": "Validasi gagal", "error": { "code": "VALIDATION_ERROR", "details": [] } }
```

## WhatsApp QR

Lihat QR di log:

```bash
docker compose logs -f bemunair_wa_engine
```

Atau via web/API:

```bash
curl http://localhost:3001/api/connect -H "Authorization: Bearer your_internal_api_key_here"
curl http://localhost:3001/api/qr -H "Authorization: Bearer your_internal_api_key_here"
```

## Test

```bash
make test
make test-coverage
```

Catatan lingkungan ini: `go test ./...` berhasil. Coverage Go dengan toolchain terunduh gagal karena tool `covdata` tidak tersedia di cache toolchain lokal. WA Engine `pnpm test` berhasil: 3 file, 7 test.

## Dokumentasi API

File docs ada di `docs/api/`: `index.json`, `overview.md`, `auth.md`, `users.md`, `content-submissions.md`, `letter-submissions.md`, `medinfo-pj-queue.md`, `letter-templates.md`, `wa-engine.md`.

Untuk menambah modul dokumentasi, buat file `docs/api/<slug>.md`, lalu tambahkan `{ "slug": "<slug>", "title": "...", "order": n }` ke `docs/api/index.json`.
