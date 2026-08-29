# Markdown

`xaligo render markdown` finds fenced `xal` blocks, renders each frame through
the same SVG artifact pipeline as `xaligo render --format svg`, and writes a
new Markdown document containing image references.

```bash
xaligo render markdown guide.md
xaligo render markdown guide.md --output output/guide.md --svg-dir output/images
```

The default output is `<source>.embedded.md`; generated SVG files are written
beside the source. `--output` and `--svg-dir` override those locations.
Markdown is an orchestration workflow, not a separate geometry or encoder
implementation, so layout, routing, themes, services, and frame behavior match
direct SVG rendering.
