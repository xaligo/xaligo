---
applyTo: ".github/instructions/manual/**"
---

# 02.03 Agent guide: Common commands

## Common commands

```bash
# Build and test
go build ./...
go test ./...

# Build native adapters and test npm packaging
make build
make build-exporter
npm ci --ignore-scripts
npm run test:npm-installer

# Render and validate
.bin/xaligo validate docs/src/examples/samples/sample.xal
.bin/xaligo render docs/src/examples/samples/sample.xal --format svg -o output/sample.svg
.bin/xaligo render docs/src/examples/samples/sample.xal --format pptx -o output/sample.pptx
.bin/xaligo render markdown docs/src/examples/embedded-xal.md
.bin/xaligo serve docs/src/examples/samples/sample.xal --mode network

# Clean generated artifacts
make clean
```

Native PPTX export is compiled into the engine static library. It consumes the
resolved `BuildPPTXPlan` JSON through a C ABI and creates PPTX with the Rust
`pptx` crate; no runtime exporter artifact is required.
