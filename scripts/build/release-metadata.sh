#!/usr/bin/env bash
set -euo pipefail

source "$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)/common.sh"

usage() {
  cat <<'EOF'
Usage:
  RELEASE_EVENT_NAME=push RELEASE_RUN_NUMBER=42 scripts/build/release-metadata.sh
  RELEASE_EVENT_NAME=workflow_dispatch scripts/build/release-metadata.sh

Environment:
  RELEASE_EVENT_NAME       push for a main prerelease, or workflow_dispatch for a stable release.
  RELEASE_RUN_NUMBER       Positive GitHub run number. Required for a main prerelease.
  RELEASE_INPUT_VERSION    Optional workflow input. When set, it must match VERSION.
  RELEASE_VERSION_FILE     VERSION file path. Defaults to the repository VERSION.
  RELEASE_PACKAGE_JSON     package.json path. Defaults to the repository package.json.
EOF
}

if [[ "${1:-}" == "--help" || "${1:-}" == "-h" ]]; then
  usage
  exit 0
fi

require_command node

ROOT="$(repo_root)"
VERSION_FILE="${RELEASE_VERSION_FILE:-${ROOT}/VERSION}"
PACKAGE_JSON="${RELEASE_PACKAGE_JSON:-${ROOT}/package.json}"
EVENT_NAME="${RELEASE_EVENT_NAME:-}"
RUN_NUMBER="${RELEASE_RUN_NUMBER:-}"
INPUT_VERSION="${RELEASE_INPUT_VERSION:-}"

if [[ ! -s "$VERSION_FILE" ]]; then
  printf 'ERROR: release VERSION file not found or empty: %s\n' "$VERSION_FILE" >&2
  exit 1
fi
if [[ ! -s "$PACKAGE_JSON" ]]; then
  printf 'ERROR: release package.json not found or empty: %s\n' "$PACKAGE_JSON" >&2
  exit 1
fi

SOURCE_VERSION="$(sed -n '1{s/^v//;p;q;}' "$VERSION_FILE")"
MANIFEST_VERSION="$(
  node -e 'const fs = require("node:fs"); const value = JSON.parse(fs.readFileSync(process.argv[1], "utf8")).version; if (typeof value !== "string") process.exit(1); process.stdout.write(value);' "$PACKAGE_JSON"
)"

if [[ ! "$SOURCE_VERSION" =~ ^(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)$ ]]; then
  printf 'ERROR: invalid VERSION value: %s. Must be X.Y.Z\n' "$SOURCE_VERSION" >&2
  exit 1
fi
if [[ "$MANIFEST_VERSION" != "$SOURCE_VERSION" ]]; then
  printf 'ERROR: package.json version must match VERSION: %s != %s\n' "$MANIFEST_VERSION" "$SOURCE_VERSION" >&2
  exit 1
fi
if [[ -n "$INPUT_VERSION" && "$INPUT_VERSION" != "$SOURCE_VERSION" ]]; then
  printf 'ERROR: requested version %s does not match VERSION (%s)\n' "$INPUT_VERSION" "$SOURCE_VERSION" >&2
  exit 1
fi

BASE_VERSION="$SOURCE_VERSION"
case "$EVENT_NAME" in
  push)
    if [[ ! "$RUN_NUMBER" =~ ^[1-9][0-9]*$ ]]; then
      printf 'ERROR: main prerelease requires a positive RELEASE_RUN_NUMBER: %s\n' "$RUN_NUMBER" >&2
      exit 1
    fi
    NATIVE_VERSION_VALUE="${BASE_VERSION}-main.${RUN_NUMBER}"
    NPM_VERSION="${NATIVE_VERSION_VALUE}"
    DEB_VERSION="${BASE_VERSION}~main.${RUN_NUMBER}"
    RPM_VERSION="${BASE_VERSION}"
    RPM_RELEASE="0.main.${RUN_NUMBER}"
    RELEASE_TAG="main-${RUN_NUMBER}"
    RELEASE_TITLE="xaligo ${BASE_VERSION} main ${RUN_NUMBER}"
    PRERELEASE="true"
    NPM_TAG="next"
    ;;
  workflow_dispatch)
    if [[ -n "$RUN_NUMBER" && ! "$RUN_NUMBER" =~ ^[1-9][0-9]*$ ]]; then
      printf 'ERROR: RELEASE_RUN_NUMBER must be positive when set: %s\n' "$RUN_NUMBER" >&2
      exit 1
    fi
    NATIVE_VERSION_VALUE="$BASE_VERSION"
    NPM_VERSION="$BASE_VERSION"
    DEB_VERSION="$BASE_VERSION"
    RPM_VERSION="$BASE_VERSION"
    RPM_RELEASE="1"
    RELEASE_TAG="v${BASE_VERSION}"
    RELEASE_TITLE="xaligo v${BASE_VERSION}"
    PRERELEASE="false"
    NPM_TAG="latest"
    ;;
  *)
    printf 'ERROR: unsupported RELEASE_EVENT_NAME: %s\n' "$EVENT_NAME" >&2
    exit 1
    ;;
esac

validate_native_version "$NATIVE_VERSION_VALUE"
validate_native_version "$NPM_VERSION"
validate_deb_version "$DEB_VERSION"
validate_rpm_version "$RPM_VERSION"
validate_rpm_release "$RPM_RELEASE"

printf 'base_version=%s\n' "$BASE_VERSION"
printf 'native_version=%s\n' "$NATIVE_VERSION_VALUE"
printf 'npm_version=%s\n' "$NPM_VERSION"
printf 'deb_version=%s\n' "$DEB_VERSION"
printf 'rpm_version=%s\n' "$RPM_VERSION"
printf 'rpm_release=%s\n' "$RPM_RELEASE"
printf 'release_tag=%s\n' "$RELEASE_TAG"
printf 'release_title=%s\n' "$RELEASE_TITLE"
printf 'prerelease=%s\n' "$PRERELEASE"
printf 'npm_tag=%s\n' "$NPM_TAG"
