# .xal Overview

`.xal` is an XML DSL for diagrams. Every canonical language version uses a
`<xaligo>` envelope, document-wide data, and identified frames.

```xml
<xaligo version="1">
  <data>...</data>
  <frames gap="48">
    <frame id="overview" width="1440" height="900">...</frame>
    <frame id="detail" width="1440" height="900">...</frame>
  </frames>
</xaligo>
```

Historical root `<frame>` and `<frames>` documents remain compatible, but
produce a warning recommending migration to `<xaligo version="1">`. Omitting
the version from the canonical envelope defaults to V1 with a warning. V2 uses
`<xaligo version="2">`; `<frame version="2">` and `<frames version="2">` are
not valid V2 roots. Dispatch reads the root version once, and a V1 parser
rejects version 2 before interpreting nested syntax.

V2 deliberately retains the concise V1 authoring profile. Existing
`<row>`/`<col>`, AWS group tags, numeric catalog items, ports, and
`<connections>` lower directly to the typed V2 model; generic V2 parameters
such as `weight`, explicit `padding`, and direct `<line>` tags remain available
when needed.

The root `version="1"` above selects the DSL. A non-empty `version` on an
identified `<frame>` directly inside `<frames>` instead identifies that page's
visible content revision and does not select another language. Page frames can
also expose `title` and arbitrary key/value entries in a configurable
[frame metadata band](layout.md#frame-metadata).

Important rules:

- Layout coordinates use pixels.
- An identified child frame is one SVG artifact or PPTX slide by default.
  `--combine-frames` preserves the compatibility canvas or slide.
- Frame metadata is page-owned decoration inset from the selected logical frame
  edge and both row ends by its resolved `row-gap` (4 pixels by default). The
  same value separates wrapped rows. Its full-width reservation strip still
  starts at the outer logical frame edge, excludes normal items, text, lines,
  and labels, and follows that page projection. It also supplies the inward
  normal inset for safe page-link terminals; a frame without metadata uses 4
  pixels instead.
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
