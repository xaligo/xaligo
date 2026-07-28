---
applyTo: ".github/instructions/manual/**"
---

# 06.01.02 Roadmap: Common DSL and Go Core

### Common DSL and Go Core

- Keep `.xal` as the single source DSL for every visual mode and export format.
- Keep Go as the core parser, validation, layout, routing, and rendering engine.
- VS Code, browser preview, and exporters must consume public core APIs instead
  of reimplementing parsing or layout.
- Preserve the pipeline boundary:

```text
.xal -> parser -> layout/shared model -> mode renderer -> format encoder
```
