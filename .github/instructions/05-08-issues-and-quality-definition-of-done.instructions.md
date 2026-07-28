---
applyTo: ".github/instructions/manual/**"
---

# 05.08 Issues and quality: Definition of Done

## Definition of Done

A feature-quality slice is done only when:

- The specification and documentation match the implementation.
- Representative valid samples pass validation and render successfully.
- Representative invalid inputs fail with source-positioned, user-correctable
  diagnostics.
- Regression tests cover the observable behavior being protected.
- Design quality has been reviewed for hierarchy, spacing, typography, shape
  language, connector readability, and contrast.
- Cross-format checks have been run when the feature affects shared scene,
  routing, plan, pagination, or renderer contracts.
- Determinism, cancellation, resource limits, compatibility, accessibility,
  and artifact reproducibility have been considered and tested where relevant.
- Required security and repository checks have passed.
- The slice is committed without unrelated working-tree changes.
