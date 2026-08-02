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
- The Rust V2 engine under `external/engine` implements C ABI v2, parent-indexed
  typed concept transfer, nested fixed/flex/grid/absolute layout, ports,
  generic straight/orthogonal routing, safe SVG normalization, and a
  deterministic generic SVG projection. The three-method Go `EngineUsecase`
  boundary remains source-compatible with the original flat prototype.
- The generic V1/V2 frontend retains source spans while lowering common
  concepts to `EngineDocumentSpec`. Native `version="2"` rendering resolves
  through Rust and projects the same `DocumentPlan` to SVG or the existing
  PPTX exporter. Full V1 profile parity remains on the frozen V1 render path.
- ABI field indexes are generated for Go and Rust from one CSV schema. Native
  calls use a C-owned atomic cancellation handle, and engine failures surface
  through a structured Go diagnostic contract.
- The embedded SQLite SVG registry and builtin generic icon profile are
  available to the V2 use cases.
- `xaligo rag index|search|watch` provides the durable project FTS5 boundary.
  Its initial corpus is deliberately restricted to Markdown below `docs/`;
  `.xal` concepts are analyzed only when an editor or agent explicitly supplies
  a document.
- `xaligo lsp` implements an LSP 3.18 stdio session over the same project
  analysis boundary. Open-document diagnostics and symbols stay in memory;
  saving an explicitly opened `.xal` document may update its concept rows
  without changing the docs-only initial RAG corpus.
- `xaligo mcp` implements the stateless MCP 2026-07-28 contract over
  newline-delimited stdio and localhost-only Streamable HTTP POST. Its tools
  compose the existing diagnostics, render, project, and icon use cases;
  `index_docs` retains the same docs-Markdown-only discovery boundary.

Important gaps:

- Full builtin/AWS/UML profile normalization has not yet replaced the current
  V1 source-tag path; the shared frontend currently covers generic concepts.
- The V1 SVG/PPTX path still depends on a renderer-shaped compatibility scene;
  it must migrate to an immutable renderer-neutral resolved document.
- Numeric domains are checked before layout, but a typed normalized parameter
  structure has not yet replaced all repeated reads from `Node.Attrs`.
- Catalog-derived label measurement and final connector geometry still need to
  join the same diagnostic geometry stage for complete validate/render parity.
- Cross-encoder visual regression coverage for SVG, PPTX, and Markdown remains
  limited.
