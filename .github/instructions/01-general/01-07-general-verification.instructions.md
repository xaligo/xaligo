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
go test ./...
go build ./...
npm ci --ignore-scripts
npm run build --workspace=@xaligo/xaligo-external
npm --prefix external/pptx-exporter run build:pptx-exporter-wasm
git diff --check
```
