# Command Line

## Render

```bash
xaligo render <file.xal> --format <format> -o <output>
```

Supported formats:

| Format | Output |
|---|---|
| `svg` | Standalone SVG; one file per frame by default |
| `pptx` | PowerPoint presentation; one slide per frame by default |

The default is `svg`. Retired format names are rejected as unknown formats.
Markdown is handled by `xaligo render markdown`; it embeds SVG artifacts and is
not a separate `--format` value.

### Frames and physical pages

Identified child frames are physical pages for SVG and PPTX. They
are emitted in source order.

| Format | Multiple-frame default | `--combine-frames` |
|---|---|---|
| SVG | Separate `<output-stem>-<safe-frame-id>.svg` files | One SVG canvas |
| PPTX | One slide per frame in one presentation | One diagram slide |

Live SVG preview uses the combined canvas so every frame remains visible in
one browser view. Markdown follows the SVG artifact mapping.

For a one-frame SVG document, `-o` is the exact output filename. For several
frames, `-o output/diagram.svg` produces names such as
`output/diagram-overview.svg` and `output/diagram-detail.svg`. Safe frame IDs
retain ASCII letters, digits, `_`, and `-`; other character runs become `-`.
Leading and trailing `-` are removed, an empty result falls back to the frame's
source order, and a filename collision is an error. SVG output is not wrapped
in an implicit ZIP archive.

Common render flags:

| Flag | Description |
|---|---|
| `--mode standard|network|aws` | Accepted V1 rendering mode. These three values currently have identical V1 rendering semantics and share the same resolved 2D pipeline |
| `--theme light|dark` | Output theme |
| `--services <csv>` | Service metadata and label overrides |
| `--svg-legend-position top|right|bottom|left` | SVG legend placement |
| `--arrow-style thin|standard|triangle|stealth|arrow|diamond|oval|none` | SVG/PPTX plan default used when a connection omits its arrowhead. Explicit DSL arrowheads on normal/traffic connections and explicit stroke widths take precedence; routes require effective arrowheads to be `none` |
| `--combine-frames` | Combine all frames onto one SVG canvas or PPTX slide |

Any mode, format, theme, orientation, paper size, arrow-style option, or SVG
legend-position value outside its documented enum returns an error.

`--arrow-style` belongs to the shared physical plan used by SVG and PPTX.

Physical-page and PPTX flags:

| Flag | Description |
|---|---|
| `--paper A5|A4|A3|A2|A1|Letter|Legal|Tabloid` | Physical page/slide paper size |
| `--orientation portrait|landscape` | Physical page/slide orientation |
| `--paper-margin <inches>` | Margin applied before fitting |
| `--paper-margin-top/right/bottom/left <inches>` | Per-side margin override |
| `--px-per-inch <number>` | Layout scaling base |
| `--title`, `--author`, `--company`, `--subject` | PPTX presentation metadata |
| `--compression true|false` | PPTX compression |

`--title` sets package-level PPTX metadata. It is unrelated to the visible
`title` attribute on a page `<frame>` and never creates a frame tag.

## Structural Diff

```bash
xaligo diff before.xal after.xal -o output/architecture
```

The command compares parsed `.xal` data structures rather than source lines and
writes two SVG images:

- `output/architecture-removed.svg` renders the old document and highlights
  removed elements plus the old side of modified or moved elements in pale red.
- `output/architecture-added.svg` renders the new document and highlights
  added elements plus the new side of modified or moved elements in pale green.

XML formatting, comments, attribute order, parser-private attributes, and the
equivalent V1 forms with an omitted version or `version="1"` do not create a
diff when that version is on the document root. A `version` on an identified
child frame is visible page content and is compared normally, including the
literal value `1`. Matching prefers unique `name`, `ref`, and `id` values, then exact
subtrees, followed by deterministic order-aware structural matching. Give
elements an explicit `id`, `name`, or `ref` when moves must remain identifiable
across different parents.

`-o` is an output prefix, not a single output filename. A trailing `.svg` is
removed before `-removed.svg` and `-added.svg` are appended. The default prefix
is `output`. `--theme`, `--mode`, and `--px-per-inch` are applied identically to
both images. No difference is a successful result and still produces two
unhighlighted SVGs.

See the [structural diff sample](samples.md#structural-diff) for a complete
before/after pair and the generated images.

## Generate

Generate a starter `.xal` hierarchy:

```bash
xaligo generate xal -o generated.xal --paper A4 --orientation landscape
```

Useful generation flags:

| Flag | Default | Description |
|---|---:|---|
| `--clouds`, `--accounts`, `--regions`, `--azs` | `1`, `1`, `1`, `2` | AWS hierarchy counts |
| `--az-layout grid|staggered` | `grid` | Availability Zone placement |
| `--subnets` | `2` | Number of subnets |
| `--spacing vertical|horizontal|both` | `both` | Spacing direction |
| `--start top|left` | `top` | Drawing start position |
| `--paper` | `A4` | Paper size |
| `--orientation` | `landscape` | Page orientation |

Only `--output` is required. The generated file uses the canonical
`<xaligo version="1"><frames>...</frames></xaligo>` document envelope.

## SVG Icon Registry

The native CLI manages namespaced SVGs in the local `xaligo-assets.db` SQLite
registry. The first icon operation installs 13 domain-neutral icons in the
`builtin` namespace. Registrations are size-limited, safety-checked, and
normalized by the in-process Rust engine before SQLite indexes their name,
description, tags, and aliases with FTS5.

```bash
xaligo icon list --namespace builtin
xaligo icon search 'database OR storage'
xaligo icon get builtin:database -o database.svg
xaligo icon add router.svg --name network:router \
  --description 'Generic network router' --tag network --tag routing \
  --alias gateway --license MIT --source local
xaligo icon remove network:router
xaligo icon namespaces
```

Stable identities use `namespace:name`. `icon add` updates an existing identity
atomically, including its tags, aliases, and search row. `icon get` writes SVG
to standard output unless `-o` names a file. `icon search` accepts an FTS5
query and all list/search commands accept `--limit` up to 100. Configure the
database path with `paths.assets_db` in `etc/resources/aws/app.yaml`.

## Other Commands

| Command | Description |
|---|---|
| `xaligo diff <before.xal> <after.xal> -o <prefix>` | Render paired structural-diff SVGs |
| `xaligo validate <file.xal>` | Validate syntax, layout, and connection references |
| `xaligo render markdown <file.md>` | Embed rendered `xal` code blocks as SVG images into a Markdown file |
| `xaligo serve <file.xal\|file.md>` | Serve a live preview; `.xal` previews one combined SVG, `.md`/`.markdown` previews the full document with diagrams embedded inline; `--port` overrides the configured `serve.port` (default `8080`), and `--paper`/`--orientation` fix the preview to a physical page size |
| `xaligo icon <add|get|search|remove|list|namespaces>` | Manage the embedded SQLite SVG registry |
| `xaligo init [-o <dir>]` | Generate a sample `.xal` file |
| `xaligo version` | Print version |
