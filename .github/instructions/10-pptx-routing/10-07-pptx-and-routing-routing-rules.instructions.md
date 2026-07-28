---
applyTo: ".github/instructions/manual/**"
---

# 10.07 PPTX and routing: Routing Rules

## Routing Rules

- Route calculation is in `internal/usecase/v1/engine/route_*`.
- Obstacles include image and text rectangles from the Excalidraw scene.
- Start/end rectangles are excluded from obstacle checks for that connection.
- Binding `gap` from Excalidraw arrows must be honored in PPTX routing.
- If any obstacle-free candidate exists, obstacle-hitting candidates must not be
  selected.
- Lines on an obstacle boundary count as collision.
- Existing routed paths are included in scoring so later lines avoid overlap and
  near-parallel crowding.
- Excalidraw output also feeds previously routed lines back into the shared
  router so matching X/Y lanes are offset before export.
- Visible container borders are reserved routing paths. Connectors may cross a
  frame boundary, but parallel paths prefer the configured line margin.
- Previously placed line lanes are used as candidate offsets, so `--arrow-margin`
  affects routes that would otherwise share the same position.
- Final PPTX drawing order is:
  1. anchor backgrounds and containers/shapes
  2. route lines, traffic lines, and line-jump masks
  3. automatic junction markers
  4. icons and labels

This order prevents lines from visually covering icons even at endpoints.

Excalidraw output mirrors this readability rule with editable JSON elements:
each item image and label is grouped with a small 5x5 white anchor grid. The
grid is drawn above connector lines and below the icon/label so labels remain
readable without hiding the connector endpoint. Excalidraw routing treats group
header tags, item icons, and labels as obstacles where possible, and serializes
arrowhead sizes as `"s"` for dense diagrams.
