# PowerPoint Output

PPTX output is generated from a Go draw plan and written by the configured
statically linked Rust `pptx` exporter.

```bash
xaligo render diagram.xal --format pptx -o diagram.pptx \
  --services services.csv \
  --paper A3 --orientation landscape \
  --paper-margin-top 0.75 --paper-margin-bottom 0.75
```

Each identified child frame becomes one diagram slide in source order. Because
PowerPoint presentations use one common slide size, xaligo uses the largest
resolved frame-page width and height and centers smaller pages without scaling
them independently. A document with one frame still produces one slide.

Use `--combine-frames` to preserve the former one-canvas, one-diagram-slide
layout:

```bash
xaligo render diagram.xal --format pptx -o diagram.pptx --combine-frames
```

The page frame is a logical slide boundary and is not drawn as a visible
rectangle. Its edge supplies each cross-frame side and frame-anchor tangent
coordinate; the drawable stub terminal is on the parallel inward inset line.
For a smaller frame centered on the common slide, that inset is measured from
the logical frame edge before slide centering, not from the physical slide edge.

PPTX-specific behavior:

- Paper fitting supports named paper sizes and per-side margins.
- Routing avoids icon and label obstacles where possible.
- Existing routed lanes influence later paths to reduce overlap.
- Anchor backgrounds are drawn before lines, while icons and labels are drawn
  above lines.
- Line-jump masks are drawn at interior crossings.
- Legend slide(s) are added after all frame/diagram slides when `--services`
  is provided.
- Cross-frame connections are page links: the source slide contains
  `to <destination frame ID>` and the destination slide contains
  `from <source frame ID>`. Endpoint and logical frame anchors can select
  different sides. The terminal inset is the resolved metadata `row-gap`, or
  4 layout pixels without metadata; zero keeps the outer edge. The page-link
  label stays 4 layout pixels inward and at least 4 layout pixels along the
  edge from that final terminal.

The PPTX exporter should only translate the resolved draw plan into a
presentation. Geometry and routing decisions belong on the Go side.
