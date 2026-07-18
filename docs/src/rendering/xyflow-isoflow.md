# XYFlow and Isoflow

## XYFlow

XYFlow output exports nodes and edges for React Flow / XYFlow-style editors:

```bash
xaligo render diagram.xal --format xyflow -o diagram.xyflow.json
```

It includes nested group nodes, icon data URLs, labels, connection handles,
route/traffic metadata, layer order, line styles, and arrow markers.
All frames remain in one graph document; `--combine-frames` has no effect.

V1 connections may target an item, AWS group, rectangle, port, or identified
child frame. XYFlow emits a node for each rendered endpoint kind and resolves
edges through a stable scene-element-to-node index. Connector `data` retains
the logical arrowhead names, bends, coordinate scale, grid, normalized fixed
points, and whether each endpoint anchor was explicit. The standard XYFlow
marker is used where it can represent the arrowhead; the logical name in
`data` remains authoritative for the other V1 arrowhead shapes.

The canonical scene records each node's semantic kind and semantic parent from
the resolved `Box` tree. XYFlow uses those IDs directly, so equal-size children,
visible overflow, and pure layout containers do not change the intended
hierarchy. Geometric containment is used only when reading a legacy scene that
does not contain semantic metadata. An overflowing node keeps its semantic
`parentId` but omits XYFlow's parent `extent`, so the target editor does not
clamp V1 `overflow="visible"` geometry back inside the parent.

A V1 connection between frames is represented by two editable local stubs in
the canonical scene for page-oriented formats. The source projection runs from
its endpoint to the logical frame edge with `to <destination frame ID>`; the
destination projection runs from its logical frame edge to the endpoint with
`from <source frame ID>`. Both stubs carry one logical connection ID, the two
original endpoint/frame IDs, and routing metadata including manual bends. The
XYFlow adapter uses that metadata to emit one edge, including source and
destination frame metadata, instead of exporting the two page projections.

## Isoflow

Isoflow output exports an Isoflow-compatible model:

```bash
xaligo render diagram.xal --format isoflow -o diagram.isoflow.json
```

The model follows the upstream Isoflow shape with `items`, `views`, `icons`,
`colors`, and view `rectangles` / `connectors`.
All frames remain in one model document; `--combine-frames` has no effect.

Items keep their existing IDs, order, and tile reservations. When a connection
references an AWS group, rectangle, port, or identified child frame, the
adapter additionally emits a generic Isoflow model/view item for that endpoint.
Two cross-frame page-link stubs are likewise combined by their shared logical
ID into one connector; neither `to ...` nor `from ...` becomes an Isoflow item.
For a same-frame connection with explicit bends, resolved interior points are
encoded as native Isoflow tile anchors. A built-in isometric fallback icon is
registered whenever an endpoint or source item has no target icon, so every
Isoflow item references a valid icon. The document uses the upstream
`fitToScreen` field.

Isoflow's connector schema has no arbitrary metadata field. Consequently,
V1-only connector kind, arrowhead names, original scale/grid values, and
Excalidraw fixed points are not added as private JSON extensions. The canonical
V1 scene remains the source of those semantics; Isoflow is a capability-based
projection and must not be used as an intermediate format when a future V2
frontend renders V1 input.
