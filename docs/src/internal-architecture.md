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

V2 moves generic measurement, layout, and routing into the Rust workspace at
`external/engine`. Native Go calls a Rust `staticlib` through the versioned C
ABI and cgo:

```text
Go use case -> cgo -> C ABI -> Rust staticlib
```

Rust calculates only generic frames, groups, captures, items, ports, text,
spacers, and lines. Builtin, AWS, and UML profiles are declarative concept data,
defaults, constraints, styles, aliases, icons, and attribution. Profiles cannot
register calculation callbacks or output encoders.

The target V2 response is an immutable, renderer-neutral `ResolvedDocument`.
SVG and PPTX consume it without recomputing dimensions, grids, anchors, paths,
text policy, or icon placement. The V1 compatibility frontend will eventually
lower directly to the same model, removing the temporary renderer-shaped
scene.

## Assets

Generic icons are built in. AWS and UML icon data remain separable profiles.
The SQLite registry owns SVG storage, search, aliases, tags, checksums, and
license metadata. The Rust engine receives resolved icon metadata/content and
does not access SQLite, repositories, filesystems, or the network.

## Deliberately absent outputs

Excalidraw, PDF, Excel/XLSX, XYFlow, and Isoflow encoders, aliases,
dependencies, generated assets, and browser globals are retired. Adding a new
format requires an explicit product-scope decision; plugins cannot add one.
