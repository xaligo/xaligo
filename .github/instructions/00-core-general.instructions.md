---
applyTo: "**"
---

# General

Read `README.md` here and matching instructions before acting; re-check when
scope changes. Preserve unrelated worktree changes.

Pipeline: `.xal -> parser -> layout -> shared scene/plan -> encoder`.
Orchestration belongs in `internal/usecase`, synchronous V1 calculations in
`internal/usecase/v1/engine`, and I/O/encoders in `internal/repository`.
Keep mode separate from format and geometry/routing renderer-neutral.

Return context-wrapped errors; core code must not panic. Add focused regression
tests. Do not commit generated binaries, dependencies, caches, `output`, WASM,
or TypeScript `dist`; tracked documentation SVGs are allowed with their `.xal`.

Full project map, assets, release rules, commands, and APIs: `reference.md`
sections `01–02`.
