# SVG Output

SVG output is suitable for previews, documentation, and web publishing.

```bash
xaligo render diagram.xal --format svg -o diagram.svg
```

With `--services`, SVG can draw a legend:

```bash
xaligo render diagram.xal --format svg -o diagram.svg \
  --services services.csv \
  --svg-legend-position bottom
```

SVG rendering uses the shared orthogonal router and includes:

- Route and traffic layer ordering.
- Line-jump masks at interior crossings.
- Automatic junction markers for shared route endpoints.
- Distinct marker geometry for V1 `arrow`, `triangle`, `stealth`, `diamond`,
  and `oval` arrowheads.
- Canvas/viewBox bounds expanded from the resolved stroke width and
  stroke-scaled marker geometry.
- Service legends when metadata is provided.
