---
applyTo: ".github/instructions/manual/**"
---

# 10.08.03 PPTX and routing: Route Connectors

### Route Connectors

Frozen V1 routes are headless in every format. Their effective
`start-arrowhead` and `end-arrowhead` must both resolve to `none` after
`<connections>` defaults and child aliases are merged. A non-`none` value is a
validation error rather than a renderer-specific circular endpoint.

Small circular route connector nodes remain a future versioned feature; they
must use a renderer-neutral connector-node concept instead of overloading V1
arrowheads.

Conceptual shape:

```text
[EC2] -- o -------- o -- [RDS]
```

Future behavior may render explicit connector nodes in SVG/PPTX and equivalent
editable shapes in Excalidraw. It is not part of the V1 compatibility profile.
