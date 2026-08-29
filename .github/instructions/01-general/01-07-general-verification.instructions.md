---
applyTo: ".github/instructions/manual/**"
---

# 01.07 General: Verification

## Verification

Set up the repository-pinned security scanner and npm audit metadata once:

```bash
make security-setup
```

Run the security gate before every commit, followed by the relevant tests and
builds:

```bash
make security-check
make test-engine
go test ./...
go build ./...
npm ci --ignore-scripts
cargo test --manifest-path test/unit/external/engine/Cargo.toml --locked
cargo test --manifest-path test/unit/external/exporter/Cargo.toml --locked
make build-exporter
git diff --check
```
