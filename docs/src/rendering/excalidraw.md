# Excalidraw Output

Excalidraw output is editable JSON:

```bash
xaligo render diagram.xal --format excalidraw -o diagram.excalidraw
```

Key behavior:

- Connectors are emitted as elbowed Excalidraw arrows.
- Bindings use fixed edge points so arrows stay attached when opened in
  Excalidraw.
- Item icons and labels are grouped with a small 5x5 white anchor grid.
- Anchor grid cells are drawn above connector lines and below item content so
  labels remain readable while endpoints stay visible.
- Previously routed lines are fed back into the router so later lines can avoid
  exact or near-exact X/Y lane overlaps.
- Group header tags, item icons, and labels are treated as routing obstacles
  where possible.
- Arrowhead sizes are serialized as Excalidraw's smallest supported size
  (`s`) for dense diagrams.

Excalidraw remains an editable drawing tool. Its editor may re-interpret
elbowed connectors after manual edits, so SVG and PPTX can be more exact for
final presentation output.
