# Layout

## Frame

`<frame>` defines one page canvas and its top-level layout. `<frames>` wraps
multiple identified `<frame>` children; it lays them out horizontally by
default and vertically with `layout="vertical"` in the complete logical scene.
SVG, PPTX, PDF, and Excel project each child frame to one file, slide, page, or
worksheet by default. `--combine-frames` preserves that complete scene as one
physical canvas/page. SVG, PPTX, PDF, and Excel do not draw the page frame's
outline, including in combined compatibility output. A default page-local SVG
uses the exact frame rectangle as its canvas and clip boundary; PDF and Excel
inherit that strict page/image crop. Combined SVG keeps marker-safe bounds
expansion. Excalidraw keeps editable frame objects but gives them transparent
strokes.

| Attribute | Default | Description |
|---|---|---|
| `version` | `1` with warning when omitted | On `<xaligo>` or a legacy root, selects DSL V1. On an identified direct child `<frame>`, any non-empty value is the visible page revision |
| `title` |  | Page title shown in the frame metadata band |
| `width` | `1280` | Frame width in pixels |
| `height` | `720` | Frame height in pixels |
| `class` |  | Vuetify-style spacing class |
| `layout` | vertical | Use `horizontal` for side-by-side children |
| `gap` | `16` | Gap between child elements |
| `item-size` | render context, normally `32` | Max icon size for `<item>` elements. A root value makes output deterministic across native and embedded renderers |
| `margin`, `margin-*` |  | Content inset in pixels |
| `content-width`, `content-height` |  | Override usable content size |
| `align` |  | `top|middle|bottom` plus `left|center|right`; item grids additionally support `spread` |

Document-root and page-frame versions have different roles. Use
`<xaligo version="1">` to select the V1 language. A `version` on an identified
`<frame>` directly inside the root `<frames>` is page content such as
`version="2026.07"`; it is displayed in the metadata band and participates in
structural diff. A legacy root `<frame version="1">` still means DSL V1 and is
not displayed as a page revision.

## Frame metadata

Set a frame `title`, a child-frame `version`, or add direct `<metadata>` to
add a key/value tag band near the selected edge of the outer frame border box;
the default position is the top. The resolved `row-gap` is also the page-edge
inset at the selected top/bottom edge and both row ends, so the default `4`
leaves a 4-pixel selected-edge gutter and 4-pixel row-end gutters. Padding,
content margins, and the content box do not replace or add to that inset. An
existing frame with only an `id` remains visually unchanged; once the band is
enabled its non-empty built-ins appear as `id`, `title`, and `version`,
followed by custom entries in source order.

```xml
<frame id="architecture" title="AWS Architecture" version="1.0.0"
       width="720" height="420"
       margin-top="52" margin-right="24"
       margin-bottom="52" margin-left="24">
  <metadata align="right" font-family="helvetica">
    <entry key="owner" value="Platform Engineering" />
    <entry key="status" value="Approved" break-before="true"
           width="180" key-width="56" />
  </metadata>
  ...
</frame>
```

`<metadata>` is non-layout configuration and may appear at most once as a
direct frame child. It accepts only empty `<entry key="..." value="..." />`
children; keys and values must be non-empty. This frame-local spelling is
separate from document-level `<xaligo><metadata>`.

| Attribute | Default | Description |
|---|---|---|
| `position` | `top` | `top` or `bottom` |
| `align` | `left` | `left`, `center`, or `right`, applied to each row separately |
| `font-family` | `virgil` | One of the nine supported presentation fonts |
| `font-size` | `12` | Positive pixels; tag height is `ceil(font-size × 1.2) + 4` |
| `color` | `#64748b` | Value text and, unless overridden, key text |
| `key-color` | `color` | Key text |
| `background-color` | `transparent` | Value-cell fill |
| `key-background-color` | `#f8fafc` | Key-cell fill |
| `border-color` | `#cbd5e1` | Cell borders, drawn with a fixed `0.75`-pixel stroke |
| `width` | auto | Total width applied to every tag |
| `key-width` | auto | Key-cell width applied to every tag |
| `gap` | `8` | Horizontal space between tags |
| `row-gap` | `4` | Space between wrapped rows and the same-sized inset from the selected top/bottom edge and both row ends |

Each `<entry>` also accepts `break-before="true|false"`. Its default is
`false`; `true` starts that entry on a new row when a preceding tag exists.

An entry-level `width` or `key-width` overrides its metadata-level default.
Omit either attribute for auto sizing; `auto` is descriptive, not a literal
V1 numeric value. A fixed total width must still leave positive key and value
cells. Colors accept `#RRGGBB` or `transparent`; fonts are `virgil`,
`helvetica`, `cascadia`, `assistant`, `excalifont`, `nunito`, `lilita-one`,
`comic-shanns`, or `liberation-sans`.

Tags retain source order. Greedy left-to-right packing fills each row within
`frame width - 2 * row-gap`, producing the minimum number of rows without
reordering; `break-before` can introduce an earlier row boundary. The usable
width must remain positive. Metadata `align` then positions each completed row
independently: left and right stop one `row-gap` inside the corresponding outer
frame edge, while center still uses the frame center.
This metadata page-edge inset is measured from the logical frame edge and is
independent of PPTX export `--paper-margin*` options or common-slide centering.

For a top band, its top edge is `frame.y + row-gap`; for a bottom band, its
bottom edge is `frame.y + frame.height - row-gap`. The metadata-side
reservation strip still spans the full frame width from the outer logical
frame edge to the corresponding boundary of the final content box. Its depth
is at least `row-gap + complete band height + 8` pixels; the complete band
height already includes the gaps between multiple rows. A closer content
boundary moves inward, while a boundary already farther inward is retained.
Normal items and text, local and UML connector paths and labels, and page-link
paths and labels never enter this strip; `overflow="visible"` does not override
the exclusion.
A page link without an explicit frame terminal remaps a reserved top/bottom
edge to the nearest safe edge and clamps a left/right terminal outside the
strip. An explicit `src-frame-side`, `dst-frame-side`, `src-frame-anchor`, or
`dst-frame-anchor` that selects the reservation is a validation error.
SVG, PPTX, PDF, Excel, and Excalidraw display the page-owned band; XYFlow and
Isoflow omit it as page decoration. See the
[frame metadata example](../examples/frame-metadata.md) for top, bottom,
left/right/center alignment, explicit row breaks, and auto or fixed widths.

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
