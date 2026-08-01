---
applyTo: ".github/instructions/manual/**"
---

# 04.09 Feature catalog: Group 9 — Validation, Diagnostics, Diff & External Integration (`XAL-9xxxxxx`)

## Group 9 — Validation, Diagnostics, Diff & External Integration (`XAL-9xxxxxx`)

| ID | Feature | Status | Summary |
|---|---|---|---|
| XAL-9000010 | Shared diagnostics use case | Implemented | One `DiagnosticsUsecase` (`Validate`/`Diagnose`) reused by both `xaligo validate` and the render pipeline's pre-flight checks. |
| XAL-9000020 | Source-positioned error reporting | Implemented | Validation and parse errors carry enough position/context information to be directly user-correctable. |
| XAL-9000030 | Strict-error vs. warning-fallback classification | Implemented | The `.xal` spec's documented split between hard validation errors and compatibility warning fallbacks (e.g., empty `align`, unknown nested attributes). |
| XAL-9000040 | Structural document diff engine | Implemented | `xaligo diff` compares parsed `.xal` data structures rather than source lines or formatting. |
| XAL-9000050 | Diff element matching strategy | Implemented | Matching prefers unique `id`/`name`/`ref`, then exact subtrees, then deterministic order-aware structural matching. |
| XAL-9000060 | Diff visual output | Implemented | Paired `-removed.svg`/`-added.svg` images highlighting removed/added elements and the old/new side of modified or moved elements. |
| XAL-9000070 | `@xaligo/xaligo-external` WASM/TypeScript package | Implemented | Exposes the PPTX draw-plan API through WebAssembly for browser/Node.js callers. |
| XAL-9000080 | TypeScript PPTX byte-output pipeline | Implemented | The external package consumes `BuildPPTXPlan` through WASM and produces PPTX bytes via PptxGenJS. |
| XAL-9000090 | `cmd/wasm` JavaScript/WASM adapter | Implemented | Exposes SVG rendering, PPTX-plan construction, and diagnostics to JavaScript/WASM hosts alongside the native CLI entry point. |
| XAL-9000100 | V1/V2 compatibility golden tests | Planned | Golden tests comparing V1-compatibility and native V2 engine output at the neutral-model and resolved-geometry boundaries once the V2 engine exists. |
| XAL-9000110 | Cross-renderer visual regression suite | Planned | Representative visual regression coverage across SVG, PPTX, and Markdown SVG embedding; currently limited per the roadmap's documented gaps. |
| XAL-9000120 | Render determinism and concurrency-safety guarantees | Planned | Byte-stable output for identical source/options/assets/environment, and parallel-job safety without shared mutable render state (Q13.1/Q13.2 backlog). |
