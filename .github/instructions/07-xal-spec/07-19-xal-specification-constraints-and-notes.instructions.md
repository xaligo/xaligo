---
applyTo: ".github/instructions/manual/**"
---

# 07.19 XAL specification: Constraints and Notes

## Constraints and Notes

- The canonical root is `<xaligo version="1">`. Legacy `<frame>` and
  `<frames>` roots are accepted with a warning. Direct children of `<frames>`
  must be identified `<frame>` tags. V2 uses `<scene version="2">`, which is
  intentionally rejected by V1.
- Both self-closing (`<card title="..." />`) and regular (`<card title="..."></card>`) forms are supported.
- The sum of `span` values in direct children of `<row>` must not exceed 12.
  Excess is a validation error rather than implicit overflow to the right.
- `.xal` files must be saved in UTF-8.
