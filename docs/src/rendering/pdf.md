# PDF Output

PDF output creates one document whose pages follow identified child frames in
source order.

```bash
xaligo render diagram.xal --format pdf -o diagram.pdf
```

Each frame is first rendered through the shared SVG/page plan and then placed
as vector artwork on one PDF page. The PDF page uses the SVG's intrinsic canvas
dimensions, including marker-safe overflow padding, and converts them to
physical units at 96 pixels per inch without non-uniform scaling. A one-frame
document produces one page. The logical frame defines page projection but its
outline is not drawn.

Use `--combine-frames` to preserve the compatibility canvas as one PDF page:

```bash
xaligo render diagram.xal --format pdf -o diagram.pdf --combine-frames
```

Cross-frame connections remain page links. The source page contains
`to <destination frame ID>` and the destination page contains
`from <source frame ID>`; no line is drawn through an inter-page gap.
