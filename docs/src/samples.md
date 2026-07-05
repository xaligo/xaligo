# Samples

## Hybrid Enterprise Architecture

![Hybrid enterprise architecture sample](../images/complex-hybrid-architecture.svg)

Source: [`examples/complex-hybrid-architecture.xal`](../../examples/complex-hybrid-architecture.xal)

```bash
xaligo render examples/complex-hybrid-architecture.xal \
  --format svg \
  --mode network \
  -o output/complex-hybrid-architecture.svg
```

## Isoflow Export

![Isoflow editor rendering the hybrid architecture sample](../images/isoflow-complex-hybrid-architecture.png)

Use the same `.xal` source to generate an Isoflow-compatible model:

```bash
xaligo render examples/complex-hybrid-architecture.xal \
  --format isoflow \
  --mode network \
  -o output/complex-hybrid-architecture.isoflow.json
```

## Route and Traffic Separation

Use `kind="route"` for structural paths without arrowheads, then add
`kind="traffic"` connections over the same endpoints for directional flows.
Traffic lines are drawn beside the matching route lane when possible.

Source: [`examples/route-traffic.xal`](../../examples/route-traffic.xal)

```bash
xaligo render examples/route-traffic.xal \
  --format svg \
  --mode network \
  -o output/route-traffic.svg
```

## Generated AWS Hierarchy

Generate a starter AWS hierarchy and render it:

```bash
xaligo generate xal --clouds 1 --accounts 1 --regions 2 --azs 2 \
  --az-layout staggered --subnets 2 --spacing both --start top \
  --paper A4 --orientation landscape -o output/infra.xal

xaligo render output/infra.xal \
  --format excalidraw \
  -o output/infra.excalidraw \
  --services examples/services.csv
```
