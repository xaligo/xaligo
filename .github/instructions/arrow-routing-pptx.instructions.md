---
applyTo: "**/*.{go,ts,xal,md}"
---

# xaligo — PPTX Routing / Legend Preconditions

This file is the current source of truth for PPTX export geometry.

## Brainstorm Reference

- ChatGPT share: https://chatgpt.com/share/6a35c5b9-4528-83e8-aff9-bc37907a4d80
- The share page may not be accessible from automated tooling. Keep the concrete
  decisions below authoritative for implementation.

## Confirmed Decisions

- PPTX export is an A3-landscape-first workflow for the current AWS sample.
- The PPTX export implementation should be compiled to WASM and invoked from
  the Go repository layer.
- Do not use `goja` or V8 for PPTX export execution.
- Avoid a long-term Node.js subprocess dependency for repository-layer PPTX
  export. Node may remain a development/build tool only while the WASM exporter
  is being prepared.
- All PPTX geometry and routing decisions are computed by the Go use-case
  pipeline before the exporter boundary.
- Each identified child frame becomes one diagram slide in source order by
  default. `--combine-frames` is the explicit compatibility path for the
  former single-slide canvas.
- A presentation has one common slide size. Multi-frame PPTX uses the largest
  resolved page width and height and centers smaller frame pages without
  scaling them independently.
- The PPTX drawing/export layer must not make independent layout/routing
  decisions.
- Lines must not visually cover icons or labels.
- If any obstacle-free route exists, obstacle-hitting routes must be rejected.
- Item labels are 8pt at the default 96 PPI and scale with the same effective
  PPI/paper-fit transform as item icons.
- Item icons should remain visually consistent with their labels; avoid shrinking
  icons merely to satisfy a cramped row when layout whitespace controls can be used.
- Legend belongs on separate PPTX slide(s), not outside the diagram page.
- Legend slide layout is fixed to 4 columns and contains icon, abbreviation, and
  official service name.
- DSL must support empty grid cells and both inner/outer whitespace controls.

## Current Pipeline

```text
.xal DSL
  -> Go parse and numeric-domain validation (typed normalization is the target)
  -> resolved layout and canonical scene
  -> ordered page-oriented Go document plan (neutral-schema migration remains)
  -> internal repository encoder (SVG), or
  -> Go repository -> WASM command -> external controller -> use case -> repository
  -> SVG | .pptx
```

Geometry belongs on the Go side. The WASM export module should only translate
the resolved plan into PPTX bytes. Excalidraw-compatible JSON may be one scene
serialization, but it is not the target architecture name or ownership boundary
for the shared plan.

## Go / WASM Boundary

The adopted integration style is Go invoking a WASM-compiled PPTX exporter from
the repository layer.

Implementation preconditions:

- Go owns CLI/controller/repository orchestration.
- WASM must be called from `internal/repository/powerpoint.go`, not directly from
  controller or command packages.
- The exporter must be compiled to WASM before repository-layer execution.
- Go forwards user-facing PPTX options to the WASM exporter through a typed
  options structure or JSON bridge.
- The WASM exporter consumes the resolved shared Go plan and returns PPTX
  bytes or writes them through a repository-controlled output path.
- The WASM exporter must not perform independent geometry, layout, or routing.
- The external WASI command calls its controller, the controller calls the
  external use case, and only the use case calls the external PPTX repository.
  Command/controller code must not bypass this path.
- Go repository/controller code must not implement PPTX/OOXML drawing or zip
  writing directly. Keep Go as the adapter that builds the plan, invokes the
  WASM exporter, and persists the returned bytes.
- If existing TypeScript/PptxGenJS code cannot be compiled into a practical WASM
  exporter, replace that drawing layer with a WASM-compatible PPTX writer rather
  than introducing `goja` or V8.

Other integration styles are not the current implementation target:

| Style | Status |
|---|---|
| stdin/stdout JSON-RPC | Candidate for long-running/high-volume workflows |
| HTTP API | Candidate for service/BFF separation |
| gRPC | Candidate for high-performance typed service boundaries |
| Node.js subprocess | Temporary fallback only; not the target architecture |
| Embedded JS engine (`goja`, V8) | Not a target for PPTX export |

Do not spend implementation time replacing the repository-layer exporter with
`goja` or V8 unless that architecture is explicitly re-approved.

## Ownership

| Area | Owner |
|---|---|
| DSL parse/layout | `internal/usecase/v1/engine/parse_*`, `internal/usecase/v1/engine/layout_*` |
| Canonical scene and item metadata | `internal/usecase/v1/engine/scene_*` |
| Plan geometry, text layout, paper scaling, routing, legend data | `internal/usecase/v1/engine/plan_*`, `internal/usecase/v1/engine/route_*` |
| WASM exporter invocation from Go | `internal/repository/powerpoint.go` |
| WASM-compatible PPTX drawing/export | `external` TypeScript package and implementation |
| PPTX WASI command entry | `external/command.ts` |
| Public browser/JavaScript API bridge | `cmd/wasm/main.go` |

## Paper / Scaling

- PPTX export supports `--paper`, `--orientation`, and paper-margin fitting
  flags.
- A3 landscape is generated with:

```bash
.bin/xaligo render docs/src/examples/samples/sample.xal \
  --format pptx \
  --services docs/src/examples/samples/services.csv \
  -o out.pptx \
  --paper A3 \
  --orientation landscape \
  --paper-margin-top 0.75 \
  --paper-margin-bottom 0.75
```

- The shared Go plan resolves paper size and computes one layout-pixel transform.
- The page-oriented plan is built after the full scene and cross-frame page
  links are resolved. Its frame projections preserve source order.
- Shape coordinates, font sizes, strokes, padding, and routing geometry use
  that same transform. `--px-per-inch 144` must not scale text independently
  from its containing shape.
- `--paper-margin N` applies an inch-based margin to every side before fitting
  the diagram to the selected paper.
- `--paper-margin-top`, `--paper-margin-right`, `--paper-margin-bottom`, and
  `--paper-margin-left` override the all-side value for individual sides.
- Paper margins do not change the slide size; they reduce the available fit
  area and centre the diagram within that inset area.
- The `paper-frame` element remains the content frame for scaling.
- Root `<frame margin="N">` or `class="ma-N"` is content outer whitespace: it
  insets diagram content without shrinking the paper frame itself.

## Routing Rules

- Route calculation is in `internal/usecase/v1/engine/route_*`.
- Obstacles include image and text rectangles from the Excalidraw scene.
- Start/end rectangles are excluded from obstacle checks for that connection.
- Binding `gap` from Excalidraw arrows must be honored in PPTX routing.
- If any obstacle-free candidate exists, obstacle-hitting candidates must not be
  selected.
- Lines on an obstacle boundary count as collision.
- Existing routed paths are included in scoring so later lines avoid overlap and
  near-parallel crowding.
- Excalidraw output also feeds previously routed lines back into the shared
  router so matching X/Y lanes are offset before export.
- Visible container borders are reserved routing paths. Connectors may cross a
  frame boundary, but parallel paths prefer the configured line margin.
- Previously placed line lanes are used as candidate offsets, so `--arrow-margin`
  affects routes that would otherwise share the same position.
- Final PPTX drawing order is:
  1. anchor backgrounds and containers/shapes
  2. route lines, traffic lines, and line-jump masks
  3. automatic junction markers
  4. icons and labels

This order prevents lines from visually covering icons even at endpoints.

Excalidraw output mirrors this readability rule with editable JSON elements:
each item image and label is grouped with a small 5x5 white anchor grid. The
grid is drawn above connector lines and below the icon/label so labels remain
readable without hiding the connector endpoint. Excalidraw routing treats group
header tags, item icons, and labels as obstacles where possible, and serializes
arrowhead sizes as `"s"` for dense diagrams.

### Cross-frame page links

A connection between different frames is a page link in page-oriented output;
it is never one line crossing the inter-frame canvas. The shared scene emits
two axis-aligned local stubs for Excalidraw, SVG, PPTX, PDF, and Excel:

- source endpoint to the source frame's logical page edge, with the exact label
  `to <destination frame ID>`; and
- destination frame's logical page edge to the destination endpoint, with the
  exact label `from <source frame ID>`.

Angle brackets are literal punctuation, so a link from `overview` to `detail`
renders `to <detail>` and `from <overview>`. The shared scene, not the PPTX
exporter, selects endpoint binding and logical frame terminal geometry. The
endpoint uses `src-side`/`dst-side` and `src-anchor`/`dst-anchor`. A cross-frame
connection may independently select its logical terminal with
`src-frame-side`/`dst-frame-side` or the more specific
`src-frame-anchor`/`dst-frame-anchor`. Every side has five anchors at
10/30/50/70/90 percent along the edge. The endpoint- and frame-adjacent route
segments remain perpendicular to their respective sides even when the two
sides differ.

Frame-terminal precedence is explicit frame anchor, explicit frame side,
legacy endpoint anchor, endpoint side, then automatic nearest-side selection.
The automatic side minimizes the perpendicular distance from the endpoint
visual envelope to the four logical frame edges. Ties prefer a tied side facing
the remote frame, then `top`, `right`, `bottom`, `left`. Frame-terminal
attributes are cross-frame-only; using them on a same-frame connection is a
validation error.

A frame metadata reservation strip is a final safety constraint. The visible
metadata rows are inset from their selected vertical edge and both horizontal
edges by the resolved `row-gap`, while the reservation itself remains
full-width from the outer logical frame edge to the content boundary. When no
explicit frame-terminal attribute is present, a selected reserved top/bottom
edge is remapped to the nearest safe edge and a left/right terminal is clamped
outside the full-width strip before the orthogonal dogleg is built. An
explicit frame side or anchor that selects the metadata edge, or an exact
left/right anchor inside the strip, is a validation error instead of being
moved. Neither the local path nor its label may enter the reservation strip.

The unconstrained terminal uses the endpoint binding's coordinate parallel to
the border. If that coordinate enters a 24-layout-px corner gutter, the
terminal is clamped and a two-bend orthogonal dogleg bridges the coordinate
difference; the segments at both the endpoint and logical page edge remain
perpendicular to their selected side. A border shorter than 96 layout pixels
uses one quarter of its length as an adaptive gutter. If an unconstrained
terminal and the endpoint coincide on the border, the terminal shifts by up to
24 layout pixels within the available gutter range so the line remains
visible. An explicit frame anchor remains at its exact slot; its orthogonal
local stub supplies the visible separation.
Manual bends remain connector metadata and
do not steer page-local stubs. Both stubs retain one logical connector ID;
XYFlow and Isoflow reconstruct one graph edge from that metadata rather than
exporting the two page projections.

The `to <...>` / `from <...>` label is placed just inside the owning page with
a 4-layout-pixel inward gap and a minimum 4-layout-pixel tangent gap from the
terminal. Candidate placement chooses the closest tangent position that avoids
the endpoint envelope and metadata reservation; tiny pages use a clamped
fallback rather than increasing the normal label distance.

The page edge is geometric, not a visible rectangle: SVG, PPTX, PDF, and Excel
omit page-frame outlines in both default and combined output.

Default PPTX output places the source and destination stubs on their respective
frame slides. `--combine-frames` places both stubs on the compatibility slide
but never draws a replacement line across the frame gap.

## Advanced Routing Features

### Line Jumps

Excalidraw does not provide reliable built-in line jumps/bridges for this
workflow. The shared draw plan therefore implements them for SVG/PPTX.

Current approach:

- Detect line segment intersections after routing.
- Determine which line is visually above the other by layer/kind/order.
- Render jumps as a 6px background-colored mask below the upper line in
  SVG/PPTX output. The mask uses the uppermost opaque container background at
  the crossing. A curved arc may replace the rectangular mask later.
- For Excalidraw output, approximate with normal lines or supported shape
  primitives when necessary.

SVG preview and PPTX can support line jumps more accurately than Excalidraw JSON.

### Route / Traffic Separation

Network diagrams distinguish structural route lines from traffic-flow lines.

Implemented model:

| Kind | Meaning | Visual Direction |
|---|---|---|
| `route` | Physical/logical connection path | Thin, lower layer, no arrowheads, shortest orthogonal route |
| `traffic` | Communication flow over a route | Offset beside a matching route, higher layer, directional arrow/style |

Potential DSL forms:

```xml
<connection src="A" dst="B" kind="route" />
<connection src="A" dst="B" kind="traffic" />
```

or future shorthand:

```text
A -> B
A => B
```

Routing orders routes below normal connections and traffic. When a traffic line
shares the same endpoints as a route line, the traffic line follows a nearby
parallel lane instead of drawing directly on top of the route.

### Route Connectors

Frozen V1 routes are headless in every format. Their effective
`start-arrowhead` and `end-arrowhead` must both resolve to `none` after
`<connections>` defaults and child aliases are merged. A non-`none` value is a
validation error rather than a renderer-specific circular endpoint.

Small circular route connector nodes remain a future versioned feature; they
must use a renderer-neutral connector-node concept instead of overloading V1
arrowheads.

Conceptual shape:

```text
[EC2] -- o -------- o -- [RDS]
```

Future behavior may render explicit connector nodes in SVG/PPTX and equivalent
editable shapes in Excalidraw. It is not part of the V1 compatibility profile.

## Connector Style Options

`xaligo render --format pptx` forwards all PPTX routing options:

| Flag | Meaning |
|---|---|
| `--arrow-style` | `thin`, `standard`, `triangle`, `stealth`, `arrow`, `diamond`, `oval`, `none` |
| `--arrow-stub` | Pixel stub before the first/last bend |
| `--arrow-margin` | Pixel margin reserved around existing line lanes |
| `--px-per-inch` | Layout scaling base, default 96 |
| `--paper` | Named slide paper size: `A5`, `A4`, `A3`, `A2`, `A1`, `Letter`, `Legal`, `Tabloid` |
| `--orientation` | `portrait` or `landscape`; auto-fit when omitted |
| `--paper-margin` | Inch margin applied to all sides before paper fitting |
| `--paper-margin-top/right/bottom/left` | Inch margin override for one side |

`--arrow-style` is a Plan-level default. A connection's explicit or inherited
DSL arrowhead and stroke width take precedence; `kind="route"` remains
headless. The `thin` and `standard` presets may supply a default line width only
when the DSL did not supply `stroke-width` or its `width` alias.

Every numeric render option must be finite. `--px-per-inch`, arrow stub and
margin values, and paper margins reject negative values; the internal zero
value selects the documented default. Validation happens before scene/plan
construction so `NaN` or infinity cannot first fail during JSON encoding.
Paper size, orientation, and arrow style are closed enums. Paper margins require
a named paper size, and their effective left/right and top/bottom sums must
leave a strictly positive content area in the selected (or at least one
automatic) orientation.

## Group Header Tags

- Group header tag labels are single-line in every output whose text engine can
  represent that policy. The shared draw plan marks their semantic role,
  wrapping, fitting, clipping, line height, and padding; the TS drawing layer
  consumes those values rather than inferring behavior from an element ID.
- Excalidraw scene generation must reserve conservative tag label width before
  PPTX export. `groupLabelCharW` is intentionally larger than the average
  Excalidraw text metric so PowerPoint no-wrap text stays inside the tag
  background.
- When changing group tag font size, font family, padding, or tag geometry,
  update both the scene width estimate and the group-header regression tests.

## Item Labels

- Item icon size defaults to 32px in native CLI config.
- Item label font is 8pt at the default 96 PPI. At another effective PPI it is
  `10.666...px * 72 / effectivePPI` points so its ratio to the icon is stable.
- Excalidraw font size for item labels is `8pt * 96 / 72 = 10.666...px`.
- Item label boxes are 14px high.
- Do not shrink label boxes to text metrics if it breaks PowerPoint placement.

## Layout / Whitespace

Supported whitespace controls:

| Syntax | Behavior |
|---|---|
| `<spacer />` / `<blank />` | Empty layout slot, not rendered |
| `<item />` | Empty item-grid slot, not rendered |
| `class="pa-4"` | Inner padding, Vuetify-style 8px unit |
| `class="ma-4"` | Outer margin; on root frame this becomes page-edge content whitespace |
| `margin="N"` and `margin-*` | Pixel margin |
| `content-width="N"` / `content-height="N"` | Shrinks usable inner layout area |
| `align="top-left"` etc. | Aligns the usable content area or item grid |
| `width="N"` / `height="N"` | Fixed child size, except root frame is the paper/content frame |

Fixed children are reserved before flexible `row`/`col` allocation. The
resolved size advances the sibling cursor and must remain inside the parent's
content box unless the source explicitly uses `overflow="visible"`. Layout
overflow is diagnosed before plan construction; SVG or PPTX clipping is not a
substitute for a valid layout.

For item grids, horizontal `spread` is also supported.

## Legend Pages

PPTX export adds legend slides after all frame/diagram slides when `--services`
is provided.

- Legend data is derived from `services.csv`.
- Only services actually used in the scene are included.
- The legend contains icon, abbreviation, and official name.
- Legend layout is fixed to 4 columns per slide.
- Additional legend slides may be created when entries exceed one slide.
- The diagram slide should not include an outside-frame legend; the PPTX legend
  belongs on separate slides.

## Verification Checklist

Before considering PPTX routing/layout changes complete:

```bash
go test ./...
make build
make build-wasm
npm run build --workspace @xaligo/xaligo-external
.bin/xaligo render docs/src/examples/samples/sample.xal --format pptx --services docs/src/examples/samples/services.csv -o out.pptx --paper A3 --orientation landscape --arrow-style thin
unzip -t out.pptx
```

For icon-overlap regressions, inspect the resolved PPTX XML and ensure routed
custom geometry does not intersect target icon/label rectangles.
