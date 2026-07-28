---
applyTo: ".github/instructions/manual/**"
---

# 11.05 Diagram creation: Command Reference

## Command Reference

| Command | Description |
|---|---|
| `grep -i "<name>" etc/resources/aws/service-index.csv` | Search for a service ID |
| `xaligo render <xal> --format excalidraw -o <out> --services <csv>` | Convert .xal → .excalidraw with legend |
| `xaligo render <xal> --format svg -o <out.svg> --services <csv> --svg-legend-position right` | Convert .xal → SVG with a service legend |
| `xaligo render <xal> --format pptx -o <out.pptx> --services <csv> --paper A3 --orientation landscape` | Convert .xal → PPTX when the WASI exporter is configured |
| `xaligo render <xal> --format pdf -o <out.pdf>` | Convert .xal → PDF with one frame per page by default |
| `xaligo render <xal> --format excel -o <out.xlsx>` | Convert .xal → Excel with one frame SVG per worksheet by default |
| `xaligo add service --list <csv> --file <excalidraw>` | Add service icons to an existing file |
| `xaligo render <xal> -o <excalidraw>` | Convert .xal → .excalidraw without legend |
