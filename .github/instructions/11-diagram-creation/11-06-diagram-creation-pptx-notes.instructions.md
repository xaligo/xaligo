---
applyTo: ".github/instructions/manual/**"
---

# 11.06 Diagram creation: PPTX Notes

## PPTX Notes

- Native CLI export uses the Rust `pptx` adapter compiled into the single
  executable; no runtime exporter file is required.
- Diagram and legend icons are native SVG media. SVG-capable PowerPoint or a
  compatible viewer is required; legacy raster-only viewers do not display the
  icons because independently rasterized copies are not embedded.
- PPTX export adds separate legend slide(s) after all frame/diagram slides.
- Legend pages use 4 columns and show icon, abbreviation, and official name.
- Use `--paper A3 --orientation landscape --paper-margin-top 0.75 --paper-margin-bottom 0.75`
  for the current large AWS sample.
- Connector routing is resolved in the shared Go pipeline and avoids icon/label
  obstacles before the native Rust exporter boundary.
- Group header tag labels are intentionally single-line in PPTX output; keep
  tag background width and label width in sync when adjusting tag text metrics.
- Group headers use proportional half-width glyph estimates and one font-size
  unit for East Asian wide/full-width glyphs, keeping Japanese and mixed-width
  labels compact and aligned across SVG and PPTX. Item labels count full-width
  characters as two display columns when calculating wrapping.
- AWS boundary resources such as `<vpc-endpoint>` use the same resolved
  border-centered icon bounds in SVG and PPTX. Keep them out of normal child
  flow and preserve their connection endpoint ID when changing projection.
- Keep `docs/src/examples/samples/sample.xal` and `docs/src/examples/samples/services.csv` in sync so the legend
  includes every diagram service.
