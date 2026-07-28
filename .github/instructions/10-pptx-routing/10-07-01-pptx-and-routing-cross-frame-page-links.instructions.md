---
applyTo: ".github/instructions/manual/**"
---

# 10.07.01 PPTX and routing: Cross-frame page links

### Cross-frame page links

A connection between different frames is a page link in page-oriented output;
it is never one line crossing the inter-frame canvas. The shared scene emits
two axis-aligned local stubs for Excalidraw, SVG, PPTX, PDF, and Excel:

- source endpoint to the source frame's page-terminal inset line, with the
  exact label `to <destination frame ID>`; and
- destination frame's page-terminal inset line to the destination endpoint,
  with the exact label `from <source frame ID>`.

Angle brackets are literal punctuation, so a link from `overview` to `detail`
renders `to <detail>` and `from <overview>`. The shared scene, not the PPTX
exporter, selects endpoint binding and logical frame terminal geometry. The
endpoint uses `src-side`/`dst-side` and `src-anchor`/`dst-anchor`. A cross-frame
connection may independently select its logical terminal with
`src-frame-side`/`dst-frame-side` or the more specific
`src-frame-anchor`/`dst-frame-anchor`. Every side has five anchors at
10/30/50/70/90 percent along the outer frame extent. The endpoint- and
frame-terminal-adjacent route segments remain perpendicular to their
respective sides even when the two sides differ.

Frame-terminal precedence is explicit frame anchor, explicit frame side,
legacy endpoint anchor, endpoint side, then automatic nearest-side selection.
The first two are fixed choices. Without either one, the legacy/automatic result
is a preferred side: the renderer keeps it when safe or selects the nearest
safe side from the endpoint's rendered visual envelope when it is unsafe. Ties
prefer a tied side facing the remote frame, then `top`, `right`, `bottom`,
`left`. Validation checks that at least one safe candidate exists but does not
predict the automatic side from layout `Box` geometry. Frame-terminal
attributes are cross-frame-only; using them on a same-frame connection is a
validation error.

A frame metadata reservation strip is a final safety constraint. The visible
metadata rows are inset from their selected vertical edge and both horizontal
edges by the resolved `row-gap`, while the reservation itself remains
full-width from the outer logical frame edge to the content boundary. Without
an explicit frame-terminal attribute, the renderer filters unsafe sides before
its visual nearest-side choice; an unsafe normal preference is remapped to the
nearest safe candidate. A left/right terminal is clamped outside the full-width
strip before the orthogonal dogleg is built. An explicit frame side or anchor
that selects the metadata edge, or an exact left/right anchor inside the strip,
is a validation error instead of being moved. Neither the local path nor its
label may enter the reservation strip.
If an explicit frame side is vertically opposite the metadata edge, its actual
terminal must remain outside the same strip: `top` against bottom metadata, or
`bottom` against top metadata. Moving that explicit terminal into the strip is
a source-positioned validation error. For an automatic page terminal, the same
conflict excludes that candidate rather than immediately rejecting the
connection. A safe explicit `left` or `right` terminal is allowed even if the
unused top/bottom inset line would intersect the reservation.

An automatic terminal first uses the endpoint binding's coordinate parallel to
the outer logical frame edge. If that coordinate enters a 24-layout-px corner
gutter, it is clamped and a two-bend orthogonal dogleg bridges the difference.
A border shorter than 96 layout pixels uses one quarter of its length as an
adaptive gutter. An explicit frame anchor instead retains its exact
10/30/50/70/90-percent tangent slot.

After side, tangent, and reservation handling, the drawable page terminal is
shifted inward from the outer logical frame edge along that side's normal. A
frame with metadata uses the resolved metadata `row-gap` on all four terminal
sides, regardless of whether the band is at the top or bottom. A frame without
metadata uses 4 layout pixels. A metadata `row-gap` of zero leaves the terminal
on the outer edge. This step does not change the resolved tangent coordinate.
For an explicit frame side or anchor, the inset must be strictly smaller than
that side's normal frame dimension: height for `top`/`bottom`, width for
`left`/`right`. A reservation conflict on that actual side is likewise an
error. Both failures are source-positioned at the connection. For an automatic
page terminal, the same tests classify each candidate side; an unsafe preferred
side is remapped and only an empty safe candidate set is an error. The inset is
applied exactly, without an implicit clamp.

The segments at both the endpoint and page-terminal inset line remain
perpendicular to their selected side; an explicit frame anchor's orthogonal
local stub supplies visible separation. If an unconstrained final inset
terminal and the endpoint coincide, the terminal shifts by up to 24 layout
pixels along the parallel axis within the available gutter range so the line
remains visible. On a left/right side with metadata, that preferred range also
keeps 8 layout pixels from the reservation. If a very small safe region cannot
retain both the corner gutter and that clearance, the shift falls back to the
entire non-reserved interval, may touch its boundary, and never leaves the
frame or enters the strip.

If an owning frame has metadata with resolved `row-gap="0"`, the endpoint
resolves to that frame itself, and an explicit frame anchor coincides with the
resolved endpoint point, that connection is a source-positioned validation
error. An explicit endpoint anchor supplies its side and slot; an explicit
endpoint side or automatically resolved side uses the center slot (`side-3`).
The author must use a different endpoint/frame anchor or a positive `row-gap`;
xaligo must not move the fixed frame-anchor tangent coordinate or emit an
invisible zero-length stub.

Manual bends remain connector metadata and do not steer page-local stubs. Both
stubs retain one logical connector ID; XYFlow and Isoflow reconstruct one graph
edge from that metadata rather than exporting the two page projections.

The `to <...>` / `from <...>` label is placed from the final inset terminal
with a 4-layout-pixel inward gap and a minimum 4-layout-pixel tangent gap.
Candidate placement chooses the closest tangent position that avoids
the endpoint envelope and metadata reservation; tiny pages use a clamped
fallback rather than increasing the normal label distance.

The outer logical page edge and the parallel terminal inset line are geometric,
not visible rectangles: SVG, PPTX, PDF, and Excel omit page-frame outlines in
both default and combined output.

Default PPTX output places the source and destination stubs on their respective
frame slides. `--combine-frames` places both stubs on the compatibility slide
but never draws a replacement line across the frame gap.
