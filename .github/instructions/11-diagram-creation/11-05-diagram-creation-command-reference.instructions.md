---
applyTo: ".github/instructions/manual/**"
---

# 11.05 Diagram creation: Command Reference

## Command Reference

| Command | Description |
|---|---|
| `grep -i "<name>" etc/resources/aws/service-index.csv` | Search for a service ID |
| `xaligo render <xal> --format svg -o <out.svg> --services <csv> --svg-legend-position right` | Convert .xal → SVG with a service legend |
| `xaligo render <xal> --format pptx -o <out.pptx> --services <csv> --paper A3 --orientation landscape` | Convert .xal → PPTX when the WASI exporter is configured |
| `xaligo render <xal> -o <out.svg>` | Convert .xal → SVG using the default format |
| `xaligo render markdown <input.md>` | Replace fenced `xal` blocks with generated SVG image references |
