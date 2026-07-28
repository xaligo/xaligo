---
applyTo: ".github/instructions/manual/**"
---

# 10.14 PPTX and routing: Verification Checklist

## Verification Checklist

Before considering PPTX routing/layout changes complete:

```bash
go test ./...
make build
make build-wasm
npm run build --workspace @xaligo/xaligo-external
.bin/xaligo render docs/src/examples/samples/sample.xal --format pptx --services docs/src/examples/samples/services.csv -o out.pptx --paper A3 --orientation landscape --arrow-style thin
unzip -t out.pptx
```

For icon-overlap regressions, inspect the resolved PPTX XML and ensure routed
custom geometry does not intersect target icon/label rectangles.
