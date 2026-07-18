---
applyTo: "**/*.{go,ts,xal,md}"
---

# xaligo — Roadmap Preconditions

This roadmap is a planning precondition for future implementation work. Prefer
changes that move xaligo toward a diagram-as-code platform with a clean render
pipeline, SVG-first preview capability, network-diagram primitives, and
eventual VS Code / PPTX integration.

Implementation guidance:

- Keep the core pipeline separable as `.xal -> parser -> layout -> renderer`.
- Treat Excalidraw, SVG, and PPTX as output renderers over a shared model where
  possible.
- Prioritize SVG renderer and network diagram primitives before advanced PPTX
  feature polish when the choice is otherwise ambiguous.
- Route/traffic separation, route connectors, orthogonal routing, edge offsets,
  layer routing, junctions, and line jumps are roadmap features, not one-off
  export hacks.
- Live preview and VS Code integration should build on `xaligo render` /
  `xaligo validate`, not separate hidden pipelines.

## Product Architecture Preconditions

### V1 Structured-Diagram Profile

The table, relational-database, and UML design extends the V1 engine. Canonical
V1 uses `<xaligo version="1">`; historical root `<frame>` and `<frames>`
documents remain compatible but emit a migration warning. This does not replace
the reserved V2 `<scene version="2">` contract.

The target document shape uses `<xaligo>` as a document envelope, a document-
wide `<data>` registry, and `<frames>` containing identified `<frame>`
components. Data definitions are reusable across frames. General tables, RDB
schemas, and UML diagrams have separate semantic frontends and layout engines;
they may share renderer-neutral primitives and output encoders but must not be
forced through one diagram-specific processor.

Keep these semantic distinctions:

- `<table>` is general tabular data, `<database>`/`<entity>` is relational
  schema meaning, and `<grid>` is visual layout.
- Pipe and explicit tag syntax lower to the same typed model for a given
  component; imported files enter that model through an import adapter.
- `<uml>` is the common UML component. Exactly one diagram-kind child such as
  `<class-diagram>` or `<sequence-diagram>` selects its processor; the frame
  does not carry a UML kind.
- Imports are resolved before semantic validation, retain provenance, and do
  not execute arbitrary commands. Inline/tag overrides are explicit and
  deterministic.

The user-facing design is documented in
`docs/src/design/structured-diagrams.md`.

### Common DSL and Go Core

- Keep `.xal` as the single source DSL for every visual mode and export format.
- Keep Go as the core parser, validation, layout, routing, and rendering engine.
- VS Code, browser preview, and exporters must consume public core APIs instead
  of reimplementing parsing or layout.
- Preserve the pipeline boundary:

```text
.xal -> parser -> layout/shared model -> mode renderer -> format encoder
```

### V1 Compatibility and V2 Input

Keep root `<frame>` and `<frames>` as legacy V1 compatibility inputs. Canonical
V1 uses `<xaligo version="1">` with a document-wide `<data>` registry and
identified frames. Legacy roots emit a migration warning. V2 uses a distinct
`<scene version="2">` root; do not place `version="2"` on a V1 root. This is a
reject-safe boundary: existing V1 readers reject V2 without having to know any
V2 syntax.

V2 must render both native V2 documents and the frozen V1 profile. Implement
that compatibility in the V2 side only: a V1 compatibility frontend and the
native V2 frontend each lower directly to the same typed, version-neutral
model. Keep the existing V1 engine independent of V2.

The compatibility path is complete only when it preserves V1 defaults,
fallback/error behavior, unknown nested-tag handling, connection-group
inheritance, anchor aliases, numeric catalog-ID range, and render-context item
size. Golden tests must compare V1 and V2-engine output at the neutral-model and
resolved-geometry boundaries across native and embedded targets.

Do not implement compatibility by changing root tags as strings, reparsing,
retrying parsers after syntax errors, serializing through the V1 scene, or
calling the full V1 renderer before V2. Root dispatch reads the first start
element once and selects exactly one frontend; renderers and encoders remain
shared downstream.

### Mode and Format Are Independent

`mode` selects visual and layout semantics. `format` selects serialization or
the target integration. Do not encode a visual mode as a file format or assume
that one format has only one mode.

Target modes:

| Mode | Visual/layout intent |
|---|---|
| `standard` | Normal two-dimensional architecture diagrams |
| `network` | Route, traffic, circular connector, and topology-oriented diagrams |
| `aws` | AWS official-icon-oriented architecture diagrams |
| `aws-2.5d` | Cloudcraft/legacy AWS-reference-style oblique diagrams |
| `topology` | Instana/SkyWalking-style dependency topology |

Target formats:

| Format | Primary use |
|---|---|
| `svg` | Portable output and live preview |
| `excalidraw` | Editable Excalidraw scene |
| `pptx` | Editable presentation export |
| `xyflow` | React Flow/GUI editor integration |
| `isoflow` | Isometric/2.5D integration |

Target CLI shape:

```bash
xaligo render input.xal --mode network --format svg -o output.svg
xaligo render input.xal --mode aws-2.5d --format pptx -o output.pptx
```

Backward compatibility: omitting `--mode` must retain the current standard/AWS
behavior until an explicit default-mode migration is released.

Current V1 status: `standard`, `network`, and `aws` are accepted but have no
semantic difference; they execute the same resolved 2D pipeline. Treat them as
compatibility inputs until a versioned implementation introduces distinct
mode semantics. `aws-2.5d` and `topology` remain recognized but return a
not-implemented error.

### Shared Rendering APIs

The shared in-repository use-case boundary should support at least:

```go
RenderSVG()
RenderExcalidraw()
RenderPPTX()
RenderXYFlow()
RenderIsoflow()
```

Prefer a shared extensible API underneath the convenience functions:

```go
Render(ctx, input, RenderOptions{Mode: mode, Format: format})
Validate(ctx, input)
```

### Rendering Correctness Gate

New renderer features are gated by a shared geometry and text contract. Fixes
must be made at the earliest shared stage that owns the information, not as
format-specific clipping or coordinate adjustments.

The required order is:

1. Parse numeric layout attributes into finite, typed values and validate their
   domains with source positions.
2. Make validation and rendering execute the same geometry invariants.
3. Resolve fixed-size children before flexible weights, then record content
   boxes and explicit overflow state in the resolved layout.
4. Move item-grid selection and occupancy into resolved layout so items and
   other children cannot unknowingly occupy the same region; scene construction
   only emits the already resolved cells.
5. Carry renderer-neutral text layout, semantic role, and glyph-overflow policy
   through the draw plan.
6. Apply the same output transform to geometry and typography at every PPI and
   paper-fit setting.
7. Consolidate format dispatch in one use case and migrate the shared scene and
   plan to format-neutral names and schemas. Compatibility aliases may preserve
   public APIs, but the canonical schema must not remain Excalidraw- or
   PPTX-shaped.

Completion requires regression coverage for validation/render agreement,
finite resolved coordinates, parent/content containment, fixed-plus-flex
siblings, mixed item/rectangle groups, item offsets, connector numeric values,
empty numeric attributes, long labels across output formats including editable
Excalidraw metadata, overlapping ports, and non-96 PPI.

## Delivery Phases

### Phase 1: Basic Output

Status: complete.

- Stabilize `xaligo render` and `xaligo validate`.
- Complete the SVG renderer as the primary preview surface.
- Add shared Light and Dark themes.
- Extract stable renderer-facing shared use cases.

### Phase 2: Network Diagram Features

Status: headless V1 routes, the remaining routing steps, and textual connection
shorthands have initial shared implementations. Explicit circular connector
nodes remain future versioned work. Continue with hardening and cross-renderer
visual regression coverage.

Implement shared model/routing concepts in this order where dependencies allow:

1. Headless V1 route connectors; add explicit circular connector nodes only in
   a future versioned model.
2. Orthogonal Routing.
3. Route/Traffic separation.
4. Edge Offset.
5. Line Jump.
6. Layer Routing.
7. Junction generation.

These features must be shared across renderers where possible, rather than
implemented as PPTX-only corrections.

### Phase 3: Live Preview

Status: initial implementation complete. `xaligo serve` polls `.xal` sources,
renders through the public SVG API, reports source-positioned diagnostics, and
publishes SSE reload events. Browser polish remains.

- Add `xaligo serve` on top of public render/validate APIs. (implemented)
- Watch `.xal` files and automatically re-render. (implemented)
- Serve an SVG-first browser preview with incremental refresh. (implemented)
- Keep the protocol reusable by the VS Code extension.

## VS Code Extension Preconditions

The VS Code extension is developed in a separate repository. This repository
owns the reusable Go/WASM APIs and HTTP/SSE preview protocol only; do not add
extension packaging or VS Code-specific parser/rendering forks here.

The extension target includes:

- `.xal` syntax highlighting.
- Validation and source-positioned diagnostics. (Go and TypeScript/WASM APIs implemented)
- Live Preview and a Preview Panel.
- SVG preview first; Excalidraw, XYFlow, Isoflow, and 2.5D views later.

The extension must call the same validation/render pipeline as the CLI. Do not
create an extension-only parser, layout engine, or hidden preview format.

## AWS 2.5D Mode

`mode: aws-2.5d` targets Cloudcraft and legacy AWS-reference-style oblique
architecture diagrams. It is a visual mode, not a standalone file format.

Required concepts:

- `plane` / `zone` layout primitives.
- Isometric-style nodes and routing.
- AWS node presets including `route53`, `cloudfront`, `elb`, `ec2`, `rds`, and
  `s3`.
- AWS Legacy / Cloudcraft-like themes.

Implement the first version in the native SVG renderer. WebView or GUI work may
learn from compatible 2.5D OSS projects, but the core representation must remain
usable without a specific UI framework.

## Export Roadmap

Primary formats remain SVG, Excalidraw, and PPTX. Add:

- XYFlow export for React Flow-style GUI editors. (initial implementation complete)
- Isoflow export for isometric and 2.5D integrations. (initial upstream model export complete)

Both exports should consume the shared resolved model; they must not become
alternative parsers for `.xal`.

## Long-Term Product Position

Position xaligo between PlantUML, Excalidraw, draw.io, Cloudcraft, and
Instana-style topology tools:

- Diagram as Code.
- Strong AWS and network diagram support.
- 2D, 2.5D, and topology views from one DSL.
- Comfortable VS Code authoring.
- SVG, PPTX, Excalidraw, XYFlow, and Isoflow output.

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
- `xaligo render --format excalidraw|svg|pptx` is implemented.
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
- Stable Go use cases in `internal/usecase` expose `Render`, `RenderExcalidraw`, `RenderSVG`, `RenderPPTX`,
  `RenderXYFlow`, `RenderIsoflow`, and `Validate`; CLI SVG/Excalidraw/validation
  use the same pipeline.
- CLI, preview, and WASM adapters now use the same render use case. Embedded
  environments inject an `AssetSource` instead of reimplementing parser,
  layout, or scene construction.
- Isoflow exports shared group borders as view rectangles and produces stable
  icon ordering.
- Frozen V1 routes are headless across Excalidraw, SVG, PPTX, XYFlow, and
  Isoflow. Circular route connector nodes remain a future versioned feature.
- Node/PptxGenJS can still generate `out.pptx` as a temporary development path,
  but it is not the long-term repository-layer architecture.

Important gaps:

- `external/wasm/xaligo.wasm` is the PPTX exporter WASM artifact.
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

## Rebaselined Implementation Order

Use this order when starting new roadmap work from the current repository state:

1. Complete the shared geometry/text correctness gate and its cross-renderer
   regression tests.
2. Move mixed item-grid occupancy into resolved layout and finish neutral
   scene/plan naming.
3. Complete the repository-layer WASM PPTX exporter contract by providing
   `xaligo.wasm`; keep Go free of PPTX/OOXML writer code.
4. Harden shared network routing with cross-renderer visual regression tests.
5. Build the VS Code preview on the reusable HTTP/SSE protocol exposed by
   `xaligo serve`.

## v0.1 Foundation

Status: complete. CLI and shared use cases share parser/layout/render paths for
validation, Excalidraw, and SVG.

### Rendering Engine Refactoring

Separate the core pipeline into:

```text
.xal
 ↓
parser
 ↓
layout
 ↓
renderer
```

### Public API

```go
RenderExcalidraw()
RenderSVG()
RenderPPTX()
RenderXYFlow()
RenderIsoflow()
```

Current target API shape:

```go
Render(ctx, input, RenderOptions{Mode: mode, Format: format}) ([]byte, error)
Validate(ctx, input) error
```

### CLI

```bash
xaligo render
xaligo validate
```

Required compatibility:

- Keep existing `xaligo render <input.xal> -o <out.excalidraw>` working.
- Add `xaligo render <input.xal> --format excalidraw|svg|pptx`.
- Keep format conversion under `xaligo render --format ...`; `generate` should
  remain focused on source `.xal` generation.
- `validate` must reuse parser/layout validation rather than duplicate parsing.

---

## v0.2 SVG Renderer

Status: initial renderer, route/traffic primitives, and shared Light/Dark themes
are implemented.

### SVG Export

```bash
xaligo render input.xal --format svg
```

### Supported Elements

- Node
- Group
- Label
- Route
- Traffic

### Themes

- Light
- Dark

---

## v0.3 Network Diagram Features

Status: route/traffic kinds, headless V1 routes, styling, layer order,
basic lane separation, automatic route junctions, and textual connection
shorthands are implemented across Excalidraw, SVG, and PPTX.

### Route Connector

Frozen V1 route lines have no arrowheads. A future version may add explicit
renderer-neutral circular connector nodes without changing V1 arrowhead
semantics.

```text
o------o
```

### Connector Model

```go
type Connector struct {}
```

### Orthogonal Routing

Support right-angle routing.

```text
+----+
|    |
+----+
```

### Route / Traffic Separation

#### Route

```text
o------o
```

#### Traffic

```text
======>
```

### DSL

```text
web --- db
web ==> db
```

Status: implemented using `<item name="...">`, `<item ref="...">`, or numeric
item IDs. Shorthands expand into the shared connection model during parsing.

---

## v0.4 Advanced Routing

Status: initial shared implementations are complete for edge offsets, routing
layers, frame-border clearance, and automatic fan-out/fan-in junctions.

### Edge Offset

Automatically separate overlapping routes.

```text
------
======
```

### Layer Routing

Separate routing layers.

```text
Route Layer
Traffic Layer
```

### Junction Generation

```text
      +-- DB
o-----+
      +-- Cache
```

---

## v0.5 Line Jumps

Status: rectangular background-mask jumps are implemented in the shared draw
plan for SVG/PPTX. Curved bridge arcs and an Excalidraw approximation remain.

### Bridge / Jump Lines

```text
----^----
---------
```

### Features

- Segment intersection detection
- Automatic bridge generation

---

## v0.6 Live Preview

Status: initial HTTP/SSE live preview and source-positioned parser diagnostics
implemented; VS Code integration remains.

### xaligo serve

```bash
xaligo serve
```

### Features

- File watching
- Automatic re-rendering
- Real-time updates

### Backend Stack

- Go
- Echo
- WebSocket
- fsnotify

### Frontend Stack

- templ
- HTMX

### Preview Flow

#### Initial

```text
.xal
 ↓
SVG
 ↓
Browser
```

#### Real-Time Updates

```text
File Change
 ↓
Re-render
 ↓
Server-Sent Events
 ↓
Preview Refresh
```

---

## v0.7 VS Code Extension

Status: maintained in a separate repository. This core repository now provides
the required WASM diagnostics API, source positions, stable SVG rendering, and
HTTP/SSE preview protocol.

### Language Support

```text
.xal
```

### Features

- Syntax Highlighting
- Validation
- Error Location Reporting

### Preview Panel

```text
Editor
|
+- Source
+- Preview
```

### Live Preview

```text
Save
 ↓
xaligo render
 ↓
Preview Update
```

---

## v0.8 Excalidraw Integration

Status: native Excalidraw export exists; live WebView/updateScene integration
is not started.

### Excalidraw Preview

```text
.xal
 ↓
Excalidraw JSON
 ↓
WebView
```

### Features

- updateScene() support
- Real-time synchronization

### Excalidraw Export

```bash
xaligo render --format excalidraw
```

---

## v0.9 PowerPoint Export

Status: partially implemented ahead of schedule. Go-side geometry/routing plan
generation exists, and Node/PptxGenJS can generate PPTX as a temporary
development path. The required long-term gap is `xaligo.wasm`, invoked
from the Go repository layer with resolved plan JSON.

### PPTX Export

```bash
xaligo render --format pptx
```

Compatibility during transition:

- Keep `xaligo render --format pptx` usable when a WASM exporter is available.
- Do not reintroduce repository-layer Node subprocess execution as the default.
- Do not implement PPTX/OOXML writing in Go controller/repository code.
- Keep route/traffic/theme support renderer-agnostic where possible.

### Supported Features

- Shapes
- Connectors
- Routes
- Traffic Flows
- Themes

---

## v1.0

### VS Code Marketplace Release

#### Included Features

- Live Preview
- SVG Export
- Excalidraw Export
- PPTX Export
- Route Connectors
- Orthogonal Routing
- Traffic Layers
- Edge Offset
- Line Jumps

---

## Future Vision

### AWS Architecture Mode

```text
AWS Icons
Auto Layout
Route Layer
Traffic Layer
```

### Network Diagram Mode

```text
L2
L3
Route
Traffic
```

### Infrastructure as Diagram

```text
Diagram as Code
+
VS Code
+
Git
+
CI/CD
```

---

## Project Goal

Create a Diagram as Code platform positioned between:

```text
PlantUML
      +
Excalidraw
      +
draw.io
      +
Cloudcraft
      +
Instana-style Topology
```

with a strong focus on:

- AWS Architecture Diagrams
- Network Topology Diagrams
- Infrastructure Documentation
- Diagram-Driven Development
- Multi-mode 2D / 2.5D / Topology Rendering
- SVG / PPTX / Excalidraw / XYFlow / Isoflow Export
