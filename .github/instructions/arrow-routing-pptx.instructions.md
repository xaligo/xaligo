---
applyTo: "**/*.{go,ts,xal,md}"
---

# xaligo — PPTX Routing / Legend Preconditions

This file is the current source of truth for PPTX export geometry.

## Current Pipeline

```text
.xal DSL
  -> Go parser/layout
  -> Excalidraw scene JSON
  -> Go/WASM pptxplan.BuildPlanJSON
  -> TypeScript PptxGenJS drawing layer
  -> .pptx
```

Geometry belongs on the Go side. TypeScript should only translate the resolved
plan into PptxGenJS calls.

## Ownership

| Area | Owner |
|---|---|
| DSL parse/layout | `internal/parser`, `internal/layout` |
| Excalidraw scene and item metadata | `internal/excalidraw/scene.go` |
| PPTX geometry, paper scaling, routing, legend data | `internal/pptxplan` |
| Node invocation from Go | `internal/repository/pptx.go` |
| PptxGenJS drawing | `packages/xaligo/src/pptx.ts` |
| WASM bridge | `cmd/wasm/main.go` |

## Paper / Scaling

- PPTX export supports `--paper` and `--orientation`.
- A3 landscape is generated with:

```bash
.bin/xaligo generate pptx \
  --xal examples/sample.xal \
  --services examples/services.csv \
  -o out.pptx \
  --paper A3 \
  --orientation landscape
```

- Go `pptxplan` resolves paper size and computes the pixel-to-inch conversion.
- The `paper-frame` element remains the content frame for scaling.
- Root `<frame margin="N">` or `class="ma-N"` is content outer whitespace: it
  insets diagram content without shrinking the paper frame itself.

## Routing Rules

- Route calculation is in `internal/pptxplan/routing.go`.
- Obstacles include image and text rectangles from the Excalidraw scene.
- Start/end rectangles are excluded from obstacle checks for that connection.
- Binding `gap` from Excalidraw arrows must be honored in PPTX routing.
- If any obstacle-free candidate exists, obstacle-hitting candidates must not be
  selected.
- Lines on an obstacle boundary count as collision.
- Existing routed paths are included in scoring so later lines avoid overlap and
  near-parallel crowding.
- Previously placed line lanes are used as candidate offsets, so `--arrow-margin`
  affects routes that would otherwise share the same position.
- Final PPTX drawing order is:
  1. containers/shapes
  2. routed lines
  3. icons and labels

This order prevents lines from visually covering icons even at endpoints.

## Connector Style Options

`xaligo generate pptx` forwards all PPTX routing options:

| Flag | Meaning |
|---|---|
| `--arrow-style` | `thin`, `standard`, `triangle`, `stealth`, `arrow`, `diamond`, `oval`, `none` |
| `--arrow-stub` | Pixel stub before the first/last bend |
| `--arrow-margin` | Pixel margin reserved around existing line lanes |
| `--px-per-inch` | Layout scaling base, default 96 |

## Item Labels

- Item icon size defaults to 32px in native CLI config.
- Item label font is 8pt in PPTX output.
- Excalidraw font size for item labels is `8pt * 96 / 72 = 10.666...px`.
- Item label boxes are 14px high.
- Do not shrink label boxes to text metrics if it breaks PowerPoint placement.

## Layout / Whitespace

Supported whitespace controls:

| Syntax | Behavior |
|---|---|
| `<spacer />` / `<blank />` | Empty layout slot, not rendered |
| `<item />` | Empty item-grid slot, not rendered |
| `class="pa-4"` | Inner padding, Vuetify-style 8px unit |
| `class="ma-4"` | Outer margin; on root frame this becomes page-edge content whitespace |
| `margin="N"` and `margin-*` | Pixel margin |
| `content-width="N"` / `content-height="N"` | Shrinks usable inner layout area |
| `align="top-left"` etc. | Aligns the usable content area or item grid |
| `width="N"` / `height="N"` | Fixed child size, except root frame is the paper/content frame |

For item grids, horizontal `spread` is also supported.

## Legend Pages

PPTX export adds legend slides after the diagram slide when `--services` is
provided.

- Legend data is derived from `services.csv`.
- Only services actually used in the scene are included.
- The legend contains icon, abbreviation, and official name.
- Legend layout is fixed to 4 columns per slide.
- Additional legend slides may be created when entries exceed one slide.
- The diagram slide should not include an outside-frame legend; the PPTX legend
  belongs on separate slides.

## Verification Checklist

Before considering PPTX routing/layout changes complete:

```bash
go test ./...
make build
make build-wasm
npm run build --workspace packages/xaligo
.bin/xaligo generate pptx --xal examples/sample.xal --services examples/services.csv -o out.pptx --paper A3 --orientation landscape --arrow-style thin
unzip -t out.pptx
```

For icon-overlap regressions, inspect the resolved PPTX XML and ensure routed
custom geometry does not intersect target icon/label rectangles.
