# Rectangles and Ports

`<rectangle>` creates a general-purpose labeled box.

```xml
<rectangle id="service" title="Service" width="180" height="100" font-size="18">
  <port id="service-in" side="left" title="in" font-size="9" />
  <port id="service-out" side="right" title="out" font-size="10" />
</rectangle>
```

`<port>` creates a small labeled rectangle inside a side of its parent
rectangle. Multiple ports on the same side are spaced evenly.

| Attribute | Target | Description |
|---|---|---|
| `id` | rectangle, port | Required connection reference ID |
| `width`, `height` | rectangle, port | Size in layout pixels |
| `title` or text | rectangle, port | Label text |
| `font-size` | rectangle, port | Label font size |
| `side` | port | `top`, `right`, `bottom`, or `left` |
| `x`, `y` | port | Optional position relative to parent top-left |

Port `x` and `y` values are clamped so the port remains inside the parent
rectangle.
