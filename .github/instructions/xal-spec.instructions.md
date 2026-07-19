---
applyTo: "**/*.xal"
---

# xaligo DSL (.xal) Specification

## Overview

`.xal` is a Vue-style layout DSL with XML syntax. Canonical V1 documents use a
`<xaligo>` envelope containing document-wide data and one `<frames>` page
collection. Historical `<frame>` and `<frames>` roots remain readable but emit
a migration warning.
The parser uses `encoding/xml` and handles attributes, nested tags, and text content.

## V1 Compatibility Profile and Version Boundary

Canonical V1 source explicitly sets `version="1"` on `<xaligo>`. An
unversioned `<xaligo>` defaults to V1 with a warning. A `version` value other
than `1` is invalid. Legacy `<frame>` and `<frames>` roots accept the historical
V1 version rules but always emit a warning recommending the canonical envelope.
This document-root `version` selects the DSL and is not visible page metadata.
By contrast, a non-empty `version` on an identified `<frame>` that is a direct
child of the document-root `<frames>` is that page's visible content revision;
it does not select a language version. Structural diff ignores only the
document-root DSL version and compares child-frame content revisions normally.

V2 uses a distinct, reject-safe root:

```xml
<scene version="2">
  ...
</scene>
```

`<scene>` requires `version="2"`; an unversioned `<scene>` is invalid. A V1
reader recognizes `<xaligo>`, `<frame>`, and `<frames>`, but rejects a
V2 document at the root instead of partially rendering V2 syntax as V1. Do not
use `<frame version="2">` or `<frames version="2">`.

A V2 implementation must accept this V1 profile as input, preserve its
defaults and compatibility behavior, and lower it directly to the shared typed
model. It must not rewrite V1 XML into V2 XML, parse the document twice, or
invoke V1 through a serialized intermediate representation. V1 has no
dependency on, and no obligation to understand, V2.

Canonical V1 source uses lowercase XML tag names, attribute names, and enum
tokens exactly as documented here. Historical case-insensitive or directional
aliases that are not listed in this specification are accepted implementation
details, not part of the frozen compatibility profile. A V2 compatibility
frontend canonicalizes the documented V1 values once at its input boundary.

## Root Tag

```xml
<xaligo version="1">
  <data>
    <!-- reusable definitions -->
  </data>
  <frames gap="48">
    <frame id="overview" width="1440" height="900" class="pa-4">
      ...
    </frame>
  </frames>
</xaligo>
```

`<xaligo>` permits document-level metadata, imports, data, and styles, and
requires exactly one `<frames>`. Give every child `<frame>` a stable `id`.

```xml
<xaligo version="1">
<frames gap="48">
  <frame id="overview" width="1440" height="900">
    ...
  </frame>
  <frame id="detail" width="1440" height="900">
    ...
  </frame>
</frames>
</xaligo>
```

| Attribute | Type | Default | Description |
|---|---|---|---|
| `version` | string | `"1"` with warning when omitted | On `<xaligo>` or a legacy root, selects V1 and only `"1"` is accepted. On an identified direct child `<frame>`, a non-empty value is the visible page revision |
| `title` | string | — | On a page `<frame>`, enables the metadata band and supplies its built-in `title` tag |
| `width` | float | `1280` | Frame width (px) |
| `height` | float | `720` | Frame height (px) |
| `class` | string | — | Spacing class |
| `layout` | string | — | Set to `"horizontal"` to arrange children horizontally |
| `gap` | float | `16` | Gap between child elements (px) |
| `item-size` | float | render-context default, normally `32` | Max icon size (px) applied to all `<item>` elements in this file. Overrides the native `item.icon_size` or embedded asset-source value |
| `margin` / `margin-*` | float | — | DSL content whitespace in pixels. On root `<frame>`, paper-frame size is preserved and content is inset. This is separate from PPTX CLI `--paper-margin*` flags, which are inch-based export fitting margins |
| `content-width` / `content-height` | float | — | Shrink usable inner layout area |
| `align` | string | — | Align usable content area (`top|middle|bottom` + `left|center|right`) |
| `overflow` | string | `error` | Child containment policy: `error` or `visible` |

Legacy input may still use root `<frame>` or `<frames>`. It remains renderable,
but diagnostics recommend wrapping identified frames in the canonical
`<xaligo version="1"><frames>...</frames></xaligo>` envelope.

`<frames>` accepts `gap` and optional `layout="vertical"`. Without
`layout="vertical"`, frames are arranged horizontally. A `<frame>` inside
`<frames>` requires a non-empty `id`.

### Frame and physical-page contract

An identified child `<frame>` is the V1 physical page unit. Frames are emitted
in source order after the complete document scene and all cross-frame links
have been resolved.

| Format | Default mapping |
|---|---|
| SVG | One `.svg` artifact per frame |
| PPTX | One slide per frame |
| PDF | One page per frame |
| Excel | One worksheet per frame, containing the frame's SVG image |
| Excalidraw, XYFlow, Isoflow | One logical document containing all frames |

SVG, PPTX, PDF, and Excel omit the page-frame outline in both default and
`--combine-frames` output. Frame geometry remains authoritative for page size,
cropping, endpoint ownership, and the outer logical page edge used to select a
cross-frame page-link side and tangent anchor. The drawable frame terminal may
sit on a parallel inward inset line. A default page-local SVG uses the exact
frame rectangle as its canvas and clip boundary; PDF pages and Excel page
images inherit that strict crop. Combined SVG compatibility output retains
marker-safe bounds expansion. Excalidraw retains editable frame structure with
transparent page-frame strokes.

For a document with one child frame, SVG writes exactly the requested output
path. For multiple child frames, an output request such as `diagram.svg`
produces `diagram-<safe-frame-id>.svg` for each frame. The safe ID retains ASCII
letters, digits, `_`, and `-`; every run of other characters becomes one `-`,
leading and trailing `-` are removed, and an empty result falls back to
`frame-<source-order>`. Two IDs that resolve to the same output filename are an
error. SVG does not create an implicit archive.

`--combine-frames` is the explicit compatibility option for page-oriented
formats. It restores the historical single canvas, single slide, single PDF
page, or single Excel worksheet. It does not change Excalidraw, XYFlow, or
Isoflow because those formats are already single logical documents.

### Frame metadata tag band

An identified page frame may expose `id`, `title`, a page-content `version`,
and arbitrary key/value entries as a two-cell tag band. The resolved metadata
`row-gap`, 4 layout pixels by default, is both the space between wrapped rows
and the metadata page-edge inset at the selected vertical edge and both
horizontal edges. Frame padding, content margins, and the content box never
replace or add to that inset. The band is enabled when the page frame has a
non-empty `title`, a child-frame content `version`, or a direct `<metadata>`
child. Existing identified frames that have only an `id` remain visually
unchanged. Once the band is enabled, non-empty built-ins are emitted in stable
`id`, `title`, `version` order, followed by `<entry>` children in source order.

```xml
<frame id="aws-architecture" title="AWS Architecture" version="1.0.0"
       width="720" height="400"
       margin-top="52" margin-right="24"
       margin-bottom="52" margin-left="24">
  <metadata position="top" align="right" font-family="helvetica">
    <entry key="owner" value="Platform Engineering" />
    <entry key="status" value="Approved" break-before="true"
           width="180" key-width="56" />
  </metadata>
  <rectangle id="diagram" title="Page content" />
</frame>
```

`<metadata>` is a non-layout direct child of a page `<frame>` and may occur at
most once. This context is distinct from document-level
`<xaligo><metadata>`. It contains only empty `<entry>` children; every entry
requires non-empty `key` and `value` attributes. Duplicate keys are retained.

| Attribute | Target | Default | Contract |
|---|---|---|---|
| `position` | `metadata` | `top` | Closed enum `top|bottom` |
| `align` | `metadata` | `left` | Closed enum `left|center|right`; applied independently to each resolved row |
| `font-family` | `metadata` | `virgil` | `virgil|helvetica|cascadia|assistant|excalifont|nunito|lilita-one|comic-shanns|liberation-sans` |
| `font-size` | `metadata` | `12` | Positive layout pixels; tag height is exactly `ceil(font-size * 1.2) + 4` |
| `color` | `metadata` | `#64748b` | Value text color; also key text color unless `key-color` is set |
| `key-color` | `metadata` | value of `color` | Key text color |
| `background-color` | `metadata` | `transparent` | Value-cell fill |
| `key-background-color` | `metadata` | `#f8fafc` | Key-cell fill |
| `border-color` | `metadata` | `#cbd5e1` | Cell border color; the cell stroke is fixed at `0.75` layout pixels |
| `width` | `metadata`, `entry` | auto | Positive total key/value tag width. An entry value overrides the metadata-level default |
| `key-width` | `metadata`, `entry` | auto | Positive key-cell width smaller than total width. An entry value overrides the metadata-level default |
| `gap` | `metadata` | `8` | Non-negative horizontal gap between tags |
| `row-gap` | `metadata` | `4` | Non-negative gap between wrapped rows and the same-sized inset from the selected top/bottom edge and both horizontal page edges |
| `break-before` | `entry` | `false` | Closed boolean `true|false`; `true` starts this entry on a new row when a preceding tag exists |

Colors use `#RRGGBB` or `transparent`. Auto width measures both cells with the
selected font and full-width-rune-aware metrics. Omit `width` or `key-width`
to request auto sizing; the literal string `auto` is not a V1 numeric value.
Fixed widths use no-wrap shrink-to-fit with clipping as the final overflow
guard. Tags preserve input order and use greedy left-to-right packing against
the usable width `frame.width - 2 * row-gap`, which produces the minimum row
count without reordering. The usable width must remain positive.
`break-before="true"` forces a row boundary before that custom entry. The
metadata `align` is then applied to each row separately against that same
usable width: left starts at `frame.x + row-gap`, right ends at
`frame.x + frame.width - row-gap`, and center still uses the frame center.

For `position="top"`, the band starts at `frame.y + row-gap`; for
`position="bottom"`, it ends at `frame.y + frame.height - row-gap`. The
metadata-side reservation strip spans the full frame width from the outer
logical frame edge to the corresponding boundary of the final content box.
Its depth is never less than `row-gap` plus the complete band height plus the
fixed 8-pixel content gap: if the normal content boundary is closer, it is
moved inward to that minimum; if it is already farther inward, it is retained.
Normal items and their text, local connector paths and labels, UML connector
paths and labels, and cross-frame page-link paths and labels cannot enter this
strip. `overflow="visible"` never overrides this page-decoration exclusion.
The frame's outer page size and invisible logical edge do not change.
The inset is measured from that logical frame edge before any common PPTX slide
centering and is unrelated to the export-only `--paper-margin*` options.
For a cross-frame page link, the same resolved `row-gap` is also the inward
normal inset of safe frame terminals on all four sides, independent of metadata
`position`; zero retains the outer edge. A frame without metadata instead uses
a 4-layout-pixel terminal inset. An explicit `src/dst-frame-side` or
`src/dst-frame-anchor` requires the inset to be strictly smaller than its
specified side's normal frame dimension: height for `top`/`bottom`, width for
`left`/`right`. Without an explicit frame terminal, validation requires at
least one side that satisfies this inset bound and the metadata reservation;
the shared scene later selects among those safe sides from rendered visual
geometry. The inset is never implicitly clamped.

The shared layout and presentation scene own this geometry. SVG, PPTX, PDF,
Excel, and Excalidraw render the owning frame's tags; per-frame projection
cannot leak another page's band, and combined output retains every band. The
entire reservation strip, rather than only the tag cells, is a hard exclusion
zone for normal rendered geometry. XYFlow and Isoflow omit the band because it
is page decoration rather than a graph node or endpoint.

## Numeric and Geometry Contract

Numeric attributes are validated before layout. A numeric value must be a
finite base-10 number; `NaN`, positive or negative infinity, an empty numeric
value, and malformed trailing text are errors. The current implementation
validates the source attributes and then reads those validated values during
layout; replacing the string attribute map with a typed normalized layout
specification is a separate roadmap item.

The following domain rules apply:

| Attributes | Required domain |
|---|---|
| `width`, `height`, `content-width`, `content-height`, `item-size`, `font-size`, `key-width` | greater than `0` when specified |
| `row`, `col` | greater than `0` when specified |
| `span` | greater than `0` and at most `12`; flexible sibling spans in one `<row>` must total at most `12` |
| `gap`, `row-gap`, margins, spacing-class padding | greater than or equal to `0` |
| `scale`, `coordinate-scale`, `grid`, `stroke-width` | greater than `0` when specified |
| `x`, `y`, `dx`, `dy`, bend coordinates | any finite value, subject to the containing geometry rule |

An omitted attribute uses its documented default. An explicitly empty
`align` is treated as omitted; it must not produce an invalid-alignment warning.
Unknown non-empty enum values remain errors or source-positioned warnings as
specified by that attribute.

V1 intentionally distinguishes strict values from compatibility fallbacks:

| Input | V1 behavior |
|---|---|
| Invalid `overflow`, connection side, or connection anchor | Validation error |
| Unknown `layout`, connection `kind`, stroke style, arrowhead, or arrowhead-size value | Validation error |
| Unknown render mode, format, theme, paper/orientation, arrow-style option, or SVG legend position | Render-option error. The CLI normalizes `xlsx` to `excel` before validation |
| Recognized but unavailable render mode (`aws-2.5d` or `topology`) | Not-implemented error |
| Empty `align` | Omitted; defaults to `top-left` |
| Malformed or unknown non-empty `align` | Warning; each unsupported component keeps its `top` or `left` default |
| Unknown nested attribute or malformed/unrecognized spacing-class token | Ignored; a recognized numeric negative spacing class remains an error |

These fallbacks are part of V1 compatibility, not a mechanism for opting into
V2. The distinct V2 root prevents new V2 constructs from being silently
treated as V1 extensions.

`validate` and every render format use the same normalized values and resolved
geometry checks. Successfully validated input must not later produce `NaN`,
`Inf`, a negative drawable size, or an output serialization error caused
by geometry.

### Fixed and flexible child allocation

For a vertical parent, an explicit child `height` is a fixed main-axis size;
for a horizontal parent, an explicit child `width` is fixed. The parent first
reserves fixed sizes, margins, and gaps. Children without a fixed main-axis size
divide the remaining space using their positive `row` or `col` weights. A
`<row>` uses validated `span` values against its 12-column grid.

The resolved child size is the size used both for recursive layout and for
placing the next sibling. A child cannot replace its assigned size after the
parent cursor has advanced. Explicit cross-axis sizes must fit the parent's
content box unless overflow is explicitly allowed.

Layout parents accept `overflow`:

| Value | Behavior |
|---|---|
| `error` | Default. A child outside the parent's content box is a source-positioned validation error. |
| `visible` | The child may extend outside the content box, but all coordinates and sizes must remain finite and sibling cursors still use resolved sizes. |

The policy belongs to a parent and applies only to its direct children; it is
not inherited. If fixed children consume the full main axis under `visible`,
the parent's original usable extent is used as the flex pool and the flexible
children receive their weighted sizes while all children retain source order.
Sibling cursors use each resolved size, gap, and margin, making the resulting
overflow explicit. Under the default `error` policy the same layout is
rejected.

Overflow is never silently introduced by a renderer. Clipping is a drawing and
text policy and does not make invalid layout geometry valid.

## Layout Tags

### `<container>`

Stacks children **vertically** (same behavior as `frame`). Use `layout="horizontal"` for horizontal arrangement.

```xml
<container class="pa-4" gap="16">
  ...
</container>
```

| Attribute | Type | Default | Description |
|---|---|---|---|
| `layout` | string | — | `"horizontal"` to arrange children side by side |
| `gap` | float | `16` | Gap between child elements (px) |
| `content-width` / `content-height` | float | — | Shrink usable inner layout area |
| `align` | string | — | Align usable content area |
| `overflow` | string | `error` | Child containment policy: `error` or `visible` |

### `<row>`

Lays out children **horizontally** in a 12-column grid.

```xml
<row gap="20">
  <col span="8">...</col>
  <col span="4">...</col>
</row>
```

| Attribute | Type | Default | Description |
|---|---|---|---|
| `gap` | float | `16` | Column spacing (px) |
| `overflow` | string | `error` | Child containment policy: `error` or `visible` |

> `<row>` is a **pure layout tag** — it does not render any border or label in the output.
> The `<col>` children are also pure layout containers.

An explicit child `width` is reserved before the grid share and is excluded
from `span` allocation. Among children without fixed width, an omitted `span`
defaults to `12 / number_of_flexible_children`; explicit flexible spans must
total at most `12`. Unused span leaves intentional trailing space.

### `<col>`

A vertical stack container inside `<row>`. Use `span` to set the number of columns occupied.

| Attribute | Type | Default | Description |
|---|---|---|---|
| `span` | float | `12 / num_columns` | Columns to occupy (out of 12) |
| `class` | string | — | Spacing class |
| `overflow` | string | `error` | Child containment policy: `error` or `visible` |

## Custom Leaf and Container Tags

An otherwise unknown nested tag with no layout children is a generic leaf and
is rendered as a rectangle plus text. An unknown nested tag with layout
children is a generic group/container: it receives the normal group header
insets and lays out those children vertically by default, horizontally for
`layout="horizontal"`, or with the V1 staggered layout for
`layout="staggered"`. If every child is item-like (`item`, `spacer`, or
`blank`), the children use the item-grid row behavior instead.

This rule applies only below a valid V1 root. An unknown root is always a parse
error, so `<scene version="2">` can never be mistaken for a generic V1 group.

```xml
<card title="Dashboard" />
<panel title="Main Chart" />
<text>Any label</text>
```

| Attribute | Behavior |
|---|---|
| `title` | Display label (takes priority) |
| Text content | Label when `title` is absent |
| (none) | Tag name used as label |
| `border` | Set to `"none"` to hide the border |
| `visible` | Set to `"false"` to hide only this component (border, icon, label). Children are still rendered individually. Layout space is preserved |
| `font-size` | Text font size in layout pixels |

## `<rectangle>` and `<port>` Tags

`<rectangle>` creates a general-purpose rectangle. Its label comes from
`title` or direct text content, and `font-size` controls the label size.
Unlike generic leaf tags, `<rectangle>` may contain multiple `<port>` children.

`<port>` creates a small rectangle inside a side of the parent rectangle.
Multiple ports on the same side are spaced evenly along that side. Its label
also comes from `title` or direct text content, and it supports `font-size`.

```xml
<rectangle id="service" title="Service" width="180" height="100" font-size="18">
  <port id="service-in" side="left" title="in" font-size="9" />
  <port id="service-out" side="right" title="out" font-size="10" />
</rectangle>
```

| Attribute | Target | Description |
|---|---|---|
| `id` | `rectangle`, `port` | Required connection reference ID |
| `width` / `height` | `rectangle`, `port` | Size in layout pixels |
| `title` / text content | `rectangle`, `port` | Text rendered inside the shape |
| `font-size` | `rectangle`, `port` | Text font size in layout pixels |
| `side` | `port` | Parent side: `top`, `right`, `bottom`, or `left`. Default `top` |
| `x` / `y` | `port` | Optional position relative to the parent rectangle's top-left corner. Values are clamped so the port remains inside the parent rectangle |

Port boxes must remain inside their parent rectangle. Explicit positions are
normalized before drawing, and overlapping ports on the same side are a layout
diagnostic rather than a renderer-specific accident. Port text carries the
shared text-layout policy: SVG, its PDF/Excel projections, and PPTX enforce it,
while editable Excalidraw-compatible output preserves it in metadata for
bound-text consumers.

## Resolved Text Layout

Text has both a geometry box and a semantic role. Scene and plan construction
must preserve the resolved role, wrapping, fitting, clipping, line height, and
padding instead of making each encoder infer them from generated IDs.

Built-in defaults are:

| Role | Wrap | Fit | Overflow |
|---|---|---|---|
| group header | no | shrink | clip to text box |
| ordinary label | yes | shrink | clip to text box |
| item label | yes | shrink | clip to text box |
| port label | yes | shrink | clip to port box |
| connector label | yes | shrink | clip to text box |

The default line-height multiplier is `1.2` unless the source scene carries a
valid positive value. Font sizes originate in layout pixels and are converted
with the same effective scale as the containing geometry. Changing
`--px-per-inch` or paper fitting therefore preserves the text-to-shape ratio.

An encoder may use native text fitting or deterministic line breaking, but the
visible result must obey the resolved policy. Editable Excalidraw-compatible
bound text carries the same `xaligoTextLayout` metadata and must not become a
separate layout authority. Encoders apply text policy in this order: resolve
padding, wrap when enabled, shrink when requested, then clip when
`TextLayout.overflow="clip"`.

## `<table>` Tag

`<table>` represents general tabular data and is distinct from the `<grid>`
layout primitive and relational `<entity>` definitions. It accepts either a
GFM-like pipe table or explicit `<header>`/`<row>` children containing
`<cell>` elements. One pipe header may be followed by explicit rows; a second
header is invalid. Both syntaxes normalize to the same typed rows before
layout.

```xml
<table id="services" title="Services">
  | Service | Role    | Port |
  |:--------|:--------|-----:|
  | API     | Backend | 8080 |
  <row><cell>DB</cell><cell>Database</cell><cell align="right">5432</cell></row>
</table>
```

Pipe separators require at least three hyphens per cell and use colons for
left, center, and right alignment. A literal pipe is escaped as `\|`. Every
header and row must have the same positive cell count. Tagged cell `align`
accepts `left`, `center`, `right`, or the normal vertical-horizontal values.
Unknown children, duplicate headers, malformed separators, and mismatched
column counts are positioned errors.

Table presentation attributes inherit from `<table>` to every row and cell.
Tagged `<header>`, `<row>`, and `<cell>` attributes override inherited values.
Pipe cells have no inline attributes, so style them with table attributes and
the `header-*` variants:

```xml
<table color="#172033" background-color="#ffffff" border-color="#94a3b8"
       font-family="nunito" font-size="16"
       header-color="#ffffff" header-background-color="#2563eb"
       header-font-family="cascadia" header-font-size="18">
  | Service | Port |
  |:--------|-----:|
  | API     | 8080 |
</table>
```

| Attribute | Values | Description |
|---|---|---|
| `color` | `#RRGGBB` or `transparent` | Cell text color |
| `background-color` | `#RRGGBB` or `transparent` | Cell fill color |
| `border-color` | `#RRGGBB` or `transparent` | Cell border color |
| `font-family` | `virgil`, `helvetica`, `cascadia`, `assistant`, `excalifont`, `nunito`, `lilita-one`, `comic-shanns`, or `liberation-sans` | Cell font family |
| `font-size` | positive number | Cell font size in layout pixels |
| `header-*` | corresponding value above | Pipe/tag header override declared on `<table>` |

Colors require six-digit hexadecimal notation. The style precedence is
`cell > header/row > table > built-in default`. Font-family names are carried
through the renderer-neutral scene and mapped to the corresponding output font
face; an output environment may substitute a missing installed font.

## Relational Database Tags

Reusable `<database-schema id="...">` definitions belong under `<data>`.
Frames render them through `<database data="schema-id">`. A schema contains
identified `<entity>` definitions; each entity contains typed `<column>`
definitions and optional single-column `<foreign-key>` definitions.

```xml
<database-schema id="app">
  <entity id="roles"><column name="id" type="bigint" primary-key="true" /></entity>
  <entity id="users">
    <column name="role_id" type="bigint" nullable="false" />
    <foreign-key columns="role_id" references="roles.id" />
  </entity>
</database-schema>
```

Columns support `name`, `type`, `primary-key`, `nullable`, `unique`, and
`default`. A foreign key requires one local column and one
`references="entity.column"` target and generates a relation between the
entities. Duplicate or missing schema/entity/column references and mixed
inline/data-backed database content are positioned errors. Composite keys,
indexes, checks, and import dialects remain planned V1 extensions.

## UML Tags

`<uml>` is the common V1 component for the fourteen UML diagram families. It
adapts their typed elements, compartments, and relations to xaligo's shared
layout, shape, connector, and output pipeline. Supporting a family selector
does not imply support for every UML 2.x glyph or interchange construct; the
closed V1 vocabulary and its projection limits are defined below.

### Component, identity, and layout contract

```xml
<xaligo version="1">
  <data>
    <uml-model id="domain-model">
      <class id="order" title="Order">
        <attribute>- id: UUID</attribute>
        <operation>+ confirm()</operation>
      </class>
      <interface id="repository" title="OrderRepository">
        <operation>save(order: Order)</operation>
      </interface>
      <realization src="order" dst="repository" title="persists" />
    </uml-model>
  </data>
  <frames>
    <frame id="domain" width="960" height="540">
      <uml id="model" title="Domain Model">
        <class-diagram data="domain-model" direction="right" />
      </uml>
    </frame>
  </frames>
</xaligo>
```

The following rules are normative:

- `<uml>` must be inside a frame, requires a non-empty `id`, and contains
  exactly one supported diagram-kind child. UML IDs must be unique within that
  frame. The same UML ID may be reused in a different frame.
- UML component IDs and diagram-local element IDs must not contain whitespace,
  `.` or `/`. `.` is reserved for the frame boundary; `/` is rejected so V1
  sources do not use path-like UML endpoint expressions.
- The diagram-kind child contains a non-empty set of direct element and
  relation children. Unknown diagram kinds and unknown children are errors;
  arbitrary custom tags are not generic UML elements. The activity-diagram
  exception is `<partition>`, which groups activity nodes into swimlanes.
- Every UML element requires a non-empty diagram-local `id`, unique within the
  UML component. A UML relation's `src` and `dst` use those local IDs, without
  a UML or frame prefix, and both endpoints must exist in the selected model.
- `direction` on the diagram-kind child accepts only `right` or `down`. When
  `<uml layout>` is omitted, `direction="right"`, sequence diagrams, and timing
  diagrams select horizontal xaligo layout; the other cases select vertical
  layout. This controls the V1 projection and is not a UML semantic ordering
  rule.
- The `<uml>` container is semantic and does not render an outer border or
  title for any UML diagram family. Visible diagram titles, versions, owners,
  and review state belong in frame metadata. Element labels resolve in the
  order `title`, `name`, direct text, then local `id`. UML elements default to
  `font-family="helvetica"` and `font-size="14"`;
  normal element font attributes override those defaults. An element `name`
  is display text only and never becomes a frame-level connection alias; use
  the public UML reference described below.
- User-facing Excalidraw export is disabled for UML inputs. Rendering a source
  containing `<uml>` with `--format excalidraw` or the public render dispatcher
  returns an error; use SVG, PDF, PPTX, Excel, XYFlow, or Isoflow for UML
  diagrams until editable UML Excalidraw output is redesigned.
- The compatibility tags `<element>` and `<relation>` are not part of the
  strict V1 UML profile. A model must use one of the element and relation tags
  allowed for its selected family.

UML elements are also normal xaligo connection endpoints. A frame-level
`<connection>` uses the following public references; the internal hex-scoped
scene ID is opaque and must not be written in source:

| Location | Public endpoint reference | Meaning |
|---|---|---|
| Same frame | `local-id` | UML element `local-id`; the ID must be unique among frame-level connection references |
| Another frame | `frame-id.local-id` | The same UML element reached across a frame boundary |

For example, element `order` in `<uml id="model">` inside frame `overview` is
`order` to a normal connection in that frame and `overview.order` to a normal
connection in another frame. Omitting the `frame-id.` prefix for a cross-frame
endpoint is an unresolved-reference error. If another UML element, rectangle,
port, table, database, entity, group, or other connection endpoint in the same
frame exposes the same public ID/reference, parsing fails with a duplicate frame
reference error. UML-native relations continue to use `src="order"`, not either
public connection form.

### Diagram-kind vocabulary

| Diagram kind | Allowed elements | Allowed relations | Additional V1 semantic checks |
|---|---|---|---|
| `class-diagram` | `package`, `class`, `interface`, `enumeration` | `association`, `aggregation`, `composition`, `generalization`, `realization`, `dependency` | Requires one classifier. `package` groups classifiers and package-local relations using the general-group visual language. `grid="N"` on `<class-diagram>` or package selects the package/classifier column count. When the class diagram contains only packages and omits `grid`, the package grid is computed from frame aspect ratio and empty-cell count. Packages expand to their assigned grid cell. Classifiers accept `stereotype`, `abstract="true|false"`, and `static="true|false"`; stereotypes and enabled modifiers are visible classifier-header metadata. Classifiers use compact multi-row placement, a cyan header with white text, separate attribute/operation body regions, white bodies, and deep-blue relation lines by default. Aggregation/composition are class to class; generalization joins equal classifier kinds; realization is class to interface. |
| `object-diagram` | `object` | `link`, `dependency` | Requires one object. Every relation endpoint is an object. |
| `component-diagram` | `component`, `interface`, `port`, `artifact` | `dependency`, `realization`, `association`, `assembly`, `delegation` | Requires one component. Component boxes render with a cyan header, white body, left-aligned white component name, and no two-rectangle adornment. A component may contain only child `interface` entries; each child interface renders as a small white boundary port box with the interface text inside and outside circle symbols, clears the component header, and sits mostly inside the component with only a small outside protrusion. Matching component associations keep the interface user endpoint on the source component's nearest non-interface-side anchor, selected from 15 top/right/bottom component anchors, add a left-bulging caller-side semicircle as a forked connection endpoint, bind the destination endpoint to a same-named interface circle, and approach the circle horizontally from the outside-left side of left-side interface circles at the circle center height; the line ends at the semicircle's left bend, and the semicircle radius is 2px larger than and center-aligned with the interface circle. When multiple associations target the same component interface, the destination renders one circle per incoming association with enough vertical spacing to prevent caller-side semicircle overlap and groups them back to the interface box with a bracket-style stem. Component diagrams use compact grid placement by default. Realization is component to interface. Assembly uses port/interface endpoints and includes a port. A port requires a component owner and renders on that owner boundary. Delegation is port to component/port. |
| `deployment-diagram` | `node`, `artifact`, `component` | `deployment`, `communication-path`, `dependency` | Requires one node. Deployment is artifact/component to node; communication-path is node to node. |
| `package-diagram` | `package`, `class`, `interface`, `component` | `dependency`, `package-import`, `package-merge` | Requires one package. Import/merge are package to package. |
| `composite-structure-diagram` | `structure`, `collaboration`, `part`, `port`, `component` | `connector`, `assembly`, `delegation`, `dependency` | Requires a part or port. Parts/ports require typed owners. Connector endpoints are parts/ports, assembly is port to port, and delegation starts at a port. |
| `profile-diagram` | `profile`, `stereotype`, `metaclass` | `extension`, `reference`, `generalization` | Requires a profile and stereotype. Extension is stereotype to metaclass; generalization is stereotype to stereotype. |
| `use-case-diagram` | `actor`, `use-case`, `system-boundary` | `association`, `include`, `extend`, `generalization` | Requires a use case. Association joins actor and use-case; include/extend join use-cases; generalization joins equal actor/use-case kinds. |
| `activity-diagram` | `initial`, `final`, `activity`, `action`, `decision`, `merge`, `fork`, `join`, `object-node`; optional `partition` containers | `control-flow`, `object-flow` | Requires an activity/action. Control-flow excludes object-node. Object-flow requires an object-node endpoint. Initial/final direction and control-node degrees are validated. `lanes="vertical|horizontal"`, `theme="xaligo"`, partition swimlanes, and loop flow metadata are supported. |
| `state-machine-diagram` | `initial`, `final`, `state`, `history`, `choice`, `fork`, `join`; optional layout-only `container` with `row` and `col` children | `transition` | Requires a state. Initial/final direction and pseudostate degrees are validated. `state` compartments render state actions and activities: `entry`, `do`, `exit`, `internal`, `region`, and `note`. The state name renders in a cyan header with white text; the white body is split into rows, and each row has a key/value column divider. Transitions render `event [guard] / action-or-effect` labels. Initial/final/choice/history pseudostates keep compact proportions; final states render the standard inner dot. A state-machine `<container>` may group elements by child `<row>` and `<col>` tags; these layout-only tags assign grid rows and columns without becoming UML elements. When no container column is supplied, the row layout reuses nearby connected columns where possible before assigning the next free column. State-machine shapes use the class-diagram xaligo palette by default: deep-blue borders/text/relations, white bodies, cyan state-name headers, and a deep-blue initial dot. Relation child `<bend x="..." y="..." />` tags are supported on class, activity, and state-machine relation tags and are used to steer orthogonal connector routes. State-machine transition routing treats intermediate state and pseudostate bodies as obstacles where possible, keeps same-frame route points inside the frame bounds, and may choose larger outside detours for distant states or bent routes that would otherwise cross a state body. UML relation labels are shifted away from endpoint items when a default label position would overlap them. |
| `sequence-diagram` | `participant`, `lifeline` | `message`, `return-message`, `create-message`, `destroy-message` | Requires a participant/lifeline. Participants and lifelines render as top headers with dashed vertical lifeline axes, not full container boxes. Every message has a diagram-unique order that controls top-to-bottom anchoring. `message` is synchronous by default; `message mode="async"` renders an asynchronous open-arrow call. `message`, `create-message`, and `destroy-message` draw destination-lifeline activation bars; `return-message` is a dashed response connector and does not start a new activation. Non-participant sources must already be active before sending a non-return message or return. Activation bars extend through the related messages up to the matching return from that lifeline; a fully contained activation is merged into its covering bar. `destroy-message` also draws a stop marker at the destination lifeline and its label must clearly describe destruction, deletion, disposal, removal, or termination. Create/destroy cannot be self messages. |
| `communication-diagram` | `object`, `participant` | `link`, `message` | Requires two participants, one link, and one message. Every message has a unique order and matching unordered link pair. |
| `interaction-overview-diagram` | `initial`, `final`, `interaction`, `decision`, `fork`, `join` | `control-flow` | Requires an interaction. Initial/final direction and control-node degrees are validated. |
| `timing-diagram` | `lifeline`, `time-state` | `transition`, `occurrence`, `duration` | Requires a lifeline and time-state. Time-state intervals do not overlap per owner. Transition joins chronological states of one owner; occurrence has `at`; duration joins time-states. |

The endpoint contracts above are closed. An admitted relation with an endpoint
pair not described by its row is a validation error.

### Activity partitions and swimlanes

An `activity-diagram` may use `lanes="vertical"` or `lanes="horizontal"` to
render responsibility swimlanes. With lanes enabled, `theme="xaligo"` applies the
xaligo logo palette: `#08b8ea` for lane headers and primary actions,
`#052d6e` for activity text/borders, and `#04b79f` reserved for follow-up
object-flow/success accents. The only accepted V1 values are
`lanes="vertical|horizontal"` and `theme="xaligo"`; other values are errors.

Each `<partition>` must be a direct child of `<activity-diagram>` and must
have a stable `id` and non-empty `title`. A partition contains only activity
element children, not relations. Relations remain direct children of the
diagram and keep using normal local element IDs:

```xml
<activity-diagram direction="down" lanes="vertical" theme="xaligo">
  <partition id="customer" title="Customer">
    <initial id="start" />
    <action id="enter-pin" title="Enter PIN" tone="primary" />
  </partition>
  <partition id="atm" title="ATM">
    <action id="request-pin" title="Request PIN" />
    <decision id="pin-valid" title="PIN valid?" />
  </partition>
  <control-flow src="enter-pin" dst="request-pin" />
  <control-flow src="pin-valid" dst="request-pin" guard="invalid PIN" route="loop" />
</activity-diagram>
```

A direct activity element may instead use `lane="partition-id"` to join a
declared partition. A nested element may omit `lane` or repeat the enclosing
partition ID; any other `lane` value is a validation error. Partition IDs use
the same local UML identifier restrictions as element IDs and must be unique
within the activity diagram.

Partitioned activities with `lanes="vertical"` are laid out with equal-width
vertical lanes and a top lane-header band. With `lanes="horizontal"`, partitions
become equal-height horizontal lanes with a left lane-header band. In both
orientations, each node is centered within its owning lane. `tone="primary"` on
`activity` or `action` uses the primary xaligo fill and white text. The partition
ID and title are retained in editable-scene custom data. The UML activity
container border and automatic diagram-kind title are omitted because page frame
metadata supplies the visible page title. Normal frame-level
`<connection>` elements may target activity elements by public UML reference
such as `action-id`, including `frame-id.action-id` for cross-frame page links.

### Ownership

`owner` is a same-diagram local element reference. Forward references are
allowed because ownership is resolved after all elements are collected, but
the referenced element must exist and have an allowed kind.

| Element and diagram | `owner` | Allowed owner kinds |
|---|---|---|
| `component-diagram/port` | required | `component` |
| `composite-structure-diagram/part` | required | `structure`, `component`, `collaboration` |
| `composite-structure-diagram/port` | required | `structure`, `part`, `component`, `collaboration` |
| `use-case-diagram/use-case` | optional | `system-boundary` |
| `timing-diagram/time-state` | required | `lifeline` |

Every other use of `owner` is invalid. Ownership is retained as semantic
metadata and a stable reference. The V1 shared layout is flat and does not
promise that an owned shape is spatially nested inside its owner.

### Element compartments

An element's direct child tags are ordered text compartments. Each compartment
must have non-whitespace direct text, `title`, or `name` and must not contain
child elements. Nested UML elements and relations are not compartments. The
typed compartment vocabulary is:

| Element | Allowed typed compartments |
|---|---|
| `class` | `attribute`, `operation`, `constraint`, `note` |
| `interface` | `operation`, `constraint`, `note` |
| `enumeration` | `literal`, `operation`, `note` |
| `object` | `slot`, `note` |
| `component` | `interface` |
| `node`, `artifact` | `property`, `responsibility`, `note` |
| `package` | `responsibility`, `note` |
| `structure` | `property`, `provided-interface`, `required-interface`, `note` |
| `part` | `property`, `responsibility`, `note` |
| `profile` | `constraint`, `note` |
| `stereotype` | `property`, `constraint`, `note` |
| `metaclass` | `property`, `note` |
| `actor` | `responsibility`, `note` |
| `use-case`, `activity`, `action` | `responsibility`, `constraint`, `note` |
| `state` | `entry`, `do`, `exit`, `internal`, `region`, `note` |
| `interaction` | `note` |
| `time-state` | `region`, `constraint`, `note` |

Elements absent from this table do not accept compartments. The generic
`<compartment>` child is a compatibility spelling accepted wherever a typed
compartment is allowed; new source should use the typed tag because its meaning
survives future semantic processing. Compartment source order is preserved,
but compartments are not independent connection endpoints.

### Relation attributes, order, and time

Every relation requires `src` and `dst`. `title` or `label` supplies its
visible text. When omitted, `event` supplies the label. `guard` is allowed only
on flows/transitions and is appended as `[guard]`. `effect` or `action` is
appended as `/ effect-or-action`, with `effect` taking priority when both are
present. `src-multiplicity` and `dst-multiplicity` are allowed only on
association, aggregation, composition, and link and are appended in
source-to-destination order. Relation color and normal connector side, anchor,
stroke-width, bend, scale, and grid attributes use the `<connection>` rules.
`kind`, `stroke-style`, `arrowhead`, `start-arrowhead`, and `end-arrowhead` are
invalid because the UML relation kind owns line and marker semantics.
`route="loop"` is allowed only on `control-flow` and `object-flow`; it is
retained as UML route metadata for renderers and editors.

Sequence and communication message kinds require `order`. Its canonical form
is one or more positive decimal integers without leading zeroes, separated by
dots, for example `1`, `2`, or `1.1`. The complete order string must be unique
across all messages in one diagram. Numeric order is prepended to the rendered
label and assigns top-to-bottom connector anchors on participant/lifeline
shapes. It does not reorder declared elements or create activation boxes or a
separate interaction axis. Sequence message anchors always use a vertical
element edge so the ordering remains vertical: explicit `top` is normalized to
`left`, explicit `bottom` to `right`, and an explicit anchor slot is superseded
by the normalized order position.

In a communication diagram, a message is valid only when a `<link>` connects
the same two endpoints. Link direction is ignored for this structural check,
so a link `a -> b` also supports a message `b -> a`.

Timing attributes use finite base-10 numbers in one caller-chosen unit. Unit
suffixes such as `20ms` are invalid; put the unit in a label instead.

| Timing construct | Required attributes | Domain |
|---|---|---|
| `time-state` | `owner`, `from`, `to` | `owner` is a lifeline; `from >= 0`; `to > from` |
| `occurrence` | `src`, `dst`, `at` | `at >= 0`; both endpoints are lifelines or time-states |
| `duration` | `src`, `dst`; optional `from` and `to` pair | Both endpoints are time-states; when supplied, `from >= 0` and `to > from` |

### Relation projection

UML relations lower to the shared orthogonal connector model with the
following fixed semantic defaults:

| Projection | Relation kinds |
|---|---|
| Dashed line with destination triangle | `dependency`, `realization`, `package-import`, `package-merge`, `reference`, `include`, `extend`, `return-message`, `deployment` |
| Solid line with destination triangle | `generalization`, `control-flow`, `object-flow`, `transition`, `message`, `create-message`, `destroy-message`, `delegation`, `extension` |
| Source diamond | `aggregation`, `composition` |
| No destination arrowhead | `association`, `link`, `occurrence`, `duration`, `communication-path`, `assembly`, `connector` |

The visible relation label is placed near the routed connector midpoint and
the UML diagram/relation kind is retained as semantic metadata where the target
format can carry it.

### Reusable UML models

Reusable definitions use `<uml-model id="...">` directly below document
`<data>`. A diagram-kind child selects one with `data="model-id"`:

```xml
<data>
  <uml-model id="order-objects">
    <object id="customer" title="customer: Customer">
      <slot>name = Alice</slot>
    </object>
    <object id="order" title="order42: Order">
      <slot>status = Confirmed</slot>
    </object>
    <link src="customer" dst="order" title="placed" />
  </uml-model>
</data>
<frames>
  <frame id="snapshot">
    <uml id="runtime"><object-diagram data="order-objects" direction="right" /></uml>
  </frame>
</frames>
```

`<uml-model>` requires a document-unique ID. A missing model, duplicate model
ID, or a diagram that combines `data` with inline children is an error. The
model itself does not declare a UML family; after expansion, all of its element,
compartment, ownership, relation, order, and endpoint rules are validated
against the selecting diagram kind. One reusable model is therefore reusable
across selectors only when every child belongs to each selector's closed
vocabulary.

### Deliberately lossy V1 projection

V1 preserves the selected UML family, element kind, relation kind, owner, and
relation label in the shared semantic scene, then projects them into the
capabilities common to xaligo outputs:

- `use-case`, `initial`, and `final` become ellipses;
- `decision`, `merge`, `choice`, and `history` become diamonds;
- every other element becomes an editable rectangle whose ordered
  compartments are flattened into its visible text;
- every relation becomes a shared orthogonal connector with a separate label;
  aggregation and composition currently share the same diamond projection;
- sequence order is retained in labels and metadata and controls top-to-bottom
  message anchors, but V1 does not draw dashed lifelines, activations, combined
  fragments, or a separate vertical event axis;
- timing intervals and occurrences are validated and retained, but V1 does not
  draw a continuous time axis, waveforms, or proportional time geometry; and
- `owner` records semantic containment without requiring spatial nesting.

State-machine diagrams may set `show-element-names="false"` to hide the
compartment element names such as `entry`, `do`, `exit`, `internal`, `note`,
and `region` without hiding state titles, compartment values, transition
labels, or the underlying shapes. The value is inherited by state-machine
elements unless an element sets its own `show-element-names` value. Other UML
diagram families do not accept the diagram-level attribute.

SVG, Excalidraw, PPTX, PDF, Excel, XYFlow, and Isoflow all consume this same resolved
geometry. Excalidraw-compatible output carries xaligo UML custom data for
editing. XYFlow retains UML node and relation fields in node/edge `data` and
records the projected node shape. Isoflow projects every connected UML shape
to a labeled generic endpoint icon because its upstream connector schema has no
arbitrary UML data field. Another target schema may omit marker details it
cannot represent. An encoder must use native target constructs where available
and must not add private schema-breaking fields. The output is not XMI and is not a
lossless UML interchange representation.

## `<item>` Tag

A leaf element that places an AWS service icon inside a container.
Specify a positive signed 32-bit decimal ID from `service-catalog.csv` as the
`id` attribute (`1` through `2147483647`).
The icon is rendered to fit within the specified size (`item-size`).

The effective item size is resolved from the root `item-size` when present;
otherwise it comes from the render context. Native configuration and the
canonical embedded-asset profile default to `32` layout pixels. Callers that
provide a custom asset source may intentionally choose another value. For
cross-environment reproducibility, declare `item-size="32"` (or another fixed
positive value) on the document root.

```xml
<public-subnet id="public-subnet" title="Public Subnet">
  <item id="1178" />   <!-- with icon -->
  <item />             <!-- spacer: no icon, only a layout slot -->
  <item id="1189" />   <!-- with icon -->
</public-subnet>
```

| Attribute | Type | Required | Description |
|---|---|---|---|
| `id` | positive int32 | — | Decimal service ID `1..2147483647` from `service-catalog.csv`. Omitted or empty → treated as spacer; zero, signs, non-decimal syntax, and out-of-range values are invalid |
| `dx` / `dy` | float | — | Relative icon offset in pixels from the icon's normal layout `x,y` position. The moved icon rectangle must remain inside the parent frame/group border |

> If no icon is found for the given `id`, rendering skips the item and emits a warning rather than failing the document.

## `<spacer>` / `<blank>` Tags

Dedicated empty layout tags, usable as alternatives to `<item />`.
They occupy layout slots but render no icon, label, border, or text.

```xml
<public-subnet id="public-subnet" title="Public Subnet">
  <item id="1178" />
  <spacer />          <!-- empty slot: no icon -->
  <blank />           <!-- empty slot: no icon -->
  <item id="1189" />
</public-subnet>
```

No attributes (`id` is ignored if specified).

## `<connection>` Tag

Draws an **elbowed arrow** between `<item>` elements or group borders.
Must be written as a direct child of `<frame>` or inside a frame-level
`<connections>` tag, **outside** layout tags.
Use the same catalog IDs as `<item id="N">`, or assign `id`, `name`, or `ref`
to an AWS/group tag, for `src` / `dst`.

```xml
<frame width="1122" height="794" class="pa-4">
  <aws-cloud id="cloud" title="AWS Cloud">
    <public-subnet id="public" title="Public Subnet">
      <item id="1178" />
      <item id="1189" />
    </public-subnet>
  </aws-cloud>

  <!-- connections go last, as direct children of frame or inside <connections> -->
  <connections grid="8">
    <connection src="1178" dst="1189" />
    <connection src="public" dst="cloud" kind="route" />
  </connections>
</frame>
```

### `<connections>` Tag

`<connections>` is an optional direct child of `<frame>` that groups
`<connection>` elements and provides shared defaults. It does not render a
shape or occupy layout space. Any per-connection attribute overrides the parent
default.

Only these non-empty group attributes are inherited:
`arrowhead-size`, `kind`, `color`, `stroke-width`, `width`, `stroke-style`,
`start-arrowhead`, `end-arrowhead`, `arrowhead`, `scale`,
`coordinate-scale`, and `grid`. Endpoint identity and geometry are deliberately
not inherited: every child must supply its own `src` and `dst`, and
`src-side`, `dst-side`, `src-anchor`, `dst-anchor`, `src-frame-side`,
`dst-frame-side`, `src-frame-anchor`, `dst-frame-anchor`, bends, points, and
via data remain child-local. Defaults are applied to a connection snapshot
during scene construction; the parsed child node is not mutated.

`stroke-width`/`width`, `end-arrowhead`/`arrowhead`, and
`coordinate-scale`/`scale` are semantic alias pairs. If a child supplies either
name, neither name is inherited from the parent. When the child itself supplies
both, the first canonical name in each pair takes precedence.

`<connections>` may contain only `<connection>` children. A misspelled or
otherwise unknown child is a validation error rather than a silently skipped
connector.

```xml
<connections kind="traffic" color="#2563eb" grid="8" scale="1">
  <connection src="web" dst="app" />
  <connection src="app" dst="db" color="#059669" />
</connections>
```

| Attribute | Type | Required | Description |
|---|---|---|---|
| `src` | string | ✓ | Catalog ID, or `id`/`name`/`ref` of the arrow start item, AWS group, rectangle, port, or identified child frame |
| `dst` | string | ✓ | Catalog ID, or `id`/`name`/`ref` of the arrow end item, AWS group, rectangle, port, or identified child frame |
| `src-side` / `dst-side` | string | — | Optional endpoint side: `top`, `right`, `bottom`, or `left` |
| `src-anchor` / `dst-anchor` | string | — | Optional edge anchor. Each side has five inset positions (`top-1` through `top-5`, etc.) for 20 unique perimeter anchors |
| `src-frame-side` / `dst-frame-side` | string | — | Cross-frame-only logical page side, independent of the endpoint side; the drawable terminal uses that side's inward inset line |
| `src-frame-anchor` / `dst-frame-anchor` | string | — | Cross-frame-only logical page side and tangent slot. Uses `top|right|bottom|left-1..5`, or a side plus numeric/named slot; inward inset does not change the slot coordinate |
| `arrowhead-size` | string | — | V1 fixed arrowhead size: `"s"` (small). This is the default; `m` and `l` are not V1 values because V1 cannot preserve them across all render formats |
| `kind` | string | — | `connection` for the normal connector, `route` for a structural path without arrows, or `traffic` for directional flow drawn beside a matching route |
| `color` | `#RRGGBB` | — | Six-digit hexadecimal stroke color override. Named, short, and alpha colors are invalid in V1 so every format preserves the same color |
| `stroke-width` / `width` | float | — | Positive stroke width override; `width` is the compatibility alias |
| `stroke-style` | string | — | `solid`, `dashed`, or `dotted` |
| `start-arrowhead` / `end-arrowhead` | string | — | Independently set either end to `none`, `arrow`, `triangle`, `stealth`, `diamond`, or `oval`. An effective `kind="route"` permits only `none` |
| `arrowhead` | string | — | Backward-compatible alias for `end-arrowhead`; an effective route permits only `none` |
| `bends` / `points` / `via` | string | — | Backward-compatible inline coordinate list. Prefer child tags for multiple bend coordinates |
| `scale` / `coordinate-scale` | float | — | Positive multiplier applied to bend coordinates before routing. Default `1` |
| `grid` | float | — | Positive per-connection snap grid in layout pixels. Defaults to the router grid |

Default connections and `kind="traffic"` use a thin 1px line with
`start-arrowhead="none"` and a slender `stealth` end arrowhead. `kind="route"`
uses `start-arrowhead="none"` and `end-arrowhead="none"` by default. Default
colors are `#1e1e1e` for normal connections, `#64748b` for routes, and
`#2563eb` for traffic. A route is always headless in V1: after applying
`<connections>` defaults and child alias overrides, any effective non-`none`
`start-arrowhead`, `end-arrowhead`, or `arrowhead` is a source-positioned
validation error. Explicit `none` is accepted. Explicit `stroke-width`, color,
and stroke style are preserved for every kind; non-route arrowhead attributes
are also preserved.

For SVG, PPTX, PDF, and Excel Plan output, the render option `arrow-style`
supplies the global arrowhead (and, for `thin`/`standard`, width) only when the connection
does not explicitly set that semantic value. Explicit DSL or inherited group
values take precedence, and `kind="route"` remains headless. Excalidraw,
XYFlow, and Isoflow V1 output consume the resolved DSL scene rather than this
Plan-only option.

When a connection references endpoints in different frames, the shared scene
represents it as a page link instead of drawing one line across the inter-frame
gap. SVG, PPTX, PDF, Excel, and Excalidraw derive exactly two local stubs:

- the source stub runs from the source endpoint to the page-terminal inset line
  of its owning frame and has the exact label `to <destination frame ID>`; and
- the destination stub runs from the page-terminal inset line of its owning
  frame to the destination endpoint and has the exact label
  `from <source frame ID>`.

Angle brackets in those forms are rendered as literal punctuation. For a
connection from frame `overview` to frame `detail`, the visible strings are
therefore `to <detail>` and `from <overview>`.

Endpoint binding and logical frame-terminal geometry are separate. The
endpoint uses `src-anchor`/`dst-anchor`, then `src-side`/`dst-side`, then its
normal automatic binding. The logical page terminal uses
`src-frame-anchor`/`dst-frame-anchor`, then
`src-frame-side`/`dst-frame-side` as fixed choices. With neither frame-terminal
attribute, the legacy endpoint anchor, endpoint side, or normal nearest-border
result is only the preferred page side. The renderer keeps it when safe;
otherwise it chooses the nearest safe side from the endpoint's rendered visual
envelope. An item's envelope is the union of its icon and external label; other
endpoints use their rendered shape. Distance ties prefer a tied side facing the
remote frame, then `top`, `right`, `bottom`, `left`.

UML relations use a shared endpoint anchor profile across Excalidraw, SVG,
PPTX, PDF, and Excel plan geometry. Sequence-diagram messages keep their
chronological lifeline position. Other UML rectangle-like endpoints snap to the
five fixed inset anchors on each side (`side-1` through `side-5`), giving 20
perimeter anchors in total. UML diamond-like endpoints such as `choice`,
`decision`, `merge`, and `history` snap to the four shape vertices; an explicit
anchor on a diamond chooses the side vertex rather than an inset side slot.

A side is safe when the resolved inset fits its normal frame dimension, it is
not the metadata edge, and an actual `top`/`bottom` terminal opposite metadata
does not enter the reservation strip. Validation of an automatic page terminal
checks only that this candidate set is non-empty. It must not infer the chosen
automatic side from layout `Box` geometry; final selection belongs to shared
scene construction after icon and label geometry is available. If the normal
preferred side is unsafe, scene construction remaps it to the nearest safe
candidate. No safe candidate is a source-positioned validation error at the
connection.

The endpoint- and frame-terminal-adjacent segments are perpendicular to their
own selected sides, so an endpoint may leave on `right` while its local stub
terminates at the page's bottom inset line. Frame-side and frame-anchor
attributes are valid only when the resolved endpoints belong to different
frames. Using any of them on a same-frame connection is a source-positioned
validation error.

Frame metadata reservation is a final safety constraint on that choice. For an
automatic page terminal, the metadata edge and any other unsafe side are
removed before the renderer's nearest-side choice. A terminal on a safe left
or right edge is clamped along that edge so it lies outside the top/bottom
reservation strip; any resulting coordinate difference is bridged
orthogonally. An explicit frame side or anchor that selects the reserved edge,
or an exact left/right anchor whose point lies inside the strip, is a
source-positioned validation error instead of being moved. Page-link paths and
labels remain outside the full strip.

When an explicit frame side is vertically opposite the metadata edge, its
actual terminal must remain outside the reservation strip. For bottom metadata
with explicit side `top`, the actual top terminal may not enter below the
strip's top boundary. For top metadata with explicit side `bottom`, the actual
bottom terminal may not enter above the strip's bottom boundary. A violation is
a source-positioned validation error at the connection. For an automatic page
terminal, the same conflict makes that candidate unsafe instead of immediately
rejecting the connection. A safe explicit `left` or `right` terminal remains
valid even if a hypothetical top/bottom inset line would enter the strip.

The parallel coordinate is resolved against the selected outer logical frame
edge before applying the normal inset. An explicit frame anchor keeps its exact
10/30/50/70/90-percent coordinate along the outer frame extent. An automatic
terminal's unconstrained parallel coordinate comes from the endpoint binding.
If it enters a 24-layout-px corner gutter, the parallel coordinate is clamped
and a two-bend orthogonal dogleg bridges the difference; a border shorter than
96 layout pixels uses one quarter of its length as an adaptive gutter. A
left/right terminal is also subject to the metadata reservation clamp described
above. Automatic left/right coincidence avoidance normally intersects that
corner-gutter range with an 8-layout-pixel clearance from the reservation. If a
very small non-reserved range cannot satisfy both preferences, it falls back to
the entire non-reserved interval, may touch its boundary, and never moves a
point outside the frame or inside the metadata strip.

The drawable terminal then lies on a page-terminal inset line parallel to that
outer edge. Let `i` be the resolved metadata `row-gap` when the frame has
metadata, or 4 layout pixels when it does not. The same `i` applies to every
terminal side regardless of metadata `position`; `i = 0` retains the outer
edge. An explicit `top`/`bottom` frame side requires `i < frame.height`; an
explicit `left`/`right` side requires `i < frame.width`. Failure is a
source-positioned validation error at that connection. For an automatic page
terminal, those inequalities classify candidates instead; only an empty safe
candidate set is an error. The resolved `i` is used exactly and is not reduced
to fit. With the resolved parallel coordinate represented by `u` for a
horizontal side or `v` for a vertical side, the terminal is:

```text
top:    (u, frame.y + i)
right:  (frame.x + frame.width - i, v)
bottom: (u, frame.y + frame.height - i)
left:   (frame.x + i, v)
```

The inset step changes only the normal coordinate. An explicit frame anchor
therefore retains its tangent slot and uses its local orthogonal stub for
visible separation. If an unconstrained final inset terminal would coincide
with the endpoint binding, its parallel coordinate moves by up to 24 layout
pixels within the available range so the stub remains visible. Manual bends do
not alter either local stub's geometry; bends remain logical routing metadata
for graph adapters.

There is one strict zero-inset case. When metadata is enabled with resolved
`row-gap="0"`, an endpoint resolves to its owning frame itself, and its explicit
frame anchor coincides with the resolved endpoint point, the connection is a
source-positioned validation error. An explicit endpoint anchor supplies that
point directly. An explicit endpoint side uses its center (`top` is `top-3`,
and likewise for the other sides); with neither endpoint attribute, the
automatically resolved endpoint side also uses its center. Fixed parallel
coordinates, perpendicular segments at both ends, and a visible local stub
cannot all be satisfied at that coincident point. The author must select a
different endpoint or frame anchor, or use a positive metadata `row-gap`.

Each `to <...>` / `from <...>` label is placed from the final inset terminal
with a 4-layout-pixel inward gap and a minimum 4-layout-pixel tangent gap.
Placement chooses the closest tangent position that avoids the endpoint
envelope and metadata reservation. Tiny pages clamp or shrink the label
fallback instead of moving it farther inward from that terminal.

Both scene stubs carry the same logical connector ID, original endpoint/frame
IDs, and V1 routing metadata. XYFlow and Isoflow use those fields to emit one
logical edge instead of two partial edges.

Default page-oriented export projects only the local stub belonging to each
frame: the source SVG/slide/page/worksheet contains `to <destination frame
ID>`, and the destination one contains `from <source frame ID>`.
`--combine-frames` places both local stubs on the compatibility canvas but
never reconnects them across the frame gap. Excalidraw also retains both stubs
in its one editable scene.

Output formats are projections of this resolved V1 meaning. A target schema
may not have fields for every V1 connector value; the upstream-compatible
Isoflow connector schema, for example, has no arbitrary metadata field. Such
adapters must use native constructs where available and must not add private,
schema-breaking fields. A V2 compatibility frontend consumes V1 directly and
must never use an output format as an intermediate representation.

When `src-side`, `dst-side`, `src-anchor`, and `dst-anchor` are omitted,
endpoint sides and anchor positions are calculated automatically from endpoint
geometry. Use `src-anchor` and `dst-anchor` to pin an endpoint to a specific
perimeter anchor. Cross-frame `src-frame-anchor` and `dst-frame-anchor` use the
same grammar to select the logical page side and tangent slot independently.
Each side has five positions at 10, 30, 50, 70, and 90 percent of the outer
frame extent, giving 20 unique tangent anchors. The drawable frame terminal
then moves only in the inward normal direction to the page-terminal inset line.
Corner anchors are not shared: `top-1` keeps a horizontal coordinate near the
left corner, while `left-1` keeps a vertical coordinate near the top corner.

```text
top:    top-1    top-2    top-3    top-4    top-5
right:  right-1  right-2  right-3  right-4  right-5
bottom: bottom-1 bottom-2 bottom-3 bottom-4 bottom-5
left:   left-1   left-2   left-3   left-4   left-5
```

Position numbers run left-to-right on `top` and `bottom`, and top-to-bottom on
`left` and `right`. Anchor positions are `1` through `5` from top/left to
bottom/right on the named side, inset from corners so each side owns its five
positions.
The aliases map one-to-one as `start=1`, `near=2`, `center=3`, `far=4`, and
`end=5`.

```xml
<connection src="web" dst="app"
            src-anchor="right-3"
            dst-anchor="left-3" />
<connection src="web" dst="app"
            src-side="right" src-anchor="3"
            dst-side="left" dst-anchor="3" />

<!-- The item and logical page terminal may use different sides. -->
<connection src="web" dst="detail.app"
            src-side="right" src-anchor="near"
            src-frame-side="bottom" src-frame-anchor="far"
            dst-side="left" dst-anchor="far"
            dst-frame-side="top" dst-frame-anchor="near" />
```

`src` and `dst` can also be expressed as child tags when the endpoint reference
and anchor should be declared together. The endpoint token can be tag text or
one of `id`, `ref`, `name`, or `target`.

```xml
<connection kind="traffic">
  <src anchor="right-3" frame-side="bottom" frame-anchor="far">web</src>
  <dst side="left" anchor="5" frame-anchor="top-2" ref="detail.app" />
</connection>
```

On `<src>` and `<dst>`, `frame-side` and `frame-anchor` map to the corresponding
source/destination connection attributes. A complete anchor such as
`bottom-4` supplies its side. With a separate side, slots accept `1..5` or the
aliases `start`, `near`, `center`, `far`, and `end`. Conflicting side and
complete-anchor values are validation errors for both endpoint and frame
anchors.

Excalidraw output always serializes arrowhead sizes as the smallest supported
size (`"s"`) to keep dense diagrams readable. The logical arrowhead type and
style metadata are still stored on the connector and used by SVG/PPTX export
and the SVG-based PDF/Excel projections.

Manual bend coordinates are expressed as child tags in the same Cartesian
layout coordinate space as the frame, with the origin at the upper-left of the
rendered frame and positive `x`/`y` extending right/down. SVG and PPTX route
calculations keep the connector orthogonal while forcing the route through each
listed bend in order. Excalidraw output stores the routing metadata on the
arrow; Excalidraw's own editor may still display its editable elbow connector
approximation.

```xml
<connection src="web" dst="db"
            scale="1" grid="8">
  <bend x="120" y="80" />
  <bend x="120" y="220" />
  <bend x="300" y="220" />
</connection>
```

`<point>`, `<via>`, and `<waypoint>` are accepted aliases for `<bend>`.
Coordinates can also be grouped inside `<bends>`, `<points>`, or `<path>`.

Items and group tags may define a connection reference with `id`, `name`, or
`ref`:

```xml
<item id="1178" name="web" />
<item id="1189" name="db" />
<vpc id="prod-vpc" />
web --- db
web ==> db
prod-vpc --- web
```

- `---` expands to `kind="route"`.
- `==>` expands to `kind="traffic"`.
- Operands may also be numeric item IDs or group IDs.
- Explicit `<connection src=... dst=...>` attributes resolve the same way.
- Shorthands must be direct text children of `<frame>`.
- References must be unique and must belong to an item or group with a
  non-empty ID.
- Use an explicit `<connection>` for color, width, or stroke overrides, and for
  arrowhead overrides on normal connections or traffic flows. Routes remain
  headless.

**Arrow spec:**
- `elbowed: true` — always right-angle connectors (Excalidraw "elbow connector")
- Arrowhead at end only by default. Excalidraw stores this as
  `endArrowhead: "arrow"` plus `endArrowheadSize: "s"`; xaligo metadata records
  the logical PPTX/SVG head as `stealth`.
- Stroke color `#1e1e1e`, stroke width `1px` for normal connections
- `kind="route"` defaults to `#64748b`, `1px`, lower route layer, no arrowheads
- `kind="traffic"` defaults to `#2563eb`, `1px`, higher traffic layer, directional end arrowhead
- A traffic line with the same endpoints as a route line is drawn beside that
  route in Excalidraw, SVG, and PPTX draw paths when possible.
- Start/end connect to the **edge midpoint** of the element
  - When direction is **downward**: label text element (`{id}-item-lbl`) bottom edge
  - Otherwise: icon image element (`{id}-item`) corresponding edge
- Edges are fixed with normalized coordinates via `fixedPoint`, so arrows snap correctly when the file is opened
- Arrow ID format: `conn-{src}-{dst}-{index}`
- Arrow ID is registered in `boundElements` of the bound elements
- Excalidraw item icons and labels are grouped with a 5x5 white anchor grid.
  Anchor grid cells are drawn above connectors and below the item content so
  lines do not cover icons/labels while endpoints remain visible.
- Excalidraw routing uses previously placed lines to offset exact or near-exact
  lane overlaps. Group header tags, item icons, and labels are treated as
  routing obstacles where possible.
- SVG/PPTX routing may additionally add automatic junction markers and line
  jump masks after the Excalidraw scene is built. PDF and Excel inherit the SVG
  projection. These are export-layer
  rendering features, not extra `.xal` tags.

**Edge selection logic:**

| Direction (dst relative to src) | Start edge | End edge |
|---|---|---|
| Right (dx ≥ dy) | right | left |
| Left | left | right |
| Down (dy > dx) | bottom (label) | top |
| Up | top | bottom (label) |

> If `src` / `dst` items are not rendered, a warning is emitted and the connection is skipped.

Connection endpoints must resolve to exactly one item, AWS group, rectangle,
port, or identified child frame. Numeric catalog IDs are valid only when that
ID appears once in the document; when the same service appears multiple times,
use unique `name` or `ref` values. Missing endpoints, ambiguous numeric IDs,
duplicate aliases, and `<connection>` tags nested below any tag other than
`<frame>` or its direct `<connections>` child are validation errors.

## AWS Group Tags

Like `container`, these stack children vertically, but are rendered with **AWS architecture diagram group border styles**.
Templates are in `etc/resources/aws/templates/excalidraw/` (`.excalidraw`) and `etc/resources/aws/templates/xal/` (`.xal`).
Icon SVGs are sourced from `etc/resources/aws/svg/Architecture-Group-Icons/`.

```xml
<aws-cloud id="production" title="Production Environment">
  <vpc id="vpc-main" title="vpc-0a1b2c3d">
    <private-subnet id="private-a" title="Private Subnet A">
      <card title="App Server" />
    </private-subnet>
  </vpc>
</aws-cloud>
```

| Tag | Display Name | Border Color | Style | Icon |
|---|---|---|---|---|
| `<aws-cloud>` | AWS Cloud | `#000000` | solid | AWS-Cloud-logo_32.svg |
| `<aws-cloud-alt>` | AWS Cloud | `#000000` | solid | AWS-Cloud_32.svg |
| `<region>` | Region | `#00A1C9` | dashed | Region_32.svg |
| `<availability-zone>` | Availability Zone | `#00A1C9` | dashed | — |
| `<security-group>` | Security group | `#CC0000` | dashed | — |
| `<auto-scaling-group>` | Auto Scaling group | `#E7601B` | dashed | Auto-Scaling-group_32.svg |
| `<vpc>` | Virtual private cloud (VPC) | `#8C4FFF` | solid | Virtual-private-cloud-VPC_32.svg |
| `<private-subnet>` | Private subnet | `#00A1C9` | solid | Private-subnet_32.svg |
| `<public-subnet>` | Public subnet | `#3F8624` | solid | Public-subnet_32.svg |
| `<server-contents>` | Server contents | `#7A7C7F` | solid | Server-contents_32.svg |
| `<corporate-data-center>` | Corporate data center | `#7A7C7F` | solid | Corporate-data-center_32.svg |
| `<ec2-instance-contents>` | EC2 instance contents | `#E7601B` | solid | EC2-instance-contents_32.svg |
| `<spot-fleet>` | Spot Fleet | `#E7601B` | solid | Spot-Fleet_32.svg |
| `<aws-account>` | AWS account | `#E7008A` | solid | AWS-Account_32.svg |
| `<aws-iot-greengrass-deployment>` | AWS IoT Greengrass Deployment | `#3F8624` | solid | AWS-IoT-Greengrass-Deployment_32.svg |
| `<aws-iot-greengrass>` | AWS IoT Greengrass | `#3F8624` | solid | — |
| `<elastic-beanstalk-container>` | Elastic Beanstalk container | `#E7601B` | solid | — |
| `<aws-step-functions-workflow>` | AWS Step Functions workflow | `#E7008A` | solid | — |
| `<generic-group>` | Generic group | `#AAB7B8` | dashed | Configurable with `icon-id` |

All AWS group tags require a non-empty `id`. IDs for group tags, `<rectangle>`,
and `<port>` must be unique among frame-like components. Group tags otherwise
accept the same attributes as `container` (`title`, `class`, `gap`, etc.).

`generic-group` additionally accepts `icon-id`, a positive signed 32-bit
decimal ID (`1..2147483647`) from `service-catalog.csv`. Zero, signs,
non-decimal syntax, and out-of-range values are invalid. It uses the same
embedded AWS, Tabler, and Yamaha icon catalog as `<item>` and renders a 32px
icon to the left of the title.
This matches the built-in group icon size. Every group header receives an
opaque mask matching its local background behind the icon and label, preventing
solid or dashed border strokes from crossing the header content.
Group header tag labels use the shared single-line text policy. The tag
background and label box use a conservative width estimate so no-wrap text
remains inside the tag in SVG and PowerPoint. Keep group tag text concise; if
changing group tag font, padding, or geometry, update the shared text-layout
policy, renderer width estimate, and regression tests together.
East Asian full-width characters, including Japanese labels, count as
double-width in group header and item label width estimates.

```xml
<generic-group id="network-topology" title="Network Topology" icon-id="104635">
  <item id="200036" />
</generic-group>
```

### Layout Control Attributes (shared by all containers)

Available on `frame` / `container` / `col`, all AWS group tags, and unknown
child-bearing container tags where noted.

| Attribute | Value | Description |
|---|---|---|
| `layout` | `"horizontal"` | Arrange children **horizontally** with proportional widths (use the `col` attribute for ratio) |
| `layout` | `"staggered"` | Stack children with a depth offset (AWS group tags and unknown child-bearing containers) |
| `gap` | float | Child spacing (px). Default `16` |
| `align` | `"{vertical}-{horizontal}"` | Position of content area and `<item>` icons. Item grids also support `spread`. Default item-grid alignment is `"middle-center"` |
| `content-width` / `content-height` | float | Shrink usable inner layout area, leaving whitespace |
| `width` / `height` | float | Fixed child size (root frame dimensions remain the paper/content frame) |
| `overflow` | `"error"` \| `"visible"` | Child containment policy. Default `error` |

**`align` values** — combine a vertical part and a horizontal part with `-`:

| Part | Values |
|---|---|
| vertical | `top` \| `middle` \| `bottom` |
| horizontal | `left` \| `center` \| `right` \| `spread` |

All 12 combinations are valid: `top-left`, `top-center`, `top-right`, `top-spread`, `middle-left`, `middle-center`, `middle-right`, `middle-spread`, `bottom-left`, `bottom-center`, `bottom-right`, `bottom-spread`.

> **`center` (default):** icons are packed together and the group is centred within the available area
> (equivalent to CSS `justify-content: center`).
>
> **`spread`:** icons are distributed with equal gaps between each icon and the container edges
> (equivalent to CSS `justify-content: space-evenly`).
>
> **`left` / `right`:** icons are packed at the respective edge with a fixed `8 px` gap between icons.

```xml
<!-- Icons centred vertically and horizontally inside the group (default) -->
<private-subnet id="app-tier" title="App Tier" align="middle-center">
  <item id="27" />
  <item id="547" />
</private-subnet>

<!-- Icons spread evenly across the full width -->
<generic-group id="global-services" title="Global Services" align="middle-spread">
  <item id="1179" />
  <item id="1178" />
  <item id="216" />
  <item id="227" />
</generic-group>

<!-- Icons pinned to the top-left -->
<generic-group id="security-services" title="Security" align="top-left">
  <item id="216" />
  <item id="227" />
</generic-group>
```

### Child Size Ratio Attributes

| Attribute | Direction | Description |
|---|---|---|
| `row` | Vertical (`layoutStack`) | **Height ratio** among children without explicit `height`. Default `1.0` |
| `col` | Horizontal (`layout="horizontal"`) | **Width ratio** among children without explicit `width`. Default `1.0` |

```xml
<!-- Horizontal: left 2 : right 1 width ratio -->
<vpc id="vpc-main" title="VPC" layout="horizontal">
  <public-subnet id="public-subnet" title="Public" col="2" />
  <private-subnet id="private-subnet" title="Private" col="1" />
</vpc>

<!-- Vertical: top 1 : bottom 2 height ratio -->
<region id="region-main" title="Region">
  <vpc id="vpc-a" title="VPC A" row="1" />
  <vpc id="vpc-b" title="VPC B" row="2" />
</region>
```

## Spacing Classes (`class` attribute)

Vuetify-style notation. **Unit: `spacingUnit = 8px`**.

### All-sides shorthand

| Class | Meaning |
|---|---|
| `pa-{n}` | padding all sides = n × 8px |
| `ma-{n}` | margin all sides = n × 8px |

### Axis shorthand

| Class | Meaning |
|---|---|
| `px-{n}` | padding left + right = n × 8px |
| `py-{n}` | padding top + bottom = n × 8px |
| `mx-{n}` | margin left + right = n × 8px |
| `my-{n}` | margin top + bottom = n × 8px |

### Per-side

| Class | Meaning |
|---|---|
| `pt-{n}` | padding-top |
| `pr-{n}` | padding-right |
| `pb-{n}` | padding-bottom |
| `pl-{n}` | padding-left |
| `mt-{n}` | margin-top |
| `mr-{n}` | margin-right |
| `mb-{n}` | margin-bottom |
| `ml-{n}` | margin-left |

Multiple classes are space-separated: `class="pa-4 mt-2"`

### Semantics

| Kind | Target tag | Behavior |
|---|---|---|
| `padding` | frame / container / col | Inner whitespace. Child layout starts `pad` pixels inward |
| `padding` | AWS group tags / unknown containers | **Added to** `defaultGroupTopInset(44)` / `defaultGroupSideInset(12)`. `pa-2` adds +16px below the header |
| `margin` | any child element | Read by the parent layout (`layoutStack` / `layoutRow`) and used as inter-sibling spacing (equivalent to CSS flex margin) |

## Layout Calculation Rules

1. Normalize and validate numeric attributes and enum values.
2. Resolve each parent's border box and content box after margin and padding.
3. `frame` / `container` / `col` → **vertical stack**: reserve fixed child
   heights, gaps, and margins, then divide the remainder by `row` weights.
4. `layout="horizontal"` → reserve fixed child widths, gaps, and margins, then
   divide the remainder by `col` weights.
5. `row` → **12-column grid** after validating each `span` and their total.
6. Leaf elements use the resolved `(x, y, w, h)` received from their parent;
   they do not replace the allocation after sibling placement.
7. Verify finite positive geometry and parent-content containment before scene
   construction. Respect only an explicit `overflow="visible"` exception.
8. Resolve item grids against the same occupied content area before encoding.

## Example

```xml
<frame width="1440" height="900" class="pa-4">
  <container class="pa-4">
    <row gap="20" class="mb-2">
      <col span="8" class="pa-2">
        <card title="Dashboard" />
      </col>
      <col span="4" class="pa-2">
        <card title="Summary" />
      </col>
    </row>

    <row gap="20">
      <col span="4" class="pa-2">
        <panel title="Filters" />
      </col>
      <col span="8" class="pa-2">
        <panel title="Main Chart" />
      </col>
    </row>
  </container>
</frame>
```

## Constraints and Notes

- The canonical root is `<xaligo version="1">`. Legacy `<frame>` and
  `<frames>` roots are accepted with a warning. Direct children of `<frames>`
  must be identified `<frame>` tags. V2 uses `<scene version="2">`, which is
  intentionally rejected by V1.
- Both self-closing (`<card title="..." />`) and regular (`<card title="..."></card>`) forms are supported.
- The sum of `span` values in direct children of `<row>` must not exceed 12.
  Excess is a validation error rather than implicit overflow to the right.
- `.xal` files must be saved in UTF-8.
