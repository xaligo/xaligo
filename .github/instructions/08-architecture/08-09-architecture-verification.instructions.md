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
cargo test --manifest-path external/exporter/Cargo.toml --locked
make build-wasm
```

Generated binaries, `node_modules`, `output`, and package `dist` directories are
ignored and must not be committed.
