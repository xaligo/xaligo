#!/usr/bin/env bash
set -euo pipefail

PACKAGE_NAME="xaligo"
PACKAGE_DESCRIPTION="Diagram-as-Code CLI for rendering .xal diagrams"
PACKAGE_URL="https://github.com/xaligo/xaligo"
PACKAGE_MAINTAINER="${PACKAGE_MAINTAINER:-Ryo Arima <ryo-arima@users.noreply.github.com>}"
RUNTIME_REL="usr/lib/${PACKAGE_NAME}"

repo_root() {
  local source_dir
  source_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
  cd "${source_dir}/../.." && pwd
}

package_version() {
  if [[ -n "${VERSION:-}" ]]; then
    printf '%s\n' "${VERSION#v}"
    return
  fi
  if [[ -s VERSION ]]; then
    sed -n '1{s/^v//;p;q;}' VERSION
    return
  fi
  printf 'ERROR: VERSION file not found or empty\n' >&2
  exit 1
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

build_native_binary() {
  local version output target_os target_arch
  version="$1"
  output="$2"
  target_os="$3"
  target_arch="$4"
  mkdir -p "$(dirname "$output")"
  GOOS="$target_os" GOARCH="$target_arch" CGO_ENABLED=0 \
    go build \
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
    mkdir -p external/wasm
    install -m 0644 "$PREBUILT_WASM" external/wasm/xaligo.wasm
    return
  fi
  require_command npm
  require_command javy
  build_dir="$(mktemp -d)"
  tar \
    --exclude='./node_modules' \
    --exclude='./package-lock.json' \
    --exclude='./dist' \
    --exclude='./wasm' \
    -C external -cf - . | tar -C "$build_dir" -xf -
  mkdir -p "$build_dir/wasm"
  npm --prefix "$build_dir" install --no-audit --no-fund
  npm --prefix "$build_dir" run build:pptx-exporter-wasm
  mkdir -p external/wasm
  install -m 0644 "$build_dir/wasm/xaligo.wasm" external/wasm/xaligo.wasm
  rm -rf "$build_dir"
  if [[ ! -s external/wasm/xaligo.wasm ]]; then
    printf 'ERROR: WASM exporter was not generated\n' >&2
    exit 1
  fi
}

install_runtime_files() {
  local destination
  destination="$1"
  mkdir -p \
    "$destination/etc/resources/aws" \
    "$destination/external/wasm"
  install -m 0644 etc/resources/aws/app.yaml "$destination/etc/resources/aws/app.yaml"
  install -m 0644 etc/resources/aws/service-catalog.csv "$destination/etc/resources/aws/service-catalog.csv"
  install -m 0644 etc/resources/aws/service-index.csv "$destination/etc/resources/aws/service-index.csv"
  install -m 0644 etc/resources/aws/isoflow-icons.json "$destination/etc/resources/aws/isoflow-icons.json"
  cp -R etc/resources/aws/svg "$destination/etc/resources/aws/svg"
  install -m 0644 external/wasm/xaligo.wasm "$destination/external/wasm/xaligo.wasm"
}
