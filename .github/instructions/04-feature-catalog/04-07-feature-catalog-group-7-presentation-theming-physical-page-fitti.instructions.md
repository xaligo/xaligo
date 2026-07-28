---
applyTo: ".github/instructions/manual/**"
---

# 04.07 Feature catalog: Group 7 — Presentation, Theming & Physical Page Fitting (`XAL-7xxxxxx`)

## Group 7 — Presentation, Theming & Physical Page Fitting (`XAL-7xxxxxx`)

| ID | Feature | Status | Summary |
|---|---|---|---|
| XAL-7000010 | Light/dark theme support | Implemented | `--theme light|dark` applied consistently across every output format. |
| XAL-7000020 | Physical paper size selection | Implemented | `--paper A5|A4|A3|A2|A1|Letter|Legal|Tabloid` for PPTX/PDF/Excel physical-page fitting. |
| XAL-7000030 | Paper orientation selection | Implemented | `--orientation portrait|landscape`, with auto-fit when omitted. |
| XAL-7000040 | Per-side paper margins | Implemented | `--paper-margin` (all sides) and `--paper-margin-top/right/bottom/left` (inches) applied before fitting frame content to the physical page. |
| XAL-7000050 | Layout scaling base | Implemented | `--px-per-inch` controls the pixel-to-inch scaling base used for PPTX/PDF/Excel layout. |
| XAL-7000060 | PPTX document metadata | Implemented | `--title`/`--author`/`--company`/`--subject` set PPTX package-level metadata, independent of any frame `title` attribute. |
| XAL-7000070 | PPTX compression control | Implemented | `--compression`/`--no-compression` toggles PPTX output compression. |
| XAL-7000080 | Rendering mode selection | Implemented | `--mode standard|network|aws`; all three currently share the same resolved 2D rendering pipeline. |
| XAL-7000090 | Roadmap-reserved rendering modes | Implemented | `aws-2.5d` and `topology` are recognized enum values that currently return a not-implemented error rather than an unknown-value error. |
| XAL-7000100 | Shared font-family catalog | Implemented | `virgil`, `helvetica`, `cascadia`, `assistant`, `excalifont`, `nunito`, `lilita-one`, `comic-shanns`, `liberation-sans` mapped consistently to each output format's font face. |
| XAL-7000110 | `aws-2.5d` rendering mode | Planned | Cloudcraft/legacy AWS-reference-style oblique diagrams with `plane`/`zone` layout primitives, isometric nodes/routing, AWS node presets (Route 53, CloudFront, ELB, EC2, RDS, S3), and AWS Legacy/Cloudcraft-like themes; currently a recognized but not-implemented `--mode` value. |
| XAL-7000120 | `topology` rendering mode | Planned | Instana/SkyWalking-style dependency topology view; currently a recognized but not-implemented `--mode` value. |
| XAL-7000130 | Distinct per-mode visual semantics | Planned | `standard`, `network`, and `aws` currently execute the identical resolved 2D pipeline; the roadmap targets genuinely distinct layout/visual semantics per mode. |
