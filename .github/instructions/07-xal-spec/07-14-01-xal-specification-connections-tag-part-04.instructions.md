---
applyTo: ".github/instructions/manual/**"
---

07-14-01-xal-specification-connections-tag — part 4/4

**Arrow spec:**
- `elbowed: true` — always right-angle connectors (Excalidraw "elbow connector")
- Arrowhead at end only by default. Excalidraw stores this as
  `endArrowhead: "arrow"` plus `endArrowheadSize: "s"`; xaligo metadata records
  the logical PPTX/SVG head as `stealth`.
- Stroke color `#1e1e1e`, stroke width `1px` for normal connections
- `kind="route"` defaults to `#64748b`, `1px`, lower route layer, no arrowheads
- `kind="traffic"` defaults to `#2563eb`, `1px`, higher traffic layer, directional end arrowhead
- A traffic line with the same endpoints as a route line is drawn beside that
  route in Excalidraw, SVG, and PPTX draw paths when possible.
- Start/end connect to the **edge midpoint** of the element
  - When direction is **downward**: label text element (`{id}-item-lbl`) bottom edge
  - Otherwise: icon image element (`{id}-item`) corresponding edge
- Edges are fixed with normalized coordinates via `fixedPoint`, so arrows snap correctly when the file is opened
- Arrow ID format: `conn-{src}-{dst}-{index}`
- Arrow ID is registered in `boundElements` of the bound elements
- Excalidraw item icons and labels are grouped with a 5x5 white anchor grid.
  Anchor grid cells are drawn above connectors and below the item content so
  lines do not cover icons/labels while endpoints remain visible.
- Excalidraw routing uses previously placed lines to offset exact or near-exact
  lane overlaps. Group header tags, item icons, and labels are treated as
  routing obstacles where possible.
- SVG/PPTX routing may additionally add automatic junction markers and line
  jump masks after the Excalidraw scene is built. PDF and Excel inherit the SVG
  projection. These are export-layer
  rendering features, not extra `.xal` tags.

**Edge selection logic:**

| Direction (dst relative to src) | Start edge | End edge |
|---|---|---|
| Right (dx ≥ dy) | right | left |
| Left | left | right |
| Down (dy > dx) | bottom (label) | top |
| Up | top | bottom (label) |

> If `src` / `dst` items are not rendered, a warning is emitted and the connection is skipped.

Connection endpoints must resolve to exactly one item, AWS group, rectangle,
port, or identified child frame. Numeric catalog IDs are valid only when that
ID appears once in the document; when the same service appears multiple times,
use unique `name` or `ref` values. Missing endpoints, ambiguous numeric IDs,
duplicate aliases, and `<connection>` tags nested below any tag other than
`<frame>` or its direct `<connections>` child are validation errors.
