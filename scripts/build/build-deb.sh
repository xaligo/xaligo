#!/usr/bin/env bash
set -euo pipefail

source "$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)/common.sh"

usage() {
  cat <<'EOF'
Usage:
  scripts/build/build-deb.sh

Environment:
  NATIVE_VERSION=1.2.3-42         Embedded CLI version. Defaults to VERSION, then the VERSION file.
  PACKAGE_VERSION=1.2.3~42        Debian package version. Defaults to VERSION, then the VERSION file.
  GOARCH=amd64                    Target Go architecture. Supported mappings include amd64 and arm64.
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

NATIVE_VERSION_VALUE="$(native_version)"
DEB_VERSION_VALUE="$(package_version)"
ARCH="$(deb_arch)"
OUT_DIR="$(output_dir)"
WORK_DIR="${OUT_DIR}/deb/${PACKAGE_NAME}_${DEB_VERSION_VALUE}_${ARCH}"
PACKAGE_PATH="${OUT_DIR}/${PACKAGE_NAME}_${DEB_VERSION_VALUE}_${ARCH}.deb"
RUNTIME_DIR="${WORK_DIR}/${RUNTIME_REL}"

validate_native_version "$NATIVE_VERSION_VALUE"
validate_deb_version "$DEB_VERSION_VALUE"

rm -rf "$WORK_DIR"
mkdir -p "$WORK_DIR/DEBIAN" "$WORK_DIR/usr/bin" "$WORK_DIR/usr/share/doc/${PACKAGE_NAME}"

build_linux_binary "$NATIVE_VERSION_VALUE" "$WORK_DIR/usr/bin/${PACKAGE_NAME}"
install_runtime_files "$RUNTIME_DIR"
chmod 0755 "$WORK_DIR/usr/bin/${PACKAGE_NAME}"
cat LICENSE THIRD_PARTY_LICENSES > "$WORK_DIR/usr/share/doc/${PACKAGE_NAME}/copyright"
install -m 0644 THIRD_PARTY_LICENSES "$WORK_DIR/usr/share/doc/${PACKAGE_NAME}/THIRD_PARTY_LICENSES"

cat > "$WORK_DIR/DEBIAN/control" <<EOF
Package: ${PACKAGE_NAME}
Version: ${DEB_VERSION_VALUE}
Section: utils
Priority: optional
Architecture: ${ARCH}
Maintainer: ${PACKAGE_MAINTAINER}
Homepage: ${PACKAGE_URL}
Description: ${PACKAGE_DESCRIPTION}
 xaligo renders the .xal diagram DSL to SVG and PPTX, with SVG embedding for Markdown.
EOF

mkdir -p "$OUT_DIR"
dpkg-deb --build "$WORK_DIR" "$PACKAGE_PATH"
BUILT_VERSION="$(dpkg-deb --field "$PACKAGE_PATH" Version)"
if [[ "$BUILT_VERSION" != "$DEB_VERSION_VALUE" ]]; then
  printf 'ERROR: built Debian version mismatch: %s != %s\n' "$BUILT_VERSION" "$DEB_VERSION_VALUE" >&2
  exit 1
fi
printf 'Built: %s\n' "$PACKAGE_PATH"
