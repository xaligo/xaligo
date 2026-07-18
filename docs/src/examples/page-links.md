# Cross-Frame Page Links

This example shows how a V1 connection becomes a page link when its endpoints
belong to different frames. The source endpoint is local to the frame that
declares the connection. The destination uses the required
`frame-id.endpoint-id` form.

The HTTP connection leaves the nearest logical page edge of `overview` with
`to <service-detail>` and enters `service-detail` with `from <overview>`. Its sides
are selected automatically from the endpoint envelopes.

The dashed CSV connection explicitly selects anchor `1` on each top side. The
anchor is close to a frame corner, so the local stubs use the corner-safe
orthogonal dogleg. It is labeled `to <database-detail>` in `overview` and
`from <overview>` in `database-detail`.

Default SVG export creates one file for each of the three frames. These are the
actual page-local artifacts generated from the sample:

The outer page-frame outline is intentionally absent; the invisible page edge
still anchors each incoming or outgoing page-link stub.

### `overview`

![Overview page with outgoing page links](../images/page-links-overview.svg)

### `service-detail`

![Service detail page with incoming page link](../images/page-links-service-detail.svg)

### `database-detail`

![Database detail page with incoming page link](../images/page-links-database-detail.svg)

The following image is the same source rendered with `--combine-frames` for
compatibility with the former single-canvas output.

![Cross-frame page links, combined compatibility view](../images/page-links.svg)

Source:

```xml
{{#include samples/page-links.xal}}
```

Validate and render it with:

```bash
xaligo validate docs/src/examples/samples/page-links.xal
xaligo render docs/src/examples/samples/page-links.xal \
  --format svg \
  -o output/page-links.svg
```

The render command writes:

```text
output/page-links-overview.svg
output/page-links-service-detail.svg
output/page-links-database-detail.svg
```

To reproduce the combined image shown above:

```bash
xaligo render docs/src/examples/samples/page-links.xal \
  --format svg \
  --combine-frames \
  -o output/page-links.svg
```

PPTX, PDF, and Excel use the same frame order as slides, pages, and worksheets:

```bash
xaligo render docs/src/examples/samples/page-links.xal --format pptx -o output/page-links.pptx
xaligo render docs/src/examples/samples/page-links.xal --format pdf -o output/page-links.pdf
xaligo render docs/src/examples/samples/page-links.xal --format excel -o output/page-links.xlsx
```
