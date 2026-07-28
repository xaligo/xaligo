---
applyTo: ".github/instructions/manual/**"
---

# 06.17.01 Roadmap: PPTX Export

### PPTX Export

```bash
xaligo render --format pptx
```

Compatibility during transition:

- Keep `xaligo render --format pptx` usable when a WASM exporter is available.
- Do not reintroduce repository-layer Node subprocess execution as the default.
- Do not implement PPTX/OOXML writing in Go controller/repository code.
- Keep route/traffic/theme support renderer-agnostic where possible.
