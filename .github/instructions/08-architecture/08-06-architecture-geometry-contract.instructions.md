---
applyTo: ".github/instructions/manual/**"
---

# 08.06 Architecture: Geometry contract

## Geometry contract

Layout is a constraint-resolution stage, not a best-effort drawing stage. It
must establish these postconditions before a scene or plan is constructed:

- every coordinate, length, weight, gap, margin, padding, and scale is finite;
- drawable width and height are greater than zero;
- row and column weights are greater than zero, and grid spans are in range;
- each content box is derived once from its allocated border box;
- fixed-size children consume space before flexible children are distributed;
- gaps are subtracted exactly once and cursors advance by the resolved size;
- containment or the selected overflow policy is recorded explicitly; and
- invalid geometry is returned as a source-positioned diagnostic, not dropped
  later by scene construction or exposed to an output encoder.

With `overflow="visible"`, fixed children still consume their resolved sizes
and advance the cursor. If they leave no positive remainder while flexible
children exist, the parent's original usable main-axis extent becomes the flex
pool. Children remain in source order, and every sibling cursor advances by the
resolved size plus its declared gap and margins, so the resulting overflow is
explicit. The default `overflow="error"` rejects the same input.

`Validate` and `Render` must both call this same stage. Encoders may reject an
I/O or serialization failure, but they must never be the first component to
discover `NaN`, `Inf`, a negative drawable size, or an impossible grid ratio.
