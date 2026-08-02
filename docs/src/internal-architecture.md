# Internal architecture

xaligo has one parser/layout/routing pipeline and two supported outputs: SVG and
PPTX. Markdown support is a host workflow that embeds SVG artifacts.

```text
.xal source
   -> version dispatch and parser
   -> validated generic concepts and parameters
   -> layout, text measurement, ports, and connector routing
   -> renderer-neutral document/draw plan
   -> SVG encoder
      or
   -> PPTX WASM exporter
```

![xaligo rendering pipeline](images/internal-rendering-pipeline.svg)

Diagram source: [internal-rendering-pipeline.xal](architecture/internal-rendering-pipeline.xal)

Controllers collect command options and persist returned bytes. They do not
parse, lay out, route, or own a second format switch. `RenderUsecase` owns
dispatch and accepts only `svg` and `pptx`; SVG is the default.

## Markdown

`xaligo render markdown` scans fenced `xal` blocks, invokes the SVG artifact
path for each block, writes one SVG per resolved frame, and replaces the block
with image references. The preview server follows the same rule. Markdown does
not introduce another geometry model or encoder.

## Current V1 compatibility path

The existing V1 frontend still creates an Excalidraw-shaped internal
presentation scene before building a draw plan. This is temporary compatibility
data only:

- it is not a supported output format;
- it has no CLI or browser-WASM render entry point;
- it is not persisted or edited by a public command; and
- new public contracts must use renderer-neutral names and types.

The complete document is resolved before frame projection so cross-frame links,
obstacles, routing, metadata bands, and source order remain consistent. An
identified child frame maps to one SVG artifact or one PPTX slide. With
`--combine-frames`, all frames share one SVG canvas or PPTX slide.

![orthogonal routing algorithm](images/internal-routing-algorithm.svg)

Diagram source: [internal-routing-algorithm.xal](architecture/internal-routing-algorithm.xal)

## PPTX boundary

Go owns semantic interpretation and geometry. It serializes a resolved document
plan and invokes `external/pptx-exporter/wasm/xaligo.wasm` in process. The
TypeScript/PptxGenJS adapter translates plan operations into PPTX bytes; it
does not parse `.xal` or recalculate layout and routing.

## V2 engine boundary

The V2 generic calculation core lives in the Rust workspace at
`external/engine`. Native Go calls its `staticlib` through C ABI v2 and cgo:

```text
Go use case -> cgo -> C ABI -> Rust staticlib
```

Rust calculates only generic frames, groups, captures, items, ports, text,
spacers, and lines. It now resolves nested vertical/horizontal/fixed/flex/grid
and absolute layouts, ports, straight/orthogonal routes, and the matching SVG
projection. Builtin, AWS, and UML profiles are declarative concept data,
defaults, constraints, styles, aliases, icons, and attribution. Profiles cannot
register calculation callbacks or output encoders.

The V2 response is an immutable, renderer-neutral `ResolvedDocument`. ABI input
uses contiguous pre-order records and parent indexes rather than JSON or
arbitrary maps, preserving unset values separately from explicit zero values.
The native V2 and V1 compatibility frontends still need to lower `.xal`
directly to this model; until then, the existing V1 public render path remains
independent rather than using a partial conversion.

See [V2 Generic Engine](design/v2-engine.md) for the implemented calculation,
validation, ABI, and compatibility contracts.

## Project intelligence

RAG, LSP, and MCP share the constructor-injected project use case. One `.xal`
analysis pass produces diagnostics and generic Frame,
Group, Capture, Item, Port, Line, Text, and Spacer symbols for editor and agent
requests. Markdown headings and content are indexed in SQLite/FTS5 for durable
search.

The initial RAG corpus is intentionally narrower than the analysis service:
`rag index` and `rag watch` discover only Markdown below `docs/`. Open or saved
`.xal` documents are explicit editor requests and do not broaden that discovery
rule. LSP session state remains in memory; only durable search rows are shared
between separately launched xaligo processes.

The MCP protocol layer is stateless. Each request supplies its protocol version
and client capabilities, then a deterministic tool adapter calls the existing
diagnostics, render, project, or icon use case. The stdio transport uses one
JSON object per line; Streamable HTTP uses one localhost-only `POST /mcp` per
message and validates Origin plus mirrored routing headers before dispatch.
Neither transport creates another parser, opens a second project database, or
starts a daemon or Rust sidecar.

## Assets

Generic icons are built in. AWS and UML icon data remain separable profiles.
The SQLite registry owns SVG storage, search, aliases, tags, checksums, and
license metadata. The Rust engine receives resolved icon metadata/content and
does not access SQLite, repositories, filesystems, or the network.

## Deliberately absent outputs

Excalidraw, PDF, Excel/XLSX, XYFlow, and Isoflow encoders, aliases,
dependencies, generated assets, and browser globals are retired. Adding a new
format requires an explicit product-scope decision; plugins cannot add one.
