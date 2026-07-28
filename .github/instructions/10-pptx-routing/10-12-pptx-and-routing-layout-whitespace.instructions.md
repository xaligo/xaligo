---
applyTo: ".github/instructions/manual/**"
---

# 10.12 PPTX and routing: Layout / Whitespace

## Layout / Whitespace

Supported whitespace controls:

| Syntax | Behavior |
|---|---|
| `<spacer />` / `<blank />` | Empty layout slot, not rendered |
| `<item />` | Empty item-grid slot, not rendered |
| `class="pa-4"` | Inner padding, Vuetify-style 8px unit |
| `class="ma-4"` | Outer margin; on root frame this becomes page-edge content whitespace |
| `margin="N"` and `margin-*` | Pixel margin |
| `content-width="N"` / `content-height="N"` | Shrinks usable inner layout area |
| `align="top-left"` etc. | Aligns the usable content area or item grid |
| `width="N"` / `height="N"` | Fixed child size, except root frame is the paper/content frame |

Fixed children are reserved before flexible `row`/`col` allocation. The
resolved size advances the sibling cursor and must remain inside the parent's
content box unless the source explicitly uses `overflow="visible"`. Layout
overflow is diagnosed before plan construction; SVG or PPTX clipping is not a
substitute for a valid layout.

For item grids, horizontal `spread` is also supported.
