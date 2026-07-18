# Connections

`<connection>` draws an elbowed connector between items, AWS groups,
rectangles, ports, or identified child frames in a `<frames>` document.

Connections must be direct children of `<frame>` or inside a frame-level
`<connections>` block.

```xml
<connections kind="traffic" color="#2563eb" grid="8" scale="1">
  <connection src="web" dst="app" />
  <connection src="app" dst="db" color="#059669" />
</connections>
```

The group inherits only non-empty `arrowhead-size`, `kind`, `color`,
`stroke-width`, `width`, `stroke-style`, `start-arrowhead`, `end-arrowhead`,
`arrowhead`, `scale`, `coordinate-scale`, and `grid` values. A child attribute
overrides its group value. Endpoint identity and geometry are never inherited:
each child supplies `src` and `dst`, and sides, anchors, bends, points, and via
data remain local to that child. This includes the cross-frame-only
`src-frame-*` and `dst-frame-*` terminal attributes.

Alias pairs are overridden as one setting: `stroke-width`/`width`,
`end-arrowhead`/`arrowhead`, and `coordinate-scale`/`scale`. If a child uses
either name, the corresponding parent pair is not inherited. If both names are
on the child, the first name shown here takes precedence.

`<connections>` accepts only `<connection>` children. Unknown or misspelled
child tags are validation errors.

| Attribute | Description |
|---|---|
| `src`, `dst` | Catalog ID or `id`/`name`/`ref` of an item, AWS group, rectangle, port, or identified child frame |
| `src-side`, `dst-side` | Optional endpoint side: `top`, `right`, `bottom`, or `left` |
| `src-anchor`, `dst-anchor` | Optional edge anchor: `top-1` through `top-5`, `right-1` through `right-5`, `bottom-1` through `bottom-5`, or `left-1` through `left-5` |
| `src-frame-side`, `dst-frame-side` | Cross-frame-only logical page-terminal side, independent of the endpoint side |
| `src-frame-anchor`, `dst-frame-anchor` | Cross-frame-only logical page-terminal anchor, using the same side/slot grammar |
| `kind` | `connection`, `route`, `traffic`, or omitted for the default connection |
| `color` | Six-digit hexadecimal stroke color (`#RRGGBB`) |
| `stroke-width`, `width` | Stroke width; `width` is the compatibility alias |
| `stroke-style` | `solid`, `dashed`, or `dotted` |
| `start-arrowhead`, `end-arrowhead` | `none`, `arrow`, `triangle`, `stealth`, `diamond`, or `oval`; an effective route permits only `none` |
| `arrowhead` | Alias for `end-arrowhead`; an effective route permits only `none` |
| `arrowhead-size` | V1 fixed size `s`. Other sizes are validation errors |
| `grid` | Snap grid in layout pixels |
| `scale`, `coordinate-scale` | Bend coordinate multiplier |

## Endpoint Anchors

When `src-side`, `dst-side`, `src-anchor`, or `dst-anchor` are omitted, xaligo
chooses the endpoint sides and anchor positions automatically from the endpoint
geometry.

Use `src-anchor` and `dst-anchor` to pin an endpoint to a specific perimeter
anchor. Each side has five inset positions, so the rectangle has 20 unique
anchor positions. Cross-frame `src-frame-anchor` and `dst-frame-anchor` use the
same positions to pin the logical page terminal independently.

```text
top:    top-1    top-2    top-3    top-4    top-5
right:  right-1  right-2  right-3  right-4  right-5
bottom: bottom-1 bottom-2 bottom-3 bottom-4 bottom-5
left:   left-1   left-2   left-3   left-4   left-5
```

Position numbers run left-to-right on `top` and `bottom`, and top-to-bottom on
`left` and `right`. They are exactly 10, 30, 50, 70, and 90 percent along the
named edge. Corner anchors are not shared: `top-1` sits inside the top edge
near the left corner, while `left-1` sits inside the left edge near the top
corner.

```xml
<connection src="web" dst="app"
            src-anchor="right-3"
            dst-anchor="left-3" />
```

You can also split the side and position:

```xml
<connection src="web" dst="app"
            src-side="right" src-anchor="3"
            dst-side="left" dst-anchor="3" />
```

Anchor aliases map one-to-one: `start=1`, `near=2`, `center=3`, `far=4`, and
`end=5`.

For a page link, endpoint binding and page-edge termination may deliberately
use different sides:

```xml
<connection src="web" dst="detail.app"
            src-side="right" src-anchor="near"
            src-frame-side="bottom" src-frame-anchor="far"
            dst-side="left" dst-anchor="far"
            dst-frame-side="top" dst-frame-anchor="near" />
```

Here the source leaves the item through `right-2` but reaches the source page
at `bottom-4`. The destination enters its page at `top-2` and reaches the item
through `left-4`. The segments adjacent to the endpoint and page edge remain
perpendicular to their respective sides.

Endpoints can also be written as child tags. Use this form when the endpoint
reference and anchor should stay together.

```xml
<connection kind="traffic">
  <src anchor="right-3" frame-side="bottom" frame-anchor="far">web</src>
  <dst side="left" anchor="5" frame-anchor="top-2">detail.app</dst>
</connection>
```

`id`, `ref`, `name`, or `target` attributes can provide the endpoint token when
the tag has no text content. On `<src>` and `<dst>`, `frame-side` and
`frame-anchor` map to the corresponding source/destination frame-terminal
attributes. A complete anchor supplies its side; a separate side accepts slot
`1..5` or `start|near|center|far|end`. Conflicting side and complete-anchor
values are validation errors.

## Route And Traffic

Use `kind="route"` for structural paths without arrowheads. Use
`kind="traffic"` for directional flows.

```xml
<connection src="web" dst="app" kind="route" />
<connection src="web" dst="app" kind="traffic" color="#2563eb" />
```

Routes are strictly headless in V1. After `<connections>` defaults and child
alias overrides are merged, a non-`none` `start-arrowhead`, `end-arrowhead`, or
`arrowhead` on an effective route is a source-positioned validation error.
Explicit `none` is allowed. For example, a child `arrowhead="none"` suppresses
an inherited `end-arrowhead` value because those names form one semantic alias
pair.

Traffic lines that share endpoints with a route are drawn beside the route lane
when possible.

## Cross-frame Connections

References without a dot are resolved only inside the frame containing the
connection. Use `frameId.id` to address an endpoint in another frame:

```xml
<frames>
  <frame id="overview">
    <rectangle id="web" />
    <connection src="web" dst="detail.db" />
  </frame>
  <frame id="detail">
    <rectangle id="db" />
  </frame>
</frames>
```

Every `<frame>` inside `<frames>` must have a unique, non-empty `id`; omitting
it is a validation error. Endpoint `id`, `name`, and `ref` values are scoped by
that frame, so the same local identifier may be reused in different frames.
Frame IDs and connectable local IDs must not contain a dot because the dot is
the scope delimiter.
The same qualification applies to connectable table, database, entity, group,
rectangle, port, and item IDs. An unqualified reference never falls back to a
different frame.

When endpoints belong to different child frames, Excalidraw, SVG, PPTX, PDF,
and Excel show the connection as a page link with two local stubs rather than
one line across the inter-frame canvas:

- The source stub runs from the source endpoint to the logical page edge of its
  frame and is labeled exactly `to <destination frame ID>`.
- The destination stub runs from the logical page edge of its frame to the
  destination endpoint and is labeled exactly `from <source frame ID>`.

Angle brackets are visible punctuation. The example above therefore renders
`to <detail>` in `overview` and `from <overview>` in `detail`.

Endpoint binding uses `src-anchor`/`dst-anchor`, then
`src-side`/`dst-side`, then its normal automatic side. Logical page-terminal
precedence is `src-frame-anchor`/`dst-frame-anchor`, then
`src-frame-side`/`dst-frame-side`, then the legacy endpoint anchor, then the
endpoint side, then automatic nearest-edge selection. Automatic selection uses
the logical frame edge nearest to that endpoint's visual envelope. An item
envelope includes both its icon and label. Equal-distance ties prefer a tied
side facing the other frame, then `top`, `right`, `bottom`, `left`.

Frame-terminal attributes are valid only when the resolved endpoints belong
to different frames. Using any `src-frame-*` or `dst-frame-*` attribute on a
same-frame connection is a source-positioned validation error.

Frame metadata reservation is a final safety constraint after that precedence.
Without an explicit frame-terminal attribute, xaligo remaps a reserved
top/bottom page edge to the nearest safe edge and clamps a left/right terminal
outside the full-width reservation strip. An explicit frame side or anchor
that selects the reserved edge, or an exact left/right anchor inside the strip,
is a validation error instead of being silently moved. The path and
`to <...>` / `from <...>` label cannot enter the strip.

The terminal is on the logical frame edge; SVG, PPTX, PDF, and Excel do not
draw that edge as a frame outline. An explicit frame anchor uses its exact
10/30/50/70/90-percent point. Otherwise the initial position along the edge
follows the endpoint binding. If that unconstrained position enters the normal
24-layout-pixel corner gutter, xaligo clamps the terminal and inserts a
two-bend orthogonal dogleg; the endpoint- and frame-adjacent segments remain
perpendicular to their selected sides. Borders shorter than 96 layout pixels
use an adaptive quarter-length gutter. If an unconstrained terminal would
coincide with the endpoint on the same border, it moves by up to 24 layout
pixels within the available range so the local stub remains visible. An
explicit frame anchor stays at its exact slot and uses a visible orthogonal
local stub. Manual bends are retained as logical routing metadata but do not
change these page-local stub paths.

The page-link label sits just inside the page with a 4-layout-pixel inward gap
and a minimum 4-layout-pixel tangent gap from its terminal. Candidate placement
uses the closest tangent position that avoids the endpoint envelope and
metadata strip; a tiny page clamps or shrinks the fallback instead of moving
the label farther inward.

The stubs share a logical connector ID and the original endpoint, frame, and
routing metadata. Graph-oriented XYFlow and Isoflow output use that metadata
to emit one logical edge rather than two disconnected stubs.

By default, the source and destination stubs appear on their respective SVG
files, PPTX slides, PDF pages, or Excel worksheet images. With
`--combine-frames`, both remain visible on the compatibility canvas but are not
joined across the frame gap.

## Manual Bends

Manual bend coordinates are in frame coordinates.

```xml
<connection src="web" dst="db" grid="8">
  <bend x="120" y="80" />
  <bend x="120" y="220" />
  <bend x="300" y="220" />
</connection>
```

`<point>`, `<via>`, and `<waypoint>` are aliases for `<bend>`. Coordinates can
also be grouped inside `<bends>`, `<points>`, or `<path>`.

## Text Shorthand

Text shorthand can be used directly inside `<frame>`:

```xml
web --- db
web ==> db
```

`---` expands to `kind="route"`. `==>` expands to `kind="traffic"`.

## Validation

Endpoint references must resolve to exactly one rendered item, group,
rectangle, port, or identified child frame. Missing endpoints, duplicate
aliases, ambiguous numeric IDs, and nested connection tags are validation
errors.

Invalid sides, anchors, `kind`, stroke styles, arrowheads, and arrowhead sizes
are validation errors. Effective routes also reject every non-`none`
arrowhead.
