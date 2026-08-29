---
applyTo: ".github/instructions/manual/**"
---

# 10.07 PPTX and routing: Routing Rules

## Routing Rules

- Route calculation is in `internal/usecase/v1/engine/route_*`.
- Obstacles include resolved image and text rectangles from the internal V1 scene.
- Start/end rectangles are excluded from obstacle checks for that connection.
- Resolved endpoint binding gaps must be honored in SVG/PPTX routing.
- If any obstacle-free candidate exists, obstacle-hitting candidates must not be
  selected.
- Lines on an obstacle boundary count as collision.
- Existing routed paths are included in scoring so later lines avoid overlap and
  near-parallel crowding.
- The shared route stage feeds previously routed lines back into scoring so
  matching X/Y lanes are offset before projection.
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

The current internal V1 scene groups each item image and label with a small 5x5
anchor grid. The grid is drawn above connector lines and below the icon/label
so labels remain readable without hiding the connector endpoint. Shared
routing treats group header tags, item icons, and labels as obstacles.
