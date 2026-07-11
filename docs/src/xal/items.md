# Items and Spacers

`<item>` places a catalog icon inside its parent group or container.

```xml
<public-subnet id="public" title="Public Subnet">
  <item id="1178" name="web" />
  <item />
  <item id="1189" name="edge" />
</public-subnet>
```

| Attribute | Description |
|---|---|
| `id` | Positive signed 32-bit decimal catalog ID (`1..2147483647`). Empty or omitted means spacer; zero, signs, non-decimal syntax, and out-of-range values are invalid |
| `name` | Optional connection reference |
| `ref` | Optional connection reference |
| `dx` / `dy` | Relative icon offset in pixels from the icon's normal layout `x,y` position. The moved icon rectangle must remain inside the parent frame/group border |

When no icon is found for a service ID, rendering skips that item and emits a
warning.

The root `item-size` overrides the rendering environment's item-icon size.
Without it, the V1 native and canonical embedded profile defaults to `32`
layout pixels, while a caller-supplied render context may intentionally choose
another value. Set `item-size="32"` explicitly when output must be identical
across environments.

Use `<spacer />` or `<blank />` for explicit empty grid cells:

```xml
<spacer />
<blank />
```

Spacers occupy layout slots but render no icon, label, border, or text.
