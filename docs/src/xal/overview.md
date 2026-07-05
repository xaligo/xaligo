# .xal Overview

`.xal` is an XML DSL for diagrams. The root tag is always `<frame>`.

```xml
<frame width="1440" height="900" class="pa-4">
  ...
</frame>
```

Important rules:

- Layout coordinates use pixels.
- The origin is the upper-left of the rendered frame.
- Positive `x` extends right and positive `y` extends down.
- Connections must be direct children of `<frame>` or inside frame-level
  `<connections>`.
- AWS group tags, `<rectangle>`, and `<port>` require non-empty unique IDs so
  they can be referenced by connectors.

The parser uses standard XML syntax. Escape special characters in attribute
values, for example `&amp;`.

For lookup tables, see the dedicated reference pages:

- [Icon Reference](../reference/icons/index.md)
- [Arrow Reference](../reference/arrows/index.md)
- [Frame and Border Reference](../reference/frames/index.md)
