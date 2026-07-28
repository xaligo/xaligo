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
npm run build --workspace=@xaligo/xaligo-external
npm --prefix external run build:pptx-exporter-wasm
```

Generated binaries, `node_modules`, `output`, and package `dist` directories are
ignored and must not be committed.
