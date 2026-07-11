# .xal Overview

`.xal` is an XML DSL for diagrams. V1 uses `<frame>` for one page and
`<frames>` for a multi-page document.

```xml
<frame version="1" width="1440" height="900" class="pa-4">
  ...
</frame>
```

```xml
<frames version="1" gap="48">
  <frame id="overview" width="1440" height="900">...</frame>
  <frame id="detail" width="1440" height="900">...</frame>
</frames>
```

Explicit `version="1"` is recommended for frozen V1 `<frame>` and `<frames>`
documents. Omission remains compatible and defaults to V1, but produces a
warning. No other explicit version is allowed on those roots. V2 uses the
distinct `<scene version="2">` root. This lets a V1 reader reject V2
immediately; `<frame version="2">` is not valid V2 syntax. A V2 renderer will
accept both native V2 documents and the frozen V1 profile, but V1 does not need
to understand V2.

Important rules:

- Layout coordinates use pixels.
- The origin is the upper-left of the rendered frame.
- Positive `x` extends right and positive `y` extends down.
- Connections must be direct children of `<frame>` or inside frame-level
  `<connections>`.
- AWS group tags, `<rectangle>`, and `<port>` require non-empty unique IDs so
  they can be referenced by connectors.

Unknown nested tags are a V1 extension point. A tag without layout children is
a generic rectangle-and-text leaf. A tag with layout children is a generic
group/container; an all-item child list uses the item grid. Unknown roots are
never treated this way.

V1 preserves several compatibility fallbacks: an empty `align` is omitted,
malformed `align` warns and keeps the default for unsupported components, and
unknown nested attributes are ignored. Invalid finite-number domains,
`overflow`, `layout`, connection sides/anchors/styles/arrowheads, and
render-option enums are errors. Recognized but unavailable modes report a
not-implemented error rather than silently changing mode.

The parser uses standard XML syntax. Escape special characters in attribute
values, for example `&amp;`. Use the documented lowercase spelling for tag names,
attribute names, and enum values; undocumented case or direction aliases are
not part of the frozen V1 profile.

For lookup tables, see the dedicated reference pages:

- [Icon Reference](../reference/icons/index.md)
- [Arrow Reference](../reference/arrows/index.md)
- [Frame and Border Reference](../reference/frames/index.md)
