---
applyTo: ".github/instructions/manual/**"
---

# 02.02 Agent guide: Working rules

## Working rules

- Preserve `.xal -> parser -> layout -> shared scene/plan -> encoder`.
- CLI, preview, and WASM format-rendering paths call `internal/usecase`. They
  do not build a parallel parser/layout/render pipeline. Focused `add` and
  source-generation utilities may use repositories/builders directly.
- Keep mode and format independent.
- Put cross-format routing and geometry in shared layers.
- Return wrapped errors; do not panic in core code.
- Preserve unrelated and pre-existing working-tree changes.
- Do not commit build output, dependencies, binaries, or caches. Checked-in
  documentation SVGs are regenerated from and committed with their `.xal`
  sources.
- Add focused tests with every behavior change.
