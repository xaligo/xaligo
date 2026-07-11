# Layout

## Frame

`<frame>` defines one canvas and its top-level layout. `<frames>` wraps multiple
identified `<frame>` children; it lays them out horizontally by default and
vertically with `layout="vertical"`.

| Attribute | Default | Description |
|---|---|---|
| `version` | `1` with warning when omitted | Root only. Explicit `1` is recommended and is the only accepted value |
| `width` | `1280` | Frame width in pixels |
| `height` | `720` | Frame height in pixels |
| `class` |  | Vuetify-style spacing class |
| `layout` | vertical | Use `horizontal` for side-by-side children |
| `gap` | `16` | Gap between child elements |
| `item-size` | render context, normally `32` | Max icon size for `<item>` elements. A root value makes output deterministic across native and embedded renderers |
| `margin`, `margin-*` |  | Content inset in pixels |
| `content-width`, `content-height` |  | Override usable content size |
| `align` |  | `top|middle|bottom` plus `left|center|right`; item grids additionally support `spread` |

## Container

`<container>` stacks children vertically by default. Use `layout="horizontal"`
for horizontal placement.

```xml
<container class="pa-4" gap="16">
  ...
</container>
```

## Row And Col

`<row>` and `<col>` provide a 12-column grid.

```xml
<row gap="20">
  <col span="8">...</col>
  <col span="4">...</col>
</row>
```

`<row>` and `<col>` are pure layout tags. They do not render borders or labels.

## Custom Tags

An unknown nested tag with no layout children is a generic leaf. If it has
layout children, V1 treats it as a generic group/container with group insets;
`layout="horizontal"` and `layout="staggered"` select those layouts. If every
child is an item or spacer, the tag uses the item-grid row behavior. The rule
does not apply to the document root.

Only documented `layout` values are valid for each tag. An unknown non-empty
value is a validation error; it does not silently select the default layout.

## Spacing

`class` supports Vuetify-like spacing classes such as `pa-4` and `ma-4`.
Root frame margins inset content without changing the paper frame size.
