# PDF Output

PDF output creates one document whose pages follow identified child frames in
source order.

```bash
xaligo render diagram.xal --format pdf -o diagram.pdf
```

Each frame is first rendered through the shared SVG/page plan and then placed
as vector artwork on one PDF page. The PDF page uses the SVG's intrinsic canvas
dimensions and converts them to physical units at 96 pixels per inch without
non-uniform scaling. Default page SVGs use the exact frame rectangle as their
canvas and clip boundary, so the PDF page has no marker-safe overflow padding
and a top/bottom metadata band's `row-gap` gutter reaches the physical edge
while its tag cells remain inset by that value. A one-frame document produces
one page. The logical frame defines page projection but its outline is not
drawn.

Use `--combine-frames` to preserve the compatibility canvas as one PDF page:

```bash
xaligo render diagram.xal --format pdf -o diagram.pdf --combine-frames
```

The combined compatibility page inherits the SVG renderer's marker-safe canvas
expansion.

Cross-frame connections remain page links. The source page contains
`to <destination frame ID>` and the destination page contains
`from <source frame ID>`; no line is drawn through an inter-page gap. Explicit
frame-side/frame-anchor geometry and the 4-layout-pixel terminal-label gap are
resolved before the PDF projection.
