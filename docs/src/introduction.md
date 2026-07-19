# xaligo

xaligo is a diagram-as-code engine for architecture and network diagrams. It
renders the Vue-style `.xal` XML DSL to Excalidraw, SVG, PPTX, PDF, Excel,
XYFlow, and Isoflow-compatible output.

![Hybrid enterprise architecture](images/complex-hybrid-architecture.svg)

The renderer is designed around a single shared pipeline:

```text
.xal
  -> parser
  -> layout
  -> shared scene and routing
  -> output encoder
  -> Excalidraw | SVG | PPTX | PDF | Excel | XYFlow | Isoflow
```

Core features:

- Vuetify-style layout primitives such as frames, rows, columns, spacing, and
  alignment.
- AWS architecture group components and catalog-backed service icons.
- Orthogonal connector routing with route and traffic layers.
- Manual bend points in frame coordinates.
- Group, rectangle, port, and item endpoints.
- Excalidraw-friendly editable output.
- Frame-oriented SVG, PPTX, PDF, and Excel output, with one frame mapped to one
  file, slide, page, or worksheet by default.
- SVG and PPTX output with line jumps, legends, and export-focused routing.

Use this documentation when authoring `.xal` files, rendering output, or
working on the xaligo codebase.

See [Planned Work](roadmap.md) for upcoming and exploratory features.
