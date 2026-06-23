#!/usr/bin/env bash
set -euo pipefail

source "$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)/common.sh"

usage() {
  cat <<'EOF'
Usage:
  scripts/build/build-deb.sh

Environment:
  VERSION=1.2.3        Package and embedded CLI version. Defaults to latest git tag or internal version.
  GOARCH=amd64         Target Go architecture. Supported mappings include amd64 and arm64.
  OUTPUT_DIR=output/packages
  PACKAGE_MAINTAINER="Name <email>"
EOF
}

if [[ "${1:-}" == "--help" || "${1:-}" == "-h" ]]; then
  usage
  exit 0
fi

require_command dpkg-deb

ROOT="$(repo_root)"
cd "$ROOT"

VERSION_VALUE="$(package_version)"
ARCH="$(deb_arch)"
OUT_DIR="$(output_dir)"
WORK_DIR="${OUT_DIR}/deb/${PACKAGE_NAME}_${VERSION_VALUE}_${ARCH}"
PACKAGE_PATH="${OUT_DIR}/${PACKAGE_NAME}_${VERSION_VALUE}_${ARCH}.deb"
RUNTIME_DIR="${WORK_DIR}/${RUNTIME_REL}"

rm -rf "$WORK_DIR"
mkdir -p "$WORK_DIR/DEBIAN" "$WORK_DIR/usr/bin" "$WORK_DIR/usr/share/doc/${PACKAGE_NAME}"

build_linux_binary "$VERSION_VALUE" "$WORK_DIR/usr/bin/${PACKAGE_NAME}"
build_wasm_exporter
install_runtime_files "$RUNTIME_DIR"
chmod 0755 "$WORK_DIR/usr/bin/${PACKAGE_NAME}"
install -m 0644 LICENSE "$WORK_DIR/usr/share/doc/${PACKAGE_NAME}/copyright"

cat > "$WORK_DIR/DEBIAN/control" <<EOF
Package: ${PACKAGE_NAME}
Version: ${VERSION_VALUE}
Section: utils
Priority: optional
Architecture: ${ARCH}
Maintainer: ${PACKAGE_MAINTAINER}
Homepage: ${PACKAGE_URL}
Description: ${PACKAGE_DESCRIPTION}
 xaligo renders the .xal diagram DSL to Excalidraw, SVG, PPTX, XYFlow, and Isoflow formats.
EOF

mkdir -p "$OUT_DIR"
dpkg-deb --build "$WORK_DIR" "$PACKAGE_PATH"
printf 'Built: %s\n' "$PACKAGE_PATH"
