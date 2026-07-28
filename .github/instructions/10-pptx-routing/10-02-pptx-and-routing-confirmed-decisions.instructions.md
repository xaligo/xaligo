---
applyTo: ".github/instructions/manual/**"
---

# 10.02 PPTX and routing: Confirmed Decisions

## Confirmed Decisions

- PPTX export is an A3-landscape-first workflow for the current AWS sample.
- The PPTX export implementation should be compiled to WASM and invoked from
  the Go repository layer.
- Do not use `goja` or V8 for PPTX export execution.
- Avoid a long-term Node.js subprocess dependency for repository-layer PPTX
  export. Node may remain a development/build tool only while the WASM exporter
  is being prepared.
- All PPTX geometry and routing decisions are computed by the Go use-case
  pipeline before the exporter boundary.
- Each identified child frame becomes one diagram slide in source order by
  default. `--combine-frames` is the explicit compatibility path for the
  former single-slide canvas.
- A presentation has one common slide size. Multi-frame PPTX uses the largest
  resolved page width and height and centers smaller frame pages without
  scaling them independently.
- The PPTX drawing/export layer must not make independent layout/routing
  decisions.
- Lines must not visually cover icons or labels.
- If any obstacle-free route exists, obstacle-hitting routes must be rejected.
- Item labels are 8pt at the default 96 PPI and scale with the same effective
  PPI/paper-fit transform as item icons.
- Item icons should remain visually consistent with their labels; avoid shrinking
  icons merely to satisfy a cramped row when layout whitespace controls can be used.
- Legend belongs on separate PPTX slide(s), not outside the diagram page.
- Legend slide layout is fixed to 4 columns and contains icon, abbreviation, and
  official service name.
- DSL must support empty grid cells and both inner/outer whitespace controls.
