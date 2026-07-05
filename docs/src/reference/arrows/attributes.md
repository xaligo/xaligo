# Arrow Attribute Reference

| Attribute | Values | Default |
|---|---|---|
| `kind` | `route`, `traffic`, or omitted | normal connection |
| `color` | CSS hex color | depends on `kind` |
| `stroke-width`, `width` | positive number | `1` |
| `stroke-style` | `solid`, `dashed`, `dotted` | `solid` |
| `start-arrowhead`, `end-arrowhead` | `none`, `arrow`, `triangle`, `stealth`, `diamond`, `oval` | depends on `kind` |
| `arrowhead` | same as `end-arrowhead` | alias |
| `src-side`, `dst-side` | `top`, `right`, `bottom`, `left` | automatic |
| `src-anchor`, `dst-anchor` | `side-1` through `side-5` | automatic |
| `bends`, `points`, `via` | inline coordinate list | none |
| `grid` | positive number | router default |
| `scale`, `coordinate-scale` | positive number | `1` |
