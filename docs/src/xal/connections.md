# Connections

`<connection>` draws an elbowed connector between items, AWS groups and
boundary resources such as `<vpc-endpoint>`, rectangles, ports, or identified
child frames in a `<frames>` document.

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
| `src`, `dst` | Catalog ID or `id`/`name`/`ref` of an item, AWS group/boundary resource, rectangle, port, or identified child frame |
| `src-side`, `dst-side` | Optional endpoint side: `top`, `right`, `bottom`, or `left` |
| `src-anchor`, `dst-anchor` | Optional edge anchor: `top-1` through `top-5`, `right-1` through `right-5`, `bottom-1` through `bottom-5`, or `left-1` through `left-5` |
| `src-frame-side`, `dst-frame-side` | Cross-frame-only logical page side, independent of the endpoint side; the drawable terminal uses that side's inward inset line |
| `src-frame-anchor`, `dst-frame-anchor` | Cross-frame-only page side and tangent slot, using the same side/slot grammar; normal inset preserves the slot coordinate |
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
same percentages to select the logical page side and tangent slot
independently. The drawable frame terminal then moves only in the inward normal
direction to the page-terminal inset line.

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

For a page link, endpoint binding and page-terminal side selection may
deliberately use different sides:

```xml
<connection src="web" dst="detail.app"
            src-side="right" src-anchor="near"
            src-frame-side="bottom" src-frame-anchor="far"
            dst-side="left" dst-anchor="far"
            dst-frame-side="top" dst-frame-anchor="near" />
```

Here the source leaves the item through `right-2` but reaches the source page's
inset line at the `bottom-4` tangent coordinate. The destination starts from
the `top-2` tangent coordinate and reaches the item through `left-4`. The
drawable page terminals are normally inset without changing those percentages.
The segments adjacent to the endpoint and terminal inset line remain
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

When endpoints belong to different child frames, SVG and PPTX show the
connection as a page link with two local stubs rather than one line across the
inter-frame canvas. Markdown inherits this behavior through its embedded SVG:

- The source stub runs from the source endpoint to its frame's page-terminal
  inset line and is labeled exactly `to <destination frame ID>`.
- The destination stub runs from its frame's page-terminal inset line to the
  destination endpoint and is labeled exactly `from <source frame ID>`.

Angle brackets are visible punctuation. The example above therefore renders
`to <detail>` in `overview` and `from <overview>` in `detail`.

Endpoint binding uses `src-anchor`/`dst-anchor`, then
`src-side`/`dst-side`, then its normal automatic side. Logical page-terminal
precedence starts with `src-frame-anchor`/`dst-frame-anchor`, then
`src-frame-side`/`dst-frame-side`; either is a fixed choice. Without those
attributes, the legacy endpoint anchor, endpoint side, or normal nearest-edge
result is a preferred page side. Xaligo keeps it when safe; otherwise rendering
chooses the nearest safe side from the endpoint's actual visual envelope. An
item envelope includes both its icon and label. Equal-distance ties prefer a
tied side facing the other frame, then `top`, `right`, `bottom`, `left`.

Validation does not predict this automatic side. It checks that at least one
side can contain the resolved inset without entering metadata; rendering makes
the final choice after visual geometry is available. If the usual preferred
side is unsafe, it is remapped. Only a frame with no safe side reports a
source-positioned validation error at the connection.

Frame-terminal attributes are valid only when the resolved endpoints belong
to different frames. Using any `src-frame-*` or `dst-frame-*` attribute on a
same-frame connection is a source-positioned validation error.

Frame metadata reservation is a final safety constraint after that precedence.
Without an explicit frame-terminal attribute, xaligo excludes every unsafe
side before its visual nearest-side choice and remaps an unsafe preferred side.
A selected left/right terminal is clamped outside the full-width reservation
strip. An explicit frame side or anchor that selects the reserved edge, or an
exact left/right anchor inside the strip, is a validation error instead of
being silently moved. The path and `to <...>` / `from <...>` label cannot enter
the strip.

The outer logical frame edge remains the side, tangent-anchor, and crop
reference; SVG and PPTX do not draw it as a frame outline. The
drawable terminal is on a parallel inward inset line. Its inset is the resolved
metadata `row-gap` when that frame has metadata, or 4 layout pixels when it
does not. The metadata value applies to all four terminal sides regardless of
whether the band is at the top or bottom; `row-gap="0"` places the terminal on
the outer edge. For an explicit frame side or anchor, the inset must be strictly
smaller than the frame height on `top`/`bottom`, or the frame width on
`left`/`right`; an invalid choice is reported at the connection. For an
automatic terminal, the same bounds classify safe candidates, and only an
empty candidate set is an error. Xaligo uses the resolved value directly
instead of clamping it.

An explicit terminal vertically opposite the metadata edge must remain outside
the full reservation strip: `top` against bottom metadata, or `bottom` against
top metadata. Moving it into the strip is a source-positioned validation error.
For an automatic terminal, the same conflict removes that side from the safe
candidate set. A safe explicit `left` or `right` terminal is allowed even if an
unused top/bottom inset line would intersect the strip.

The inset changes only the normal coordinate. An explicit frame anchor keeps
its exact 10/30/50/70/90-percent coordinate along the outer frame extent.
Otherwise the initial parallel coordinate follows the endpoint binding. If it
enters the normal 24-layout-pixel corner gutter, xaligo clamps that coordinate
and inserts a two-bend orthogonal dogleg; the endpoint- and
frame-terminal-adjacent segments remain perpendicular to their selected sides.
Borders shorter than 96 layout pixels use an adaptive quarter-length gutter.
If an unconstrained inset terminal would coincide with the endpoint, it moves
by up to 24 layout pixels along the parallel axis within the available range so
the local stub remains visible. An explicit frame anchor keeps its tangent slot
and uses a visible orthogonal local stub. Manual bends are retained as logical
routing metadata but do not change these page-local stub paths.

For an automatic left/right coincidence next to metadata, the preferred range
also keeps 8 layout pixels from the reservation. If a very small safe region
cannot keep both that clearance and the corner gutter, xaligo uses the full
non-reserved interval instead. The terminal may touch its boundary but never
moves outside the frame or into the metadata strip.

One zero-inset combination is invalid. If the owning frame has metadata with
`row-gap="0"`, the connection endpoint resolves to that frame itself, and an
explicit frame anchor coincides with the resolved endpoint point, validation
reports a source-positioned validation error. An explicit endpoint anchor uses
its stated slot. An explicit endpoint side uses its center, so `src-side="top"`
matches `src-frame-anchor="top-3"`; an automatically selected endpoint side
also uses its center. Choose a different endpoint/frame anchor or set a
positive `row-gap`. Xaligo cannot keep both ends perpendicular and fixed while
drawing a visible stub from one coincident point.

The page-link label is placed from the final inset terminal with a
4-layout-pixel inward gap and a minimum 4-layout-pixel tangent gap. Candidate
placement uses the closest tangent position that avoids the endpoint envelope
and metadata strip; a tiny page clamps or shrinks the fallback instead of
moving the label farther inward.

The stubs share a logical connector ID and the original endpoint, frame, and
routing metadata.

By default, the source and destination stubs appear on their respective SVG
files or PPTX slides. With `--combine-frames`, both remain visible on the
compatibility canvas but are not joined across the frame gap.

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
