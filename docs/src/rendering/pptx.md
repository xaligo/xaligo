# PowerPoint Output

PPTX output is generated from a Go draw plan and written by the configured
WASM/PptxGenJS exporter.

```bash
xaligo render diagram.xal --format pptx -o diagram.pptx \
  --services services.csv \
  --paper A3 --orientation landscape \
  --paper-margin-top 0.75 --paper-margin-bottom 0.75
```

PPTX-specific behavior:

- Paper fitting supports named paper sizes and per-side margins.
- Routing avoids icon and label obstacles where possible.
- Existing routed lanes influence later paths to reduce overlap.
- Anchor backgrounds are drawn before lines, while icons and labels are drawn
  above lines.
- Line-jump masks are drawn at interior crossings.
- Legend slide(s) are added after the diagram when `--services` is provided.

The PPTX exporter should only translate the resolved draw plan into a
presentation. Geometry and routing decisions belong on the Go side.
