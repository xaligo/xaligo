#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../../../.." && pwd)"
RESOLVER="${ROOT}/scripts/build/release-metadata.sh"
COMMON="${ROOT}/scripts/build/common.sh"
FIXTURE_DIR="$(mktemp -d)"
trap 'rm -rf "$FIXTURE_DIR"' EXIT

VERSION_FILE="${FIXTURE_DIR}/VERSION"
PACKAGE_JSON="${FIXTURE_DIR}/package.json"

write_fixture() {
  local source_version manifest_version
  source_version="$1"
  manifest_version="$2"
  printf '%s\n' "$source_version" > "$VERSION_FILE"
  printf '{"version":"%s"}\n' "$manifest_version" > "$PACKAGE_JSON"
}

resolve() {
  RELEASE_VERSION_FILE="$VERSION_FILE" \
    RELEASE_PACKAGE_JSON="$PACKAGE_JSON" \
    RELEASE_EVENT_NAME="$1" \
    RELEASE_RUN_NUMBER="${2:-}" \
    RELEASE_INPUT_VERSION="${3:-}" \
    "$RESOLVER"
}

assert_output() {
  local output key expected actual
  output="$1"
  key="$2"
  expected="$3"
  actual="$(sed -n "s/^${key}=//p" <<< "$output")"
  if [[ "$actual" != "$expected" ]]; then
    printf 'FAIL: %s: expected %s, got %s\n' "$key" "$expected" "$actual" >&2
    exit 1
  fi
}

source "$COMMON"
native_value="$(NATIVE_VERSION=1.2.3-main.42 VERSION=1.2.3 native_version)"
package_value="$(PACKAGE_VERSION=1.2.3~main.42 VERSION=1.2.3 package_version)"
package_fallback_value="$(NATIVE_VERSION=1.2.3 VERSION= PACKAGE_VERSION= package_version)"
release_value="$(PACKAGE_RELEASE=0.main.42 package_release)"
if [[ "$native_value" != "1.2.3-main.42" ]]; then
  printf 'FAIL: native version did not prefer NATIVE_VERSION: %s\n' "$native_value" >&2
  exit 1
fi
if [[ "$package_value" != "1.2.3~main.42" ]]; then
  printf 'FAIL: package version did not prefer PACKAGE_VERSION: %s\n' "$package_value" >&2
  exit 1
fi
if [[ "$package_fallback_value" != "1.2.3" ]]; then
  printf 'FAIL: package version did not fall back to NATIVE_VERSION: %s\n' "$package_fallback_value" >&2
  exit 1
fi
if [[ "$release_value" != "0.main.42" ]]; then
  printf 'FAIL: RPM release did not preserve PACKAGE_RELEASE: %s\n' "$release_value" >&2
  exit 1
fi

write_fixture "1.2.3" "1.2.3"
stable_output="$(resolve workflow_dispatch 52 1.2.3)"
assert_output "$stable_output" base_version "1.2.3"
assert_output "$stable_output" native_version "1.2.3"
assert_output "$stable_output" npm_version "1.2.3"
assert_output "$stable_output" deb_version "1.2.3"
assert_output "$stable_output" rpm_version "1.2.3"
assert_output "$stable_output" rpm_release "1"
assert_output "$stable_output" release_tag "v1.2.3"
assert_output "$stable_output" prerelease "false"
assert_output "$stable_output" npm_tag "latest"

main_output="$(resolve push 42)"
assert_output "$main_output" base_version "1.2.3"
assert_output "$main_output" native_version "1.2.3-main.42"
assert_output "$main_output" npm_version "1.2.3-main.42"
assert_output "$main_output" deb_version "1.2.3~main.42"
assert_output "$main_output" rpm_version "1.2.3"
assert_output "$main_output" rpm_release "0.main.42"
assert_output "$main_output" release_tag "main-42"
assert_output "$main_output" prerelease "true"
assert_output "$main_output" npm_tag "next"

if resolve push 0 >/dev/null 2>&1; then
  printf 'FAIL: zero main run number was accepted\n' >&2
  exit 1
fi
if resolve workflow_dispatch 52 1.2.4 >/dev/null 2>&1; then
  printf 'FAIL: mismatched requested version was accepted\n' >&2
  exit 1
fi

write_fixture "1.2.3" "1.2.4"
if resolve workflow_dispatch 52 >/dev/null 2>&1; then
  printf 'FAIL: mismatched package.json version was accepted\n' >&2
  exit 1
fi

write_fixture "01.2.3" "01.2.3"
if resolve workflow_dispatch 52 >/dev/null 2>&1; then
  printf 'FAIL: non-SemVer VERSION with a leading zero was accepted\n' >&2
  exit 1
fi

printf 'release metadata tests passed\n'
