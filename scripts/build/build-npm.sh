#!/usr/bin/env bash
set -euo pipefail

source "$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)/common.sh"

usage() {
  cat <<'EOF'
Usage:
  scripts/build/build-npm.sh

Environment:
  NATIVE_VERSION=1.2.3-42
  NPM_PACKAGE_TARGETS="darwin/amd64 darwin/arm64 linux/amd64 linux/arm64 windows/amd64 windows/arm64"
EOF
}

if [[ "${1:-}" == "--help" || "${1:-}" == "-h" ]]; then
  usage
  exit 0
fi

ROOT="$(repo_root)"
cd "$ROOT"

VERSION_VALUE="$(native_version)"
TARGETS="${NPM_PACKAGE_TARGETS:-darwin/amd64 darwin/arm64 linux/amd64 linux/arm64 windows/amd64 windows/arm64}"
OUT_DIR="bin/native"

validate_native_version "$VERSION_VALUE"

write_sha256() {
  local digest file
  file="$1"
  if command -v sha256sum >/dev/null 2>&1; then
    read -r digest _ < <(sha256sum "$file")
  elif command -v shasum >/dev/null 2>&1; then
    read -r digest _ < <(shasum -a 256 "$file")
  else
    printf 'ERROR: sha256sum or shasum is required\n' >&2
    exit 1
  fi
  if [[ ! "$digest" =~ ^[[:xdigit:]]{64}$ ]]; then
    printf 'ERROR: failed to calculate SHA-256 for %s\n' "$file" >&2
    exit 1
  fi
  printf '%s  %s\n' "$digest" "$(basename "$file")" > "${file}.sha256"
  chmod 0644 "${file}.sha256"
}

build_wasm_exporter
mkdir -p "$OUT_DIR"

for target in $TARGETS; do
  target_os="${target%/*}"
  target_arch="${target#*/}"
  suffix=""
  if [[ "$target_os" == "windows" ]]; then
    suffix=".exe"
  fi
  output="${OUT_DIR}/xaligo-${target_os}-${target_arch}${suffix}"
  build_native_binary "$VERSION_VALUE" "$output" "$target_os" "$target_arch"
  chmod 0755 "$output"
  write_sha256 "$output"
  printf 'Built: %s\n' "$output"
  printf 'Checksum: %s\n' "${output}.sha256"
done
