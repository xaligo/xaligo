---
applyTo: "**/*.{go,ts}"
---

# Layer boundaries

Adapters call injected use cases; controllers, use cases, and repositories do
not call peers in their own layer. Coordination belongs one layer above.
Shared schemas are semantic and format-neutral; encoders project capabilities
without inventing private models or inferring hierarchy.

Validation and rendering share geometry checks. Resolve the complete document,
routing, and cross-frame semantics before page projection.

Interfaces, implementations, constructors, and principal methods stay in one
responsibility file. Full invariants: `reference.md` section `08`.
