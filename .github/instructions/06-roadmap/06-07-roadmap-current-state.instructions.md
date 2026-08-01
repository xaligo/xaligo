---
applyTo: ".github/instructions/manual/**"
---

# 06.07 Roadmap: Current State

## Current State

Implemented or partially implemented:

- `.xal` parsing and Vuetify-like layout calculations live in
  `internal/usecase/v1/engine/parse_*` and `layout_*`.
- The current V1 compatibility path still creates an internal
  Excalidraw-shaped presentation scene before producing the shared draw plan.
  That scene is not a supported output and has no CLI or WASM format surface.
- Draw-plan geometry, text resolution, page projection, and connector routing
  live in `internal/usecase/v1/engine/plan_*` and `route_*`.
- `xaligo render --format svg|pptx` is the complete public render matrix.
  Omitting `--format` defaults to SVG. Retired format names are rejected.
- Identified child frames map to separate SVG artifacts or PPTX slides in
  source order. `--combine-frames` requests one combined canvas or slide.
- `xaligo render markdown` renders every fenced `xal` block through the SVG
  artifact pipeline and embeds image references in a generated Markdown file.
- `.xal` and Markdown live preview are SVG-based.
- `xaligo validate` reuses parser and layout validation. `xaligo diff` compares
  parsed V1 structures and emits paired SVG views.
- PPTX uses a shared Go document plan and the WASM/PptxGenJS exporter under
  `external/pptx-exporter`; repository code receives PPTX bytes and the CLI
  owns file persistence.
- Native CLI dependency composition is in `internal.NewRootCmd`. The browser
  WASM surface exposes SVG rendering, PPTX-plan construction, and diagnostics.
- Excalidraw persistence/editing, PDF, Excel/XLSX, XYFlow, and Isoflow output
  implementations, dependencies, generated assets, commands, aliases, and
  browser globals have been removed.
- The Rust V2 engine workspace and C ABI/static-library bridge are initialized
  under `external/engine`; generic V2 calculation migration remains underway.
- The embedded SQLite SVG registry and builtin generic icon profile are
  available to the V2 use cases.

Important gaps:

- Native V2 and frozen V1-compatibility frontends do not yet both lower
  directly to the final typed neutral model.
- The V1 SVG/PPTX path still depends on a renderer-shaped compatibility scene;
  it must migrate to an immutable renderer-neutral resolved document.
- Numeric domains are checked before layout, but a typed normalized parameter
  structure has not yet replaced all repeated reads from `Node.Attrs`.
- Catalog-derived label measurement and final connector geometry still need to
  join the same diagnostic geometry stage for complete validate/render parity.
- Cross-encoder visual regression coverage for SVG, PPTX, and Markdown remains
  limited.
