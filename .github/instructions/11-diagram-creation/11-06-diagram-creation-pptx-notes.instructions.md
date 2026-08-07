---
applyTo: ".github/instructions/manual/**"
---

# 11.06 Diagram creation: PPTX Notes

## PPTX Notes

- Native CLI export requires `xaligo.wasm`; the npm/WASM API currently
  exports through the Rust `pptx` adapter.
- Diagram and legend icons are native SVG media. SVG-capable PowerPoint or a
  compatible viewer is required; legacy raster-only viewers do not display the
  icons because independently rasterized copies are not embedded.
- PPTX export adds separate legend slide(s) after all frame/diagram slides.
- Legend pages use 4 columns and show icon, abbreviation, and official name.
- Use `--paper A3 --orientation landscape --paper-margin-top 0.75 --paper-margin-bottom 0.75`
  for the current large AWS sample.
- Connector routing is resolved in Go/WASM and avoids icon/label obstacles.
- Group header tag labels are intentionally single-line in PPTX output; keep
  tag background width and label width in sync when adjusting tag text metrics.
- Group header and item label width estimates count East Asian full-width
  characters as double-width, so Japanese and other full-width labels keep their
  text boxes aligned across SVG and PPTX.
- Keep `docs/src/examples/samples/sample.xal` and `docs/src/examples/samples/services.csv` in sync so the legend
  includes every diagram service.
