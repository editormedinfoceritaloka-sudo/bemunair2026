#!/usr/bin/env bash
set -Eeuo pipefail

GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m'

info()  { echo -e "${GREEN}[INFO]${NC} $1"; }
warn()  { echo -e "${YELLOW}[WARN]${NC} $1"; }
error() { echo -e "${RED}[ERROR]${NC} $1"; }

script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
app_dir="$(cd "$script_dir/.." && pwd)"
compose_file="$app_dir/infra/docker-compose.dev.yml"
server_dir="$app_dir/server"
server_env="$server_dir/.env"

usage() {
  cat <<'EOF'
Usage: ./scripts/seed.sh [seed|setup] [--local]

  seed     Jalankan seeder saja (default). Idempotent, aman diulang kapan pun.
  setup    Migrate lalu seed — sama seperti command default service
           'migrate' saat `docker compose up`.

  --local  Jalankan langsung di host (`go run ./cmd ...`) tanpa container,
           untuk server/db yang diakses via 127.0.0.1:3308. Default: lewat
           container dev compose (migrate service), konsisten dengan cara
           stack ini biasanya dijalankan.

Dev-only. Production tidak pernah menjalankan seeder (lihat server/cmd/main.go
dan infra/docker-compose.yml — service 'migrate' di sana cuma `./server migrate`).
Semua akun seed berbagi satu password dari SEED_USER_PASSWORD di server/.env;
tidak ada password yang pernah ditulis ke file seeder itu sendiri.
EOF
}

command="seed"
mode="docker"

for arg in "$@"; do
  case "$arg" in
    seed|setup) command="$arg" ;;
    --local) mode="local" ;;
    -h|--help)
      usage
      exit 0
      ;;
    *)
      error "Argumen tidak dikenal: $arg"
      usage
      exit 64
      ;;
  esac
done

if [[ ! -f "$server_env" ]]; then
  error "Tidak ada $server_env."
  error "Jalankan: cp server/.env.example server/.env, lalu isi SEED_USER_PASSWORD (dan secret lain)."
  exit 66
fi

if ! grep -qE '^SEED_USER_PASSWORD=.+' "$server_env"; then
  error "SEED_USER_PASSWORD belum diisi di $server_env — seeder akan gagal tanpa ini."
  exit 66
fi

if [[ "$mode" == "local" ]]; then
  info "Menjalankan 'go run ./cmd $command' langsung di host (DB diarahkan ke 127.0.0.1:3308)..."
  (
    cd "$server_dir"
    set -a
    # shellcheck disable=SC1091
    source .env
    set +a
    # .env pakai hostname Docker (DB_HOST=bemunair_db); dari host, MySQL dev
    # hanya reachable lewat port yang dipublikasikan di docker-compose.dev.yml.
    export DB_HOST=127.0.0.1
    export DB_PORT=3308
    go run ./cmd "$command"
  )
  info "Selesai."
  exit 0
fi

if ! docker compose -f "$compose_file" ps db --status running >/dev/null 2>&1 \
  || [[ -z "$(docker compose -f "$compose_file" ps db --status running -q 2>/dev/null)" ]]; then
  warn "Container db tidak terdeteksi berjalan; 'docker compose run' akan menyalakannya dulu."
fi

info "Menjalankan '$command' lewat container dev (service 'migrate', compose override command)..."
docker compose -f "$compose_file" run --rm migrate go run ./cmd "$command"
info "Selesai. Login pakai email akun seed + SEED_USER_PASSWORD dari server/.env."
