# Cross-Frame Page Links

This example shows how a V1 connection becomes a page link when its endpoints
belong to different frames. The source endpoint is local to the frame that
declares the connection. The destination uses the required
`frame-id.endpoint-id` form.

The HTTP connection leaves the nearest border of `overview` with
`to service-detail` and enters `service-detail` with `from overview`. Its sides
are selected automatically from the endpoint envelopes.

The dashed CSV connection explicitly selects anchor `1` on each top side. The
anchor is close to a frame corner, so the local stubs use the corner-safe
orthogonal dogleg. It is labeled `to database-detail` in `overview` and
`from overview` in `database-detail`.

![Cross-frame page links](../images/page-links.svg)

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
