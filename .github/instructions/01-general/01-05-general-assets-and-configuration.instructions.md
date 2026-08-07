---
applyTo: ".github/instructions/manual/**"
---

# 01.05 General: Assets and configuration

## Assets and configuration

- Configuration: `etc/resources/aws/app.yaml`
- ID lookup: `etc/resources/aws/service-index.csv`
- Full catalog: `etc/resources/aws/service-catalog.csv`
- Embedded assets: `etc/resources/aws/assets.go`
- SVG assets: `etc/resources/aws/svg`

Preserve bundled license and attribution files. Generated assets must be
refreshed through the scripts declared in the root `package.json`.

The root `package-lock.json` is the canonical npm lock. The Rust exporter uses
the committed `external/exporter/Cargo.lock`. Commit each lock with dependency
changes and use locked builds.

`VERSION` and the root `package.json` contain the next stable `X.Y.Z` version.
Release metadata is resolved by `scripts/build/release-metadata.sh`. A main
prerelease run `N` uses `X.Y.Z-N` for the embedded native CLI, exactly `X.Y.Z`
from `VERSION` for npm, `X.Y.Z~N` for Debian, and RPM `Version: X.Y.Z` with
`Release: 0.N`; a stable RPM uses release `1`. Source branch names such as
`main` must not appear in a package name or package version. Keep these values
separate so native and OS prereleases sort before the corresponding stable
package and remain valid for each package manager. Publish npm only from the
stable release workflow because npm package versions are immutable.
