---
applyTo: ".github/instructions/manual/**"
---

# 11.04 Diagram creation: Step 4 — Render SVG

## Step 4 — Render SVG

```bash
xaligo render docs/src/examples/samples/sample.xal \
  --format svg \
  -o output/sample.svg \
  --services docs/src/examples/samples/services.csv
```

`--services` is strongly recommended for this workflow. The CSV provides
icon label overrides and service metadata. SVG uses it to draw a service
legend; place that legend with
`--svg-legend-position top|right|bottom|left` (default `bottom`).

> **Note:** Create the output directory if it does not already exist.
> ```bash
> mkdir -p output
> ```

---
