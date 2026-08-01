---
applyTo: ".github/instructions/manual/**"
---

# 06.01.04 Roadmap: Mode and Format Are Independent

### Mode and Format Are Independent

`mode` selects visual and layout semantics. `format` selects serialization or
the target integration. Do not encode a visual mode as a file format or assume
that one format has only one mode.

Target modes:

| Mode | Visual/layout intent |
|---|---|
| `standard` | Normal two-dimensional architecture diagrams |
| `network` | Route, traffic, circular connector, and topology-oriented diagrams |
| `aws` | AWS official-icon-oriented architecture diagrams |
| `aws-2.5d` | Cloudcraft/legacy AWS-reference-style oblique diagrams |
| `topology` | Instana/SkyWalking-style dependency topology |

Target formats:

| Format | Primary use |
|---|---|
| `svg` | Portable output and live preview |
| `pptx` | Editable presentation export; one frame per slide by default |

Markdown is a document workflow that embeds SVG artifacts; it is not a
separate engine format. Excalidraw, PDF, Excel/XLSX, XYFlow, and Isoflow are
retired and must not be reintroduced as aliases or compatibility outputs
without a new explicit product decision.

Target CLI shape:

```bash
xaligo render input.xal --mode network --format svg -o output.svg
xaligo render input.xal --mode aws-2.5d --format pptx -o output.pptx
```

Backward compatibility: omitting `--mode` must retain the current standard/AWS
behavior until an explicit default-mode migration is released.

Current V1 status: `standard`, `network`, and `aws` are accepted but have no
semantic difference; they execute the same resolved 2D pipeline. Treat them as
compatibility inputs until a versioned implementation introduces distinct
mode semantics. `aws-2.5d` and `topology` remain recognized but return a
not-implemented error.
