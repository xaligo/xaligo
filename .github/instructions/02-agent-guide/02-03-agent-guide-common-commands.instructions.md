---
applyTo: ".github/instructions/manual/**"
---

# 02.03 Agent guide: Common commands

## Common commands

```bash
# Build and test
go build ./...
go test ./...

# Build distributable adapters
make build
make build-wasm
npm ci --ignore-scripts
npm run build --workspace=@xaligo/xaligo-external

# Render and validate
.bin/xaligo validate docs/src/examples/samples/sample.xal
.bin/xaligo render docs/src/examples/samples/sample.xal --format excalidraw -o output/sample.excalidraw
.bin/xaligo render docs/src/examples/samples/sample.xal --format svg -o output/sample.svg
.bin/xaligo render docs/src/examples/samples/sample.xal --format pdf -o output/sample.pdf
.bin/xaligo render docs/src/examples/samples/sample.xal --format excel -o output/sample.xlsx
.bin/xaligo render docs/src/examples/samples/sample.xal --format xyflow -o output/sample.xyflow.json
.bin/xaligo render docs/src/examples/samples/sample.xal --format isoflow -o output/sample.isoflow.json
.bin/xaligo serve docs/src/examples/samples/sample.xal --mode network

# Clean generated artifacts
make clean
```

Native PPTX export additionally requires the configured `xaligo.wasm` PPTX exporter.
The TypeScript package consumes `BuildPPTXPlan` through WASM and creates PPTX
with PptxGenJS.
