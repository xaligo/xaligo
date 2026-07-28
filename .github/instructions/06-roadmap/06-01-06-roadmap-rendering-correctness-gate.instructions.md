---
applyTo: ".github/instructions/manual/**"
---

# 06.01.06 Roadmap: Rendering Correctness Gate

### Rendering Correctness Gate

New renderer features are gated by a shared geometry and text contract. Fixes
must be made at the earliest shared stage that owns the information, not as
format-specific clipping or coordinate adjustments.

The required order is:

1. Parse numeric layout attributes into finite, typed values and validate their
   domains with source positions.
2. Make validation and rendering execute the same geometry invariants.
3. Resolve fixed-size children before flexible weights, then record content
   boxes and explicit overflow state in the resolved layout.
4. Move item-grid selection and occupancy into resolved layout so items and
   other children cannot unknowingly occupy the same region; scene construction
   only emits the already resolved cells.
5. Carry renderer-neutral text layout, semantic role, and glyph-overflow policy
   through the draw plan.
6. Apply the same output transform to geometry and typography at every PPI and
   paper-fit setting.
7. Consolidate format dispatch in one use case and migrate the shared scene and
   plan to format-neutral names and schemas. Compatibility aliases may preserve
   public APIs, but the canonical schema must not remain Excalidraw- or
   PPTX-shaped.

Completion requires regression coverage for validation/render agreement,
finite resolved coordinates, parent/content containment, fixed-plus-flex
siblings, mixed item/rectangle groups, item offsets, connector numeric values,
empty numeric attributes, long labels across output formats including editable
Excalidraw metadata, overlapping ports, and non-96 PPI.
