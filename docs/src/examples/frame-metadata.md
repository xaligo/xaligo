# Frame Metadata Tags

This two-page example adds visible page metadata without drawing a frame
outline. Each key/value band touches the selected top/bottom edge of the outer
frame border box and uses the full frame width for row wrapping and alignment;
padding, content margins, and the content box do not inset it. A full-width
reservation strip extends from that edge to the final content-box boundary and
is at least the band height plus 8 pixels deep. Normal items, text, connector
paths and labels, and page links stay outside it. The logical frame edge still
controls page size, cropping, and safe page-link terminals.

The `aws-architecture` page demonstrates automatic sizing and an explicit row
break with restrained default styling:

- omitting `position` selects the top;
- omitting the metadata-level `width` and `key-width` measures every tag;
- the default font size is `12`, with `#64748b` text, transparent value cells,
  `#f8fafc` key cells, and `#cbd5e1` borders;
- built-in `id`, `title`, and `version` tags precede `owner` on the first row;
- `status` uses `break-before="true"` and its own fixed width on the second row;
  and
- `align="right"` positions both completed rows independently.

Despite the frame's padding and side margins, both rows use the outer frame
width: the band's top edge touches the frame top and `align="right"` ends each
row at the frame's outer right edge. The top reservation continues to the final
content box, keeping normal content and connections below it.

![Top frame metadata with automatic widths](../images/frame-metadata-aws-architecture.svg)

The `release-notes` page uses one bottom row with `align="center"`. Its
metadata-level `width="138"` and `key-width="56"` apply to all tags, while
`status` overrides both. It also customizes the font, font size, and subtle text,
key-cell, and border colors. The row's bottom edge touches the outer frame
bottom, and the full-width bottom reservation remains free of normal geometry.

![Bottom frame metadata with fixed widths](../images/frame-metadata-release-notes.svg)

Tags always retain input order and pack greedily into the minimum number of
rows for the available width unless an entry requests a break. The cross-frame
connection omits `src-side` and `dst-side` deliberately. Although its source
port faces the reserved top edge and its destination port faces the reserved
bottom edge, page-link safety remaps them to the nearest safe edges. The
page-local `to <release-notes>` and `from <aws-architecture>` paths and labels
never enter either reservation strip; left/right terminals would likewise be
clamped beyond the strip.

Source:

```xml
{{#include samples/frame-metadata.xal}}
```

Validate and render the page-local SVG files with:

```bash
xaligo validate docs/src/examples/samples/frame-metadata.xal
xaligo render docs/src/examples/samples/frame-metadata.xal \
  --format svg \
  -o output/frame-metadata.svg
```

The render command writes:

```text
output/frame-metadata-aws-architecture.svg
output/frame-metadata-release-notes.svg
```

PPTX, PDF, and Excel preserve the same source order as two slides, pages, or
worksheets:

```bash
xaligo render docs/src/examples/samples/frame-metadata.xal --format pptx -o output/frame-metadata.pptx
xaligo render docs/src/examples/samples/frame-metadata.xal --format pdf -o output/frame-metadata.pdf
xaligo render docs/src/examples/samples/frame-metadata.xal --format excel -o output/frame-metadata.xlsx
```

Use `--combine-frames` only when one compatibility canvas, slide, PDF page, or
worksheet is required. Excalidraw always keeps both editable page objects;
XYFlow and Isoflow retain the logical connection but omit the tag cells because
they are page decoration.
