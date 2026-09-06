---
applyTo: ".github/instructions/manual/**"
---

# 07.17 XAL specification: Layout Calculation Rules

## Layout Calculation Rules

1. Normalize and validate numeric attributes and enum values.
2. Resolve each parent's border box and content box after margin and padding.
3. `frame` / `container` / `col` → **vertical stack**: reserve fixed child
   heights, gaps, and margins, then divide the remainder by `row` weights.
4. `layout="horizontal"` → reserve fixed child widths, gaps, and margins, then
   divide the remainder by `col` weights.
5. `row` → **12-column grid** after validating each `span` and their total.
6. Leaf elements use the resolved `(x, y, w, h)` received from their parent;
   they do not replace the allocation after sibling placement.
7. Verify finite positive geometry and parent-content containment before scene
   construction. Respect only an explicit `overflow="visible"` exception.
8. Resolve item grids against the same occupied content area before encoding.
9. Exclude semantic boundary resources from normal child flow, then resolve
   them against their owner's border by side, unit-interval anchor, tangential
   offset, and visual size. Their center may straddle the selected border, but
   their tangential extent must remain on that border and siblings must not
   overlap.
