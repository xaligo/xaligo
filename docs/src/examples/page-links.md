# Cross-Frame Page Links

This example shows how a V1 connection becomes a page link when its endpoints
belong to different frames. The source endpoint is local to the declaring frame;
the destination uses `frame-id.endpoint-id`.

Default SVG export creates one file for each identified frame:

### `overview`

![Overview page with outgoing page links](../images/page-links-overview.svg)

### `service-detail`

![Service detail page with incoming page link](../images/page-links-service-detail.svg)

### `database-detail`

![Database detail page with incoming page link](../images/page-links-database-detail.svg)

The combined compatibility view uses the same source with `--combine-frames`:

![Cross-frame page links, combined compatibility view](../images/page-links.svg)

Source: [samples/page-links.xal](samples/page-links.xal).

See [Anchors and Bends](../reference/arrows/anchors-bends.md) and
[Frame and Containers](../reference/frames/frame-containers.md) for page-edge
terminal, anchor, and metadata-reservation rules.

```bash
xaligo validate docs/src/examples/samples/page-links.xal
xaligo render docs/src/examples/samples/page-links.xal --format svg -o output/page-links.svg
xaligo render docs/src/examples/samples/page-links.xal --format svg --combine-frames -o output/page-links.svg
```

PPTX, PDF, and Excel use the same frame order as slides, pages, and worksheets:

```bash
xaligo render docs/src/examples/samples/page-links.xal --format pptx -o output/page-links.pptx
xaligo render docs/src/examples/samples/page-links.xal --format pdf -o output/page-links.pdf
xaligo render docs/src/examples/samples/page-links.xal --format excel -o output/page-links.xlsx
```
