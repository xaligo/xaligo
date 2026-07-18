# SVG Output

SVG output is suitable for previews, documentation, and web publishing.

```bash
xaligo render diagram.xal --format svg -o diagram.svg
```

An identified child frame is one SVG artifact by default. A one-frame document
writes exactly `diagram.svg`. If the document contains frames `overview` and
`service-detail`, the same command writes:

```text
diagram-overview.svg
diagram-service-detail.svg
```

The frame ID is made filename-safe by retaining ASCII letters, digits, `_`,
and `-`, replacing every run of other characters with one `-`, and trimming
leading and trailing `-`. An empty safe ID falls back to `frame-<source-order>`;
colliding filenames are reported as an error. No ZIP archive is created.

Use the compatibility option to keep every frame on one historical canvas:

```bash
xaligo render diagram.xal --format svg -o diagram.svg --combine-frames
```

Live preview uses this combined view automatically.

The page frame itself is not rendered as a visible outline. It defines the SVG
canvas, crop, and page-link edge only. Combined compatibility output also omits
each page-frame outline.

With `--services`, SVG can draw a legend:

```bash
xaligo render diagram.xal --format svg -o diagram.svg \
  --services services.csv \
  --svg-legend-position bottom
```

SVG rendering uses the shared orthogonal router and includes:

- Route and traffic layer ordering.
- Line-jump masks at interior crossings.
- Automatic junction markers for shared route endpoints.
- Distinct marker geometry for V1 `arrow`, `triangle`, `stealth`, `diamond`,
  and `oval` arrowheads.
- Canvas/viewBox bounds expanded from the resolved stroke width and
  stroke-scaled marker geometry.
- Service legends when metadata is provided.

The complete document scene is routed before its frame pages are projected.
A cross-frame connection therefore appears as `to <destination frame ID>` on
the source SVG and `from <source frame ID>` on the destination SVG. Combined
output keeps those two page-link stubs and does not reconnect them across the
inter-frame gap.
