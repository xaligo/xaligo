---
applyTo: ".github/instructions/manual/**"
---

# 06.09.02 Roadmap: Public API

### Public API

```go
RenderSVG()
RenderArtifacts()
RenderPPTX()
RenderTerminal()
```

Current target API shape:

```go
Render(ctx, input, RenderOptions{Mode: mode, Format: format}) ([]byte, error)
Validate(ctx, input) error
```

`svg` and `pptx` accept V1 and V2. `terminal` accepts V2 only. Markdown uses the
SVG artifact API. The V1 scene builder is internal and must not be exposed as a
format convenience API.
