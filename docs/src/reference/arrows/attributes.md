# Arrow Attribute Reference

| Attribute | Values | Default |
|---|---|---|
| `kind` | `connection`, `route`, `traffic`, or omitted | normal connection |
| `color` | six-digit `#RRGGBB` | depends on `kind` |
| `stroke-width`, `width` | positive number | `1` |
| `stroke-style` | `solid`, `dashed`, `dotted` | `solid` |
| `start-arrowhead`, `end-arrowhead` | `none`, `arrow`, `triangle`, `stealth`, `diamond`, `oval`; routes permit only `none` | depends on `kind` |
| `arrowhead` | same as `end-arrowhead`; routes permit only `none` | alias |
| `arrowhead-size` | `s` (V1 only) | `s` |
| `src-side`, `dst-side` | `top`, `right`, `bottom`, `left` | automatic |
| `src-anchor`, `dst-anchor` | `side-1` through `side-5` | automatic |
| `bends`, `points`, `via` | inline coordinate list | none |
| `grid` | positive number | router default |
| `scale`, `coordinate-scale` | positive number | `1` |

Anchor position aliases are `start=1`, `near=2`, `center=3`, `far=4`, and
`end=5`.

A `<connections>` group inherits only non-empty `arrowhead-size`, `kind`,
`color`, `stroke-width`, `width`, `stroke-style`, `start-arrowhead`,
`end-arrowhead`, `arrowhead`, `scale`, `coordinate-scale`, and `grid` values.
Child values override them. Endpoints, sides, anchors, and bend/point/via data
are never inherited.

Unknown non-empty `kind`, `stroke-style`, and arrowhead values are validation
errors. V1 also rejects `arrowhead-size="m"` and `"l"` instead of silently
rendering them as small. After group defaults and child semantic aliases are
merged, an effective `kind="route"` rejects every non-`none` start or end
arrowhead; explicit `none` is valid.
