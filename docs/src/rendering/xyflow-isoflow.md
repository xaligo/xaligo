# XYFlow and Isoflow

## XYFlow

XYFlow output exports nodes and edges for React Flow / XYFlow-style editors:

```bash
xaligo render diagram.xal --format xyflow -o diagram.xyflow.json
```

It includes nested group nodes, icon data URLs, labels, connection handles,
route/traffic metadata, layer order, line styles, and arrow markers.

## Isoflow

Isoflow output exports an Isoflow-compatible model:

```bash
xaligo render diagram.xal --format isoflow -o diagram.isoflow.json
```

The model follows the upstream Isoflow shape with `items`, `views`, `icons`,
`colors`, and view `rectangles` / `connectors`.
