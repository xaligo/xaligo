# Layout

## Frame

`<frame>` defines the canvas size and the top-level layout.

| Attribute | Default | Description |
|---|---|---|
| `width` | `1280` | Frame width in pixels |
| `height` | `720` | Frame height in pixels |
| `class` |  | Vuetify-style spacing class |
| `layout` | vertical | Use `horizontal` for side-by-side children |
| `gap` | `16` | Gap between child elements |
| `item-size` | config value | Max icon size for `<item>` elements |
| `margin`, `margin-*` |  | Content inset in pixels |
| `content-width`, `content-height` |  | Override usable content size |
| `align` |  | `top|middle|bottom` plus `left|center|right|spread` |

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

## Spacing

`class` supports Vuetify-like spacing classes such as `pa-4` and `ma-4`.
Root frame margins inset content without changing the paper frame size.
