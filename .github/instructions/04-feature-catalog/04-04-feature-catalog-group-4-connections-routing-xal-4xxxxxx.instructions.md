---
applyTo: ".github/instructions/manual/**"
---

# 04.04 Feature catalog: Group 4 — Connections & Routing (`XAL-4xxxxxx`)

## Group 4 — Connections & Routing (`XAL-4xxxxxx`)

| ID | Feature | Status | Summary |
|---|---|---|---|
| XAL-4000010 | `<connection>` orthogonal connector | Implemented | Elbowed arrow between items, groups, rectangles, ports, or frames, declared as a direct child of `<frame>` or `<connections>`. |
| XAL-4000020 | `<connections>` grouping tag | Implemented | Optional wrapper providing shared connector defaults (color, stroke, kind, arrowheads, scale, grid) inherited by its `<connection>` children. |
| XAL-4000030 | Endpoint binding by ID/name/ref | Implemented | `src`/`dst` resolve against a catalog ID, or the `id`, `name`, or `ref` of an item, AWS group, rectangle, port, or identified child frame. |
| XAL-4000040 | Explicit endpoint side selection | Implemented | `src-side`/`dst-side` pin a connector endpoint to a specific `top`/`right`/`bottom`/`left` side. |
| XAL-4000050 | Perimeter anchor slots | Implemented | `src-anchor`/`dst-anchor` select one of 20 fixed inset positions (5 per side) around an endpoint's perimeter. |
| XAL-4000060 | Route layer (`kind="route"`) | Implemented | Headless structural paths with no arrowheads, validated so every effective start/end arrowhead resolves to `none`. |
| XAL-4000070 | Traffic layer (`kind="traffic"`) | Implemented | Directional flow lines that share a lane beside a matching route when endpoints coincide. |
| XAL-4000080 | Manual bend points | Implemented | Explicit bend/via child tags or the `bends`/`points`/`via` inline coordinate alias in frame coordinates. |
| XAL-4000090 | Arrowhead style selection | Implemented | Independent `start-arrowhead`/`end-arrowhead` (and `arrowhead` alias) selection from `none|arrow|triangle|stealth|diamond|oval`. |
| XAL-4000100 | Stroke style, width, and color overrides | Implemented | `stroke-style`, `stroke-width`/`width`, and six-digit hex `color`, with documented per-kind defaults and alias precedence rules. |
| XAL-4000110 | Per-connection scale and grid snapping | Implemented | Positive `scale`/`coordinate-scale` bend-coordinate multiplier and per-connection `grid` snap size. |
| XAL-4000120 | Automatic route junctions | Implemented | Shared route trunks automatically fan out into junction points instead of independent overlapping lines. |
| XAL-4000130 | Cross-frame page links | Implemented | A connection whose endpoints span two frames renders as two local stubs labeled `to <frame>` / `from <frame>` instead of an impossible cross-page line. |
| XAL-4000140 | Cross-frame terminal side/anchor control | Implemented | `src/dst-frame-side` and `src/dst-frame-anchor` fix the logical page side and tangent slot used by a cross-frame page-link stub. |
| XAL-4000150 | Automatic safe-side selection | Implemented | Automatic page-terminal selection picks the nearest safe frame side from rendered visual geometry, honoring the frame metadata reservation strip. |
| XAL-4000160 | Lane offsetting for overlapping connectors | Implemented | Parallel connectors sharing the same X/Y lane are offset from one another to remain individually legible. |
| XAL-4000170 | Plan-level `--arrow-style` default | Implemented | A render-option arrowhead/width default applied to SVG/PPTX/PDF/Excel Plan output only when a connection omits its own explicit value. |
| XAL-4000180 | Explicit circular connector nodes | Planned | A future versioned routing model with dedicated circular/loopback connector nodes, distinct from today's headless V1 route connectors. |
| XAL-4000190 | Line Jump routing | Planned | Phase-2 roadmap routing feature: a small arc/hop where two unrelated crossing connectors overlap, keeping crossing lines visually distinguishable. |
| XAL-4000200 | Layer Routing | Planned | Phase-2 roadmap routing feature that groups related connectors into coherent routing layers/lanes instead of resolving each connector independently. |
