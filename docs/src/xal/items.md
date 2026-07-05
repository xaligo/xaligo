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
| `id` | Service catalog ID. Empty or omitted means spacer |
| `name` | Optional connection reference |
| `ref` | Optional connection reference |

When no icon is found for a service ID, rendering skips that item and emits a
warning.

Use `<spacer />` or `<blank />` for explicit empty grid cells:

```xml
<spacer />
<blank />
```

Spacers occupy layout slots but render no icon, label, border, or text.
