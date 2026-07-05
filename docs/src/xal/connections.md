# Connections

`<connection>` draws an elbowed connector between items, groups, rectangles, or
ports.

Connections must be direct children of `<frame>` or inside a frame-level
`<connections>` block.

```xml
<connections kind="traffic" color="#2563eb" grid="8" scale="1">
  <connection src="web" dst="app" />
  <connection src="app" dst="db" color="#059669" />
</connections>
```

| Attribute | Description |
|---|---|
| `src`, `dst` | Catalog ID or `id`/`name`/`ref` endpoint |
| `src-side`, `dst-side` | Optional endpoint side: `top`, `right`, `bottom`, or `left` |
| `src-anchor`, `dst-anchor` | Optional edge anchor: `top-1` through `top-5`, `right-1` through `right-5`, `bottom-1` through `bottom-5`, or `left-1` through `left-5` |
| `kind` | `route`, `traffic`, or default connection |
| `color` | Stroke color |
| `stroke-width` | Stroke width |
| `stroke-style` | `solid`, `dashed`, or `dotted` |
| `start-arrowhead`, `end-arrowhead` | `none`, `arrow`, `triangle`, `stealth`, `diamond`, or `oval` |
| `arrowhead` | Alias for `end-arrowhead` |
| `arrowhead-size` | Logical size `s`, `m`, or `l` |
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

`start`, `near`, `center`, `far`, and `end` are accepted aliases for anchor
positions `1`, `2`, `3`, `4`, and `5`.

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

Traffic lines that share endpoints with a route are drawn beside the route lane
when possible.

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
rectangle, or port. Missing endpoints, duplicate aliases, ambiguous numeric
IDs, and nested connection tags are validation errors.
