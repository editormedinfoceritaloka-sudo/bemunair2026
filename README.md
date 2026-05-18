# bemunair2026

Platform digital untuk BEMUNAIR 2026 — sistem informasi dan manajemen kegiatan organisasi mahasiswa.

---

## 🧱 Tech Stack

| Layer | Teknologi |
|---|---|
| Frontend | SvelteKit + TailwindCSS |
| Backend | Go + Gin |
| Database | MySQL 8.0 |
| Reverse Proxy | Nginx |
| Containerization | Docker + Docker Compose |

---

## 📁 Struktur Project

```
bemunair2026/
├── client/          # SvelteKit (frontend)
├── server/          # Go + Gin (backend/API)
├── infra/           # Docker Compose, Nginx config
│   ├── nginx/
│   │   └── nginx.conf
│   ├── docker-compose.yml
│   └── dev.sh
├── test/
└── README.md
```

---

## 🌐 Arsitektur

```
Browser
   │
   ▼
Nginx :80
   ├── /          → client (SvelteKit) :3000
   └── /api/      → server (Go Gin)   :8080  [strip /api prefix]
                          │
                          ▼
                     MySQL :3306
```

Semua service berjalan dalam Docker internal network `bemunair_net`. Hanya Nginx yang expose port ke host.

---

## 🚀 Menjalankan Project

### Prasyarat

- Docker Engine
- Docker Compose v2.22+

### Langkah

```bash
# 1. Clone repository
git clone <repo-url>
cd bemunair2026

# 2. Salin file environment
cp client/.env.example client/.env
cp server/.env.example server/.env

# 3. Masuk ke folder infra
cd infra

# 4. Jalankan dengan watch mode (development)
./dev.sh up

# Atau tanpa script helper
docker compose watch
```

### Akses

| Service | URL |
|---|---|
| Frontend | http://localhost |
| API | http://localhost/api/ |
| MySQL | localhost:3308 |

---

## 🛠️ Script Helper

```bash
cd infra

./dev.sh up              # Start dengan watch mode
./dev.sh down            # Stop semua container
./dev.sh build           # Build ulang tanpa cache
./dev.sh logs [service]  # Lihat logs
./dev.sh reset           # Reset + hapus database
./dev.sh ps              # Status container
```

---

## ⚙️ Environment Variables

### `client/.env`

```env
PUBLIC_API_URL=http://server:8080
PUBLIC_APP_NAME=bemunair2026
PORT=3000
```

### `server/.env`

```env
APP_ENV=development
APP_PORT=8080
APP_SECRET_KEY=your_secret_key_here

DB_HOST=db
DB_PORT=3306
DB_USER=bemunair
DB_PASSWORD=bemunair_password
DB_NAME=bemunair_db
DATABASE_URL=bemunair:bemunair_password@tcp(db:3306)/bemunair_db?charset=utf8mb4&parseTime=True&loc=Local

ALLOWED_ORIGINS=http://localhost
```

---

## 📡 API Endpoints

| Method | Endpoint | Deskripsi |
|---|---|---|
| GET | `/api/ping` | Health check |

> Dokumentasi API lengkap menyusul.

---

## 📦 Docker Watch Mode

Watch mode memungkinkan perubahan kode langsung ter-reflect di container tanpa rebuild manual.

| Service | Trigger | Action |
|---|---|---|
| client | Perubahan `src/` | sync |
| client | Perubahan `package.json` | rebuild |
| server | Perubahan file `.go` | sync + restart |
| server | Perubahan `go.mod` / `go.sum` | rebuild |

---

## 🤝 Kontribusi

1. Fork repository ini
2. Buat branch fitur: `git checkout -b feat/nama-fitur`
3. Commit perubahan: `git commit -m "feat: deskripsi singkat"`
4. Push branch: `git push origin feat/nama-fitur`
5. Buat Pull Request

---

## 📄 Lisensi

MIT License — BEMUNAIR 2026