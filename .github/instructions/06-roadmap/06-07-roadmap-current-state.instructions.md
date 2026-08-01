---
applyTo: ".github/instructions/manual/**"
---

# 06.07 Roadmap: Current State

## Current State

The repository is already beyond a blank v0.1 baseline in several areas.

Implemented or partially implemented:

- `.xal` XML-style parser exists in `internal/usecase/v1/engine/parse_*`.
- Vuetify-like layout calculations exist in `internal/usecase/v1/engine/layout_*`.
- Canonical scene construction exists in `internal/usecase/v1/engine/scene_*`.
  Rendered graph nodes carry Box-tree-derived semantic kind and parent IDs;
  XYFlow uses geometric containment only for legacy scenes without that data.
- Native CLI exists with `render`, `generate`, `add`, `init`, and `version`.
- `render --format excalidraw` supports `services.csv` abbreviation/legend
  workflows.
- Draw-plan geometry and routing exist in `internal/usecase/v1/engine/plan_*`
  and `internal/usecase/v1/engine/route_*`.
- PPTX routing already includes obstacle avoidance, binding gap handling,
  arrow margin/lane avoidance, A3 paper options, item label sizing, and legend
  slide data.
- Repository-layer PPTX export has been redirected toward a WASM exporter
  adapter in `internal/repository/powerpoint.go`.
- `xaligo render --format excalidraw|svg|pptx|pdf|excel` is implemented;
  `xlsx` is accepted as an alias for `excel`.
- Identified child frames map to separate SVG artifacts, PPTX slides, PDF
  pages, and Excel worksheets in source order. `--combine-frames` retains the
  compatibility single-canvas form. Excalidraw, XYFlow, and Isoflow remain one
  logical document.
- Default page-local SVG uses the exact frame rectangle as its canvas and clip
  boundary, inherited by PDF and Excel page images. The combined compatibility
  canvas retains marker-safe bounds expansion.
- Page frames support a shared top/bottom metadata tag band for built-in
  `id`, `title`, content `version`, and arbitrary key/value entries. The band
  uses the resolved `row-gap` (4 pixels by default) as both its inter-row
  spacing and its metadata page-edge inset at the selected vertical edge and
  both horizontal edges. Wrapping and per-row left/center/right alignment use
  `frame width - 2 * row-gap`. Its full-width reservation strip still starts at
  the outer logical frame edge, reaches the final content-box boundary, and is
  at least
  `row-gap + complete band height + 8` pixels deep; normal items, text,
  local/UML lines and labels, and page links remain outside it. Explicit page
  sides reject normal-dimension or reservation conflicts. Automatic page links
  filter unsafe candidates, remap a preferred side with rendered visual
  geometry, and fail only when no safe side exists; side terminals are clamped
  beyond the strip. The band also supports auto/fixed widths, explicit row
  breaks, and font/color styling. It projects with its owning physical page;
  graph adapters omit it as page decoration.
- Cross-frame connections independently support item endpoint
  `src/dst-side|anchor` and logical page-terminal
  `src/dst-frame-side|anchor` geometry. Explicit frame anchors use five fixed
  tangent slots per edge, then place the drawable terminal on a parallel inward
  inset line. The inset is the resolved metadata `row-gap`, or 4 pixels when
  metadata is absent; zero `row-gap` retains the outer edge. Terminals remain
  perpendicular to the selected side, do not clamp the inset, reject unsafe
  explicit geometry, remap unsafe automatic preferences to the nearest safe
  visual side, and keep `to <...>` / `from <...>` labels 4 layout pixels from
  the final inset terminal.
- `xaligo render --format xyflow` and TypeScript/WASM `renderXYFlow()` export
  nested React Flow-compatible nodes and edges. V1 item, AWS group, rectangle,
  port, and identified child-frame endpoints are retained; cross-frame stubs
  are combined into one logical edge with routing metadata.
- `xaligo render --format isoflow`, Go `RenderIsoflow`, and TypeScript/WASM
  `renderIsoflow()` export an upstream Isoflow-compatible model from the shared
  scene. V1 non-item endpoints and logical cross-frame connectors are retained,
  and same-frame explicit bends use native tile anchors.
- `xaligo validate` reuses parser and layout validation.
- `xaligo diff` compares parsed V1 structures and emits paired SVG views: the
  old document with removed/previous values highlighted pale red and the new
  document with added/current values highlighted pale green.
- The SVG encoder is implemented in `internal/repository/svg.go` over the
  shared draw plan, including distinct V1 arrow, triangle, stealth, diamond,
  and oval marker geometry.
- Shared `light` and `dark` themes are implemented for Excalidraw, SVG, and
  PPTX via `xaligo render --theme`.
- Stable Go use cases in `internal/usecase` expose `Render`, `RenderExcalidraw`,
  `RenderSVG`, `RenderArtifacts`, `RenderPPTX`, `RenderPDF`, `RenderExcel`,
  `RenderXYFlow`, `RenderIsoflow`, and `Validate`; CLI SVG/Excalidraw/validation
  use the same pipeline.
- CLI, preview, and WASM adapters now use the same render use case. Embedded
  environments inject an `AssetSource` instead of reimplementing parser,
  layout, or scene construction.
- Isoflow exports shared group borders as view rectangles and produces stable
  icon ordering.
- Frozen V1 routes are headless across Excalidraw, SVG, PPTX, PDF, Excel,
  XYFlow, and Isoflow. Circular route connector nodes remain a future versioned
  feature.
- Node/PptxGenJS can still generate `out.pptx` as a temporary development path,
  but it is not the long-term repository-layer architecture.

Important gaps:

- `external/pptx-exporter/wasm/xaligo.wasm` is the generated PPTX exporter
  WASM artifact.
- Cross-renderer visual regression coverage is still limited.
- Numeric domains are checked before layout, but a typed normalized layout
  structure has not yet replaced repeated reads from `Node.Attrs`.
- Item-grid minimum-cell and item-offset checks now run from `Build` through the
  same solver used by scene construction. Selected cells and catalog-derived
  label measurements still need to become first-class resolved-layout data
  before mixed item/rectangle groups are fully supported.
- `Diagnose` proves parser, resolved-box, minimum item-grid, and item-offset
  invariants. Catalog-derived label measurement and final connector geometry
  must join the same geometry stage for complete validate/render agreement.
- Compatibility names that expose Excalidraw or PPTX in otherwise shared scene
  and plan APIs have aliases, but the underlying schemas must still migrate to
  format-neutral data without breaking public callers.
- Renderer capabilities are still implicit. In particular, the compatible
  Isoflow connector schema cannot carry arbitrary V1 kind, arrowhead,
  fixed-point, or original scale/grid metadata. A typed capability/projection
  contract remains necessary before adding more output formats.
