# Cross-Frame Page Links

This example shows how a V1 connection becomes a page link when its endpoints
belong to different frames. The source endpoint is local to the frame that
declares the connection. The destination uses the required
`frame-id.endpoint-id` form.

The HTTP connection selects the nearest outer logical page side of `overview`
and terminates on that side's inward page-terminal line with
`to <service-detail>`. It starts from its selected inset line in
`service-detail` with `from <overview>`. Its sides are selected automatically
from the endpoint envelopes.

The dashed CSV connection uses the child endpoint form to demonstrate that
item and page-edge anchors are independent. It leaves the source port through
the `near` slot on `right` (`right-2`) and exits the overview page through the
`far` slot on `bottom` (`bottom-4`). On the destination page it enters through
the `near` slot on `top` (`top-2`) and reaches the port through the `far` slot
on `left` (`left-4`). The frame anchors keep their 70% and 30% tangent
coordinates while the actual terminal moves only inward normally. Each end
segment is perpendicular to its own selected side. It is labeled
`to <database-detail>` in `overview` and `from <overview>` in
`database-detail`.

None of these frames enables metadata, so their page terminals use the default
4-layout-pixel inset on every side. A metadata-enabled frame instead uses its
resolved `row-gap`; `row-gap="0"` retains the outer logical frame edge. For
an explicit frame side or anchor, the inset must be strictly smaller than the
frame height on `top`/`bottom`, or the frame width on `left`/`right`, and an
actual terminal must avoid the metadata reservation. Invalid explicit geometry
is reported at the connection source position; the inset is never clamped.

Without an explicit frame terminal, validation only requires one safe side.
Rendering keeps a safe endpoint-side preference or chooses the nearest safe
side from actual icon and label geometry when that preference is unsafe. Only
an empty candidate set is an error. A safe `left`/`right` candidate remains
available even if a hypothetical top/bottom line would enter metadata.

At zero inset, a link whose endpoint is its owning frame cannot give an
explicit frame anchor the same resolved point. An explicit endpoint anchor uses
its own side and slot; a bare endpoint side and an automatically selected
endpoint side both use their center (`side-3`) for this check.

Each page-link label is placed from the final inset terminal, 4 layout pixels
farther inward and at least 4 layout pixels sideways. It uses the closest
sideways position that remains outside endpoint envelopes and metadata
reservation strips.

Default SVG export creates one file for each of the three frames. These are the
actual page-local artifacts generated from the sample:

The outer page-frame outline is intentionally absent. Its invisible edge still
defines side selection and the frame-anchor tangent coordinate, while each
incoming or outgoing stub ends on the parallel inset line. Each default SVG is
strictly cropped to the frame rectangle, so no extra marker-safe padding is
added around these three page artifacts.

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
