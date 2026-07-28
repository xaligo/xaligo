---
applyTo: ".github/instructions/manual/**"
---

# 10.06 PPTX and routing: Paper / Scaling

## Paper / Scaling

- PPTX export supports `--paper`, `--orientation`, and paper-margin fitting
  flags.
- A3 landscape is generated with:

```bash
.bin/xaligo render docs/src/examples/samples/sample.xal \
  --format pptx \
  --services docs/src/examples/samples/services.csv \
  -o out.pptx \
  --paper A3 \
  --orientation landscape \
  --paper-margin-top 0.75 \
  --paper-margin-bottom 0.75
```

- The shared Go plan resolves paper size and computes one layout-pixel transform.
- The page-oriented plan is built after the full scene and cross-frame page
  links are resolved. Its frame projections preserve source order.
- Shape coordinates, font sizes, strokes, padding, and routing geometry use
  that same transform. `--px-per-inch 144` must not scale text independently
  from its containing shape.
- `--paper-margin N` applies an inch-based margin to every side before fitting
  the diagram to the selected paper.
- `--paper-margin-top`, `--paper-margin-right`, `--paper-margin-bottom`, and
  `--paper-margin-left` override the all-side value for individual sides.
- Paper margins do not change the slide size; they reduce the available fit
  area and centre the diagram within that inset area.
- The `paper-frame` element remains the content frame for scaling.
- Root `<frame margin="N">` or `class="ma-N"` is content outer whitespace: it
  insets diagram content without shrinking the paper frame itself.
