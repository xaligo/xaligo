---
applyTo: ".github/instructions/manual/**"
---

# 02.04 Agent guide: Shared Use-Case APIs

## Shared Use-Case APIs

Use constructor-injected components from `internal/usecase` instead of
assembling parser, layout, and encoder packages in adapters. Every direct
use-case file owns its `XxxUsecase` interface, private implementation,
`NewXxxUsecase` constructor, and receiver methods. The principal APIs are:

```go
renderUsecase := NewRenderUsecase(...)
renderUsecase.Render(ctx, source, options)
renderUsecase.RenderSVG(ctx, source, options)
renderUsecase.RenderArtifacts(ctx, source, options)
renderUsecase.RenderPPTX(ctx, source, options)

diagnosticsUsecase := NewDiagnosticsUsecase()
diagnosticsUsecase.Validate(ctx, source)
diagnosticsUsecase.Diagnose(ctx, source)
```

`RenderOptions.Assets` is only needed by embedded or virtual-filesystem
adapters. Native callers should leave it nil.

The only public render formats are `svg` and `pptx`. `BuildScene` is a
temporary internal V1 compatibility stage used to produce their shared draw
plan; it is not a public output API. Markdown rendering is an orchestration
flow over `RenderArtifacts`, not a third engine encoder.
