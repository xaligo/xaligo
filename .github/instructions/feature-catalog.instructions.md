---
applyTo: "**"
---

# xaligo Feature Catalog

This file is the authoritative, ID-addressable catalog of xaligo's supported
and planned features. Read it as a precondition alongside
`agent-guide.instructions.md` to understand what the product already does (and
what it has already committed to doing) before proposing new work, filing a
roadmap entry, or judging whether a request is a bug fix, an extension of an
existing feature, genuinely new scope, or an already-tracked planned item.

Each row is a stable 7-digit feature ID (`XAL-GNNNNNN`), so implementation
notes, commit messages, roadmap entries, and issue reports can reference a
feature without repeating its full description. IDs are never reused or
renumbered once assigned; a removed feature's ID is retired, not recycled.

Every row carries a `Status` column:

- `Implemented` — shipped and available today.
- `Planned` — not yet implemented; tracked in `roadmap.instructions.md` and/or
  the `quality-improvement.instructions.md` Q05 backlog, which is the
  authoritative source for its exact sequencing and scope.
- `Excluded unless justified` — a considered but deliberately unsupported
  capability that stays out of scope until a non-substitutable use case is
  identified.

Do not remove or renumber a row when its status changes; update its `Status`
and Summary in place instead (e.g., `Planned` -> `Implemented` once a feature
ships).

- The leading digit `G` is the group (major capability area); see the section
  headers below.
- The remaining 6 digits are a per-group sequence number in steps of 10
  (`000010`, `000020`, ...), leaving room to insert a new fine-grained feature
  between two existing ones without renumbering the group.
- Add a new feature at the next free step within its group, or open a new
  group (next leading digit) for a capability area not covered below.

## Group 1 — Core DSL & Document Envelope (`XAL-1xxxxxx`)

| ID | Feature | Status | Summary |
|---|---|---|---|
| XAL-1000010 | Vue-style `.xal` XML DSL | Implemented | Vue-style layout DSL with XML syntax, parsed with `encoding/xml`, handling attributes, nested tags, and text content. |
| XAL-1000020 | Canonical V1 envelope | Implemented | `<xaligo version="1">` containing document-wide `<data>` and exactly one `<frames>` page collection. |
| XAL-1000030 | Legacy root compatibility | Implemented | Historical `<frame>`/`<frames>` document roots remain readable and renderable but emit a migration warning recommending the canonical envelope. |
| XAL-1000040 | V2 reject-safe root boundary | Implemented | `<scene version="2">` is a distinct root; a V1 reader rejects it outright instead of partially rendering V2 syntax as V1. |
| XAL-1000050 | Shared V1 pipeline | Implemented | One `.xal -> parser -> layout -> shared scene/plan -> encoder` pipeline shared by every output format, with no per-format parallel parser or layout path. |
| XAL-1000060 | `<frames>` page collection | Implemented | Groups page `<frame>` children with `gap` and optional `layout="vertical"`; every child `<frame>` requires a non-empty `id`. |
| XAL-1000070 | Frame sizing and spacing attributes | Implemented | `width`, `height`, `class`, `margin`/`margin-*`, `content-width`/`content-height`, `align`, and `item-size` on a page or nested frame. |
| XAL-1000080 | Numeric and geometry validation contract | Implemented | Finite-number and per-attribute domain checks (positive sizes, span ≤ 12, non-negative gaps, etc.) applied before layout, shared by `validate` and every render format. |
| XAL-1000090 | Child containment overflow policy | Implemented | `overflow="error"` (default, rejects out-of-bounds children) vs `overflow="visible"` (permits them while keeping sibling cursors deterministic). |
| XAL-1000100 | Frame metadata tag band | Implemented | `<metadata>`/`<entry>` key-value tag band exposing `id`, `title`, page-content `version`, and custom entries as a reserved strip other content must avoid. |
| XAL-1000110 | Page-content `version` attribute | Implemented | A non-empty `version` on an identified child `<frame>` is the page's visible content revision, distinct from the document-root DSL version selector. |
| XAL-1000120 | Document-level `<data>` container | Implemented | Reusable definitions (UML models, database schemas) declared once under `<data>` and referenced by `data="id"` from one or more frames. |
| XAL-1000130 | Resolved shared text-layout policy | Implemented | Role-based wrap/fit/overflow/line-height rules (group header, ordinary label, item label, port label, connector label) applied consistently by every encoder. |
| XAL-1000140 | V2 `<scene version="2">` native engine | Planned | A distinct, reject-safe V2 root and native frontend that lowers directly to the same typed version-neutral model as V1, sharing renderers/encoders downstream; V1 remains a frozen compatibility profile inside V2. |
| XAL-1000150 | Typed normalized layout structure | Planned | Replace repeated reads from `Node.Attrs` with first-class resolved-layout data so validation and rendering share one geometry source of truth. |
| XAL-1000160 | Format-neutral shared scene/plan schema | Planned | Migrate the canonical scene/plan schema away from Excalidraw/PPTX-shaped fields entirely, keeping only compatibility aliases for existing public callers. |

## Group 2 — Layout & Composition Primitives (`XAL-2xxxxxx`)

| ID | Feature | Status | Summary |
|---|---|---|---|
| XAL-2000010 | `<container>` primitive | Implemented | Vertical-by-default stack container; `layout="horizontal"` arranges children side by side. |
| XAL-2000020 | `<row>`/`<col>` 12-column grid | Implemented | Responsive grid primitive; `<col span="N">` occupies `N` of 12 columns, with flexible spans auto-dividing remaining columns. |
| XAL-2000030 | Fixed vs. flexible child allocation | Implemented | Explicit main-axis `width`/`height` reserves fixed space first; remaining children divide leftover space by `row`/`col` weight. |
| XAL-2000040 | Generic leaf tag rendering | Implemented | An unknown childless tag renders as a rectangle plus text, using `title`, text content, or the tag name as its label. |
| XAL-2000050 | Generic group/container tag rendering | Implemented | An unknown tag with layout children becomes a titled group, laid out vertically, horizontally, or with the V1 staggered layout. |
| XAL-2000060 | Item-grid row behavior | Implemented | A group whose children are all item-like (`item`, `spacer`, `blank`) automatically switches to item-grid row layout. |
| XAL-2000070 | `<rectangle>` general-purpose shape | Implemented | Titled/labeled rectangle that may contain multiple `<port>` children, unlike other generic leaf tags. |
| XAL-2000080 | `<port>` side-anchored sub-rectangle | Implemented | Small labeled rectangle bound to a side of its parent `<rectangle>`, with optional explicit `x`/`y` clamped inside the parent. |
| XAL-2000090 | `<item>` AWS service-icon leaf | Implemented | Places a catalog service icon by numeric ID, with configurable `item-size` and `dx`/`dy` offset; missing icons warn and skip instead of failing. |
| XAL-2000100 | `<spacer>`/`<blank>` empty slots | Implemented | Dedicated empty layout tags that occupy a grid slot without rendering an icon, border, or label. |
| XAL-2000110 | Custom leaf display toggles | Implemented | `border="none"` hides a leaf/group border; `visible="false"` hides one component's border/icon/label while preserving its layout space. |

## Group 3 — AWS Architecture Primitives & Icon Catalog (`XAL-3xxxxxx`)

| ID | Feature | Status | Summary |
|---|---|---|---|
| XAL-3000010 | `aws-cloud` group tag | Implemented | Top-level AWS Cloud boundary grouping primitive with automatic containment layout. |
| XAL-3000020 | `region` group tag | Implemented | AWS Region boundary grouping primitive. |
| XAL-3000030 | `vpc` group tag | Implemented | VPC boundary grouping primitive. |
| XAL-3000040 | `availability-zone` group tag | Implemented | Availability Zone boundary grouping primitive. |
| XAL-3000050 | `public-subnet` / `private-subnet` group tags | Implemented | Subnet-scoped grouping primitives for AZ-local resources. |
| XAL-3000060 | `security-group` group tag | Implemented | Grouping primitive for resources sharing an EC2 security group. |
| XAL-3000070 | `auto-scaling-group` group tag | Implemented | Grouping primitive for an EC2 Auto Scaling group. |
| XAL-3000080 | `generic-group` tag | Implemented | Non-AWS-specific logical grouping primitive for content that does not match a dedicated AWS group tag. |
| XAL-3000090 | AWS Architecture-Service-Icons catalog | Implemented | Bundled AWS service icon set looked up by numeric catalog ID from `service-catalog.csv`. |
| XAL-3000100 | Tabler icon catalog | Implemented | Imported and cataloged Tabler icon set, refreshed via `npm run import:tabler-icons`. |
| XAL-3000110 | Yamaha icon catalog | Implemented | Imported and cataloged Yamaha icon set, refreshed via `npm run import:yamaha-icons`. |
| XAL-3000120 | Isoflow icon manifest | Implemented | `isoflow-icons.json` manifest mapping catalog icons to Isoflow-compatible icon references, refreshed via `npm run generate:isoflow-icons`. |
| XAL-3000130 | Service ID lookup catalogs | Implemented | `service-index.csv` for quick ID lookup and `service-catalog.csv` for the full catalog with embedded icon data. |
| XAL-3000140 | `services.csv` label overrides | Implemented | `id,OfficialName,Abbreviation,Summary,Usage,Notes` per-diagram service list driving icon abbreviation overrides. |
| XAL-3000150 | Service legend generation | Implemented | SVG and PPTX output render a service legend derived from `services.csv`, with configurable SVG legend position. |
| XAL-3000160 | Icon license and attribution bundling | Implemented | Bundled license and attribution files for AWS, Tabler, and Yamaha icon sets are preserved alongside the generated catalogs. |

## Group 4 — Connections & Routing (`XAL-4xxxxxx`)

| ID | Feature | Status | Summary |
|---|---|---|---|
| XAL-4000010 | `<connection>` orthogonal connector | Implemented | Elbowed arrow between items, groups, rectangles, ports, or frames, declared as a direct child of `<frame>` or `<connections>`. |
| XAL-4000020 | `<connections>` grouping tag | Implemented | Optional wrapper providing shared connector defaults (color, stroke, kind, arrowheads, scale, grid) inherited by its `<connection>` children. |
| XAL-4000030 | Endpoint binding by ID/name/ref | Implemented | `src`/`dst` resolve against a catalog ID, or the `id`, `name`, or `ref` of an item, AWS group, rectangle, port, or identified child frame. |
| XAL-4000040 | Explicit endpoint side selection | Implemented | `src-side`/`dst-side` pin a connector endpoint to a specific `top`/`right`/`bottom`/`left` side. |
| XAL-4000050 | Perimeter anchor slots | Implemented | `src-anchor`/`dst-anchor` select one of 20 fixed inset positions (5 per side) around an endpoint's perimeter. |
| XAL-4000060 | Route layer (`kind="route"`) | Implemented | Headless structural paths with no arrowheads, validated so every effective start/end arrowhead resolves to `none`. |
| XAL-4000070 | Traffic layer (`kind="traffic"`) | Implemented | Directional flow lines that share a lane beside a matching route when endpoints coincide. |
| XAL-4000080 | Manual bend points | Implemented | Explicit bend/via child tags or the `bends`/`points`/`via` inline coordinate alias in frame coordinates. |
| XAL-4000090 | Arrowhead style selection | Implemented | Independent `start-arrowhead`/`end-arrowhead` (and `arrowhead` alias) selection from `none|arrow|triangle|stealth|diamond|oval`. |
| XAL-4000100 | Stroke style, width, and color overrides | Implemented | `stroke-style`, `stroke-width`/`width`, and six-digit hex `color`, with documented per-kind defaults and alias precedence rules. |
| XAL-4000110 | Per-connection scale and grid snapping | Implemented | Positive `scale`/`coordinate-scale` bend-coordinate multiplier and per-connection `grid` snap size. |
| XAL-4000120 | Automatic route junctions | Implemented | Shared route trunks automatically fan out into junction points instead of independent overlapping lines. |
| XAL-4000130 | Cross-frame page links | Implemented | A connection whose endpoints span two frames renders as two local stubs labeled `to <frame>` / `from <frame>` instead of an impossible cross-page line. |
| XAL-4000140 | Cross-frame terminal side/anchor control | Implemented | `src/dst-frame-side` and `src/dst-frame-anchor` fix the logical page side and tangent slot used by a cross-frame page-link stub. |
| XAL-4000150 | Automatic safe-side selection | Implemented | Automatic page-terminal selection picks the nearest safe frame side from rendered visual geometry, honoring the frame metadata reservation strip. |
| XAL-4000160 | Lane offsetting for overlapping connectors | Implemented | Parallel connectors sharing the same X/Y lane are offset from one another to remain individually legible. |
| XAL-4000170 | Plan-level `--arrow-style` default | Implemented | A render-option arrowhead/width default applied to SVG/PPTX/PDF/Excel Plan output only when a connection omits its own explicit value. |
| XAL-4000180 | Explicit circular connector nodes | Planned | A future versioned routing model with dedicated circular/loopback connector nodes, distinct from today's headless V1 route connectors. |
| XAL-4000190 | Line Jump routing | Planned | Phase-2 roadmap routing feature: a small arc/hop where two unrelated crossing connectors overlap, keeping crossing lines visually distinguishable. |
| XAL-4000200 | Layer Routing | Planned | Phase-2 roadmap routing feature that groups related connectors into coherent routing layers/lanes instead of resolving each connector independently. |

## Group 5 — Content Elements: Tables, Databases, UML (`XAL-5xxxxxx`)

| ID | Feature | Status | Summary |
|---|---|---|---|
| XAL-5000010 | `<table>` GFM-like pipe syntax | Implemented | Markdown-style pipe table with header/alignment row, normalized to the same typed rows as tag syntax. |
| XAL-5000020 | `<table>` explicit tag syntax | Implemented | `<header>`/`<row>`/`<cell>` children as an alternative to pipe syntax, normalized identically before layout. |
| XAL-5000030 | Table style inheritance | Implemented | Table-level `color`/`background-color`/`border-color`/`font-*` and `header-*` overrides, with `cell > header/row > table > built-in` precedence. |
| XAL-5000040 | Imported table schemas | Implemented | Table structure imported from an external schema source instead of authored inline in `.xal`. |
| XAL-5000050 | `<database-schema>` reusable definitions | Implemented | Document-level `<data>` schema definitions referenced by one or more frames via `<database data="schema-id">`. |
| XAL-5000060 | `<entity>`/`<column>` relational definitions | Implemented | Typed entity/column definitions supporting `name`, `type`, `primary-key`, `nullable`, `unique`, and `default`. |
| XAL-5000070 | `<foreign-key>` relation generation | Implemented | Single-column foreign key with a `references="entity.column"` target that generates an inter-entity relation. |
| XAL-5000080 | Relational key/nullability styling | Implemented | Primary/foreign key and nullability/data-type visual styling on entity column rows. |
| XAL-5000090 | Crow-foot relation notation | Implemented | Relation endpoints render with crow-foot cardinality notation on database diagrams. |
| XAL-5000100 | `<uml>` shared diagram component | Implemented | Common V1 component adapting UML families to xaligo's shared layout, shape, connector, and output pipeline. |
| XAL-5000110 | UML class diagrams | Implemented | `class`, `interface`, `enumeration` elements with `association`, `aggregation`, `composition`, `generalization`, `realization`, `dependency` relations. |
| XAL-5000120 | UML component diagrams | Implemented | `component`, `interface`, `port`, `artifact` elements with `dependency`, `realization`, `association`, `assembly`, `delegation` relations. |
| XAL-5000130 | UML activity diagrams | Implemented | `initial`/`final`/`activity`/`action`/`decision`/`merge`/`fork`/`join`/`object-node` elements with `control-flow`/`object-flow` relations. |
| XAL-5000140 | UML activity partitions/swimlanes | Implemented | `<partition>` swimlanes with `lanes="vertical|horizontal"` and the `theme="xaligo"` swimlane visual theme. |
| XAL-5000150 | UML state-machine diagrams | Implemented | `state`/`history`/`choice`/`fork`/`join`/`initial`/`final` elements with `transition` relations, events, guards, and effects. |
| XAL-5000160 | UML sequence diagrams | Implemented | `participant`/`lifeline` elements with `message`/`return-message`/`create-message`/`destroy-message` relations and explicit diagram-unique `order`. |
| XAL-5000170 | UML element compartments | Implemented | Typed ordered text compartments per element kind (`attribute`, `operation`, `constraint`, `note`, `entry`/`do`/`exit`, etc.). |
| XAL-5000180 | UML relation attributes | Implemented | Shared relation `label`/`guard`/`route` hint and `src-multiplicity`/`dst-multiplicity` on association-family relations. |
| XAL-5000190 | UML relation projection | Implemented | Fixed lowering of UML relation kinds to the shared orthogonal connector model (dashed/solid line, source diamond, no-arrowhead). |
| XAL-5000200 | Reusable `<uml-model>` definitions | Implemented | Document-level `<data>` UML models selected by one or more diagram-kind children via `data="model-id"`. |
| XAL-5000210 | Public UML connection endpoint references | Implemented | `uml-id/local-id` (same frame) and `frame-id.uml-id/local-id` (cross-frame) public endpoint forms for normal `<connection>` tags. |
| XAL-5000220 | Deliberately lossy V1 UML projection | Implemented | Documented, intentional projection limits (no dashed lifelines/activations/combined fragments, flattened compartments) applied consistently across every output format. |
| XAL-5000230 | UML timing diagrams | Planned | Accepted new UML family (Q05.17): `timing-diagram` selector with lifeline/state-timeline elements, state/value-change events, duration constraints, and time-axis layout. |

## Group 6 — Output Formats & Rendering (`XAL-6xxxxxx`)

| ID | Feature | Status | Summary |
|---|---|---|---|
| XAL-6000010 | Excalidraw output | Implemented | Editable Excalidraw JSON scene with groups, connectors, and metadata tags. |
| XAL-6000020 | SVG output | Implemented | Standalone SVG, one file per frame by default or one combined canvas with `--combine-frames`. |
| XAL-6000030 | PowerPoint (PPTX) output | Implemented | PPTX presentation built from a shared Go draw plan and executed by a configured WASM/PptxGenJS exporter, one slide per frame by default. |
| XAL-6000040 | PDF output | Implemented | PDF document, one page per frame by default, inheriting SVG's strict crop boundary. |
| XAL-6000050 | Excel output | Implemented | Excel workbook, one worksheet per frame by default, each embedding that frame's rendered SVG image. |
| XAL-6000060 | XYFlow (React Flow) output | Implemented | XYFlow-compatible JSON nodes and edges for React Flow-based viewers. |
| XAL-6000070 | Isoflow output | Implemented | Isoflow-compatible model JSON for isometric network diagrams. |
| XAL-6000080 | Frame-to-physical-page mapping contract | Implemented | Shared default mapping of one identified child frame to one SVG file/PPTX slide/PDF page/Excel worksheet. |
| XAL-6000090 | `--combine-frames` compatibility mode | Implemented | Restores the historical single-canvas/page form for SVG, PPTX, PDF, and Excel; Excalidraw/XYFlow/Isoflow are unaffected since they are already single documents. |
| XAL-6000100 | Safe frame-ID output naming | Implemented | Multi-frame SVG output derives `<stem>-<safe-frame-id>.svg` names, with deterministic collision detection. |
| XAL-6000110 | SVG legend placement | Implemented | `--svg-legend-position top|right|bottom|left` controls where the services.csv-derived legend renders. |
| XAL-6000120 | PPTX legend slides | Implemented | Dedicated legend slide(s) appended after diagram slides, using a 4-column icon/abbreviation/official-name layout. |
| XAL-6000130 | PDF page cropping | Implemented | PDF pages inherit the exact per-frame SVG canvas and clip boundary as their strict crop. |
| XAL-6000140 | Excel worksheet SVG embedding | Implemented | Each Excel worksheet embeds its frame's rendered SVG image rather than reconstructing native shapes. |
| XAL-6000150 | Single-logical-document output invariance | Implemented | Excalidraw, XYFlow, and Isoflow always remain one logical document regardless of `--combine-frames`. |
| XAL-6000160 | Isoflow generic-endpoint projection | Implemented | UML and other shapes without an Isoflow-native equivalent project to a labeled generic endpoint icon. |
| XAL-6000170 | Oversized-frame page tiling | Planned | Generic tiling of one oversized frame across several physical pages/slides, distinct from the existing one-frame-per-page mapping. |
| XAL-6000180 | Renderer capability/projection contract | Planned | A typed per-format capability declaration (e.g., which formats can carry arbitrary UML/connector metadata) so unsupported projections become explicit instead of implicit. |

## Group 7 — Presentation, Theming & Physical Page Fitting (`XAL-7xxxxxx`)

| ID | Feature | Status | Summary |
|---|---|---|---|
| XAL-7000010 | Light/dark theme support | Implemented | `--theme light|dark` applied consistently across every output format. |
| XAL-7000020 | Physical paper size selection | Implemented | `--paper A5|A4|A3|A2|A1|Letter|Legal|Tabloid` for PPTX/PDF/Excel physical-page fitting. |
| XAL-7000030 | Paper orientation selection | Implemented | `--orientation portrait|landscape`, with auto-fit when omitted. |
| XAL-7000040 | Per-side paper margins | Implemented | `--paper-margin` (all sides) and `--paper-margin-top/right/bottom/left` (inches) applied before fitting frame content to the physical page. |
| XAL-7000050 | Layout scaling base | Implemented | `--px-per-inch` controls the pixel-to-inch scaling base used for PPTX/PDF/Excel layout. |
| XAL-7000060 | PPTX document metadata | Implemented | `--title`/`--author`/`--company`/`--subject` set PPTX package-level metadata, independent of any frame `title` attribute. |
| XAL-7000070 | PPTX compression control | Implemented | `--compression`/`--no-compression` toggles PPTX output compression. |
| XAL-7000080 | Rendering mode selection | Implemented | `--mode standard|network|aws`; all three currently share the same resolved 2D rendering pipeline. |
| XAL-7000090 | Roadmap-reserved rendering modes | Implemented | `aws-2.5d` and `topology` are recognized enum values that currently return a not-implemented error rather than an unknown-value error. |
| XAL-7000100 | Shared font-family catalog | Implemented | `virgil`, `helvetica`, `cascadia`, `assistant`, `excalifont`, `nunito`, `lilita-one`, `comic-shanns`, `liberation-sans` mapped consistently to each output format's font face. |
| XAL-7000110 | `aws-2.5d` rendering mode | Planned | Cloudcraft/legacy AWS-reference-style oblique diagrams with `plane`/`zone` layout primitives, isometric nodes/routing, AWS node presets (Route 53, CloudFront, ELB, EC2, RDS, S3), and AWS Legacy/Cloudcraft-like themes; currently a recognized but not-implemented `--mode` value. |
| XAL-7000120 | `topology` rendering mode | Planned | Instana/SkyWalking-style dependency topology view; currently a recognized but not-implemented `--mode` value. |
| XAL-7000130 | Distinct per-mode visual semantics | Planned | `standard`, `network`, and `aws` currently execute the identical resolved 2D pipeline; the roadmap targets genuinely distinct layout/visual semantics per mode. |

## Group 8 — CLI Commands & Developer Tooling (`XAL-8xxxxxx`)

| ID | Feature | Status | Summary |
|---|---|---|---|
| XAL-8000010 | `xaligo render` command | Implemented | Renders a `.xal` source file into any supported output format via `--format`. |
| XAL-8000020 | `xaligo validate` command | Implemented | Validates `.xal` syntax, layout, and connection references without producing output. |
| XAL-8000030 | `xaligo serve` command | Implemented | Serves a live-reloading SVG preview over HTTP, polling the source file and re-rendering on change. |
| XAL-8000040 | `xaligo diff` command | Implemented | Renders paired removed/added structural-diff SVGs comparing two `.xal` documents. |
| XAL-8000050 | `xaligo generate xal` command | Implemented | Scaffolds a `.xal` file with a configurable AWS Cloud/Account/Region/VPC/AZ/Subnet hierarchy. |
| XAL-8000060 | `xaligo add service` (single mode) | Implemented | Appends one named AWS service icon and a legend entry into an existing `.excalidraw` file. |
| XAL-8000070 | `xaligo add service --list` (batch mode) | Implemented | Bulk-adds every service in a `services.csv` list into an existing `.excalidraw` file. |
| XAL-8000080 | `xaligo init` command | Implemented | Writes a minimal starter `sample.xal` file to help new users get started. |
| XAL-8000090 | `xaligo version` command | Implemented | Reports the resolved build/release version, with an embedded-value, env-var, and `VERSION`-file fallback chain. |
| XAL-8000100 | Man-page-style detailed CLI help | Implemented | Every command and subcommand exposes a detailed `Long` description and runnable `Examples` in its `--help` output, mirrored in the CLI reference documentation. |
| XAL-8000110 | Structured CLI logging | Implemented | Stable `share.NewMCode`-tagged debug/info/error log lines throughout controllers and use cases for traceable CLI diagnostics. |
| XAL-8000120 | VS Code extension integration | Planned | A separate-repository VS Code extension providing `.xal` syntax highlighting, a live Preview Panel, and source-positioned diagnostics, built on this repository's reusable render/validate APIs and HTTP/SSE preview protocol. |
| XAL-8000130 | Multi-view live preview | Planned | Extending `xaligo serve`/preview beyond SVG-first to Excalidraw, XYFlow, Isoflow, and 2.5D preview views. |
| XAL-8000140 | `xaligo render markdown` command | Implemented | Reads a Markdown file, renders every fenced ` ```xal ` code block to SVG through the shared render pipeline, and writes a new Markdown file with a `![](path.svg)` image reference per rendered frame in place of each code block; generated SVGs default beside the source Markdown file and `--svg-dir`/`--output` override the locations. |

## Group 9 — Validation, Diagnostics, Diff & External Integration (`XAL-9xxxxxx`)

| ID | Feature | Status | Summary |
|---|---|---|---|
| XAL-9000010 | Shared diagnostics use case | Implemented | One `DiagnosticsUsecase` (`Validate`/`Diagnose`) reused by both `xaligo validate` and the render pipeline's pre-flight checks. |
| XAL-9000020 | Source-positioned error reporting | Implemented | Validation and parse errors carry enough position/context information to be directly user-correctable. |
| XAL-9000030 | Strict-error vs. warning-fallback classification | Implemented | The `.xal` spec's documented split between hard validation errors and compatibility warning fallbacks (e.g., empty `align`, unknown nested attributes). |
| XAL-9000040 | Structural document diff engine | Implemented | `xaligo diff` compares parsed `.xal` data structures rather than source lines or formatting. |
| XAL-9000050 | Diff element matching strategy | Implemented | Matching prefers unique `id`/`name`/`ref`, then exact subtrees, then deterministic order-aware structural matching. |
| XAL-9000060 | Diff visual output | Implemented | Paired `-removed.svg`/`-added.svg` images highlighting removed/added elements and the old/new side of modified or moved elements. |
| XAL-9000070 | `@xaligo/xaligo-external` WASM/TypeScript package | Implemented | Exposes the PPTX draw-plan API through WebAssembly for browser/Node.js callers. |
| XAL-9000080 | TypeScript PPTX byte-output pipeline | Implemented | The external package consumes `BuildPPTXPlan` through WASM and produces PPTX bytes via PptxGenJS. |
| XAL-9000090 | `cmd/wasm` JavaScript/WASM adapter | Implemented | Exposes the native Go rendering engine to JavaScript/WASM hosts alongside the native CLI entry point. |
| XAL-9000100 | V1/V2 compatibility golden tests | Planned | Golden tests comparing V1-compatibility and native V2 engine output at the neutral-model and resolved-geometry boundaries once the V2 engine exists. |
| XAL-9000110 | Cross-renderer visual regression suite | Planned | Representative visual regression coverage across the full SVG/Excalidraw/PPTX/PDF/Excel/XYFlow/Isoflow format matrix; currently limited per the roadmap's documented gaps. |
| XAL-9000120 | Render determinism and concurrency-safety guarantees | Planned | Byte-stable output for identical source/options/assets/environment, and parallel-job safety without shared mutable render state (Q13.1/Q13.2 backlog). |
