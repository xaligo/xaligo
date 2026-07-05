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
| `kind` | `route`, `traffic`, or default connection |
| `color` | Stroke color |
| `stroke-width` | Stroke width |
| `stroke-style` | `solid`, `dashed`, or `dotted` |
| `start-arrowhead`, `end-arrowhead` | `none`, `arrow`, `triangle`, `stealth`, `diamond`, or `oval` |
| `arrowhead` | Alias for `end-arrowhead` |
| `arrowhead-size` | Logical size `s`, `m`, or `l` |
| `grid` | Snap grid in layout pixels |
| `scale`, `coordinate-scale` | Bend coordinate multiplier |

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
