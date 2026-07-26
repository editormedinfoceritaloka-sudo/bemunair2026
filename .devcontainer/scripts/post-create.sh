#!/usr/bin/env bash

set -Eeuo pipefail

PROJECT_ROOT="/workspace"

echo "========================================"
echo "Configuring BEM UNAIR development"
echo "========================================"

echo "Enabling Corepack..."
corepack enable

install_node_dependencies() {
  local directory="$1"

  if [[ ! -f "${directory}/package.json" ]]; then
    echo "Skipping ${directory}: package.json not found."
    return
  fi

  echo "Installing dependencies in ${directory}..."
  cd "${directory}"

  if [[ -f pnpm-lock.yaml ]]; then
    pnpm install --frozen-lockfile
  else
    pnpm install
  fi
}

echo "Preparing node_modules directories..."

sudo mkdir -p \
  "${PROJECT_ROOT}/client/node_modules" \
  "${PROJECT_ROOT}/wa-engine/node_modules"

sudo chown -R node:node \
  "${PROJECT_ROOT}/client/node_modules" \
  "${PROJECT_ROOT}/wa-engine/node_modules"

install_node_dependencies "${PROJECT_ROOT}/client"
install_node_dependencies "${PROJECT_ROOT}/wa-engine"

if [[ -f "${PROJECT_ROOT}/server/go.mod" ]]; then
  echo "Downloading Go modules..."

  cd "${PROJECT_ROOT}/server"
  go mod download
fi

echo
echo "========================================"
echo "Development workspace is ready"
echo "========================================"
echo
echo "Nginx App : http://localhost:8081"
echo "SvelteKit : http://localhost:3000"
echo "Go API    : http://localhost:8080"
echo "WA Engine : http://localhost:3001"
echo "MySQL     : localhost:3308"
echo
