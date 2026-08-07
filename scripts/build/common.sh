#!/usr/bin/env bash
set -euo pipefail

PACKAGE_NAME="xaligo"
PACKAGE_DESCRIPTION="Diagram-as-Code CLI for rendering .xal diagrams"
PACKAGE_URL="https://github.com/xaligo/xaligo"
PACKAGE_MAINTAINER="${PACKAGE_MAINTAINER:-Ryo Arima <ryo-arima@users.noreply.github.com>}"
RUNTIME_REL="usr/lib/${PACKAGE_NAME}"
ENGINE_DIR="external/engine"
ENGINE_PACKAGE="xaligo-engine-ffi"
ENGINE_LINK_ARCHIVE="external/engine/lib/libxaligo_engine.a"
NATIVE_BUILD_TAGS="xaligo_engine sqlite_fts5 sqlite_omit_load_extension"

repo_root() {
  local source_dir
  source_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
  cd "${source_dir}/../.." && pwd
}

version_from_repository() {
  if [[ -s VERSION ]]; then
    sed -n '1{s/^v//;p;q;}' VERSION
    return
  fi
  printf 'ERROR: VERSION file not found or empty\n' >&2
  return 1
}

native_version() {
  if [[ -n "${NATIVE_VERSION:-}" ]]; then
    printf '%s\n' "${NATIVE_VERSION#v}"
    return
  fi
  if [[ -n "${VERSION:-}" ]]; then
    printf '%s\n' "${VERSION#v}"
    return
  fi
  version_from_repository
}

package_version() {
  if [[ -n "${PACKAGE_VERSION:-}" ]]; then
    printf '%s\n' "${PACKAGE_VERSION#v}"
    return
  fi
  if [[ -n "${VERSION:-}" ]]; then
    printf '%s\n' "${VERSION#v}"
    return
  fi
  if [[ -n "${NATIVE_VERSION:-}" ]]; then
    printf '%s\n' "${NATIVE_VERSION#v}"
    return
  fi
  version_from_repository
}

package_release() {
  printf '%s\n' "${PACKAGE_RELEASE:-1}"
}

validate_native_version() {
  local value
  value="$1"
  if [[ ! "$value" =~ ^(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)(-[0-9A-Za-z-]+(\.[0-9A-Za-z-]+)*)?(\+[0-9A-Za-z-]+(\.[0-9A-Za-z-]+)*)?$ ]]; then
    printf 'ERROR: invalid native version: %s\n' "$value" >&2
    return 1
  fi
}

validate_deb_version() {
  local value
  value="$1"
  if [[ ! "$value" =~ ^[0-9][0-9A-Za-z.+:~_-]*$ ]]; then
    printf 'ERROR: invalid Debian package version: %s\n' "$value" >&2
    return 1
  fi
}

validate_rpm_version() {
  local value
  value="$1"
  if [[ ! "$value" =~ ^[0-9A-Za-z.+_~^]+$ ]]; then
    printf 'ERROR: invalid RPM Version: %s\n' "$value" >&2
    return 1
  fi
}

validate_rpm_release() {
  local value
  value="$1"
  if [[ ! "$value" =~ ^[0-9A-Za-z.+_~^]+$ ]]; then
    printf 'ERROR: invalid RPM Release: %s\n' "$value" >&2
    return 1
  fi
}

go_arch() {
  if [[ -n "${GOARCH:-}" ]]; then
    printf '%s\n' "$GOARCH"
    return
  fi
  go env GOARCH
}

go_os() {
  printf '%s\n' "${GOOS:-linux}"
}

deb_arch() {
  case "$(go_arch)" in
    amd64) printf 'amd64\n' ;;
    arm64) printf 'arm64\n' ;;
    386) printf 'i386\n' ;;
    arm) printf 'armhf\n' ;;
    *) printf '%s\n' "$(go_arch)" ;;
  esac
}

rpm_arch() {
  case "$(go_arch)" in
    amd64) printf 'x86_64\n' ;;
    arm64) printf 'aarch64\n' ;;
    386) printf 'i386\n' ;;
    *) printf '%s\n' "$(go_arch)" ;;
  esac
}

output_dir() {
  printf '%s\n' "${OUTPUT_DIR:-output/packages}"
}

rust_target() {
  case "$1/$2" in
    linux/amd64) printf 'x86_64-unknown-linux-gnu\n' ;;
    linux/arm64) printf 'aarch64-unknown-linux-gnu\n' ;;
    linux/386) printf 'i686-unknown-linux-gnu\n' ;;
    linux/arm) printf 'armv7-unknown-linux-gnueabihf\n' ;;
    darwin/amd64) printf 'x86_64-apple-darwin\n' ;;
    darwin/arm64) printf 'aarch64-apple-darwin\n' ;;
    windows/amd64) printf 'x86_64-pc-windows-gnu\n' ;;
    windows/arm64) printf 'aarch64-pc-windows-gnullvm\n' ;;
    *)
      printf 'ERROR: unsupported Rust engine target: %s/%s\n' "$1" "$2" >&2
      return 1
      ;;
  esac
}

build_engine_staticlib() {
  local target_os target_arch target_triple target_root archive candidate
  target_os="$1"
  target_arch="$2"
  target_triple="$(rust_target "$target_os" "$target_arch")"
  target_root="${CARGO_TARGET_DIR:-${ENGINE_DIR}/target}"

  require_command cargo
  cargo build \
    --manifest-path "${ENGINE_DIR}/Cargo.toml" \
    --package "${ENGINE_PACKAGE}" \
    --release \
    --locked \
    --target "$target_triple"

  archive=""
  for candidate in \
    "${target_root}/${target_triple}/release/libxaligo_engine.a" \
    "${target_root}/${target_triple}/release/xaligo_engine.lib"; do
    if [[ -s "$candidate" ]]; then
      archive="$candidate"
      break
    fi
  done
  if [[ -z "$archive" ]]; then
    printf 'ERROR: Rust engine archive was not generated for %s\n' "$target_triple" >&2
    return 1
  fi
  mkdir -p "$(dirname "$ENGINE_LINK_ARCHIVE")"
  cp "$archive" "$ENGINE_LINK_ARCHIVE"
}

build_native_binary() {
  local version output target_os target_arch
  version="$1"
  output="$2"
  target_os="$3"
  target_arch="$4"
  build_engine_staticlib "$target_os" "$target_arch"
  mkdir -p "$(dirname "$output")"
  GOOS="$target_os" GOARCH="$target_arch" CGO_ENABLED=1 \
    go build \
      -tags "$NATIVE_BUILD_TAGS" \
      -buildvcs=false \
      -trimpath \
      -ldflags "-X github.com/xaligo/xaligo/internal/controller.version=${version}" \
      -o "$output" \
      ./cmd
}

build_linux_binary() {
  build_native_binary "$1" "$2" "$(go_os)" "$(go_arch)"
}

require_command() {
  local command_name
  command_name="$1"
  if ! command -v "$command_name" >/dev/null 2>&1; then
    printf 'ERROR: required command not found: %s\n' "$command_name" >&2
    exit 1
  fi
}

build_wasm_exporter() {
  local build_dir
  if [[ -n "${PREBUILT_WASM:-}" ]]; then
    if [[ ! -s "$PREBUILT_WASM" ]]; then
      printf 'ERROR: prebuilt WASM exporter not found: %s\n' "$PREBUILT_WASM" >&2
      exit 1
    fi
    mkdir -p external/exporter/wasm
    install -m 0644 "$PREBUILT_WASM" external/exporter/wasm/xaligo.wasm
    return
  fi
  require_command cargo
  build_dir="$(mktemp -d)"
  mkdir -p "$build_dir/external/exporter"
  tar \
    --exclude='./target' \
    --exclude='./wasm' \
    -C external/exporter -cf - . | tar -C "$build_dir/external/exporter" -xf -
  mkdir -p "$build_dir/external/exporter/wasm"
  cargo build --manifest-path "$build_dir/external/exporter/Cargo.toml" --package xaligo-pptx-exporter --bin xaligo-exporter --target wasm32-wasip1 --release --locked
  install -m 0644 "$build_dir/external/exporter/target/wasm32-wasip1/release/xaligo-exporter.wasm" "$build_dir/external/exporter/wasm/xaligo.wasm"
  mkdir -p external/exporter/wasm
  install -m 0644 "$build_dir/external/exporter/wasm/xaligo.wasm" external/exporter/wasm/xaligo.wasm
  rm -rf "$build_dir"
  if [[ ! -s external/exporter/wasm/xaligo.wasm ]]; then
    printf 'ERROR: WASM exporter was not generated\n' >&2
    exit 1
  fi
}

install_runtime_files() {
  local destination
  destination="$1"
  mkdir -p \
    "$destination/etc/resources/aws" \
    "$destination/external/exporter/wasm"
  install -m 0644 etc/resources/aws/app.yaml "$destination/etc/resources/aws/app.yaml"
  install -m 0644 etc/resources/aws/service-catalog.csv "$destination/etc/resources/aws/service-catalog.csv"
  install -m 0644 etc/resources/aws/service-index.csv "$destination/etc/resources/aws/service-index.csv"
  cp -R etc/resources/aws/svg "$destination/etc/resources/aws/svg"
  install -m 0644 external/exporter/wasm/xaligo.wasm "$destination/external/exporter/wasm/xaligo.wasm"
}
