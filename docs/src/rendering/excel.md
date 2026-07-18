# Excel Output

Excel output creates one `.xlsx` workbook with one worksheet per identified
child frame, in source order.

```bash
xaligo render diagram.xal --format excel -o diagram.xlsx
```

`xlsx` is accepted as a format alias:

```bash
xaligo render diagram.xal --format xlsx -o diagram.xlsx
```

The alias is normalized by the CLI. In-repository API callers use
`FormatExcel` / `excel`.

Each worksheet contains the frame's rendered SVG image at `A1`, preserving its
intrinsic canvas size and aspect ratio without non-uniform scaling. Diagram
shapes are not converted into editable
spreadsheet cells. A one-frame document produces one worksheet. The page-frame
outline is omitted from the embedded SVG. The default image is strictly
cropped to the exact frame rectangle, so metadata on a top/bottom edge reaches
the corresponding image edge.

Worksheet names come from frame IDs. Excel-invalid control characters,
backslashes, and `:`, `/`, `?`, `*`, `[`, and `]` become `_`; a leading or
trailing apostrophe also becomes `_`. Names are limited to 31 UTF-16 code
units. Empty names fall back to `Frame <source-order>`, and case-insensitive
duplicates receive ` (2)`, ` (3)`, and so on.

Use `--combine-frames` to place the compatibility canvas in one worksheet:

```bash
xaligo render diagram.xal --format excel -o diagram.xlsx --combine-frames
```

Cross-frame connections retain their local page-link labels in the worksheet
images: `to <destination frame ID>` on the source frame and
`from <source frame ID>` on the destination frame. The shared scene resolves
independent item/frame anchors and the 4-layout-pixel terminal-label gap before
embedding either image.
