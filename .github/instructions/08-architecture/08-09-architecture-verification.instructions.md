---
applyTo: ".github/instructions/manual/**"
---

# 08.09 Architecture: Verification

## Verification

Run after structural changes:

```bash
go test ./...
go build ./...
npm ci --ignore-scripts
cargo test --manifest-path test/unit/external/engine/Cargo.toml --locked
cargo test --manifest-path test/unit/external/exporter/Cargo.toml --locked
make build-exporter
```

Generated binaries, `node_modules`, `output`, and package `dist` directories are
ignored and must not be committed.
