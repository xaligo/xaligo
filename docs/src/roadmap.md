# Planned Work

xaligo is being developed as a diagram-as-code platform for architecture,
network, and operational visualization workflows. The items below are planned
or under consideration and may change as the core renderer evolves.

![xaligo Roadmap](images/xaligo-roadmap.png)

## Rendering and Scale

- Performance improvements for large diagrams.
- Rendering support for larger architecture maps.
- Page splitting for diagrams that do not fit on a single page.

### Phase 1 Implementation Plan

1. Add shared page-splitting analysis for the resolved draw plan.
   This creates page tiles and records which drawing operations intersect each
   tile without changing SVG or PPTX output yet.
2. Connect the page metadata to SVG export so large diagrams can be emitted as
   page-sized SVG files or page groups.
3. Connect the same metadata to PPTX export so oversized diagrams can become
   multiple slides while reusing the existing plan geometry.
4. Add large-diagram regression samples and performance benchmarks around
   parser, layout, scene construction, routing, plan building, and output
   encoding.
5. Optimize the slowest measured stages, with preference for shared caches and
   data structures that benefit every renderer.

Initial implementation status:

- Shared page-splitting analysis has started in the Go usecase layer.
- `SplitPlanPagesChecked` rejects non-finite slide, operation, point, page, and
  overlap geometry and assigns stable unique IDs to anonymous operations.
- Rendering output is unchanged until SVG and PPTX are wired to the page
  metadata.

## Rendering Correctness Foundation

The shared renderer now rejects non-finite and invalid layout numbers, resolves
fixed children before flex ratios, records content boxes and explicit overflow,
detects parent and port overlap violations, and gives SVG/PPTX a common text
layout and PPI transform. CLI format dispatch also goes through one use-case
entry point. See [Internal Architecture and Algorithms](internal-architecture.md).

The next structural steps are:

1. Store a typed normalized layout specification instead of repeatedly reading
   numeric strings from the syntax tree.
2. Store the shared item-grid solver's selected cells in resolved layout and
   include mixed item/rectangle occupancy; minimum-cell and item-offset
   preflight already run during `Build`.
3. Move catalog-derived intrinsic label measurement and final connector
   geometry into the same validation pass used by render.
4. Replace the remaining Excalidraw-shaped canonical scene and presentation-
   shaped physical plan schema with genuinely format-neutral models, retaining
   compatibility aliases at public boundaries.
5. Extend cross-format regression coverage for editable text behavior, item
   offsets, connector values, and non-default PPI/paper fitting.

## V1 Compatibility and V2

`<xaligo version="1">` is the canonical V1 envelope. Historical root `<frame>`
and `<frames>` documents remain compatible but emit a migration warning. V2 will use
the distinct `<scene version="2">` root, allowing an existing V1 reader to
reject V2 safely without understanding it.

The V2 renderer will also accept V1 input through a V2-owned compatibility
frontend. Native V2 and compatible V1 input will each be parsed once and
lowered directly to the same typed, renderer-neutral model. The design avoids
XML rewriting, parser retry, serialized scene round-trips, and running the full
V1 renderer inside V2. V1 itself remains independent of V2.

Golden compatibility tests will cover roots and defaults, unknown nested tags,
strict versus fallback enum behavior, connection inheritance and anchors,
signed-32-bit catalog IDs, item-size render contexts, and equivalent resolved
geometry across native and embedded targets.

## Input and Output Formats

- Excel export and Excel-friendly workflows.
- Import from existing diagram formats and conversion into `.xal`.
- Better round-tripping between generated output and `.xal` source.

The V1 structured-diagram profile includes a document-wide data registry,
general tables, relational schema/ER views, and all fourteen UML diagram-kind
components. It deliberately
keeps their semantic processors separate while reusing neutral drawing and
encoding contracts. See [Structured Diagrams: Tables, Databases, and
UML](design/structured-diagrams.md). Its `<xaligo version="1">` envelope is the
canonical V1 syntax; legacy root documents continue to render with warnings.

## Editing and Automation

- A dedicated UI for authoring and editing diagrams.
- MCP interfaces so AI agents and tools can inspect, generate, and update
  diagrams through xaligo.
- GUI-to-`.xal` workflows, including configuration changes driven from visual
  edits.

## Runtime Visualization

- Visualization of running systems.
- Agent and server components for collecting system state and rendering it as
  diagrams.

## Advanced Views

- 3D modeling support for richer architecture and infrastructure
  visualization.
