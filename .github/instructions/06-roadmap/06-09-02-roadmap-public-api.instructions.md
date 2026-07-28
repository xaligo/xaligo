---
applyTo: ".github/instructions/manual/**"
---

# 06.09.02 Roadmap: Public API

### Public API

```go
RenderExcalidraw()
RenderSVG()
RenderArtifacts()
RenderPPTX()
RenderPDF()
RenderExcel()
RenderXYFlow()
RenderIsoflow()
```

Current target API shape:

```go
Render(ctx, input, RenderOptions{Mode: mode, Format: format}) ([]byte, error)
Validate(ctx, input) error
```
