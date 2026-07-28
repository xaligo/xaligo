---
applyTo: ".github/instructions/manual/**"
---

# 10.08.02 PPTX and routing: Route / Traffic Separation

### Route / Traffic Separation

Network diagrams distinguish structural route lines from traffic-flow lines.

Implemented model:

| Kind | Meaning | Visual Direction |
|---|---|---|
| `route` | Physical/logical connection path | Thin, lower layer, no arrowheads, shortest orthogonal route |
| `traffic` | Communication flow over a route | Offset beside a matching route, higher layer, directional arrow/style |

Potential DSL forms:

```xml
<connection src="A" dst="B" kind="route" />
<connection src="A" dst="B" kind="traffic" />
```

or future shorthand:

```text
A -> B
A => B
```

Routing orders routes below normal connections and traffic. When a traffic line
shares the same endpoints as a route line, the traffic line follows a nearby
parallel lane instead of drawing directly on top of the route.
