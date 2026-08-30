---
applyTo: ".github/instructions/manual/**"
---

# 06.17.01 Roadmap: PPTX Export

### PPTX Export

```bash
xaligo render --format pptx
```

Native export contract:

- Keep `xaligo render --format pptx` on the statically linked Rust exporter C
  ABI; no runtime WASM artifact is required.
- Do not introduce repository-layer Node or WASM subprocess execution.
- Do not implement PPTX/OOXML writing in Go controller/repository code.
- Keep route/traffic/theme support renderer-agnostic where possible.
