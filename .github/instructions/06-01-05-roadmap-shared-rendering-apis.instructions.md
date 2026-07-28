---
applyTo: ".github/instructions/manual/**"
---

# 06.01.05 Roadmap: Shared Rendering APIs

### Shared Rendering APIs

The shared in-repository use-case boundary should support at least:

```go
RenderSVG()
RenderArtifacts()
RenderExcalidraw()
RenderPPTX()
RenderPDF()
RenderExcel()
RenderXYFlow()
RenderIsoflow()
```

Prefer a shared extensible API underneath the convenience functions:

```go
Render(ctx, input, RenderOptions{Mode: mode, Format: format})
Validate(ctx, input)
```
