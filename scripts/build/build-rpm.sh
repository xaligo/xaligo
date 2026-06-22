#!/usr/bin/env bash
set -euo pipefail

source "$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)/common.sh"

usage() {
  cat <<'EOF'
Usage:
  scripts/build/build-rpm.sh

Environment:
  VERSION=1.2.3        Package and embedded CLI version. Defaults to latest git tag or internal version.
  GOARCH=amd64         Target Go architecture. Supported mappings include amd64 and arm64.
  OUTPUT_DIR=output/packages
EOF
}

if [[ "${1:-}" == "--help" || "${1:-}" == "-h" ]]; then
  usage
  exit 0
fi

require_command rpmbuild

ROOT="$(repo_root)"
cd "$ROOT"

VERSION_VALUE="$(package_version)"
ARCH="$(rpm_arch)"
OUT_DIR="$(output_dir)"
BUILD_ROOT="${OUT_DIR}/rpm/buildroot"
RPM_TOP="${OUT_DIR}/rpm/rpmbuild"
SPEC_PATH="${RPM_TOP}/SPECS/${PACKAGE_NAME}.spec"
BINARY_PATH="${BUILD_ROOT}/usr/bin/${PACKAGE_NAME}"
ABS_BINARY_PATH="${ROOT}/${BINARY_PATH}"
ABS_LICENSE_PATH="${ROOT}/${BUILD_ROOT}/usr/share/doc/${PACKAGE_NAME}/LICENSE"

rm -rf "$BUILD_ROOT" "$RPM_TOP"
mkdir -p "$BUILD_ROOT/usr/bin" "$BUILD_ROOT/usr/share/doc/${PACKAGE_NAME}" "$RPM_TOP/BUILD" "$RPM_TOP/BUILDROOT" "$RPM_TOP/RPMS" "$RPM_TOP/SOURCES" "$RPM_TOP/SPECS" "$RPM_TOP/SRPMS"

build_linux_binary "$VERSION_VALUE" "$BINARY_PATH"
chmod 0755 "$BINARY_PATH"
install -m 0644 LICENSE "$BUILD_ROOT/usr/share/doc/${PACKAGE_NAME}/LICENSE"

cat > "$SPEC_PATH" <<EOF
Name: ${PACKAGE_NAME}
Version: ${VERSION_VALUE}
Release: 1%{?dist}
Summary: ${PACKAGE_DESCRIPTION}
License: MIT
URL: ${PACKAGE_URL}

%description
xaligo renders the .xal diagram DSL to Excalidraw, SVG, PPTX, XYFlow, and Isoflow formats.

%install
mkdir -p %{buildroot}/usr/bin
mkdir -p %{buildroot}/usr/share/doc/%{name}
install -m 0755 ${ABS_BINARY_PATH} %{buildroot}/usr/bin/%{name}
install -m 0644 ${ABS_LICENSE_PATH} %{buildroot}/usr/share/doc/%{name}/LICENSE
chmod 0644 %{buildroot}/usr/share/doc/%{name}/LICENSE

%files
/usr/bin/%{name}
%doc /usr/share/doc/%{name}/LICENSE

%changelog
* Tue Jun 23 2026 ${PACKAGE_MAINTAINER} - ${VERSION_VALUE}-1
- Package xaligo CLI.
EOF

rpmbuild --target "$ARCH" --define "_topdir ${ROOT}/${RPM_TOP}" --define "_buildrootdir ${ROOT}/${RPM_TOP}/BUILDROOT" -bb "$SPEC_PATH"
mkdir -p "$OUT_DIR"
find "$RPM_TOP/RPMS" -type f -name '*.rpm' -exec cp {} "$OUT_DIR" \;
printf 'Built RPM packages into: %s\n' "$OUT_DIR"