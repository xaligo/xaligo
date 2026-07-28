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
renderUsecase.RenderPDF(ctx, source, options)
renderUsecase.RenderExcel(ctx, source, options)

diagnosticsUsecase := NewDiagnosticsUsecase()
diagnosticsUsecase.Validate(ctx, source)
diagnosticsUsecase.Diagnose(ctx, source)
```

`RenderOptions.Assets` is only needed by embedded or virtual-filesystem
adapters. Native callers should leave it nil.
