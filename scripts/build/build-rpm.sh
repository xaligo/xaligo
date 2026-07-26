#!/usr/bin/env bash
set -euo pipefail

source "$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)/common.sh"

usage() {
  cat <<'EOF'
Usage:
  scripts/build/build-rpm.sh

Environment:
  NATIVE_VERSION=1.2.3-42        Embedded CLI version. Defaults to VERSION, then the VERSION file.
  PACKAGE_VERSION=1.2.3          RPM Version. Defaults to VERSION, then the VERSION file.
  PACKAGE_RELEASE=0.42           RPM Release. Defaults to 1.
  GOARCH=amd64                   Target Go architecture. Supported mappings include amd64 and arm64.
  OUTPUT_DIR=output/packages
EOF
}

if [[ "${1:-}" == "--help" || "${1:-}" == "-h" ]]; then
  usage
  exit 0
fi

require_command rpmbuild
require_command rpm

ROOT="$(repo_root)"
cd "$ROOT"

NATIVE_VERSION_VALUE="$(native_version)"
RPM_VERSION_VALUE="$(package_version)"
RPM_RELEASE_VALUE="$(package_release)"
ARCH="$(rpm_arch)"
OUT_DIR="$(output_dir)"
BUILD_ROOT="${OUT_DIR}/rpm/buildroot"
RPM_TOP="${OUT_DIR}/rpm/rpmbuild"
SPEC_PATH="${RPM_TOP}/SPECS/${PACKAGE_NAME}.spec"
BINARY_PATH="${BUILD_ROOT}/usr/bin/${PACKAGE_NAME}"
ABS_BINARY_PATH="${ROOT}/${BINARY_PATH}"
ABS_LICENSE_PATH="${ROOT}/${BUILD_ROOT}/usr/share/doc/${PACKAGE_NAME}/LICENSE"
ABS_THIRD_PARTY_LICENSES_PATH="${ROOT}/${BUILD_ROOT}/usr/share/doc/${PACKAGE_NAME}/THIRD_PARTY_LICENSES"
RUNTIME_PATH="${BUILD_ROOT}/${RUNTIME_REL}"
ABS_RUNTIME_PATH="${ROOT}/${RUNTIME_PATH}"

validate_native_version "$NATIVE_VERSION_VALUE"
validate_rpm_version "$RPM_VERSION_VALUE"
validate_rpm_release "$RPM_RELEASE_VALUE"
if [[ "$NATIVE_VERSION_VALUE" != "$RPM_VERSION_VALUE" && -z "${PACKAGE_RELEASE:-}" ]]; then
  printf 'ERROR: PACKAGE_RELEASE is required when native and RPM versions differ\n' >&2
  exit 1
fi

rm -rf "$BUILD_ROOT" "$RPM_TOP"
mkdir -p "$BUILD_ROOT/usr/bin" "$BUILD_ROOT/usr/share/doc/${PACKAGE_NAME}" "$RPM_TOP/BUILD" "$RPM_TOP/BUILDROOT" "$RPM_TOP/RPMS" "$RPM_TOP/SOURCES" "$RPM_TOP/SPECS" "$RPM_TOP/SRPMS"

build_linux_binary "$NATIVE_VERSION_VALUE" "$BINARY_PATH"
build_wasm_exporter
install_runtime_files "$RUNTIME_PATH"
chmod 0755 "$BINARY_PATH"
install -m 0644 LICENSE "$BUILD_ROOT/usr/share/doc/${PACKAGE_NAME}/LICENSE"
install -m 0644 THIRD_PARTY_LICENSES "$BUILD_ROOT/usr/share/doc/${PACKAGE_NAME}/THIRD_PARTY_LICENSES"

cat > "$SPEC_PATH" <<EOF
Name: ${PACKAGE_NAME}
Version: ${RPM_VERSION_VALUE}
Release: ${RPM_RELEASE_VALUE}%{?dist}
Summary: ${PACKAGE_DESCRIPTION}
License: MIT AND Zlib
URL: ${PACKAGE_URL}

%description
xaligo renders the .xal diagram DSL to Excalidraw, SVG, PPTX, XYFlow, and Isoflow formats.

%install
mkdir -p %{buildroot}/usr/bin
mkdir -p %{buildroot}/usr/share/doc/%{name}
mkdir -p %{buildroot}/usr/lib/%{name}
install -m 0755 ${ABS_BINARY_PATH} %{buildroot}/usr/bin/%{name}
install -m 0644 ${ABS_LICENSE_PATH} %{buildroot}/usr/share/doc/%{name}/LICENSE
install -m 0644 ${ABS_THIRD_PARTY_LICENSES_PATH} %{buildroot}/usr/share/doc/%{name}/THIRD_PARTY_LICENSES
cp -a ${ABS_RUNTIME_PATH}/. %{buildroot}/usr/lib/%{name}/
chmod 0644 %{buildroot}/usr/share/doc/%{name}/LICENSE
chmod 0644 %{buildroot}/usr/share/doc/%{name}/THIRD_PARTY_LICENSES

%files
/usr/bin/%{name}
/usr/lib/%{name}
%doc /usr/share/doc/%{name}/LICENSE
%doc /usr/share/doc/%{name}/THIRD_PARTY_LICENSES

%changelog
* Tue Jun 23 2026 ${PACKAGE_MAINTAINER} - ${RPM_VERSION_VALUE}-${RPM_RELEASE_VALUE}
- Package xaligo CLI.
EOF

rpmbuild --target "$ARCH" --define "_topdir ${ROOT}/${RPM_TOP}" --define "_buildrootdir ${ROOT}/${RPM_TOP}/BUILDROOT" -bb "$SPEC_PATH"
mkdir -p "$OUT_DIR"
BUILT_COUNT=0
while IFS= read -r package_path; do
  BUILT_VERSION="$(rpm -qp --queryformat '%{VERSION}' "$package_path")"
  BUILT_RELEASE="$(rpm -qp --queryformat '%{RELEASE}' "$package_path")"
  if [[ "$BUILT_VERSION" != "$RPM_VERSION_VALUE" ]]; then
    printf 'ERROR: built RPM Version mismatch: %s != %s\n' "$BUILT_VERSION" "$RPM_VERSION_VALUE" >&2
    exit 1
  fi
  if [[ "$BUILT_RELEASE" != "$RPM_RELEASE_VALUE" && "$BUILT_RELEASE" != "${RPM_RELEASE_VALUE}."* ]]; then
    printf 'ERROR: built RPM Release mismatch: %s does not start with %s\n' "$BUILT_RELEASE" "$RPM_RELEASE_VALUE" >&2
    exit 1
  fi
  cp "$package_path" "$OUT_DIR"
  BUILT_COUNT=$((BUILT_COUNT + 1))
done < <(find "$RPM_TOP/RPMS" -type f -name '*.rpm' -print)
if [[ "$BUILT_COUNT" -eq 0 ]]; then
  printf 'ERROR: rpmbuild produced no RPM package\n' >&2
  exit 1
fi
printf 'Built RPM packages into: %s\n' "$OUT_DIR"
