---
applyTo: ".github/instructions/manual/**"
---

# 06.01.05 Roadmap: Shared Rendering APIs

### Shared Rendering APIs

The shared in-repository use-case boundary should support at least:

```go
RenderSVG()
RenderArtifacts()
RenderPPTX()
RenderTerminal()
```

Prefer a shared extensible API underneath the convenience functions:

```go
Render(ctx, input, RenderOptions{Mode: mode, Format: format})
Validate(ctx, input)
```

`BuildScene()` may remain internal while V1 compatibility still lowers
through the legacy scene schema. It is not an output format or public product
contract. Markdown rendering composes `RenderArtifacts()`.
Terminal rendering consumes the resolved V2 document directly and rejects V1.
