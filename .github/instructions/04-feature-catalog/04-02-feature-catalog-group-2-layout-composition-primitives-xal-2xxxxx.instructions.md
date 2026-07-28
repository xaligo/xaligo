---
applyTo: ".github/instructions/manual/**"
---

# 04.02 Feature catalog: Group 2 — Layout & Composition Primitives (`XAL-2xxxxxx`)

## Group 2 — Layout & Composition Primitives (`XAL-2xxxxxx`)

| ID | Feature | Status | Summary |
|---|---|---|---|
| XAL-2000010 | `<container>` primitive | Implemented | Vertical-by-default stack container; `layout="horizontal"` arranges children side by side. |
| XAL-2000020 | `<row>`/`<col>` 12-column grid | Implemented | Responsive grid primitive; `<col span="N">` occupies `N` of 12 columns, with flexible spans auto-dividing remaining columns. |
| XAL-2000030 | Fixed vs. flexible child allocation | Implemented | Explicit main-axis `width`/`height` reserves fixed space first; remaining children divide leftover space by `row`/`col` weight. |
| XAL-2000040 | Generic leaf tag rendering | Implemented | An unknown childless tag renders as a rectangle plus text, using `title`, text content, or the tag name as its label. |
| XAL-2000050 | Generic group/container tag rendering | Implemented | An unknown tag with layout children becomes a titled group, laid out vertically, horizontally, or with the V1 staggered layout. |
| XAL-2000060 | Item-grid row behavior | Implemented | A group whose children are all item-like (`item`, `spacer`, `blank`) automatically switches to item-grid row layout. |
| XAL-2000070 | `<rectangle>` general-purpose shape | Implemented | Titled/labeled rectangle that may contain multiple `<port>` children, unlike other generic leaf tags. |
| XAL-2000080 | `<port>` side-anchored sub-rectangle | Implemented | Small labeled rectangle bound to a side of its parent `<rectangle>`, with optional explicit `x`/`y` clamped inside the parent. |
| XAL-2000090 | `<item>` AWS service-icon leaf | Implemented | Places a catalog service icon by numeric ID, with configurable `item-size` and `dx`/`dy` offset; missing icons warn and skip instead of failing. |
| XAL-2000100 | `<spacer>`/`<blank>` empty slots | Implemented | Dedicated empty layout tags that occupy a grid slot without rendering an icon, border, or label. |
| XAL-2000110 | Custom leaf display toggles | Implemented | `border="none"` hides a leaf/group border; `visible="false"` hides one component's border/icon/label while preserving its layout space. |
| XAL-2000120 | `<capture>` annotation group tag | Implemented | Border-only structural child container that participates in normal nested layout; connectable by id/name/ref like any other group tag, including cross-frame page-link stubs, without AWS/architectural semantics. |
