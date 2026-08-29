# xaligo

xaligo is a diagram-as-code engine for architecture, network, and UML diagrams.
It renders the Vue-style `.xal` XML DSL to SVG and PPTX and embeds SVG output
in Markdown.

![Hybrid enterprise architecture](images/complex-hybrid-architecture.svg)

The renderer is designed around a single shared pipeline:

```text
.xal
  -> parser
  -> layout
  -> generic layout and routing
  -> renderer-neutral document plan
  -> SVG | PPTX
```

Core features:

- Vuetify-style layout primitives such as frames, rows, columns, spacing, and
  alignment.
- AWS architecture group components and catalog-backed service icons.
- Orthogonal connector routing with route and traffic layers.
- Manual bend points in frame coordinates.
- Group, rectangle, port, and item endpoints.
- Frame-oriented SVG and PPTX output, with one frame mapped to one file or
  slide by default.
- Markdown rendering that embeds those same SVG artifacts.
- SVG and PPTX output with line jumps, legends, and export-focused routing.

Use this documentation when authoring `.xal` files, rendering output, or
working on the xaligo codebase.

See [Planned Work](roadmap.md) for upcoming and exploratory features.
