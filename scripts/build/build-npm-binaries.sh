#!/usr/bin/env bash
set -euo pipefail

source "$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)/common.sh"

usage() {
  cat <<'EOF'
Usage:
  scripts/build/build-npm-binaries.sh

Environment:
  VERSION=1.2.3
  NPM_PACKAGE_TARGETS="darwin/amd64 darwin/arm64 linux/amd64 linux/arm64 windows/amd64 windows/arm64"
EOF
}

if [[ "${1:-}" == "--help" || "${1:-}" == "-h" ]]; then
  usage
  exit 0
fi

ROOT="$(repo_root)"
cd "$ROOT"

VERSION_VALUE="$(package_version)"
TARGETS="${NPM_PACKAGE_TARGETS:-darwin/amd64 darwin/arm64 linux/amd64 linux/arm64 windows/amd64 windows/arm64}"
OUT_DIR="npm/bin/native"

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
  printf 'Built: %s\n' "$output"
done

