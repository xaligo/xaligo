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
data remain local to that child.

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
anchor positions.

```text
top:    top-1    top-2    top-3    top-4    top-5
right:  right-1  right-2  right-3  right-4  right-5
bottom: bottom-1 bottom-2 bottom-3 bottom-4 bottom-5
left:   left-1   left-2   left-3   left-4   left-5
```

Position numbers run left-to-right on `top` and `bottom`, and top-to-bottom on
`left` and `right`. Corner anchors are not shared: `top-1` sits slightly inside
the top edge near the left corner, while `left-1` sits slightly inside the left
edge near the top corner.

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

Endpoints can also be written as child tags. Use this form when the endpoint
reference and anchor should stay together.

```xml
<connection kind="traffic">
  <src anchor="right-3">web</src>
  <dst side="left" anchor="5">app</dst>
</connection>
```

`id`, `ref`, `name`, or `target` attributes can provide the endpoint token when
the tag has no text content.

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

When endpoints belong to different child frames, editable page-oriented output
contains one local stub per frame, labeled `to <frame-id>` and
`from <frame-id>`. The stubs share a logical connector ID and the original
endpoint/routing metadata. Graph-oriented XYFlow and Isoflow output use that
metadata to emit one logical edge rather than two disconnected stubs.

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
