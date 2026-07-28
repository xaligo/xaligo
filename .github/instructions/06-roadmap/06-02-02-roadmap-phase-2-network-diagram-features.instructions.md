---
applyTo: ".github/instructions/manual/**"
---

# 06.02.02 Roadmap: Phase 2: Network Diagram Features

### Phase 2: Network Diagram Features

Status: headless V1 routes, the remaining routing steps, and textual connection
shorthands have initial shared implementations. Explicit circular connector
nodes remain future versioned work. Continue with hardening and cross-renderer
visual regression coverage.

Implement shared model/routing concepts in this order where dependencies allow:

1. Headless V1 route connectors; add explicit circular connector nodes only in
   a future versioned model.
2. Orthogonal Routing.
3. Route/Traffic separation.
4. Edge Offset.
5. Line Jump.
6. Layer Routing.
7. Junction generation.

These features must be shared across renderers where possible, rather than
implemented as PPTX-only corrections.
