---
applyTo: ".github/instructions/manual/**"
---

# 06.09.02 Roadmap: Public API

### Public API

```go
RenderSVG()
RenderArtifacts()
RenderPPTX()
```

Current target API shape:

```go
Render(ctx, input, RenderOptions{Mode: mode, Format: format}) ([]byte, error)
Validate(ctx, input) error
```

Only `svg` and `pptx` are valid `RenderOptions.Format` values. Markdown uses
the SVG artifact API. The V1 scene builder is internal and must not be exposed
as a format convenience API.
