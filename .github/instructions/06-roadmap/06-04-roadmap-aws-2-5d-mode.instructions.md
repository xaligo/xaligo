---
applyTo: ".github/instructions/manual/**"
---

# 06.04 Roadmap: AWS 2.5D Mode

## AWS 2.5D Mode

`mode: aws-2.5d` targets Cloudcraft and legacy AWS-reference-style oblique
architecture diagrams. It is a visual mode, not a standalone file format.

Required concepts:

- `plane` / `zone` layout primitives.
- Isometric-style nodes and routing.
- AWS node presets including `route53`, `cloudfront`, `elb`, `ec2`, `rds`, and
  `s3`.
- AWS Legacy / Cloudcraft-like themes.

Implement the first version in the native SVG renderer. WebView or GUI work may
learn from compatible 2.5D OSS projects, but the core representation must remain
usable without a specific UI framework.
