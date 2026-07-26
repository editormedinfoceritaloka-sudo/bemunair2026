#!/usr/bin/env bash
set -Eeuo pipefail

if [[ $# -ne 4 ]]; then
  echo "Usage: $0 <release-sha> <server-image> <client-image> <wa-engine-image>" >&2
  exit 64
fi

release_sha="$1"
server_image="$2"
client_image="$3"
wa_engine_image="$4"

script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
app_dir="$(cd "$script_dir/.." && pwd)"
production_env="$script_dir/.env.production"
release_env="$script_dir/.release.env"
previous_env="$script_dir/.release.env.previous"
backup_dir="$app_dir/backups"

if [[ ! -f "$production_env" ]]; then
  echo "Missing $production_env. Copy .env.production.example and fill every secret first." >&2
  exit 66
fi

umask 077
mkdir -p "$backup_dir"

if [[ -f "$release_env" ]]; then
  cp "$release_env" "$previous_env"
fi

cat >"$release_env" <<EOF
SERVER_IMAGE=$server_image
CLIENT_IMAGE=$client_image
WA_ENGINE_IMAGE=$wa_engine_image
EOF

compose=(docker compose --env-file "$production_env" --env-file "$release_env" -f "$script_dir/docker-compose.yml")

rollback() {
  exit_code=$?
  if [[ $exit_code -eq 0 ]]; then
    return
  fi

  echo "Deployment failed. Restoring previous application images..." >&2
  if [[ -f "$previous_env" ]]; then
    cp "$previous_env" "$release_env"
    "${compose[@]}" up -d --remove-orphans || true
  else
    echo "No previous release metadata is available for automatic rollback." >&2
  fi
  exit "$exit_code"
}
trap rollback ERR

"${compose[@]}" config --quiet
"${compose[@]}" pull
"${compose[@]}" up -d db

for attempt in {1..30}; do
  if "${compose[@]}" exec -T db sh -c 'mysqladmin ping -h 127.0.0.1 -u"$MYSQL_USER" -p"$MYSQL_PASSWORD" --silent' >/dev/null 2>&1; then
    break
  fi
  if [[ $attempt -eq 30 ]]; then
    echo "Database did not become ready in time." >&2
    exit 1
  fi
  sleep 2
done

# Backup dibuat sebelum image migrasi baru dijalankan.
backup_file="$backup_dir/pre-deploy-${release_sha}-$(date -u +%Y%m%dT%H%M%SZ).sql"
"${compose[@]}" exec -T db sh -c 'MYSQL_PWD="$MYSQL_PASSWORD" exec mysqldump --single-transaction --quick --lock-tables=false -u"$MYSQL_USER" "$MYSQL_DATABASE"' >"$backup_file"
chmod 600 "$backup_file"

"${compose[@]}" up -d --remove-orphans

http_port="$(sed -n 's/^HTTP_PORT=//p' "$production_env" | tail -n 1)"
http_port="${http_port:-8081}"
for attempt in {1..45}; do
  if curl --fail --silent --show-error "http://127.0.0.1:${http_port}/healthz" >/dev/null \
    && curl --fail --silent --show-error "http://127.0.0.1:${http_port}/healthz/client" >/dev/null \
    && "${compose[@]}" exec -T wa-engine node -e "fetch('http://127.0.0.1:3001/health').then(r => process.exit(r.ok ? 0 : 1)).catch(() => process.exit(1))" >/dev/null 2>&1; then
    trap - ERR
    echo "Release $release_sha is healthy."
    "${compose[@]}" ps
    exit 0
  fi
  sleep 4
done

echo "Application health check failed after deployment." >&2
exit 1
