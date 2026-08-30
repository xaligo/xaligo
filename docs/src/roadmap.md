# Planned Work

xaligo is being developed as a diagram-as-code platform for architecture,
network, and operational visualization workflows. The items below are planned
or under consideration and may change as the core renderer evolves.

![xaligo Roadmap](images/xaligo-roadmap.png)

## Rendering and Scale

- Performance improvements for large diagrams.
- Rendering support for larger architecture maps.
- Tiling one oversized frame across several physical pages.

### Frame pagination status

Identified child frames are now physical page boundaries. The full scene is
resolved first and then projected in source order:

| Output | Default frame mapping |
|---|---|
| SVG | One file per frame |
| PPTX | One slide per frame |

`--combine-frames` requests one SVG canvas or PPTX slide. Default SVG uses the
exact frame rectangle as its canvas and clip boundary, while combined SVG keeps
marker-safe bounds. Markdown embeds the same SVG artifacts.

Page frames can also add a top/bottom metadata band for built-in `id`, `title`,
content `version`, and arbitrary key/value tags. The resolved `row-gap`
(4 pixels by default) supplies both the inter-row spacing and the metadata
page-edge inset at the selected vertical edge and both horizontal edges; rows
use `frame width - 2 * row-gap`. A full-width reservation strip still starts
at the outer logical frame edge, extends to the final content-box boundary,
remains at least
`row-gap + complete band height + 8` pixels deep, and excludes normal items,
text, connector geometry and labels, and page links. Auto/fixed widths,
typography, colors, ordered greedy row wrapping, explicit row breaks, per-row
alignment, page ownership, and safe page-link edge selection are resolved
before the physical formats project one frame per page. Cross-frame links can
select item and logical page-edge anchors independently, preserve the frame
anchor's tangent coordinate, and place the drawable terminal on a parallel
inward line. That inset is the resolved metadata `row-gap`, or 4 pixels when
metadata is absent; zero retains the outer edge and the value is never clamped.
Links reject unsafe explicit geometry. Without an explicit frame terminal,
unsafe candidates are filtered and rendering chooses the nearest safe side
from actual visual geometry; only an empty candidate set is an error. Labels
remain 4 layout pixels from the final inset terminal.

This frame pagination is separate from generic tiling. Remaining scale work is
to split one oversized frame into multiple tiles, add large-diagram regression
samples and benchmarks, and optimize the slowest measured shared stages.

## Rendering Correctness Foundation

The shared renderer now rejects non-finite and invalid layout numbers, resolves
fixed children before flex ratios, records content boxes and explicit overflow,
detects parent and port overlap violations, and gives SVG/PPTX a common text
layout and PPI transform. CLI format
dispatch also goes through one use-case entry point. See
[Internal Architecture and Algorithms](internal-architecture.md).

The next structural steps are:

1. Store a typed normalized layout specification instead of repeatedly reading
   numeric strings from the syntax tree.
2. Extend the V2 resolved adaptive item grid beyond all-item groups to mixed
   item/rectangle occupancy; minimum-cell and item-offset preflight already run
   during V1 `Build`.
3. Move catalog-derived intrinsic label measurement and final connector
   geometry into the same validation pass used by render.
4. Replace the remaining internal Excalidraw-shaped V1 compatibility scene and
   presentation-shaped plan schema with genuinely format-neutral models.
5. Extend SVG/PPTX/Markdown regression coverage for text behavior, item
   offsets, connector values, and non-default PPI/paper fitting.

## V1 Compatibility and V2

`<xaligo version="1">` is the canonical V1 envelope. Historical root `<frame>`
and `<frames>` documents remain compatible but emit a migration warning. V2
uses `<xaligo version="2">`; a V1 reader rejects the unsupported document
version before interpreting nested syntax.

The V2-owned frontend can normalize both language versions, while the public
V1 render path remains frozen. V2 preserves the concise V1 authoring profile
and lowers it, together with explicit generic V2 extensions, directly to the
same typed, renderer-neutral model. Each source is parsed once. The design
avoids XML rewriting, parser retry, serialized intermediate round-trips, and
running the full V1 renderer inside V2. V1 itself remains independent of V2.

The embedded Rust calculation core, C ABI v2, authoring-profile frontend, and
SVG/PPTX projection are implemented for the
generic Frame, Group, Capture, Item, Port, Line, Text, and Spacer concepts. It
supports nested fixed/flex, fixed-column and adaptive item grids, absolute
placement, ports, generic routing, deterministic SVG projection, V1-style
frame metadata and AWS/catalog defaults, and shared icon placement. Remaining
work includes multi-page V2 documents and broader profile coverage. See
[V2 Generic Engine](design/v2-engine.md).

Golden compatibility tests will cover roots and defaults, unknown nested tags,
strict versus fallback enum behavior, connection inheritance and anchors,
signed-32-bit catalog IDs, item-size render contexts, and equivalent resolved
geometry across native and embedded targets.

## Input and Output Formats

SVG and PPTX are the cross-version engine-output set. V2 additionally exposes
terminal text; Markdown embeds SVG artifacts.
Import from existing diagram formats may be considered separately, but it must
normalize into `.xal`/the generic model and cannot add a hidden render path.

The V1 structured-diagram profile includes a document-wide data registry,
general tables, relational schema/ER views, and the supported class, component,
activity, state-machine, and sequence UML diagram components. It deliberately
keeps their semantic processors separate while reusing neutral drawing and encoding
contracts. See [Structured Diagrams: Tables, Databases, and
UML](design/structured-diagrams.md). Its `<xaligo version="1">` envelope is the
canonical V1 syntax; legacy root documents continue to render with warnings.

## Editing and Automation

- A dedicated UI for authoring and editing diagrams.
- Add explicitly authorized model-assisted editing operations through the
  existing application-service boundary.
- GUI-to-`.xal` workflows, including configuration changes driven from visual
  edits.

## Runtime Visualization

- Visualization of running systems.
- Agent and server components for collecting system state and rendering it as
  diagrams.

## Advanced Views

- 3D modeling support for richer architecture and infrastructure
  visualization.
