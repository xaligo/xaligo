# Detailed instruction reference

Read sections in numeric order when the whole reference is required. For normal tasks, use `README.md` and open only the named sections.

1. `01 General`
2. `02 Agent guide`
3. `03 Development flow`
4. `04 Feature catalog`
5. `05 Issues and quality`
6. `06 Roadmap`
7. `07 XAL specification`
8. `08 Architecture`
9. `09 Coding`
10. `10 PPTX and routing`
11. `11 Diagram creation`

---

# 01 General


## Project

`xaligo` is a Go CLI and WebAssembly/TypeScript package that converts the
`.xal` diagram DSL to Excalidraw, SVG, PPTX, PDF, Excel, XYFlow, and Isoflow
outputs.

```text
module: github.com/xaligo/xaligo
Go:     1.26
```

Read `development-flow.instructions.md` for task slicing, verification, and
local commit workflow; `roadmap.instructions.md` for product direction;
`xal-spec.instructions.md` for DSL behavior; and
`architecture.instructions.md` for implementation boundaries. Read
`coding.instructions.md` before changing Go or TypeScript source.

## Directory structure

```text
xaligo/
├── cmd/
│   ├── main.go                  native CLI entry point
│   └── wasm/main.go             JavaScript/WASM adapter
├── internal/
│   ├── command.go               root Cobra command assembly
│   ├── controller/              CLI flags and file-I/O adapters
│   ├── entity/                  internal structures; independent entity layer
│   ├── usecase/
│   │   ├── render.go            RenderUsecase, dispatch, and orchestration
│   │   ├── diff.go              DiffUsecase, structural comparison, and paired SVG orchestration
│   │   ├── diagnostics.go       DiagnosticsUsecase and shared validation
│   │   ├── scene_io.go          SceneIOUsecase for editable scene persistence
│   │   ├── catalog.go           CatalogUsecase for service metadata
│   │   ├── export.go            ExportUsecase for persisted PPTX output
│   │   ├── parser.go            ParserUsecase over the V1 engine
│   │   ├── layout.go            LayoutUsecase over the V1 engine
│   │   ├── element.go           ElementUsecase over the V1 engine
│   │   ├── pagination.go        PaginationUsecase over the V1 engine
│   │   ├── plan.go              PlanUsecase over the V1 engine
│   │   ├── scene.go             SceneUsecase and repository-port adapter
│   │   ├── theme.go             ThemeUsecase over the V1 engine
│   │   └── v1/engine/           synchronous V1 parser/layout/scene/plan logic
│   ├── repository/              filesystem and output-format adapters
│   └── config/                  project configuration
├── test/
│   ├── unit/                    unit tests mirroring the project tree
│   └── integration/             black-box use-case/adapter tests
├── external/                    TypeScript external adapter layer
│   ├── index.ts                 package API composition boundary
│   ├── command.ts               TypeScript CLI entry point
│   ├── controller/              CLI argument and file-I/O adapters
│   ├── entity/                  TypeScript API and PPTX plan types
│   ├── repository/              WASM, PptxGenJS, and package adapters
│   └── usecase/                 independent TypeScript application use cases
├── etc/resources/aws/           catalogs, templates, icons, attribution
├── docs/src/examples/samples/   example .xal and services CSV files
├── scripts/                     asset/catalog generation scripts
├── docs/src/images/             documentation and README gallery assets
├── Makefile
├── go.mod / go.sum
└── README.md
```

The repository root contains no Go source files. Executable adapters belong in
`cmd`; application implementation belongs in `internal`. This repository's
external integration boundary is the CLI, HTTP/SSE preview protocol, and WASM
adapter rather than an importable public Go package.

## Architecture rules

- Preserve `.xal -> parser -> layout -> shared scene/plan -> encoder`.
- Format-rendering adapters call `internal/usecase`; they do not create
  parallel parser or layout pipelines.
- Input/output-format-specific encoding and persistence belong to
  `internal/repository`; use-case filenames describe processing, not formats.
- `internal/entity` owns structures exchanged between layers and contains no
  application orchestration. Shared value helpers such as theme names and
  service labels may live here when they are renderer-independent.
- Calculation and orchestration belong under `internal/usecase`. Synchronous
  calculations live in `internal/usecase/v1/engine`; repository
  I/O, cancellation, stage ordering, and concurrency control remain in the
  parent `internal/usecase` package.
- `v1/engine` must not import repositories, start goroutines, own worker
  pools, or interpret contexts. Parallel execution is a caller-owned policy
  over independent engine jobs.
- Keep mode (visual semantics) independent from format (serialization).
- Keep cross-format routing and geometry in shared layers.
- `cmd` imports command/adapters only; business logic stays outside entry points.
- Native and embedded environments differ through `usecase.AssetSource`, not
  through duplicated render implementations.
- Go constructs PPTX draw plans; the configured WASM/PptxGenJS adapter writes
  PPTX bytes. Do not add a second OOXML writer.
- Return context-wrapped errors. Do not panic in core code.

## Testing rules

- Put unit tests under `test/unit`, mirroring the source tree they cover.
- Put black-box tests of exported APIs and adapters in `test/integration`.
- Prefer externally observable behavior over package-private helper assertions
  when moving tests outside implementation packages.
- Add focused coverage for behavior changes and preserve regression tests.

## Assets and configuration

- Configuration: `etc/resources/aws/app.yaml`
- ID lookup: `etc/resources/aws/service-index.csv`
- Full catalog: `etc/resources/aws/service-catalog.csv`
- Embedded assets: `etc/resources/aws/assets.go`
- SVG assets: `etc/resources/aws/svg`
- Isoflow manifest: `etc/resources/aws/isoflow-icons.json`

Preserve bundled license and attribution files. Generated assets must be
refreshed through the scripts declared in the root `package.json`.

The root `package-lock.json` is the canonical lock for the npm workspace.
Commit it with dependency changes, use `npm ci --ignore-scripts` for
reproducible builds, and do not commit a separate `external/package-lock.json`.

`VERSION` and the root `package.json` contain the next stable `X.Y.Z` version.
Release metadata is resolved by `scripts/build/release-metadata.sh`. A main
prerelease run `N` uses `X.Y.Z-N` for the embedded native CLI, exactly `X.Y.Z`
from `VERSION` for npm, `X.Y.Z~N` for Debian, and RPM `Version: X.Y.Z` with
`Release: 0.N`; a stable RPM uses release `1`. Source branch names such as
`main` must not appear in a package name or package version. Keep these values
separate so native and OS prereleases sort before the corresponding stable
package and remain valid for each package manager. Publish npm only from the
stable release workflow because npm package versions are immutable.

## Conventions

- Run `gofmt` on changed Go files.
- Use lowercase single-word package names.
- Organize files by implementation responsibility, not declaration kind.
  Interfaces and constructors/factories belong in the file containing the
  concrete implementation they describe or create; do not add declaration-only
  `interface*` or `constructor*` files.
- Put a component's interface, unexported concrete type, and constructor near
  the beginning of its responsibility file, before its implementation methods.
  Do not collect unrelated component methods in a package-wide facade.
- Do not repeat a package layer in Go filenames. Use `<component>.go` and
  `<component>_<detail>.go`; express the layer in the exported interface and
  constructor suffix instead.
- Wrap errors with `fmt.Errorf("context: %w", err)`.
- Represent Excalidraw elements as `map[string]interface{}` for format
  compatibility.
- Do not commit binaries, dependencies, caches, `output`, WASM artifacts, or
  TypeScript `dist` output. Checked-in documentation SVGs generated from
  `docs/src/architecture/*.xal` are the explicit exception; commit each SVG
  with its `.xal` source.

## Verification

Set up the repository-pinned security scanner and npm audit metadata once:

```bash
make security-setup
```

Run the security gate before every commit, followed by the relevant tests and
builds:

```bash
make security-check
go test ./...
go build ./...
npm ci --ignore-scripts
npm run build --workspace=@xaligo/xaligo-external
npm --prefix external run build:pptx-exporter-wasm
git diff --check
```

---

# 02 Agent guide


Use this file as the repository working agreement. Read the following
preconditions before changing code:

1. `feature-catalog.instructions.md` — ID-addressable catalog of xaligo's
   supported features; check it before treating a request as new scope.
2. `development-flow.instructions.md` — task slicing, verification, and local
   commit workflow.
3. `issues.instructions.md` — open quality and feature-hardening issues,
  design quality, and completion gates.
4. `roadmap.instructions.md` — product and pipeline direction.
5. `xal-spec.instructions.md` — authoritative `.xal` behavior.
6. `architecture.instructions.md` — package boundaries and dependency rules.
7. `coding.instructions.md` — mandatory file and identifier conventions.

## Project summary

- Go 1.26 module: `github.com/xaligo/xaligo`
- CLI entry point: `cmd/main.go`
- PPTX exporter WASM entry point: `external/command.ts`
- TypeScript package and implementation: `external`
- Shared application boundary: `internal/usecase`
- Generated CLI: `.bin/xaligo`
- Generated PPTX exporter WASM: `external/wasm/xaligo.wasm`

## Working rules

- Preserve `.xal -> parser -> layout -> shared scene/plan -> encoder`.
- CLI, preview, and WASM format-rendering paths call `internal/usecase`. They
  do not build a parallel parser/layout/render pipeline. Focused `add` and
  source-generation utilities may use repositories/builders directly.
- Keep mode and format independent.
- Put cross-format routing and geometry in shared layers.
- Return wrapped errors; do not panic in core code.
- Preserve unrelated and pre-existing working-tree changes.
- Do not commit build output, dependencies, binaries, or caches. Checked-in
  documentation SVGs are regenerated from and committed with their `.xal`
  sources.
- Add focused tests with every behavior change.

## Common commands

```bash
# Build and test
go build ./...
go test ./...

# Build distributable adapters
make build
make build-wasm
npm ci --ignore-scripts
npm run build --workspace=@xaligo/xaligo-external

# Render and validate
.bin/xaligo validate docs/src/examples/samples/sample.xal
.bin/xaligo render docs/src/examples/samples/sample.xal --format excalidraw -o output/sample.excalidraw
.bin/xaligo render docs/src/examples/samples/sample.xal --format svg -o output/sample.svg
.bin/xaligo render docs/src/examples/samples/sample.xal --format pdf -o output/sample.pdf
.bin/xaligo render docs/src/examples/samples/sample.xal --format excel -o output/sample.xlsx
.bin/xaligo render docs/src/examples/samples/sample.xal --format xyflow -o output/sample.xyflow.json
.bin/xaligo render docs/src/examples/samples/sample.xal --format isoflow -o output/sample.isoflow.json
.bin/xaligo serve docs/src/examples/samples/sample.xal --mode network

# Clean generated artifacts
make clean
```

Native PPTX export additionally requires the configured `xaligo.wasm` PPTX exporter.
The TypeScript package consumes `BuildPPTXPlan` through WASM and creates PPTX
with PptxGenJS.

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

## Asset workflow

- Quick ID lookup: `etc/resources/aws/service-index.csv`
- Full catalog: `etc/resources/aws/service-catalog.csv`
- Embedded asset declaration: `etc/resources/aws/assets.go`
- AWS/Tabler/Yamaha SVGs: `etc/resources/aws/svg`
- Isoflow icon manifest: `etc/resources/aws/isoflow-icons.json`

Use `npm run import:tabler-icons`, `npm run import:yamaha-icons`, or
`npm run generate:isoflow-icons` to refresh generated catalogs. Preserve the
bundled license and attribution files.

## Services CSV

The accepted columns are:

```text
id,OfficialName,Abbreviation,Summary,Usage,Notes
```

Pass its in-memory bytes through `RenderOptions.ServicesCSV`, or use
`--services` in the CLI. Catalog IDs and abbreviations are shared by all
renderers.

## Completion checklist

1. Format changed Go files with `gofmt`.
2. Run `go test ./...` and `go build ./...`.
3. For shared render use-case or asset changes, cross-build `cmd/wasm`.
4. For TypeScript-facing changes, build `external` via `npm run build --workspace=@xaligo/xaligo-external`.
5. Run `git diff --check` and inspect `git status --short`.
6. Update the DSL spec, architecture, README, or roadmap when their contract
   changed.

Unit tests belong in `test/unit`, mirroring the source tree they cover.
Black-box API and adapter tests belong in `test/integration`. Prefer testing
observable behavior over exposing package-private helpers only for tests.

---

# 03 Development flow


Read this file before planning, editing, reviewing, generating, testing, or
committing repository changes. The goal is a traceable history made of the
smallest coherent changes that remain safe to review, test, revert, and build.

## Authorization and Scope

- An implementation or change request authorizes incremental local commits for
  its in-scope changes unless the user explicitly requests no commits. A review,
  diagnosis, explanation, or status request remains read-only and does not
  authorize a commit.
- A local commit does not authorize pushing, force-pushing, publishing,
  tagging, opening a pull request, or rewriting existing history. Those actions
  require an explicit request.
- Treat durable user decisions as repository preconditions. Record them in the
  applicable instruction file during the same task instead of leaving them
  only in conversation history.

## Before Changing Files

1. Read `general.instructions.md` and every instruction whose `applyTo` matches
   the current scope. Re-evaluate matching instructions whenever the scope
   expands.
2. Use RTK for repository operations when an RTK wrapper exists, especially
  noisy commands such as `git`, `rg`, `go`, `npm`, tests, diffs, logs, and
  dependency inspection. Prefer `rtk <tool>` or another RTK-filtered command
  shape to reduce token-heavy output while preserving enough diagnostic
  signal. If RTK is missing in a local environment, install it with
  `curl -fsSL https://raw.githubusercontent.com/rtk-ai/rtk/refs/heads/master/install.sh | sh`
  before continuing unless the task is urgent or the environment disallows
  network installs.
3. Inspect `git status --short`, the unstaged diff, and the staged diff before
   editing. Existing changes belong to the user unless the current task clearly
   created them.
4. Identify the affected contracts, implementation layers, tests,
   documentation, generated source-controlled assets, and verification
   commands before choosing commit boundaries.
5. Reproduce a reported defect or establish an observable invariant before
   changing code. Prefer a regression test that fails for the diagnosed cause.
6. Fix the earliest shared layer that owns the information. Do not hide an
   engine or contract defect by changing only one sample, diagram, controller,
   or output encoder.

When one occurrence reveals a structural pattern, audit the corresponding
controller, use-case, repository, entity, V1-engine, and external TypeScript
layers as applicable. Examples include peer-layer dependencies, misplaced
interfaces or constructors, duplicate dispatch, stale format-specific names,
and inconsistent validation/render paths.

## Change Slices

Split work into the smallest cohesive slices that provide an independently
understandable outcome. Good boundaries include:

- one precondition or coding-rule change;
- one entity or cross-layer contract change;
- one synchronous V1-engine responsibility with its focused tests;
- one root use-case component or orchestration boundary;
- one repository or output-format behavior;
- one controller, command, or composition-root migration;
- one external TypeScript command/controller/use-case/repository boundary;
- one user-facing specification or documentation topic; and
- one architecture diagram source together with its rendered documentation
  asset.

Separate a mechanical rename or move from a behavior change when each is valid
on its own. Keep a rename, its required reference updates, and any deletion of
the superseded file together when splitting them would leave broken imports,
links, or duplicate definitions.

Use dependency order where possible:

```text
contract/entity
  -> synchronous engine
  -> use-case orchestration
  -> repository/format adapter
  -> controller/command/composition root
  -> documentation and source-controlled derived assets
```

Keep implementation and its focused regression tests in the same commit.
Keep a checked-in derived documentation asset with the source that generates
it. A contract-changing implementation and its authoritative specification
must be committed together when separating them would make either commit
misleading; otherwise documentation may be a separate immediately following
commit.

Prefer every commit to build and pass its focused tests. If an interface or
constructor signature creates an inherently coupled cross-layer cutover, keep
that cutover in one atomic commit. Do not create a knowingly broken
intermediate commit, and do not add temporary aliases or compatibility facades
solely to increase the number of commits.

Do not split changes mechanically by line count or filename. A large cohesive
engine move may be one commit, while unrelated hunks in one file must be split
with patch staging.

## Commit Cadence

Repository-changing work is committed incrementally unless the user explicitly
requests no commits. Commit each completed slice as soon as its focused
verification succeeds; do not wait until the entire request is finished.
Long-running work should therefore produce several small, ordered commits
rather than one end-of-task snapshot.

Before each commit:

1. Stage explicit paths or hunks. Prefer `git add -p` for mixed files and avoid
   broad staging in a dirty worktree.
2. Inspect `git diff --cached --name-status`, `git diff --cached --stat`, and
   `git diff --cached` to confirm that the index contains one responsibility
   and no unrelated user changes.
3. Run `git diff --cached --check`.
4. Run `make security-check`. Security scanning is a mandatory commit
   precondition for every change, including documentation-only changes. Run
   `make security-setup` once after cloning or whenever the pinned scanner
   version changes.
5. Run the narrowest relevant test, build, validation, or render command.
6. Use the existing concise subject style, normally
   `<type>: <imperative summary>`, and describe the outcome rather than the
   edited filenames.
7. Re-run `git status --short` after the commit and confirm that no unintended
   path remains staged.

Do not create empty checkpoint or `WIP` commits. Do not amend, squash, rebase,
reset, or otherwise rewrite existing history unless the user explicitly asks.
Do not push commits or publish artifacts unless the user explicitly asks.

## Working-Tree Safety

- Preserve unrelated tracked and untracked changes. Never discard, reset, or
  stash them merely to simplify the current task.
- Do not stage another task's changes unless the user explicitly asks to commit
  all current changes.
- When the user explicitly asks to commit all current changes, inventory every
  modified, deleted, and untracked path, classify the complete set by cohesive
  responsibility, and commit those groups in dependency order. Do not turn
  that instruction into one catch-all stage or commit.
- If the current task overlaps an existing edit, inspect the combined diff and
  preserve both intents. Stop for direction only when they cannot be reconciled
  safely.
- When work is delegated in parallel, give workers non-overlapping edit scopes.
  One coordinating owner performs staging and commits serially so concurrent
  index changes cannot mix responsibilities.
- Do not commit dependencies, caches, binaries, package `dist`, `output`, or
  WASM build artifacts. Source-controlled documentation images are allowed
  only when the repository intentionally tracks them and their editable source
  is committed in the same slice.

## Verification Cadence

Use layered verification rather than running every expensive command after
every hunk:

- run focused unit or adapter tests before the commit that introduces the
  behavior;
- run package-level or integration tests after a structural cutover;
- run the broad relevant suite at implementation milestones and before final
  handoff; and
- rerun verification after generators or formatters in case they changed the
  worktree.

Typical final checks are selected by scope:

```bash
make security-check
go test ./... -count=1
go build ./...
npm --prefix external test
mdbook build docs
git diff --check
git status --short
```

For `.xal` changes, validate and render the affected source with the same V1
pipeline used in production. For PPTX changes, also exercise the external
exporter tests and package validation when the required WASM/tooling is
available. A sandbox, missing optional tool, or pre-existing failure must be
reported explicitly; it must not be represented as a successful check.

When a shared DSL, layout, routing, scene, or renderer contract changes, verify
the relevant format matrix and representative documentation `.xal` corpus, not
only the one source that originally exposed the defect.

## Documentation and Architecture Diagrams

- Documentation about internals links to the actual implementation files or
  functions. Re-check links after moves, renames, or line-sensitive edits.
- Design-first documentation distinguishes current behavior from a planned
  target. It must not describe an unimplemented target as the current system.
- Distinguish dependency construction from runtime data flow, and include every
  real format path relevant to the documented behavior.
- Source-controlled architecture diagrams are authored as `.xal`; do not edit
  their generated SVG as the source of truth.
- Commit a diagram's `.xal` source and rendered SVG together, then build the
  documentation.
- In the internal pipeline diagram, preserve the hierarchy
  `xaligo -> internal/external -> Main/Other`. `Main` follows
  `command -> controller -> usecase -> repository` from top to bottom; command
  is an entry point. Packages are nested groups, functions are rectangles, and
  conceptual data is represented by ports. `Other` contains entities,
  configuration, shared utilities, tools, and generated artifacts.
- Check generated diagrams for group containment, sibling overlap, port-label
  overflow, and stale or missing package paths. A local size adjustment is safe
  only when the shared layout invariant has already been verified or separately
  recorded as a structural issue.

## Completion Audit

Before handing work back:

1. Confirm every task-owned tracked and untracked change is committed.
2. Confirm the worktree has no unexpected staged or unstaged changes.
3. Review the new commit sequence with `git log --oneline` for ordering,
   granularity, and clear subjects.
4. Report the commit count or range, verification results, any intentionally
   uncommitted user changes, and whether anything was not run.

---

# 04 Feature catalog


This file is the authoritative, ID-addressable catalog of xaligo's supported
and planned features. Read it as a precondition alongside
`agent-guide.instructions.md` to understand what the product already does (and
what it has already committed to doing) before proposing new work, filing a
roadmap entry, or judging whether a request is a bug fix, an extension of an
existing feature, genuinely new scope, or an already-tracked planned item.

Each row is a stable 7-digit feature ID (`XAL-GNNNNNN`), so implementation
notes, commit messages, roadmap entries, and issue reports can reference a
feature without repeating its full description. IDs are never reused or
renumbered once assigned; a removed feature's ID is retired, not recycled.

Every row carries a `Status` column:

- `Implemented` — shipped and available today.
- `Planned` — not yet implemented; tracked in `roadmap.instructions.md` and/or
  the `issues.instructions.md` Q05 backlog, which is the
  authoritative source for its exact sequencing and scope.
- `Excluded unless justified` — a considered but deliberately unsupported
  capability that stays out of scope until a non-substitutable use case is
  identified.

Do not remove or renumber a row when its status changes; update its `Status`
and Summary in place instead (e.g., `Planned` -> `Implemented` once a feature
ships).

- The leading digit `G` is the group (major capability area); see the section
  headers below.
- The remaining 6 digits are a per-group sequence number in steps of 10
  (`000010`, `000020`, ...), leaving room to insert a new fine-grained feature
  between two existing ones without renumbering the group.
- Add a new feature at the next free step within its group, or open a new
  group (next leading digit) for a capability area not covered below.

## Group 1 — Core DSL & Document Envelope (`XAL-1xxxxxx`)

| ID | Feature | Status | Summary |
|---|---|---|---|
| XAL-1000010 | Vue-style `.xal` XML DSL | Implemented | Vue-style layout DSL with XML syntax, parsed with `encoding/xml`, handling attributes, nested tags, and text content. |
| XAL-1000020 | Canonical V1 envelope | Implemented | `<xaligo version="1">` containing document-wide `<data>` and exactly one `<frames>` page collection. |
| XAL-1000030 | Legacy root compatibility | Implemented | Historical `<frame>`/`<frames>` document roots remain readable and renderable but emit a migration warning recommending the canonical envelope. |
| XAL-1000040 | V2 reject-safe root boundary | Implemented | `<scene version="2">` is a distinct root; a V1 reader rejects it outright instead of partially rendering V2 syntax as V1. |
| XAL-1000050 | Shared V1 pipeline | Implemented | One `.xal -> parser -> layout -> shared scene/plan -> encoder` pipeline shared by every output format, with no per-format parallel parser or layout path. |
| XAL-1000060 | `<frames>` page collection | Implemented | Groups page `<frame>` children with `gap` and optional `layout="vertical"`; every child `<frame>` requires a non-empty `id`. |
| XAL-1000070 | Frame sizing and spacing attributes | Implemented | `width`, `height`, `class`, `margin`/`margin-*`, `content-width`/`content-height`, `align`, and `item-size` on a page or nested frame. |
| XAL-1000080 | Numeric and geometry validation contract | Implemented | Finite-number and per-attribute domain checks (positive sizes, span ≤ 12, non-negative gaps, etc.) applied before layout, shared by `validate` and every render format. |
| XAL-1000090 | Child containment overflow policy | Implemented | `overflow="error"` (default, rejects out-of-bounds children) vs `overflow="visible"` (permits them while keeping sibling cursors deterministic). |
| XAL-1000100 | Frame metadata tag band | Implemented | `<metadata>`/`<entry>` key-value tag band exposing `id`, `title`, page-content `version`, and custom entries as a reserved strip other content must avoid. |
| XAL-1000110 | Page-content `version` attribute | Implemented | A non-empty `version` on an identified child `<frame>` is the page's visible content revision, distinct from the document-root DSL version selector. |
| XAL-1000120 | Document-level `<data>` container | Implemented | Reusable definitions (UML models, database schemas) declared once under `<data>` and referenced by `data="id"` from one or more frames. |
| XAL-1000130 | Resolved shared text-layout policy | Implemented | Role-based wrap/fit/overflow/line-height rules (group header, ordinary label, item label, port label, connector label) applied consistently by every encoder. |
| XAL-1000140 | V2 `<scene version="2">` native engine | Planned | A distinct, reject-safe V2 root and native frontend that lowers directly to the same typed version-neutral model as V1, sharing renderers/encoders downstream; V1 remains a frozen compatibility profile inside V2. |
| XAL-1000150 | Typed normalized layout structure | Planned | Replace repeated reads from `Node.Attrs` with first-class resolved-layout data so validation and rendering share one geometry source of truth. |
| XAL-1000160 | Format-neutral shared scene/plan schema | Planned | Migrate the canonical scene/plan schema away from Excalidraw/PPTX-shaped fields entirely, keeping only compatibility aliases for existing public callers. |

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

## Group 3 — AWS Architecture Primitives & Icon Catalog (`XAL-3xxxxxx`)

| ID | Feature | Status | Summary |
|---|---|---|---|
| XAL-3000010 | `aws-cloud` group tag | Implemented | Top-level AWS Cloud boundary grouping primitive with automatic containment layout. |
| XAL-3000020 | `region` group tag | Implemented | AWS Region boundary grouping primitive. |
| XAL-3000030 | `vpc` group tag | Implemented | VPC boundary grouping primitive. |
| XAL-3000040 | `availability-zone` group tag | Implemented | Availability Zone boundary grouping primitive. |
| XAL-3000050 | `public-subnet` / `private-subnet` group tags | Implemented | Subnet-scoped grouping primitives for AZ-local resources. |
| XAL-3000060 | `security-group` group tag | Implemented | Grouping primitive for resources sharing an EC2 security group. |
| XAL-3000070 | `auto-scaling-group` group tag | Implemented | Grouping primitive for an EC2 Auto Scaling group. |
| XAL-3000080 | `generic-group` tag | Implemented | Non-AWS-specific logical grouping primitive for content that does not match a dedicated AWS group tag. |
| XAL-3000090 | AWS Architecture-Service-Icons catalog | Implemented | Bundled AWS service icon set looked up by numeric catalog ID from `service-catalog.csv`. |
| XAL-3000100 | Tabler icon catalog | Implemented | Imported and cataloged Tabler icon set, refreshed via `npm run import:tabler-icons`. |
| XAL-3000110 | Yamaha icon catalog | Implemented | Imported and cataloged Yamaha icon set, refreshed via `npm run import:yamaha-icons`. |
| XAL-3000120 | Isoflow icon manifest | Implemented | `isoflow-icons.json` manifest mapping catalog icons to Isoflow-compatible icon references, refreshed via `npm run generate:isoflow-icons`. |
| XAL-3000130 | Service ID lookup catalogs | Implemented | `service-index.csv` for quick ID lookup and `service-catalog.csv` for the full catalog with embedded icon data. |
| XAL-3000140 | `services.csv` label overrides | Implemented | `id,OfficialName,Abbreviation,Summary,Usage,Notes` per-diagram service list driving icon abbreviation overrides. |
| XAL-3000150 | Service legend generation | Implemented | SVG and PPTX output render a service legend derived from `services.csv`, with configurable SVG legend position. |
| XAL-3000160 | Icon license and attribution bundling | Implemented | Bundled license and attribution files for AWS, Tabler, and Yamaha icon sets are preserved alongside the generated catalogs. |

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

## Group 5 — Content Elements: Tables, Databases, UML (`XAL-5xxxxxx`)

| ID | Feature | Status | Summary |
|---|---|---|---|
| XAL-5000010 | `<table>` GFM-like pipe syntax | Implemented | Markdown-style pipe table with header/alignment row, normalized to the same typed rows as tag syntax. |
| XAL-5000020 | `<table>` explicit tag syntax | Implemented | `<header>`/`<row>`/`<cell>` children as an alternative to pipe syntax, normalized identically before layout. |
| XAL-5000030 | Table style inheritance | Implemented | Table-level `color`/`background-color`/`border-color`/`font-*` and `header-*` overrides, with `cell > header/row > table > built-in` precedence. |
| XAL-5000040 | Imported table schemas | Implemented | Table structure imported from an external schema source instead of authored inline in `.xal`. |
| XAL-5000050 | `<database-schema>` reusable definitions | Implemented | Document-level `<data>` schema definitions referenced by one or more frames via `<database data="schema-id">`. |
| XAL-5000060 | `<entity>`/`<column>` relational definitions | Implemented | Typed entity/column definitions supporting `name`, `type`, `primary-key`, `nullable`, `unique`, and `default`. |
| XAL-5000070 | `<foreign-key>` relation generation | Implemented | Single-column foreign key with a `references="entity.column"` target that generates an inter-entity relation. |
| XAL-5000080 | Relational key/nullability styling | Implemented | Primary/foreign key and nullability/data-type visual styling on entity column rows. |
| XAL-5000090 | Crow-foot relation notation | Implemented | Relation endpoints render with crow-foot cardinality notation on database diagrams. |
| XAL-5000100 | `<uml>` shared diagram component | Implemented | Common V1 component adapting UML families to xaligo's shared layout, shape, connector, and output pipeline. |
| XAL-5000110 | UML class diagrams | Implemented | `class`, `interface`, `enumeration` elements with `association`, `aggregation`, `composition`, `generalization`, `realization`, `dependency` relations. |
| XAL-5000120 | UML component diagrams | Implemented | `component`, `interface`, `port`, `artifact` elements with `dependency`, `realization`, `association`, `assembly`, `delegation` relations. |
| XAL-5000130 | UML activity diagrams | Implemented | `initial`/`final`/`activity`/`action`/`decision`/`merge`/`fork`/`join`/`object-node` elements with `control-flow`/`object-flow` relations. |
| XAL-5000140 | UML activity partitions/swimlanes | Implemented | `<partition>` swimlanes with `lanes="vertical|horizontal"` and the `theme="xaligo"` swimlane visual theme. |
| XAL-5000150 | UML state-machine diagrams | Implemented | `state`/`history`/`choice`/`fork`/`join`/`initial`/`final` elements with `transition` relations, events, guards, and effects. |
| XAL-5000160 | UML sequence diagrams | Implemented | `participant`/`lifeline` elements with `message`/`return-message`/`create-message`/`destroy-message` relations and explicit diagram-unique `order`. |
| XAL-5000170 | UML element compartments | Implemented | Typed ordered text compartments per element kind (`attribute`, `operation`, `constraint`, `note`, `entry`/`do`/`exit`, etc.). |
| XAL-5000180 | UML relation attributes | Implemented | Shared relation `label`/`guard`/`route` hint and `src-multiplicity`/`dst-multiplicity` on association-family relations. |
| XAL-5000190 | UML relation projection | Implemented | Fixed lowering of UML relation kinds to the shared orthogonal connector model (dashed/solid line, source diamond, no-arrowhead). |
| XAL-5000200 | Reusable `<uml-model>` definitions | Implemented | Document-level `<data>` UML models selected by one or more diagram-kind children via `data="model-id"`. |
| XAL-5000210 | Public UML connection endpoint references | Implemented | `uml-id/local-id` (same frame) and `frame-id.uml-id/local-id` (cross-frame) public endpoint forms for normal `<connection>` tags. |
| XAL-5000220 | Deliberately lossy V1 UML projection | Implemented | Documented, intentional projection limits (no dashed lifelines/activations/combined fragments, flattened compartments) applied consistently across every output format. |
| XAL-5000230 | UML timing diagrams | Planned | Accepted new UML family (Q05.17): `timing-diagram` selector with lifeline/state-timeline elements, state/value-change events, duration constraints, and time-axis layout. |

## Group 6 — Output Formats & Rendering (`XAL-6xxxxxx`)

| ID | Feature | Status | Summary |
|---|---|---|---|
| XAL-6000010 | Excalidraw output | Implemented | Editable Excalidraw JSON scene with groups, connectors, and metadata tags. |
| XAL-6000020 | SVG output | Implemented | Standalone SVG, one file per frame by default or one combined canvas with `--combine-frames`. |
| XAL-6000030 | PowerPoint (PPTX) output | Implemented | PPTX presentation built from a shared Go draw plan and executed by a configured WASM/PptxGenJS exporter, one slide per frame by default. |
| XAL-6000040 | PDF output | Implemented | PDF document, one page per frame by default, inheriting SVG's strict crop boundary. |
| XAL-6000050 | Excel output | Implemented | Excel workbook, one worksheet per frame by default, each embedding that frame's rendered SVG image. |
| XAL-6000060 | XYFlow (React Flow) output | Implemented | XYFlow-compatible JSON nodes and edges for React Flow-based viewers. |
| XAL-6000070 | Isoflow output | Implemented | Isoflow-compatible model JSON for isometric network diagrams. |
| XAL-6000080 | Frame-to-physical-page mapping contract | Implemented | Shared default mapping of one identified child frame to one SVG file/PPTX slide/PDF page/Excel worksheet. |
| XAL-6000090 | `--combine-frames` compatibility mode | Implemented | Restores the historical single-canvas/page form for SVG, PPTX, PDF, and Excel; Excalidraw/XYFlow/Isoflow are unaffected since they are already single documents. |
| XAL-6000100 | Safe frame-ID output naming | Implemented | Multi-frame SVG output derives `<stem>-<safe-frame-id>.svg` names, with deterministic collision detection. |
| XAL-6000110 | SVG legend placement | Implemented | `--svg-legend-position top|right|bottom|left` controls where the services.csv-derived legend renders. |
| XAL-6000120 | PPTX legend slides | Implemented | Dedicated legend slide(s) appended after diagram slides, using a 4-column icon/abbreviation/official-name layout. |
| XAL-6000130 | PDF page cropping | Implemented | PDF pages inherit the exact per-frame SVG canvas and clip boundary as their strict crop. |
| XAL-6000140 | Excel worksheet SVG embedding | Implemented | Each Excel worksheet embeds its frame's rendered SVG image rather than reconstructing native shapes. |
| XAL-6000150 | Single-logical-document output invariance | Implemented | Excalidraw, XYFlow, and Isoflow always remain one logical document regardless of `--combine-frames`. |
| XAL-6000160 | Isoflow generic-endpoint projection | Implemented | UML and other shapes without an Isoflow-native equivalent project to a labeled generic endpoint icon. |
| XAL-6000170 | Oversized-frame page tiling | Planned | Generic tiling of one oversized frame across several physical pages/slides, distinct from the existing one-frame-per-page mapping. |
| XAL-6000180 | Renderer capability/projection contract | Planned | A typed per-format capability declaration (e.g., which formats can carry arbitrary UML/connector metadata) so unsupported projections become explicit instead of implicit. |

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

## Group 8 — CLI Commands & Developer Tooling (`XAL-8xxxxxx`)

| ID | Feature | Status | Summary |
|---|---|---|---|
| XAL-8000010 | `xaligo render` command | Implemented | Renders a `.xal` source file into any supported output format via `--format`. |
| XAL-8000020 | `xaligo validate` command | Implemented | Validates `.xal` syntax, layout, and connection references without producing output. |
| XAL-8000030 | `xaligo serve` command | Implemented | Serves a live-reloading preview over HTTP, polling the source file and re-rendering on change; a `.xal` source previews one combined SVG canvas, a `.md`/`.markdown` source previews the full Markdown document with rendered `xal` code blocks embedded inline; `--paper`/`--orientation` fix the preview to a specific physical page size and orientation at server startup. |
| XAL-8000040 | `xaligo diff` command | Implemented | Renders paired removed/added structural-diff SVGs comparing two `.xal` documents. |
| XAL-8000050 | `xaligo generate xal` command | Implemented | Scaffolds a `.xal` file with a configurable AWS Cloud/Account/Region/VPC/AZ/Subnet hierarchy. |
| XAL-8000060 | `xaligo add service` (single mode) | Implemented | Appends one named AWS service icon and a legend entry into an existing `.excalidraw` file. |
| XAL-8000070 | `xaligo add service --list` (batch mode) | Implemented | Bulk-adds every service in a `services.csv` list into an existing `.excalidraw` file. |
| XAL-8000080 | `xaligo init` command | Implemented | Writes a minimal starter `sample.xal` file to help new users get started. |
| XAL-8000090 | `xaligo version` command | Implemented | Reports the resolved build/release version, with an embedded-value, env-var, and `VERSION`-file fallback chain. |
| XAL-8000100 | Man-page-style detailed CLI help | Implemented | Every command and subcommand exposes a detailed `Long` description and runnable `Examples` in its `--help` output, mirrored in the CLI reference documentation. |
| XAL-8000110 | Structured CLI logging | Implemented | Stable `share.NewMCode`-tagged debug/info/error log lines throughout controllers and use cases for traceable CLI diagnostics. |
| XAL-8000120 | VS Code extension integration | Planned | A separate-repository VS Code extension providing `.xal` syntax highlighting, a live Preview Panel, and source-positioned diagnostics, built on this repository's reusable render/validate APIs and HTTP/SSE preview protocol. |
| XAL-8000130 | Multi-view live preview | Planned | Extending `xaligo serve`/preview beyond SVG-first to Excalidraw, XYFlow, Isoflow, and 2.5D preview views. |
| XAL-8000140 | `xaligo render markdown` command | Implemented | Reads a Markdown file, renders every fenced ` ```xal ` code block to SVG through the shared render pipeline, and writes a new Markdown file with a `![](path.svg)` image reference per rendered frame in place of each code block; generated SVGs default beside the source Markdown file and `--svg-dir`/`--output` override the locations; `--paper`/`--orientation` fit each rendered diagram to a physical page size, matching `render --format svg`. |

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
| XAL-9000090 | `cmd/wasm` JavaScript/WASM adapter | Implemented | Exposes the native Go rendering engine to JavaScript/WASM hosts alongside the native CLI entry point. |
| XAL-9000100 | V1/V2 compatibility golden tests | Planned | Golden tests comparing V1-compatibility and native V2 engine output at the neutral-model and resolved-geometry boundaries once the V2 engine exists. |
| XAL-9000110 | Cross-renderer visual regression suite | Planned | Representative visual regression coverage across the full SVG/Excalidraw/PPTX/PDF/Excel/XYFlow/Isoflow format matrix; currently limited per the roadmap's documented gaps. |
| XAL-9000120 | Render determinism and concurrency-safety guarantees | Planned | Byte-stable output for identical source/options/assets/environment, and parallel-job safety without shared mutable render state (Q13.1/Q13.2 backlog). |

---

# 05 Issues and quality


Use this file as the authoritative list of open quality and feature-hardening
issues when planning, implementing, reviewing, or verifying feature quality
work. The goal is to raise each Phase 2 feature from "renders" to a
documented, tested, readable, and cross-format-stable user experience. Record
every current issue here; once an issue's work is verified complete, remove
its row instead of marking it done in place.

## Quality Rubric

Evaluate every feature slice against these gates before considering it done:

1. Specification correctness.
2. Validate/render agreement.
3. Layout correctness.
4. Design quality.
5. Cross-format consistency.
6. Regression tests.
7. Samples and documentation.
8. CI, security, and release readiness.
9. Determinism, performance, and resource safety.
10. Accessibility and theme readability.
11. API, CLI, and preview compatibility.

## Feature Quality Workflow

For each feature, proceed in this order:

1. Identify the authoritative contract in the DSL specification,
   architecture instructions, design documentation, examples, and tests.
2. Establish one observable invariant or failing behavior before editing.
3. Prefer a focused regression test that can fail for the diagnosed cause.
4. Fix the earliest shared layer that owns the behavior. Do not hide shared
   parser, layout, routing, scene, or plan defects inside one renderer.
5. Validate the touched slice with the narrowest useful command, then broaden
   only when the change affects shared contracts or multiple formats.
6. Update examples, documentation, and generated source-controlled assets when
   their visible behavior or contract changes.
7. Commit one coherent feature-quality slice after the required checks pass.

## Status Tracking

Maintain the feature hardening status in this file so future agents can resume
from the current quality pass without relying on conversation history.

Use these status values:

- `not-started`: the feature is in the Phase 2 scope but has not been audited.
- `auditing`: contracts, samples, tests, and visual output are being inspected.
- `needs-work`: a quality gap is known and has not been fixed.
- `in-progress`: a fix or quality improvement is actively being implemented.
- `blocked`: progress requires a decision, missing tool, or external input.

There is no `done` status value. Once a feature or task meets the Definition
of Done in this file, remove its row from the status board and backlog
instead of marking it done in place, so this file always reflects only open
issues.

Update the status board when a feature-quality slice starts, when a gap is
found, when validation changes the assessment, and when a slice is committed.
Keep each row concise but evidence-based. Prefer commit hashes, test names,
sample paths, and rendered artifact paths over prose-only status notes.

| Order | Feature | Status | Evidence | Next action |
|---:|---|---|---|---|
| 1 | Canonical V1 document envelope | in-progress | `TestV1ParseValidatesCanonicalEnvelopeHierarchy`; canonical sample validates and renders two SVG artifacts | Finish migration-warning/docs-image audit, then close Q01. |
| 2 | UML diagrams | in-progress | The supported sample matrix is frozen to class, component, activity, state-machine, and sequence; all retained samples validate under the strict profile. Timing-diagram support is accepted as an additional in-scope kind, tracked as Q05.17. | Finish per-diagram visual baselines and focused semantics, then verify cross-format parity. |
| 3 | Data registry and imports | not-started | table data, database schema, SQL import commits exist | Inventory import coverage and diagnostics after the UML precision/design pass starts. |
| 4 | Tables | not-started | `docs/src/examples/samples/tables.xal` | Audit table layout, pipe styling, text fitting, and docs image quality. |
| 5 | Relational databases | not-started | `docs/src/examples/samples/databases.xal` | Audit entity layout, key styling, relation routing, and SQL import behavior. |
| 6 | Frame-scoped references | not-started | frame-qualified endpoint support exists | Audit duplicate ID handling and user-correctable diagnostics. |
| 7 | Cross-frame page links | not-started | `docs/src/examples/samples/page-links.xal` | Audit terminal side selection, labels, insets, and invalid edge handling. |
| 8 | Frame metadata bands | not-started | `docs/src/examples/samples/frame-metadata.xal` | Audit metadata layout, reservation strips, wrapping, and collisions. |
| 9 | Page-oriented outputs | not-started | SVG/PPTX/PDF/Excel pagination commits exist | Audit default artifacts, `--combine-frames`, safe IDs, and crops. |
| 10 | Renderer matrix | not-started | SVG, Excalidraw, PPTX, PDF, Excel, XYFlow, Isoflow encoders | Define representative cross-format checks for changed shared contracts. |
| 11 | CI, release, and tooling | not-started | `VERSION`, workflows, RTK and security instructions | Audit CI gates, version policy, npm lockfile policy, and vendored dependency handling. |
| 12 | Diagnostics and error UX | not-started | `internal/usecase/diagnostics.go`, diagnostics tests | Audit severity, source positions, aggregation, cancellation, and actionable messages. |
| 13 | Determinism, concurrency, and performance | not-started | render determinism and concurrency tests | Audit repeatability, ordering, cancellation, limits, and representative performance. |
| 14 | CLI, API, and preview contracts | not-started | controllers, use-case API tests, preview tests | Audit option parity, exit behavior, live preview, and backward compatibility. |
| 15 | Themes and accessibility | not-started | theme application and renderer theme tests | Audit contrast, semantic distinction, text alternatives, and light/dark parity. |
| 16 | Structural diff | not-started | diff engine and highlight tests | Audit matching, moved/modified classification, highlights, and pagination interaction. |
| 17 | Reproducibility and artifact integrity | not-started | deterministic render and generated docs assets | Audit byte stability, clean regeneration, package contents, and stale artifacts. |
| 18 | Configuration, logging, and observability | not-started | config, options, and shared logging code | Audit defaults, precedence, redaction, error context, and noisy-output behavior. |

When a row's feature or task meets the Definition of Done in this file,
remove the row instead of marking it done, and record the closing commit in
the relevant commit history rather than in this file. If a row is `blocked`,
the next action must state the specific decision or tool needed.

## Detailed Task Backlog

Track detailed work with stable task IDs. Update task status in the nearest
feature row above when the detailed evidence changes; split or add IDs only
when the new task can be verified and committed independently.

### Q01 Canonical V1 Document Envelope

| Task | Status | Scope | Verification target |
|---|---|---|---|
| Q01.2 | not-started | Confirm legacy `<frame>` and `<frames>` inputs remain compatible and emit the intended migration warning. | diagnostics tests covering warning text and source position |
| Q01.4 | in-progress | Confirm canonical samples render as separate artifacts and combined compatibility output. | canonical sample renders `canonical-v1-envelope-overview.svg` and `canonical-v1-envelope-database-detail.svg`; combined compatibility still pending |
| Q01.5 | not-started | Review canonical-envelope docs for current behavior, command accuracy, and image freshness. | `mdbook build docs` and regenerated SVG comparison |

### Q02 Data Registry and Imports

| Task | Status | Scope | Verification target |
|---|---|---|---|
| Q02.1 | not-started | Inventory supported `<data>` children and their reference syntax. | spec/docs consistency check |
| Q02.2 | not-started | Verify relative import resolution, missing files, unsupported extensions, and provenance. | diagnostics tests with temp files |
| Q02.3 | not-started | Verify duplicate data IDs and unresolved data references across frames. | parser/validation unit tests |
| Q02.4 | not-started | Confirm imports are deterministic and do not execute commands or read outside the allowed context. | security-oriented unit tests |
| Q02.5 | not-started | Confirm imported data renders identically to equivalent inline data where applicable. | golden or structural scene comparison |

### Q03 Tables

| Task | Status | Scope | Verification target |
|---|---|---|---|
| Q03.1 | not-started | Audit standard table layout: headers, rows, empty cells, column widths, and borders. | scene/layout unit tests and SVG sample render |
| Q03.2 | not-started | Audit pipe-table parsing and styling parity with tag-based tables. | parser and render comparison tests |
| Q03.3 | not-started | Verify long English, Japanese, symbols, and narrow columns wrap or fit without overlap. | SVG text geometry assertions |
| Q03.4 | not-started | Review table visual hierarchy, spacing, typography, and docs-image polish. | design review notes plus regenerated sample SVG |
| Q03.5 | not-started | Confirm table behavior in Excalidraw and PPTX when shared table contracts change. | focused format checks |

### Q04 Relational Databases

| Task | Status | Scope | Verification target |
|---|---|---|---|
| Q04.1 | not-started | Audit entity and column layout, key styling, nullability, and data-type display. | scene/layout unit tests and database sample render |
| Q04.2 | not-started | Verify primary, foreign, and composite keys from inline and imported SQL schemas. | SQL import and schema unit tests |
| Q04.3 | not-started | Validate relation endpoints, labels, crow-foot notation, and self/multi relations. | routing/scene tests |
| Q04.4 | not-started | Check database diagrams for readable density, relation crossing, and text contrast. | SVG visual inspection and coordinate assertions where practical |
| Q04.5 | not-started | Confirm invalid schemas produce source-positioned, user-correctable diagnostics. | invalid SQL/schema tests |

### Q05 UML Diagrams

| Task | Status | Scope | Verification target |
|---|---|---|---|
| Q05.2 | in-progress | Establish a per-UML visual baseline before edits: validate and render each supported individual `uml-*.xal` sample to SVG, then identify overlap, spacing, typography, connector, and semantic-notation gaps. | all retained UML samples validate; tracked SVG baselines exist, with component and activity visual audits completed. |
| Q05.3 | in-progress | Improve activity-diagram semantic accuracy: initial/final nodes, actions, object nodes, decisions, forks, joins, merges, responsibilities, constraints, guards, `control-flow`, and `object-flow`. | vertical and horizontal partition samples validate; focused parser/scene tests cover partitions, loop routes, and activity metadata. |
| Q05.4 | in-progress | Improve activity-diagram design quality: left-to-right reading flow, diamond/bar/final-node proportions, lane spacing, label placement, and control-vs-object-flow distinction. | vertical and horizontal activity SVG baselines are regenerated; focused geometry tests cover lane placement and loop routing. |
| Q05.6 | in-progress | Improve class-diagram design quality: compartment rhythm, long member wrapping, stereotype readability, relation label spacing, crow-foot/diamond marker clarity, and dense-layout crossing reduction. | class SVG/PPTX review and text-fit tests; relation-label spacing fixed so labels avoid unrelated classifier boxes in the same frame, not only their own connection endpoints (`TestUMLRelationLabelsAvoidUnrelatedClassifierBoxes`, regenerated `docs/src/images/uml-class.svg`); compartment rhythm, long member wrapping, crow-foot/diamond marker clarity, and dense-layout crossing reduction remain to be reviewed |
| Q05.7 | not-started | Improve sequence-diagram semantic accuracy: lifelines, participants, activation bars, sync/async messages, returns, self messages, create/delete, ordering, and message labels. | `uml-sequence.xal` layout tests and route ordering assertions |
| Q05.8 | not-started | Improve sequence-diagram design quality: timeline spacing, activation contrast, message arrow style, return-line distinction, lifeline header readability, and vertical density. | sequence SVG/PPTX review and geometry assertions |
| Q05.9 | not-started | Improve state-machine semantic accuracy: initial/final states, composite states, transitions, events, guards, effects, entry/do/exit actions, and invalid transition diagnostics. | `uml-state-machine.xal` valid/invalid tests |
| Q05.10 | not-started | Improve state-machine design quality: state shape proportions, nested-state padding, transition bend points, guard/effect label placement, and final-node readability. | state-machine SVG review and collision assertions |
| Q05.11 | in-progress | Improve component-diagram semantic and design quality: boundary interfaces, shared interface-name widths, compact automatic height, explicit sizing, fan-out, and connector routing. | component parser/layout/scene/routing tests plus `uml-component.xal` and its regenerated SVG baseline |
| Q05.12 | not-started | Plan future communication-diagram support only when ordered topology semantics are explicitly required beyond sequence/state-machine diagrams. | plan entry only; no implementation, samples, or generated assets until restarted |
| Q05.13 | not-started | Keep object, use-case, profile, and interaction-overview out of the supported UML set unless a non-substitutable use case is identified. | unsupported parser diagnostics and documentation review |
| Q05.17 | not-started | Accepted addition: design and implement timing-diagram support (`timing-diagram` selector) — lifeline/state-timeline elements, state/value change events, duration constraints, and time-axis layout — as a new UML family alongside class/component/activity/state-machine/sequence. | new `uml-timing.xal` sample; parser/scene/layout/routing tests; diagram-kind vocabulary and relation-projection tables in `xal-spec.instructions.md` updated once implemented |
| Q05.14 | not-started | Normalize shared UML visual language across all diagram types: typography scale, stroke weights, marker sizes, semantic colors, label backgrounds, and light/dark contrast. | design review checklist plus theme render comparisons |
| Q05.15 | not-started | Verify UML cross-format parity for every shared scene/plan change: SVG baseline, Excalidraw editability, PPTX plan/export, PDF/Excel page projection, and XYFlow/Isoflow applicability. | focused renderer matrix for changed UML contracts |

### Q06 Frame-Scoped References

| Task | Status | Scope | Verification target |
|---|---|---|---|
| Q06.1 | not-started | Confirm local endpoint references resolve inside the owning frame first. | reference-resolution unit tests |
| Q06.2 | not-started | Confirm `frame-id.endpoint-id` resolves only to the intended frame endpoint. | cross-frame tests |
| Q06.3 | not-started | Validate duplicate IDs across frames, ambiguous endpoints, and missing frames. | diagnostics tests |
| Q06.4 | not-started | Confirm reference diagnostics include actionable source positions. | error-message assertions |
| Q06.5 | not-started | Check docs and examples explain local vs qualified reference behavior. | docs review and `mdbook build docs` |

### Q07 Cross-Frame Page Links

| Task | Status | Scope | Verification target |
|---|---|---|---|
| Q07.1 | not-started | Audit automatic page terminal side selection and safe side fallback. | routing tests and page-link sample render |
| Q07.2 | not-started | Verify explicit `src-frame-side`, `dst-frame-side`, anchors, and invalid same-frame usage. | validation tests |
| Q07.3 | not-started | Confirm terminal insets, labels, stubs, and arrowheads avoid endpoints and metadata strips. | SVG coordinate assertions |
| Q07.4 | not-started | Check page-link labels remain readable in multi-page SVG, PPTX, PDF, and Excel projections. | representative format checks |
| Q07.5 | not-started | Review page-link docs and samples for command accuracy and image freshness. | docs build and regenerated images |

### Q08 Frame Metadata Bands

| Task | Status | Scope | Verification target |
|---|---|---|---|
| Q08.1 | not-started | Audit visible `id`, `title`, `version`, and custom metadata rows. | metadata unit tests and sample render |
| Q08.2 | not-started | Verify `row-gap`, `width`, `key-width`, wrapping, row breaks, and top/bottom placement. | layout and SVG text tests |
| Q08.3 | not-started | Confirm reservation strips exclude normal content, labels, connectors, and page terminals. | validation/render agreement tests |
| Q08.4 | not-started | Audit collisions with AWS/general group headers and nested tags. | regression tests for nested headers |
| Q08.5 | not-started | Review metadata visual hierarchy, contrast, and docs polish. | design review and regenerated SVGs |

### Q09 Page-Oriented Outputs

| Task | Status | Scope | Verification target |
|---|---|---|---|
| Q09.1 | not-started | Verify default SVG emits one artifact per identified frame with stable safe IDs. | integration tests and output filename assertions |
| Q09.2 | not-started | Verify `--combine-frames` compatibility behavior for SVG, PPTX, PDF, and Excel. | CLI integration tests |
| Q09.3 | not-started | Confirm frame crop boundaries, hidden page-frame outlines, and marker safety behavior. | SVG/PDF geometry checks |
| Q09.4 | not-started | Verify filename collision handling and one-frame exact output path behavior. | artifact tests |
| Q09.5 | not-started | Confirm API documentation matches CLI and use-case behavior. | docs/API review |

### Q10 Renderer Matrix

| Task | Status | Scope | Verification target |
|---|---|---|---|
| Q10.1 | not-started | Define the minimal representative sample set for SVG, Excalidraw, PPTX, PDF, Excel, XYFlow, and Isoflow. | matrix documented in this file or docs |
| Q10.2 | not-started | Verify shared scene/plan changes are visible consistently across applicable encoders. | focused format matrix command set |
| Q10.3 | not-started | Confirm Excalidraw editability and metadata survive renderer changes. | JSON structural assertions |
| Q10.4 | not-started | Confirm PPTX plan parity and external TypeScript exporter behavior. | Go/WASM plan tests and `npm --prefix external test` |
| Q10.5 | not-started | Confirm native-only PDF/Excel dependencies stay out of browser-WASM builds. | `GOOS=js GOARCH=wasm go list -deps ./cmd/wasm` guard |

### Q11 CI, Release, and Tooling

| Task | Status | Scope | Verification target |
|---|---|---|---|
| Q11.1 | not-started | Audit GitHub Actions for Go, TypeScript, docs, version, release, and pages jobs. | workflow review and CI run results |
| Q11.2 | not-started | Confirm `VERSION`, npm package metadata, and release tags stay coherent. | version-gate reproduction |
| Q11.3 | not-started | Confirm npm lockfile policy is reflected in workflows, Makefile, and docs. | TypeScript install/test command |
| Q11.4 | not-started | Verify RTK and security-check preconditions are current and actionable. | `make security-check` and instruction review |
| Q11.5 | not-started | Confirm generated artifacts, vendored dependencies, binaries, caches, and docs images follow repository policy. | `git status`, ignore rules, and release/package review |

### Q12 Diagnostics and Error UX

| Task | Status | Scope | Verification target |
|---|---|---|---|
| Q12.1 | not-started | Audit diagnostic severity, line/column positions, element context, and stable wording. | diagnostics unit tests |
| Q12.2 | not-started | Confirm multiple independent input errors are reported without hiding the first actionable cause. | aggregate-diagnostics tests |
| Q12.3 | not-started | Confirm warnings are non-blocking while errors consistently fail validate and render. | validate/diagnose agreement tests |
| Q12.4 | not-started | Verify canceled contexts stop diagnostics and rendering with wrapped context errors. | cancellation tests |
| Q12.5 | not-started | Review CLI and preview presentation of diagnostics for concise, user-correctable output. | controller/preview tests and manual command check |

### Q13 Determinism, Concurrency, and Performance

| Task | Status | Scope | Verification target |
|---|---|---|---|
| Q13.1 | not-started | Confirm identical source, options, assets, and environment produce byte-stable output where the format allows it. | render determinism tests |
| Q13.2 | not-started | Confirm parallel jobs preserve document/page/artifact order and do not share mutable render state. | concurrency and race-enabled tests |
| Q13.3 | not-started | Verify cancellation propagates through I/O and orchestration without goroutine or temporary-file leaks. | cancellation/leak tests |
| Q13.4 | not-started | Establish representative render benchmarks for complex architecture, tables, database, UML, and multi-frame documents. | Go benchmarks with recorded baseline |
| Q13.5 | not-started | Define and test safe behavior for extreme node counts, text lengths, dimensions, ratios, and import sizes. | bounded stress tests and finite-geometry assertions |

### Q14 CLI, API, and Preview Contracts

| Task | Status | Scope | Verification target |
|---|---|---|---|
| Q14.1 | not-started | Audit CLI flags, defaults, aliases, validation, exit codes, stdout/stderr, and output path behavior. | controller/command integration tests |
| Q14.2 | not-started | Confirm constructor-injected use-case APIs and convenience methods return equivalent contracts. | use-case API tests |
| Q14.3 | not-started | Verify native and embedded asset sources produce equivalent output at matching settings. | native/embedded parity tests |
| Q14.4 | not-started | Audit live preview initial render, reload events, diagnostics, browser refresh, and file-change recovery. | preview tests and local serve check |
| Q14.5 | not-started | Confirm legacy mode/format aliases and public option defaults remain backward compatible. | compatibility tests and docs review |

### Q15 Themes and Accessibility

| Task | Status | Scope | Verification target |
|---|---|---|---|
| Q15.1 | not-started | Audit light/dark theme contrast for text, borders, fills, connectors, metadata, and diff highlights. | contrast calculation plus theme renders |
| Q15.2 | not-started | Confirm semantic distinctions do not depend on color alone when line style, marker, label, or shape can carry meaning. | renderer structural assertions and design review |
| Q15.3 | not-started | Verify text remains readable at documented output sizes and under zoom/scaling across formats. | SVG/PPTX/PDF visual checks |
| Q15.4 | not-started | Audit SVG titles/descriptions, meaningful identifiers, and document structure available to assistive tooling where supported. | SVG structure tests |
| Q15.5 | not-started | Verify full-width, combining, RTL, and uncommon glyph fallback behavior is documented or safely handled. | typography tests with representative strings |

### Q16 Structural Diff

| Task | Status | Scope | Verification target |
|---|---|---|---|
| Q16.1 | not-started | Audit added, removed, modified, moved, and unchanged classification independent of source ordering. | diff classification tests |
| Q16.2 | not-started | Confirm matching uses semantic identity rather than generated scene IDs or source lines. | fingerprint/matching tests |
| Q16.3 | not-started | Verify highlights remain behind labels, outside routing obstacles, and readable in light/dark themes. | diff scene/SVG tests |
| Q16.4 | not-started | Confirm multi-frame diff pagination and output pairing remain stable. | document-plan integration tests |
| Q16.5 | not-started | Review diff samples and docs for realistic change cases and visual clarity. | sample render and docs build |

### Q17 Reproducibility and Artifact Integrity

| Task | Status | Scope | Verification target |
|---|---|---|---|
| Q17.1 | not-started | Confirm checked-in documentation SVGs reproduce byte-for-byte or through a documented normalization step. | regenerate-and-compare command |
| Q17.2 | not-started | Detect stale generated images, missing source `.xal` files, and orphaned documentation assets. | source/asset inventory script or test |
| Q17.3 | not-started | Audit npm, release, and documentation package contents for required and forbidden files. | package dry run and archive listing |
| Q17.4 | not-started | Confirm generators are idempotent and do not reorder or rewrite unrelated files. | double-run clean-worktree check |
| Q17.5 | not-started | Verify timestamps, random IDs, map iteration, and platform paths do not make supported outputs unstable. | cross-run and cross-platform determinism checks |

### Q18 Configuration, Logging, and Observability

| Task | Status | Scope | Verification target |
|---|---|---|---|
| Q18.1 | not-started | Audit configuration defaults, file/environment/flag precedence, and invalid-value diagnostics. | config/controller tests |
| Q18.2 | not-started | Confirm logs and wrapped errors identify the failed stage without duplicating or obscuring the root cause. | error-chain assertions and command checks |
| Q18.3 | not-started | Verify normal commands keep stdout machine-usable and send diagnostics/progress to the intended stream. | CLI output tests |
| Q18.4 | not-started | Confirm logs and diagnostics do not expose imported source contents, credentials, tokens, or sensitive paths unnecessarily. | redaction/security tests |
| Q18.5 | not-started | Define optional timing/count observability for parser, layout, routing, scene, and encoder stages without changing default output. | design decision and focused tests if implemented |

## Design Quality Gate

Treat design quality as a first-class feature requirement, not a cosmetic
follow-up. A feature is not complete only because it validates and renders.

Check these visual qualities for every rendered feature:

- Visual hierarchy: titles, group labels, node labels, metadata tags, and
  connector labels should communicate reading order and importance clearly.
- Spacing and rhythm: node gaps, group padding, table cells, metadata rows,
  connector-label offsets, and page-link terminals should feel intentional and
  consistent with the diagram type.
- Typography: text size, wrapping, fitting, clipping, and full-width character
  handling should remain readable across English, Japanese, symbols, and long
  identifiers.
- Shape language: UML, table, database, AWS, metadata, and page-link elements
  should look meaningfully distinct while still belonging to one visual system.
- Connector readability: route, traffic, UML control flow, object flow,
  relation, and page-link lines should remain distinguishable and avoid
  labels, ports, metadata bands, and important shapes.
- Color and contrast: themes, diff highlights, semantic markers, and metadata
  tags should preserve readable contrast in SVG, Excalidraw, PPTX, PDF, Excel,
  XYFlow, and Isoflow projections where applicable.
- Documentation polish: examples should be good enough for users to copy and
  for documentation to present without extra manual cleanup.

## Phase 2 Quality Order

Harden Phase 2 features in this order unless a blocking defect requires a
smaller local detour:

1. Canonical V1 document envelope: `<xaligo version="1">`, `<data>`,
   `<frames>`, identified `<frame>`, legacy compatibility, and migration
   diagnostics.
2. UML diagrams: first semantic accuracy for each UML kind, then design
   polish for that kind, then shared UML visual-language and cross-format
   parity. Prioritize activity, class, sequence, state machine, timing,
   component/deployment/package/composite-structure, then the remaining UML
   samples and strict profile diagnostics.
3. Data registry and imports: reusable table data, database schemas, SQL
   imports, relative paths, provenance, duplicate IDs, and missing resources.
4. Tables: standard tables, pipe tables, styling, cell text, headers, column
   sizing, empty cells, and imported data.
5. Relational databases: entities, columns, primary keys, foreign keys,
   composite keys, SQL imports, relation routing, and notation-specific labels.
6. Frame-scoped references: local endpoints, `frame-id.endpoint-id`, duplicate
   IDs across frames, and user-correctable diagnostics.
7. Cross-frame page links: source/destination page terminals, automatic and
   explicit side selection, anchors, labels, inset behavior, and invalid
   metadata-reserved edges.
8. Frame metadata bands: visible `id`, `title`, `version`, custom metadata,
   row gaps, widths, wrapping, top/bottom placement, and collision avoidance.
9. Page-oriented outputs: one frame per SVG artifact, PPTX slide, PDF page, and
   Excel worksheet; `--combine-frames`; safe artifact IDs; crop boundaries; and
   filename collision handling.
10. Renderer matrix: SVG baseline behavior, Excalidraw editability, PPTX plan
    parity, PDF/Excel native constraints, XYFlow/Isoflow projections, and
    browser-WASM dependency boundaries.
11. CI, release, and tooling: Go, TypeScript, docs, security, version gates,
    RTK operation, npm lockfile policy, vendored dependencies, and generated
    artifact policy.
12. Diagnostics and error UX: severity, aggregation, source positions,
    cancellation, and presentation in CLI/preview adapters.
13. Determinism, concurrency, and performance: stable ordering and bytes,
    race safety, cancellation, benchmarks, and bounded extreme inputs.
14. CLI, API, and preview contracts: options, defaults, output behavior,
    native/embedded parity, reload behavior, and backward compatibility.
15. Themes and accessibility: contrast, non-color semantics, typography,
    assistive SVG structure, and international text behavior.
16. Structural diff: semantic matching, classification, highlights,
    multi-frame pairing, and sample quality.
17. Reproducibility and artifact integrity: generated asset freshness,
    package contents, generator idempotence, and platform stability.
18. Configuration, logging, and observability: precedence, wrapped errors,
    output streams, redaction, and optional stage telemetry.

## First UML Quality Slice

After the current Q01 canonical-envelope follow-ups are closed or explicitly
paused, start the UML quality pass with
`docs/src/examples/samples/uml-activity.xal`. It is the smallest UML slice that
exercises both semantic precision and visible design quality, and should
verify:

- `initial`, `action`, `object-node`, `decision`, `fork`, `join`, `merge`, and
  `final` nodes.
- `control-flow` and `object-flow` rendering and visual distinction.
- `guard`, `title`, nested `responsibility`, and nested `constraint` behavior.
- Left-to-right activity readability.
- Decision diamonds, fork/join bars, merge nodes, final nodes, and connector
  labels avoiding collisions.
- SVG output quality sufficient for documentation, with additional format
  checks added when shared contracts or encoder behavior changes.

## Definition of Done

A feature-quality slice is done only when:

- The specification and documentation match the implementation.
- Representative valid samples pass validation and render successfully.
- Representative invalid inputs fail with source-positioned, user-correctable
  diagnostics.
- Regression tests cover the observable behavior being protected.
- Design quality has been reviewed for hierarchy, spacing, typography, shape
  language, connector readability, and contrast.
- Cross-format checks have been run when the feature affects shared scene,
  routing, plan, pagination, or renderer contracts.
- Determinism, cancellation, resource limits, compatibility, accessibility,
  and artifact reproducibility have been considered and tested where relevant.
- Required security and repository checks have passed.
- The slice is committed without unrelated working-tree changes.
---

# 06 Roadmap


This roadmap is a planning precondition for future implementation work. Prefer
changes that move xaligo toward a diagram-as-code platform with a clean render
pipeline, SVG-first preview capability, network-diagram primitives, and
eventual VS Code / PPTX integration.

Implementation guidance:

- Keep the core pipeline separable as `.xal -> parser -> layout -> renderer`.
- Treat Excalidraw, SVG, PPTX, PDF, Excel, XYFlow, and Isoflow as output
  renderers/projections over a shared model where possible.
- Prioritize SVG renderer and network diagram primitives before advanced PPTX
  feature polish when the choice is otherwise ambiguous.
- Route/traffic separation, route connectors, orthogonal routing, edge offsets,
  layer routing, junctions, and line jumps are roadmap features, not one-off
  export hacks.
- Live preview and VS Code integration should build on `xaligo render` /
  `xaligo validate`, not separate hidden pipelines.

## Product Architecture Preconditions

### V1 Structured-Diagram Profile

The table, relational-database, and UML design extends the V1 engine. Canonical
V1 uses `<xaligo version="1">`; historical root `<frame>` and `<frames>`
documents remain compatible but emit a migration warning. This does not replace
the reserved V2 `<scene version="2">` contract.

The target document shape uses `<xaligo>` as a document envelope, a document-
wide `<data>` registry, and `<frames>` containing identified `<frame>`
components. Data definitions are reusable across frames. General tables, RDB
schemas, and UML diagrams have separate semantic frontends and layout engines;
they may share renderer-neutral primitives and output encoders but must not be
forced through one diagram-specific processor.

Keep these semantic distinctions:

- `<table>` is general tabular data, `<database>`/`<entity>` is relational
  schema meaning, and `<grid>` is visual layout.
- Pipe and explicit tag syntax lower to the same typed model for a given
  component; imported files enter that model through an import adapter.
- `<uml>` is the common UML component. Exactly one diagram-kind child such as
  `<class-diagram>` or `<sequence-diagram>` selects its processor; the frame
  does not carry a UML kind.
- Imports are resolved before semantic validation, retain provenance, and do
  not execute arbitrary commands. Inline/tag overrides are explicit and
  deterministic.

The user-facing design is documented in
`docs/src/design/structured-diagrams.md`.

### Common DSL and Go Core

- Keep `.xal` as the single source DSL for every visual mode and export format.
- Keep Go as the core parser, validation, layout, routing, and rendering engine.
- VS Code, browser preview, and exporters must consume public core APIs instead
  of reimplementing parsing or layout.
- Preserve the pipeline boundary:

```text
.xal -> parser -> layout/shared model -> mode renderer -> format encoder
```

### V1 Compatibility and V2 Input

Keep root `<frame>` and `<frames>` as legacy V1 compatibility inputs. Canonical
V1 uses `<xaligo version="1">` with a document-wide `<data>` registry and
identified frames. Legacy roots emit a migration warning. V2 uses a distinct
`<scene version="2">` root; do not place `version="2"` on a V1 root. This is a
reject-safe boundary: existing V1 readers reject V2 without having to know any
V2 syntax.

V2 must render both native V2 documents and the frozen V1 profile. Implement
that compatibility in the V2 side only: a V1 compatibility frontend and the
native V2 frontend each lower directly to the same typed, version-neutral
model. Keep the existing V1 engine independent of V2.

The compatibility path is complete only when it preserves V1 defaults,
fallback/error behavior, unknown nested-tag handling, connection-group
inheritance, anchor aliases, numeric catalog-ID range, and render-context item
size. Golden tests must compare V1 and V2-engine output at the neutral-model and
resolved-geometry boundaries across native and embedded targets.

Do not implement compatibility by changing root tags as strings, reparsing,
retrying parsers after syntax errors, serializing through the V1 scene, or
calling the full V1 renderer before V2. Root dispatch reads the first start
element once and selects exactly one frontend; renderers and encoders remain
shared downstream.

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
| `excalidraw` | Editable Excalidraw scene |
| `pptx` | Editable presentation export; one frame per slide by default |
| `pdf` | Paginated document export; one frame per page by default |
| `excel` (`xlsx` alias) | Workbook export; one frame SVG per worksheet by default |
| `xyflow` | React Flow/GUI editor integration |
| `isoflow` | Isometric/2.5D integration |

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

### Shared Rendering APIs

The shared in-repository use-case boundary should support at least:

```go
RenderSVG()
RenderArtifacts()
RenderExcalidraw()
RenderPPTX()
RenderPDF()
RenderExcel()
RenderXYFlow()
RenderIsoflow()
```

Prefer a shared extensible API underneath the convenience functions:

```go
Render(ctx, input, RenderOptions{Mode: mode, Format: format})
Validate(ctx, input)
```

### Rendering Correctness Gate

New renderer features are gated by a shared geometry and text contract. Fixes
must be made at the earliest shared stage that owns the information, not as
format-specific clipping or coordinate adjustments.

The required order is:

1. Parse numeric layout attributes into finite, typed values and validate their
   domains with source positions.
2. Make validation and rendering execute the same geometry invariants.
3. Resolve fixed-size children before flexible weights, then record content
   boxes and explicit overflow state in the resolved layout.
4. Move item-grid selection and occupancy into resolved layout so items and
   other children cannot unknowingly occupy the same region; scene construction
   only emits the already resolved cells.
5. Carry renderer-neutral text layout, semantic role, and glyph-overflow policy
   through the draw plan.
6. Apply the same output transform to geometry and typography at every PPI and
   paper-fit setting.
7. Consolidate format dispatch in one use case and migrate the shared scene and
   plan to format-neutral names and schemas. Compatibility aliases may preserve
   public APIs, but the canonical schema must not remain Excalidraw- or
   PPTX-shaped.

Completion requires regression coverage for validation/render agreement,
finite resolved coordinates, parent/content containment, fixed-plus-flex
siblings, mixed item/rectangle groups, item offsets, connector numeric values,
empty numeric attributes, long labels across output formats including editable
Excalidraw metadata, overlapping ports, and non-96 PPI.

## Delivery Phases

### Phase 1: Basic Output

Status: complete.

- Stabilize `xaligo render` and `xaligo validate`.
- Complete the SVG renderer as the primary preview surface.
- Add shared Light and Dark themes.
- Extract stable renderer-facing shared use cases.

### Phase 2: Network Diagram Features

Status: headless V1 routes, the remaining routing steps, and textual connection
shorthands have initial shared implementations. Explicit circular connector
nodes remain future versioned work. Continue with hardening and cross-renderer
visual regression coverage.

Implement shared model/routing concepts in this order where dependencies allow:

1. Headless V1 route connectors; add explicit circular connector nodes only in
   a future versioned model.
2. Orthogonal Routing.
3. Route/Traffic separation.
4. Edge Offset.
5. Line Jump.
6. Layer Routing.
7. Junction generation.

These features must be shared across renderers where possible, rather than
implemented as PPTX-only corrections.

### Phase 3: Live Preview

Status: initial implementation complete. `xaligo serve` polls `.xal` sources,
renders through the public SVG API, reports source-positioned diagnostics, and
publishes SSE reload events. Browser polish remains.

- Add `xaligo serve` on top of public render/validate APIs. (implemented)
- Watch `.xal` files and automatically re-render. (implemented)
- Serve an SVG-first browser preview with incremental refresh. (implemented)
- Keep the protocol reusable by the VS Code extension.

## VS Code Extension Preconditions

The VS Code extension is developed in a separate repository. This repository
owns the reusable Go/WASM APIs and HTTP/SSE preview protocol only; do not add
extension packaging or VS Code-specific parser/rendering forks here.

The extension target includes:

- `.xal` syntax highlighting.
- Validation and source-positioned diagnostics. (Go and TypeScript/WASM APIs implemented)
- Live Preview and a Preview Panel.
- SVG preview first; Excalidraw, XYFlow, Isoflow, and 2.5D views later.

The extension must call the same validation/render pipeline as the CLI. Do not
create an extension-only parser, layout engine, or hidden preview format.

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

## Export Roadmap

Primary formats include SVG, Excalidraw, PPTX, PDF, and Excel. The page-oriented
formats use identified child frames as their default physical page boundary:
SVG files, PPTX slides, PDF pages, and Excel worksheets respectively.
`--combine-frames` retains the previous single-canvas behavior. Add or continue:

- XYFlow export for React Flow-style GUI editors. (initial implementation complete)
- Isoflow export for isometric and 2.5D integrations. (initial upstream model export complete)
- Generic tiling of one oversized frame across several pages. Frame pagination
  is implemented and is distinct from this remaining tiling work.

Both exports should consume the shared resolved model; they must not become
alternative parsers for `.xal`.

## Long-Term Product Position

Position xaligo between PlantUML, Excalidraw, draw.io, Cloudcraft, and
Instana-style topology tools:

- Diagram as Code.
- Strong AWS and network diagram support.
- 2D, 2.5D, and topology views from one DSL.
- Comfortable VS Code authoring.
- SVG, PPTX, PDF, Excel, Excalidraw, XYFlow, and Isoflow output.

## Current State

The repository is already beyond a blank v0.1 baseline in several areas.

Implemented or partially implemented:

- `.xal` XML-style parser exists in `internal/usecase/v1/engine/parse_*`.
- Vuetify-like layout calculations exist in `internal/usecase/v1/engine/layout_*`.
- Canonical scene construction exists in `internal/usecase/v1/engine/scene_*`.
  Rendered graph nodes carry Box-tree-derived semantic kind and parent IDs;
  XYFlow uses geometric containment only for legacy scenes without that data.
- Native CLI exists with `render`, `generate`, `add`, `init`, and `version`.
- `render --format excalidraw` supports `services.csv` abbreviation/legend
  workflows.
- Draw-plan geometry and routing exist in `internal/usecase/v1/engine/plan_*`
  and `internal/usecase/v1/engine/route_*`.
- PPTX routing already includes obstacle avoidance, binding gap handling,
  arrow margin/lane avoidance, A3 paper options, item label sizing, and legend
  slide data.
- Repository-layer PPTX export has been redirected toward a WASM exporter
  adapter in `internal/repository/powerpoint.go`.
- `xaligo render --format excalidraw|svg|pptx|pdf|excel` is implemented;
  `xlsx` is accepted as an alias for `excel`.
- Identified child frames map to separate SVG artifacts, PPTX slides, PDF
  pages, and Excel worksheets in source order. `--combine-frames` retains the
  compatibility single-canvas form. Excalidraw, XYFlow, and Isoflow remain one
  logical document.
- Default page-local SVG uses the exact frame rectangle as its canvas and clip
  boundary, inherited by PDF and Excel page images. The combined compatibility
  canvas retains marker-safe bounds expansion.
- Page frames support a shared top/bottom metadata tag band for built-in
  `id`, `title`, content `version`, and arbitrary key/value entries. The band
  uses the resolved `row-gap` (4 pixels by default) as both its inter-row
  spacing and its metadata page-edge inset at the selected vertical edge and
  both horizontal edges. Wrapping and per-row left/center/right alignment use
  `frame width - 2 * row-gap`. Its full-width reservation strip still starts at
  the outer logical frame edge, reaches the final content-box boundary, and is
  at least
  `row-gap + complete band height + 8` pixels deep; normal items, text,
  local/UML lines and labels, and page links remain outside it. Explicit page
  sides reject normal-dimension or reservation conflicts. Automatic page links
  filter unsafe candidates, remap a preferred side with rendered visual
  geometry, and fail only when no safe side exists; side terminals are clamped
  beyond the strip. The band also supports auto/fixed widths, explicit row
  breaks, and font/color styling. It projects with its owning physical page;
  graph adapters omit it as page decoration.
- Cross-frame connections independently support item endpoint
  `src/dst-side|anchor` and logical page-terminal
  `src/dst-frame-side|anchor` geometry. Explicit frame anchors use five fixed
  tangent slots per edge, then place the drawable terminal on a parallel inward
  inset line. The inset is the resolved metadata `row-gap`, or 4 pixels when
  metadata is absent; zero `row-gap` retains the outer edge. Terminals remain
  perpendicular to the selected side, do not clamp the inset, reject unsafe
  explicit geometry, remap unsafe automatic preferences to the nearest safe
  visual side, and keep `to <...>` / `from <...>` labels 4 layout pixels from
  the final inset terminal.
- `xaligo render --format xyflow` and TypeScript/WASM `renderXYFlow()` export
  nested React Flow-compatible nodes and edges. V1 item, AWS group, rectangle,
  port, and identified child-frame endpoints are retained; cross-frame stubs
  are combined into one logical edge with routing metadata.
- `xaligo render --format isoflow`, Go `RenderIsoflow`, and TypeScript/WASM
  `renderIsoflow()` export an upstream Isoflow-compatible model from the shared
  scene. V1 non-item endpoints and logical cross-frame connectors are retained,
  and same-frame explicit bends use native tile anchors.
- `xaligo validate` reuses parser and layout validation.
- `xaligo diff` compares parsed V1 structures and emits paired SVG views: the
  old document with removed/previous values highlighted pale red and the new
  document with added/current values highlighted pale green.
- The SVG encoder is implemented in `internal/repository/svg.go` over the
  shared draw plan, including distinct V1 arrow, triangle, stealth, diamond,
  and oval marker geometry.
- Shared `light` and `dark` themes are implemented for Excalidraw, SVG, and
  PPTX via `xaligo render --theme`.
- Stable Go use cases in `internal/usecase` expose `Render`, `RenderExcalidraw`,
  `RenderSVG`, `RenderArtifacts`, `RenderPPTX`, `RenderPDF`, `RenderExcel`,
  `RenderXYFlow`, `RenderIsoflow`, and `Validate`; CLI SVG/Excalidraw/validation
  use the same pipeline.
- CLI, preview, and WASM adapters now use the same render use case. Embedded
  environments inject an `AssetSource` instead of reimplementing parser,
  layout, or scene construction.
- Isoflow exports shared group borders as view rectangles and produces stable
  icon ordering.
- Frozen V1 routes are headless across Excalidraw, SVG, PPTX, PDF, Excel,
  XYFlow, and Isoflow. Circular route connector nodes remain a future versioned
  feature.
- Node/PptxGenJS can still generate `out.pptx` as a temporary development path,
  but it is not the long-term repository-layer architecture.

Important gaps:

- `external/wasm/xaligo.wasm` is the PPTX exporter WASM artifact.
- Cross-renderer visual regression coverage is still limited.
- Numeric domains are checked before layout, but a typed normalized layout
  structure has not yet replaced repeated reads from `Node.Attrs`.
- Item-grid minimum-cell and item-offset checks now run from `Build` through the
  same solver used by scene construction. Selected cells and catalog-derived
  label measurements still need to become first-class resolved-layout data
  before mixed item/rectangle groups are fully supported.
- `Diagnose` proves parser, resolved-box, minimum item-grid, and item-offset
  invariants. Catalog-derived label measurement and final connector geometry
  must join the same geometry stage for complete validate/render agreement.
- Compatibility names that expose Excalidraw or PPTX in otherwise shared scene
  and plan APIs have aliases, but the underlying schemas must still migrate to
  format-neutral data without breaking public callers.
- Renderer capabilities are still implicit. In particular, the compatible
  Isoflow connector schema cannot carry arbitrary V1 kind, arrowhead,
  fixed-point, or original scale/grid metadata. A typed capability/projection
  contract remains necessary before adding more output formats.

## Rebaselined Implementation Order

Use this order when starting new roadmap work from the current repository state:

1. Complete the shared geometry/text correctness gate and its cross-renderer
   regression tests.
2. Move mixed item-grid occupancy into resolved layout and finish neutral
   scene/plan naming.
3. Complete the repository-layer WASM PPTX exporter contract by providing
   `xaligo.wasm`; keep Go free of PPTX/OOXML writer code.
4. Harden shared network routing with cross-renderer visual regression tests.
5. Build the VS Code preview on the reusable HTTP/SSE protocol exposed by
   `xaligo serve`.

## v0.1 Foundation

Status: complete. CLI and shared use cases share parser/layout/render paths for
validation, Excalidraw, and SVG.

### Rendering Engine Refactoring

Separate the core pipeline into:

```text
.xal
 ↓
parser
 ↓
layout
 ↓
renderer
```

### Public API

```go
RenderExcalidraw()
RenderSVG()
RenderArtifacts()
RenderPPTX()
RenderPDF()
RenderExcel()
RenderXYFlow()
RenderIsoflow()
```

Current target API shape:

```go
Render(ctx, input, RenderOptions{Mode: mode, Format: format}) ([]byte, error)
Validate(ctx, input) error
```

### CLI

```bash
xaligo render
xaligo validate
```

Required compatibility:

- Keep existing `xaligo render <input.xal> -o <out.excalidraw>` working.
- Add `xaligo render <input.xal> --format excalidraw|svg|pptx|pdf|excel`;
  accept `xlsx` as the Excel alias.
- Keep format conversion under `xaligo render --format ...`; `generate` should
  remain focused on source `.xal` generation.
- `validate` must reuse parser/layout validation rather than duplicate parsing.

---

## v0.2 SVG Renderer

Status: initial renderer, route/traffic primitives, and shared Light/Dark themes
are implemented.

### SVG Export

```bash
xaligo render input.xal --format svg
```

### Supported Elements

- Node
- Group
- Label
- Route
- Traffic

### Themes

- Light
- Dark

---

## v0.3 Network Diagram Features

Status: route/traffic kinds, headless V1 routes, styling, layer order,
basic lane separation, automatic route junctions, and textual connection
shorthands are implemented across Excalidraw, SVG, and PPTX and inherited by
the SVG-based PDF and Excel projections.

### Route Connector

Frozen V1 route lines have no arrowheads. A future version may add explicit
renderer-neutral circular connector nodes without changing V1 arrowhead
semantics.

```text
o------o
```

### Connector Model

```go
type Connector struct {}
```

### Orthogonal Routing

Support right-angle routing.

```text
+----+
|    |
+----+
```

### Route / Traffic Separation

#### Route

```text
o------o
```

#### Traffic

```text
======>
```

### DSL

```text
web --- db
web ==> db
```

Status: implemented using `<item name="...">`, `<item ref="...">`, or numeric
item IDs. Shorthands expand into the shared connection model during parsing.

---

## v0.4 Advanced Routing

Status: initial shared implementations are complete for edge offsets, routing
layers, frame-border clearance, and automatic fan-out/fan-in junctions.

### Edge Offset

Automatically separate overlapping routes.

```text
------
======
```

### Layer Routing

Separate routing layers.

```text
Route Layer
Traffic Layer
```

### Junction Generation

```text
      +-- DB
o-----+
      +-- Cache
```

---

## v0.5 Line Jumps

Status: rectangular background-mask jumps are implemented in the shared draw
plan for SVG/PPTX and therefore their PDF/Excel projections. Curved bridge arcs
and an Excalidraw approximation remain.

### Bridge / Jump Lines

```text
----^----
---------
```

### Features

- Segment intersection detection
- Automatic bridge generation

---

## v0.6 Live Preview

Status: initial HTTP/SSE live preview and source-positioned parser diagnostics
implemented; VS Code integration remains.

### xaligo serve

```bash
xaligo serve
```

### Features

- File watching
- Automatic re-rendering
- Real-time updates

### Backend Stack

- Go
- Echo
- WebSocket
- fsnotify

### Frontend Stack

- templ
- HTMX

### Preview Flow

#### Initial

```text
.xal
 ↓
SVG
 ↓
Browser
```

#### Real-Time Updates

```text
File Change
 ↓
Re-render
 ↓
Server-Sent Events
 ↓
Preview Refresh
```

---

## v0.7 VS Code Extension

Status: maintained in a separate repository. This core repository now provides
the required WASM diagnostics API, source positions, stable SVG rendering, and
HTTP/SSE preview protocol.

### Language Support

```text
.xal
```

### Features

- Syntax Highlighting
- Validation
- Error Location Reporting

### Preview Panel

```text
Editor
|
+- Source
+- Preview
```

### Live Preview

```text
Save
 ↓
xaligo render
 ↓
Preview Update
```

---

## v0.8 Excalidraw Integration

Status: native Excalidraw export exists; live WebView/updateScene integration
is not started.

### Excalidraw Preview

```text
.xal
 ↓
Excalidraw JSON
 ↓
WebView
```

### Features

- updateScene() support
- Real-time synchronization

### Excalidraw Export

```bash
xaligo render --format excalidraw
```

---

## v0.9 PowerPoint Export

Status: partially implemented ahead of schedule. Go-side geometry/routing plan
generation exists, and Node/PptxGenJS can generate PPTX as a temporary
development path. The required long-term gap is `xaligo.wasm`, invoked
from the Go repository layer with resolved plan JSON.

### PPTX Export

```bash
xaligo render --format pptx
```

Compatibility during transition:

- Keep `xaligo render --format pptx` usable when a WASM exporter is available.
- Do not reintroduce repository-layer Node subprocess execution as the default.
- Do not implement PPTX/OOXML writing in Go controller/repository code.
- Keep route/traffic/theme support renderer-agnostic where possible.

### Supported Features

- Shapes
- Connectors
- Routes
- Traffic Flows
- Themes

---

## v1.0

### VS Code Marketplace Release

#### Included Features

- Live Preview
- SVG Export
- Excalidraw Export
- PPTX Export
- Route Connectors
- Orthogonal Routing
- Traffic Layers
- Edge Offset
- Line Jumps

---

## Future Vision

### AWS Architecture Mode

```text
AWS Icons
Auto Layout
Route Layer
Traffic Layer
```

### Network Diagram Mode

```text
L2
L3
Route
Traffic
```

### Infrastructure as Diagram

```text
Diagram as Code
+
VS Code
+
Git
+
CI/CD
```

---

## Project Goal

Create a Diagram as Code platform positioned between:

```text
PlantUML
      +
Excalidraw
      +
draw.io
      +
Cloudcraft
      +
Instana-style Topology
```

with a strong focus on:

- AWS Architecture Diagrams
- Network Topology Diagrams
- Infrastructure Documentation
- Diagram-Driven Development
- Multi-mode 2D / 2.5D / Topology Rendering
- SVG / PPTX / Excalidraw / XYFlow / Isoflow Export

---

# 07 XAL specification


## Overview

`.xal` is a Vue-style layout DSL with XML syntax. Canonical V1 documents use a
`<xaligo>` envelope containing document-wide data and one `<frames>` page
collection. Historical `<frame>` and `<frames>` roots remain readable but emit
a migration warning.
The parser uses `encoding/xml` and handles attributes, nested tags, and text content.

## V1 Compatibility Profile and Version Boundary

Canonical V1 source explicitly sets `version="1"` on `<xaligo>`. An
unversioned `<xaligo>` defaults to V1 with a warning. A `version` value other
than `1` is invalid. Legacy `<frame>` and `<frames>` roots accept the historical
V1 version rules but always emit a warning recommending the canonical envelope.
This document-root `version` selects the DSL and is not visible page metadata.
By contrast, a non-empty `version` on an identified `<frame>` that is a direct
child of the document-root `<frames>` is that page's visible content revision;
it does not select a language version. Structural diff ignores only the
document-root DSL version and compares child-frame content revisions normally.

V2 uses a distinct, reject-safe root:

```xml
<scene version="2">
  ...
</scene>
```

`<scene>` requires `version="2"`; an unversioned `<scene>` is invalid. A V1
reader recognizes `<xaligo>`, `<frame>`, and `<frames>`, but rejects a
V2 document at the root instead of partially rendering V2 syntax as V1. Do not
use `<frame version="2">` or `<frames version="2">`.

A V2 implementation must accept this V1 profile as input, preserve its
defaults and compatibility behavior, and lower it directly to the shared typed
model. It must not rewrite V1 XML into V2 XML, parse the document twice, or
invoke V1 through a serialized intermediate representation. V1 has no
dependency on, and no obligation to understand, V2.

Canonical V1 source uses lowercase XML tag names, attribute names, and enum
tokens exactly as documented here. Historical case-insensitive or directional
aliases that are not listed in this specification are accepted implementation
details, not part of the frozen compatibility profile. A V2 compatibility
frontend canonicalizes the documented V1 values once at its input boundary.

## Root Tag

```xml
<xaligo version="1">
  <data>
    <!-- reusable definitions -->
  </data>
  <frames gap="48">
    <frame id="overview" width="1440" height="900" class="pa-4">
      ...
    </frame>
  </frames>
</xaligo>
```

`<xaligo>` permits document-level metadata, imports, data, and styles, and
requires exactly one `<frames>`. Give every child `<frame>` a stable `id`.

```xml
<xaligo version="1">
<frames gap="48">
  <frame id="overview" width="1440" height="900">
    ...
  </frame>
  <frame id="detail" width="1440" height="900">
    ...
  </frame>
</frames>
</xaligo>
```

| Attribute | Type | Default | Description |
|---|---|---|---|
| `version` | string | `"1"` with warning when omitted | On `<xaligo>` or a legacy root, selects V1 and only `"1"` is accepted. On an identified direct child `<frame>`, a non-empty value is the visible page revision |
| `title` | string | — | On a page `<frame>`, enables the metadata band and supplies its built-in `title` tag |
| `width` | float | `1280` | Frame width (px) |
| `height` | float | `720` | Frame height (px) |
| `class` | string | — | Spacing class |
| `layout` | string | — | Set to `"horizontal"` to arrange children horizontally |
| `gap` | float | `16` | Gap between child elements (px) |
| `item-size` | float | render-context default, normally `32` | Max icon size (px) applied to all `<item>` elements in this file. Overrides the native `item.icon_size` or embedded asset-source value |
| `margin` / `margin-*` | float | — | DSL content whitespace in pixels. On root `<frame>`, paper-frame size is preserved and content is inset. This is separate from PPTX CLI `--paper-margin*` flags, which are inch-based export fitting margins |
| `content-width` / `content-height` | float | — | Shrink usable inner layout area |
| `align` | string | — | Align usable content area (`top|middle|bottom` + `left|center|right`) |
| `overflow` | string | `error` | Child containment policy: `error` or `visible` |

Legacy input may still use root `<frame>` or `<frames>`. It remains renderable,
but diagnostics recommend wrapping identified frames in the canonical
`<xaligo version="1"><frames>...</frames></xaligo>` envelope.

`<frames>` accepts `gap` and optional `layout="vertical"`. Without
`layout="vertical"`, frames are arranged horizontally. A `<frame>` inside
`<frames>` requires a non-empty `id`.

### Frame and physical-page contract

An identified child `<frame>` is the V1 physical page unit. Frames are emitted
in source order after the complete document scene and all cross-frame links
have been resolved.

| Format | Default mapping |
|---|---|
| SVG | One `.svg` artifact per frame |
| PPTX | One slide per frame |
| PDF | One page per frame |
| Excel | One worksheet per frame, containing the frame's SVG image |
| Excalidraw, XYFlow, Isoflow | One logical document containing all frames |

SVG, PPTX, PDF, and Excel omit the page-frame outline in both default and
`--combine-frames` output. Frame geometry remains authoritative for page size,
cropping, endpoint ownership, and the outer logical page edge used to select a
cross-frame page-link side and tangent anchor. The drawable frame terminal may
sit on a parallel inward inset line. A default page-local SVG uses the exact
frame rectangle as its canvas and clip boundary; PDF pages and Excel page
images inherit that strict crop. Combined SVG compatibility output retains
marker-safe bounds expansion. Excalidraw retains editable frame structure with
transparent page-frame strokes.

For a document with one child frame, SVG writes exactly the requested output
path. For multiple child frames, an output request such as `diagram.svg`
produces `diagram-<safe-frame-id>.svg` for each frame. The safe ID retains ASCII
letters, digits, `_`, and `-`; every run of other characters becomes one `-`,
leading and trailing `-` are removed, and an empty result falls back to
`frame-<source-order>`. Two IDs that resolve to the same output filename are an
error. SVG does not create an implicit archive.

`--combine-frames` is the explicit compatibility option for page-oriented
formats. It restores the historical single canvas, single slide, single PDF
page, or single Excel worksheet. It does not change Excalidraw, XYFlow, or
Isoflow because those formats are already single logical documents.

### Frame metadata tag band

An identified page frame may expose `id`, `title`, a page-content `version`,
and arbitrary key/value entries as a two-cell tag band. The resolved metadata
`row-gap`, 4 layout pixels by default, is both the space between wrapped rows
and the metadata page-edge inset at the selected vertical edge and both
horizontal edges. Frame padding, content margins, and the content box never
replace or add to that inset. The band is enabled when the page frame has a
non-empty `title`, a child-frame content `version`, or a direct `<metadata>`
child. Existing identified frames that have only an `id` remain visually
unchanged. Once the band is enabled, non-empty built-ins are emitted in stable
`id`, `title`, `version` order, followed by `<entry>` children in source order.

```xml
<frame id="aws-architecture" title="AWS Architecture" version="1.0.0"
       width="720" height="400"
       margin-top="52" margin-right="24"
       margin-bottom="52" margin-left="24">
  <metadata position="top" align="right" font-family="helvetica">
    <entry key="owner" value="Platform Engineering" />
    <entry key="status" value="Approved" break-before="true"
           width="180" key-width="56" />
  </metadata>
  <rectangle id="diagram" title="Page content" />
</frame>
```

`<metadata>` is a non-layout direct child of a page `<frame>` and may occur at
most once. This context is distinct from document-level
`<xaligo><metadata>`. It contains only empty `<entry>` children; every entry
requires non-empty `key` and `value` attributes. Duplicate keys are retained.

| Attribute | Target | Default | Contract |
|---|---|---|---|
| `position` | `metadata` | `top` | Closed enum `top|bottom` |
| `align` | `metadata` | `left` | Closed enum `left|center|right`; applied independently to each resolved row |
| `font-family` | `metadata` | `virgil` | `virgil|helvetica|cascadia|assistant|excalifont|nunito|lilita-one|comic-shanns|liberation-sans` |
| `font-size` | `metadata` | `12` | Positive layout pixels; tag height is exactly `ceil(font-size * 1.2) + 4` |
| `color` | `metadata` | `#64748b` | Value text color; also key text color unless `key-color` is set |
| `key-color` | `metadata` | value of `color` | Key text color |
| `background-color` | `metadata` | `transparent` | Value-cell fill |
| `key-background-color` | `metadata` | `#f8fafc` | Key-cell fill |
| `border-color` | `metadata` | `#cbd5e1` | Cell border color; the cell stroke is fixed at `0.75` layout pixels |
| `width` | `metadata`, `entry` | auto | Positive total key/value tag width. An entry value overrides the metadata-level default |
| `key-width` | `metadata`, `entry` | auto | Positive key-cell width smaller than total width. An entry value overrides the metadata-level default |
| `gap` | `metadata` | `8` | Non-negative horizontal gap between tags |
| `row-gap` | `metadata` | `4` | Non-negative gap between wrapped rows and the same-sized inset from the selected top/bottom edge and both horizontal page edges |
| `break-before` | `entry` | `false` | Closed boolean `true|false`; `true` starts this entry on a new row when a preceding tag exists |

Colors use `#RRGGBB` or `transparent`. Auto width measures both cells with the
selected font and full-width-rune-aware metrics. Omit `width` or `key-width`
to request auto sizing; the literal string `auto` is not a V1 numeric value.
Fixed widths use no-wrap shrink-to-fit with clipping as the final overflow
guard. Tags preserve input order and use greedy left-to-right packing against
the usable width `frame.width - 2 * row-gap`, which produces the minimum row
count without reordering. The usable width must remain positive.
`break-before="true"` forces a row boundary before that custom entry. The
metadata `align` is then applied to each row separately against that same
usable width: left starts at `frame.x + row-gap`, right ends at
`frame.x + frame.width - row-gap`, and center still uses the frame center.

For `position="top"`, the band starts at `frame.y + row-gap`; for
`position="bottom"`, it ends at `frame.y + frame.height - row-gap`. The
metadata-side reservation strip spans the full frame width from the outer
logical frame edge to the corresponding boundary of the final content box.
Its depth is never less than `row-gap` plus the complete band height plus the
fixed 8-pixel content gap: if the normal content boundary is closer, it is
moved inward to that minimum; if it is already farther inward, it is retained.
Normal items and their text, local connector paths and labels, UML connector
paths and labels, and cross-frame page-link paths and labels cannot enter this
strip. `overflow="visible"` never overrides this page-decoration exclusion.
The frame's outer page size and invisible logical edge do not change.
The inset is measured from that logical frame edge before any common PPTX slide
centering and is unrelated to the export-only `--paper-margin*` options.
For a cross-frame page link, the same resolved `row-gap` is also the inward
normal inset of safe frame terminals on all four sides, independent of metadata
`position`; zero retains the outer edge. A frame without metadata instead uses
a 4-layout-pixel terminal inset. An explicit `src/dst-frame-side` or
`src/dst-frame-anchor` requires the inset to be strictly smaller than its
specified side's normal frame dimension: height for `top`/`bottom`, width for
`left`/`right`. Without an explicit frame terminal, validation requires at
least one side that satisfies this inset bound and the metadata reservation;
the shared scene later selects among those safe sides from rendered visual
geometry. The inset is never implicitly clamped.

The shared layout and presentation scene own this geometry. SVG, PPTX, PDF,
Excel, and Excalidraw render the owning frame's tags; per-frame projection
cannot leak another page's band, and combined output retains every band. The
entire reservation strip, rather than only the tag cells, is a hard exclusion
zone for normal rendered geometry. XYFlow and Isoflow omit the band because it
is page decoration rather than a graph node or endpoint.

## Numeric and Geometry Contract

Numeric attributes are validated before layout. A numeric value must be a
finite base-10 number; `NaN`, positive or negative infinity, an empty numeric
value, and malformed trailing text are errors. The current implementation
validates the source attributes and then reads those validated values during
layout; replacing the string attribute map with a typed normalized layout
specification is a separate roadmap item.

The following domain rules apply:

| Attributes | Required domain |
|---|---|
| `width`, `height`, `component-width`, `component-height`, `interface-width`, `content-width`, `content-height`, `item-size`, `font-size`, `key-width` | greater than `0` when specified |
| `row`, `col` | greater than `0` when specified |
| `span` | greater than `0` and at most `12`; flexible sibling spans in one `<row>` must total at most `12` |
| `gap`, `row-gap`, margins, spacing-class padding | greater than or equal to `0` |
| `scale`, `coordinate-scale`, `grid`, `stroke-width` | greater than `0` when specified |
| `x`, `y`, `dx`, `dy`, bend coordinates | any finite value, subject to the containing geometry rule |

An omitted attribute uses its documented default. An explicitly empty
`align` is treated as omitted; it must not produce an invalid-alignment warning.
Unknown non-empty enum values remain errors or source-positioned warnings as
specified by that attribute.

V1 intentionally distinguishes strict values from compatibility fallbacks:

| Input | V1 behavior |
|---|---|
| Invalid `overflow`, connection side, or connection anchor | Validation error |
| Unknown `layout`, connection `kind`, stroke style, arrowhead, or arrowhead-size value | Validation error |
| Unknown render mode, format, theme, paper/orientation, arrow-style option, or SVG legend position | Render-option error. The CLI normalizes `xlsx` to `excel` before validation |
| Recognized but unavailable render mode (`aws-2.5d` or `topology`) | Not-implemented error |
| Empty `align` | Omitted; defaults to `top-left` |
| Malformed or unknown non-empty `align` | Warning; each unsupported component keeps its `top` or `left` default |
| Unknown nested attribute or malformed/unrecognized spacing-class token | Ignored; a recognized numeric negative spacing class remains an error |

These fallbacks are part of V1 compatibility, not a mechanism for opting into
V2. The distinct V2 root prevents new V2 constructs from being silently
treated as V1 extensions.

`validate` and every render format use the same normalized values and resolved
geometry checks. Successfully validated input must not later produce `NaN`,
`Inf`, a negative drawable size, or an output serialization error caused
by geometry.

### Fixed and flexible child allocation

For a vertical parent, an explicit child `height` is a fixed main-axis size;
for a horizontal parent, an explicit child `width` is fixed. The parent first
reserves fixed sizes, margins, and gaps. Children without a fixed main-axis size
divide the remaining space using their positive `row` or `col` weights. A
`<row>` uses validated `span` values against its 12-column grid.

The resolved child size is the size used both for recursive layout and for
placing the next sibling. A child cannot replace its assigned size after the
parent cursor has advanced. Explicit cross-axis sizes must fit the parent's
content box unless overflow is explicitly allowed.

Layout parents accept `overflow`:

| Value | Behavior |
|---|---|
| `error` | Default. A child outside the parent's content box is a source-positioned validation error. |
| `visible` | The child may extend outside the content box, but all coordinates and sizes must remain finite and sibling cursors still use resolved sizes. |

The policy belongs to a parent and applies only to its direct children; it is
not inherited. If fixed children consume the full main axis under `visible`,
the parent's original usable extent is used as the flex pool and the flexible
children receive their weighted sizes while all children retain source order.
Sibling cursors use each resolved size, gap, and margin, making the resulting
overflow explicit. Under the default `error` policy the same layout is
rejected.

Overflow is never silently introduced by a renderer. Clipping is a drawing and
text policy and does not make invalid layout geometry valid.

## Layout Tags

### `<container>`

Stacks children **vertically** (same behavior as `frame`). Use `layout="horizontal"` for horizontal arrangement.

```xml
<container class="pa-4" gap="16">
  ...
</container>
```

| Attribute | Type | Default | Description |
|---|---|---|---|
| `layout` | string | — | `"horizontal"` to arrange children side by side |
| `gap` | float | `16` | Gap between child elements (px) |
| `content-width` / `content-height` | float | — | Shrink usable inner layout area |
| `align` | string | — | Align usable content area |
| `overflow` | string | `error` | Child containment policy: `error` or `visible` |

### `<row>`

Lays out children **horizontally** in a 12-column grid.

```xml
<row gap="20">
  <col span="8">...</col>
  <col span="4">...</col>
</row>
```

| Attribute | Type | Default | Description |
|---|---|---|---|
| `gap` | float | `16` | Column spacing (px) |
| `overflow` | string | `error` | Child containment policy: `error` or `visible` |

> `<row>` is a **pure layout tag** — it does not render any border or label in the output.
> The `<col>` children are also pure layout containers.

An explicit child `width` is reserved before the grid share and is excluded
from `span` allocation. Among children without fixed width, an omitted `span`
defaults to `12 / number_of_flexible_children`; explicit flexible spans must
total at most `12`. Unused span leaves intentional trailing space.

### `<col>`

A vertical stack container inside `<row>`. Use `span` to set the number of columns occupied.

| Attribute | Type | Default | Description |
|---|---|---|---|
| `span` | float | `12 / num_columns` | Columns to occupy (out of 12) |
| `class` | string | — | Spacing class |
| `overflow` | string | `error` | Child containment policy: `error` or `visible` |

## Custom Leaf and Container Tags

An otherwise unknown nested tag with no layout children is a generic leaf and
is rendered as a rectangle plus text. An unknown nested tag with layout
children is a generic group/container: it receives the normal group header
insets and lays out those children vertically by default, horizontally for
`layout="horizontal"`, or with the V1 staggered layout for
`layout="staggered"`. If every child is item-like (`item`, `spacer`, or
`blank`), the children use the item-grid row behavior instead.

This rule applies only below a valid V1 root. An unknown root is always a parse
error, so `<scene version="2">` can never be mistaken for a generic V1 group.

```xml
<card title="Dashboard" />
<panel title="Main Chart" />
<text>Any label</text>
```

| Attribute | Behavior |
|---|---|
| `title` | Display label (takes priority) |
| Text content | Label when `title` is absent |
| (none) | Tag name used as label |
| `border` | Set to `"none"` to hide the border |
| `visible` | Set to `"false"` to hide only this component (border, icon, label). Children are still rendered individually. Layout space is preserved |
| `font-size` | Text font size in layout pixels |

## `<rectangle>` and `<port>` Tags

`<rectangle>` creates a general-purpose rectangle. Its label comes from
`title` or direct text content, and `font-size` controls the label size.
Unlike generic leaf tags, `<rectangle>` may contain multiple `<port>` children.

`<port>` creates a small rectangle inside a side of the parent rectangle.
Multiple ports on the same side are spaced evenly along that side. Its label
also comes from `title` or direct text content, and it supports `font-size`.

```xml
<rectangle id="service" title="Service" width="180" height="100" font-size="18">
  <port id="service-in" side="left" title="in" font-size="9" />
  <port id="service-out" side="right" title="out" font-size="10" />
</rectangle>
```

| Attribute | Target | Description |
|---|---|---|
| `id` | `rectangle`, `port` | Required connection reference ID |
| `width` / `height` | `rectangle`, `port` | Size in layout pixels |
| `title` / text content | `rectangle`, `port` | Text rendered inside the shape |
| `font-size` | `rectangle`, `port` | Text font size in layout pixels |
| `side` | `port` | Parent side: `top`, `right`, `bottom`, or `left`. Default `top` |
| `x` / `y` | `port` | Optional position relative to the parent rectangle's top-left corner. Values are clamped so the port remains inside the parent rectangle |

Port boxes must remain inside their parent rectangle. Explicit positions are
normalized before drawing, and overlapping ports on the same side are a layout
diagnostic rather than a renderer-specific accident. Port text carries the
shared text-layout policy: SVG, its PDF/Excel projections, and PPTX enforce it,
while editable Excalidraw-compatible output preserves it in metadata for
bound-text consumers.

## Resolved Text Layout

Text has both a geometry box and a semantic role. Scene and plan construction
must preserve the resolved role, wrapping, fitting, clipping, line height, and
padding instead of making each encoder infer them from generated IDs.

Built-in defaults are:

| Role | Wrap | Fit | Overflow |
|---|---|---|---|
| group header | no | shrink | clip to text box |
| ordinary label | yes | shrink | clip to text box |
| item label | yes | shrink | clip to text box |
| port label | yes | shrink | clip to port box |
| connector label | yes | shrink | clip to text box |

The default line-height multiplier is `1.2` unless the source scene carries a
valid positive value. Font sizes originate in layout pixels and are converted
with the same effective scale as the containing geometry. Changing
`--px-per-inch` or paper fitting therefore preserves the text-to-shape ratio.

An encoder may use native text fitting or deterministic line breaking, but the
visible result must obey the resolved policy. Editable Excalidraw-compatible
bound text carries the same `xaligoTextLayout` metadata and must not become a
separate layout authority. Encoders apply text policy in this order: resolve
padding, wrap when enabled, shrink when requested, then clip when
`TextLayout.overflow="clip"`.

## `<table>` Tag

`<table>` represents general tabular data and is distinct from the `<grid>`
layout primitive and relational `<entity>` definitions. It accepts either a
GFM-like pipe table or explicit `<header>`/`<row>` children containing
`<cell>` elements. One pipe header may be followed by explicit rows; a second
header is invalid. Both syntaxes normalize to the same typed rows before
layout.

```xml
<table id="services" title="Services">
  | Service | Role    | Port |
  |:--------|:--------|-----:|
  | API     | Backend | 8080 |
  <row><cell>DB</cell><cell>Database</cell><cell align="right">5432</cell></row>
</table>
```

Pipe separators require at least three hyphens per cell and use colons for
left, center, and right alignment. A literal pipe is escaped as `\|`. Every
header and row must have the same positive cell count. Tagged cell `align`
accepts `left`, `center`, `right`, or the normal vertical-horizontal values.
Unknown children, duplicate headers, malformed separators, and mismatched
column counts are positioned errors.

Table presentation attributes inherit from `<table>` to every row and cell.
Tagged `<header>`, `<row>`, and `<cell>` attributes override inherited values.
Pipe cells have no inline attributes, so style them with table attributes and
the `header-*` variants:

```xml
<table color="#172033" background-color="#ffffff" border-color="#94a3b8"
       font-family="nunito" font-size="16"
       header-color="#ffffff" header-background-color="#2563eb"
       header-font-family="cascadia" header-font-size="18">
  | Service | Port |
  |:--------|-----:|
  | API     | 8080 |
</table>
```

| Attribute | Values | Description |
|---|---|---|
| `color` | `#RRGGBB` or `transparent` | Cell text color |
| `background-color` | `#RRGGBB` or `transparent` | Cell fill color |
| `border-color` | `#RRGGBB` or `transparent` | Cell border color |
| `font-family` | `virgil`, `helvetica`, `cascadia`, `assistant`, `excalifont`, `nunito`, `lilita-one`, `comic-shanns`, or `liberation-sans` | Cell font family |
| `font-size` | positive number | Cell font size in layout pixels |
| `header-*` | corresponding value above | Pipe/tag header override declared on `<table>` |

Colors require six-digit hexadecimal notation. The style precedence is
`cell > header/row > table > built-in default`. Font-family names are carried
through the renderer-neutral scene and mapped to the corresponding output font
face; an output environment may substitute a missing installed font.

## Relational Database Tags

Reusable `<database-schema id="...">` definitions belong under `<data>`.
Frames render them through `<database data="schema-id">`. A schema contains
identified `<entity>` definitions; each entity contains typed `<column>`
definitions and optional single-column `<foreign-key>` definitions.

```xml
<database-schema id="app">
  <entity id="roles"><column name="id" type="bigint" primary-key="true" /></entity>
  <entity id="users">
    <column name="role_id" type="bigint" nullable="false" />
    <foreign-key columns="role_id" references="roles.id" />
  </entity>
</database-schema>
```

Columns support `name`, `type`, `primary-key`, `nullable`, `unique`, and
`default`. A foreign key requires one local column and one
`references="entity.column"` target and generates a relation between the
entities. Duplicate or missing schema/entity/column references and mixed
inline/data-backed database content are positioned errors. Composite keys,
indexes, checks, and import dialects remain planned V1 extensions.

## UML Tags

`<uml>` is the common V1 component for the supported core UML diagram families. It
adapts their typed elements, compartments, and relations to xaligo's shared
layout, shape, connector, and output pipeline. Supporting a family selector
does not imply support for every UML 2.x glyph or interchange construct; the
closed V1 vocabulary and its projection limits are defined below.

### Component, identity, and layout contract

```xml
<xaligo version="1">
  <data>
    <uml-model id="domain-model">
      <class id="order" title="Order">
        <attribute>- id: UUID</attribute>
        <operation>+ confirm()</operation>
      </class>
      <interface id="repository" title="OrderRepository">
        <operation>save(order: Order)</operation>
      </interface>
      <realization src="order" dst="repository" title="persists" />
    </uml-model>
  </data>
  <frames>
    <frame id="domain" width="960" height="540">
      <uml id="model" title="Domain Model">
        <class-diagram data="domain-model" direction="right" />
      </uml>
    </frame>
  </frames>
</xaligo>
```

The following rules are normative:

- `<uml>` must be inside a frame, requires a non-empty `id`, and contains
  exactly one supported diagram-kind child. UML IDs must be unique within that
  frame. The same UML ID may be reused in a different frame.
- UML component IDs and diagram-local element IDs must not contain whitespace,
  `.` or `/`. `.` is reserved for the frame boundary and `/` for the UML
  boundary in public connection references.
- The diagram-kind child contains a non-empty set of direct element and
  relation children. Unknown diagram kinds and unknown children are errors;
  arbitrary custom tags are not generic UML elements.
- Every UML element requires a non-empty diagram-local `id`, unique within the
  UML component. A UML relation's `src` and `dst` use those local IDs, without
  a UML or frame prefix, and both endpoints must exist in the selected model.
- `direction` on the diagram-kind child accepts only `right` or `down`. When
  `<uml layout>` is omitted, `direction="right"` and sequence diagrams
  diagrams select horizontal xaligo layout; the other cases select vertical
  layout. This controls the V1 projection and is not a UML semantic ordering
  rule.
- When `<uml title>` is omitted, the selector name without `-diagram` is used.
  Element labels resolve in the order `title`, `name`, direct text, then local
  `id`. UML elements default to `font-family="helvetica"` and `font-size="14"`;
  normal element font attributes override those defaults. An element `name`
  is display text only and never becomes a frame-level connection alias; use
  the public UML reference described below.
- A class-diagram classifier with a non-empty `stereotype` renders
  `«stereotype»` as a separate first header line. `abstract="true"` and
  `static="true"` append `{abstract}` and `{static}` to the classifier-name
  header line. These lines remain one graphical header even when the
  classifier has no compartments.
- The compatibility tags `<element>` and `<relation>` are not part of the
  strict V1 UML profile. A model must use one of the element and relation tags
  allowed for its selected family.

UML elements are also normal xaligo connection endpoints. A frame-level
`<connection>` uses the following public references; the internal hex-scoped
scene ID is opaque and must not be written in source:

| Location | Public endpoint reference | Meaning |
|---|---|---|
| Same frame | `uml-id/local-id` | Element `local-id` in UML component `uml-id` |
| Another frame | `frame-id.uml-id/local-id` | The same UML element reached across a frame boundary |

For example, element `order` in `<uml id="model">` inside frame `overview` is
`model/order` to a normal connection in that frame and
`overview.model/order` to a normal connection in another frame. Omitting the
`frame-id.` prefix for a cross-frame endpoint is an unresolved-reference
error. UML-native relations continue to use `src="order"`, not either public
connection form.

### Diagram-kind vocabulary

| Diagram kind | Allowed elements | Allowed relations | Additional V1 semantic checks |
|---|---|---|---|
| `class-diagram` | `class`, `interface`, `enumeration` | `association`, `aggregation`, `composition`, `generalization`, `realization`, `dependency` | Requires one classifier. Aggregation/composition are class to class; generalization joins equal classifier kinds; realization is class to interface. |
| `component-diagram` | `component`, `interface`, `port`, `artifact` | `dependency`, `realization`, `association`, `assembly`, `delegation` | Requires one component. Realization is component to interface. Assembly requires at least one port endpoint. Delegation starts at a port. |
| `activity-diagram` | `initial`, `final`, `activity`, `action`, `decision`, `merge`, `fork`, `join`, `object-node` | `control-flow`, `object-flow` | Requires an activity/action. Control-flow excludes object-node. Object-flow requires an object-node endpoint. Initial/final direction and control-node degrees are validated. |
| `state-machine-diagram` | `initial`, `final`, `state`, `history`, `choice`, `fork`, `join` | `transition` | Requires a state. Initial/final direction and pseudostate degrees are validated. |
| `sequence-diagram` | `participant`, `lifeline` | `message`, `return-message`, `create-message`, `destroy-message` | Requires a participant/lifeline. Every message has a diagram-unique order. Create/destroy cannot be self messages. |

The endpoint contracts above are closed. An admitted relation with an endpoint
pair not described by its row is a validation error.

### Component diagram sizing

Component boxes use automatic height by default. The resolved height reserves
the component header, one compact row for every declared boundary interface,
and additional vertical space when multiple incoming component associations
bind to the same destination interface. Interface groups are packed from the
header downward; unused diagram height is not redistributed into component
rows.

`component-width` and `component-height` on `<component-diagram>` set positive
diagram-wide defaults. `width` and `height` on an individual `<component>`
override the corresponding diagram default. When neither height attribute is
present, automatic height remains active. An explicit height is authoritative
and must be large enough for its interface and connection visuals to remain
inside the component.

A positive `interface-width` on `<component>` sets one common width for every
interface-name box in that component. It is not accepted on individual
`<interface>` children. The configured width must leave enough horizontal
space for the component's interface descriptions when they are present.

### Activity partitions

An `activity-diagram` may group activity elements in swimlanes with direct
`<partition id="..." title="...">` children. A partition may contain only
activity elements allowed by the `activity-diagram` row above. Partition IDs
are diagram-local identifiers and must be unique. A nested element may repeat
`lane="partition-id"`; when present, it must match the enclosing partition.
Elements may also be declared directly under the diagram with
`lane="partition-id"`, but the referenced partition must exist.

`lanes="vertical|horizontal"` is accepted only on `activity-diagram`.
`theme="xaligo"` selects the supported activity swimlane visual theme. Other
diagram families reject `lanes` and `theme`.

### Ownership

`owner` is reserved except for component `port` elements. A component port
requires `owner="component-id"`, and that owner must reference a component in
the same diagram.

### Element compartments

An element's direct child tags are ordered text compartments. Each compartment
must have non-whitespace direct text, `title`, or `name` and must not contain
child elements. Nested UML elements and relations are not compartments. The
typed compartment vocabulary is:

| Element | Allowed typed compartments |
|---|---|
| `class` | `attribute`, `operation`, `constraint`, `note` |
| `interface` | `operation`, `constraint`, `note` |
| `enumeration` | `literal`, `operation`, `note` |
| `component` | `interface`, `provided-interface`, `required-interface`, `property`, `constraint`, `note` |
| `artifact` | `property`, `responsibility`, `note` |
| `activity`, `action` | `responsibility`, `constraint`, `note` |
| `state` | `entry`, `do`, `exit`, `region`, `note` |

Elements absent from this table do not accept compartments. The generic
`<compartment>` child is a compatibility spelling accepted wherever a typed
compartment is allowed; new source should use the typed tag because its meaning
survives future semantic processing. Compartment source order is preserved,
but compartments are not independent connection endpoints. In a class diagram,
adjacent structural (`attribute` or `literal`) compartments and adjacent
behavioral compartments may share one graphical section; a transition between
those kinds starts a new section without reordering either kind. Every
newline-separated compartment line contributes to the classifier's intrinsic
height.

### Relation attributes, order, and time

Every relation requires `src` and `dst`. `title` or `label` supplies its
visible text. `guard` is allowed only on flows/transitions and is appended as
`[guard]`. `route` is retained as UML relation metadata for flow/transition
routing hints such as `route="loop"`. `src-multiplicity` and
`dst-multiplicity` are allowed only on association, aggregation, and
composition and are appended in
source-to-destination order. Relation color and normal connector side, anchor,
stroke-width, bend, scale, and grid attributes use the `<connection>` rules.
`kind`, `stroke-style`, `arrowhead`, `start-arrowhead`, and `end-arrowhead` are
invalid because the UML relation kind owns line and marker semantics.

Sequence message kinds require `order`. Its canonical form
is one or more positive decimal integers without leading zeroes, separated by
dots, for example `1`, `2`, or `1.1`. The complete order string must be unique
across all messages in one diagram. Numeric order is prepended to the rendered
label and assigns top-to-bottom connector anchors on participant/lifeline
shapes. It does not reorder declared elements or create activation boxes or a
separate interaction axis. Sequence message anchors always use a vertical
element edge so the ordering remains vertical: explicit `top` is normalized to
`left`, explicit `bottom` to `right`, and an explicit anchor slot is superseded
by the normalized order position.

### Relation projection

UML relations lower to the shared orthogonal connector model with the
following fixed semantic defaults:

| Projection | Relation kinds |
|---|---|
| Dashed line with destination triangle | `dependency`, `realization`, `return-message` |
| Solid line with destination triangle | `generalization`, `control-flow`, `object-flow`, `transition`, `message`, `create-message`, `destroy-message` |
| Source diamond | `aggregation`, `composition` |
| No destination arrowhead | `association` |

The visible relation label is placed near the routed connector midpoint and
the UML diagram/relation kind is retained as semantic metadata where the target
format can carry it.

### Reusable UML models

Reusable definitions use `<uml-model id="...">` directly below document
`<data>`. A diagram-kind child selects one with `data="model-id"`:

```xml
<data>
  <uml-model id="order-domain">
    <class id="customer" title="Customer" />
    <class id="order" title="Order" />
    <association src="customer" dst="order" title="places" />
  </uml-model>
</data>
<frames>
  <frame id="domain">
    <uml id="model"><class-diagram data="order-domain" direction="right" /></uml>
  </frame>
</frames>
```

`<uml-model>` requires a document-unique ID. A missing model, duplicate model
ID, or a diagram that combines `data` with inline children is an error. The
model itself does not declare a UML family; after expansion, all of its element,
compartment, ownership, relation, order, and endpoint rules are validated
against the selecting diagram kind. One reusable model is therefore reusable
across selectors only when every child belongs to each selector's closed
vocabulary.

### Deliberately lossy V1 projection

V1 preserves the selected UML family, element kind, relation kind, owner, and
relation label in the shared semantic scene, then projects them into the
capabilities common to xaligo outputs:

- `initial` and `final` become ellipses;
- `decision`, `merge`, `choice`, and `history` become diamonds;
- every other element becomes an editable rectangle whose ordered
  compartments are flattened into its visible text;
- every relation becomes a shared orthogonal connector with a separate label;
  aggregation and composition currently share the same diamond projection;
- sequence order is retained in labels and metadata and controls top-to-bottom
  message anchors, but V1 does not draw dashed lifelines, activations, combined
  fragments, or a separate vertical event axis;
- current V1 UML elements do not support semantic ownership.

SVG, Excalidraw, PPTX, PDF, Excel, XYFlow, and Isoflow all consume this same resolved
geometry. Excalidraw-compatible output carries xaligo UML custom data for
editing. XYFlow retains UML node and relation fields in node/edge `data` and
records the projected node shape. Isoflow projects every connected UML shape
to a labeled generic endpoint icon because its upstream connector schema has no
arbitrary UML data field. Another target schema may omit marker details it
cannot represent. An encoder must use native target constructs where available
and must not add private schema-breaking fields. The output is not XMI and is not a
lossless UML interchange representation.

## `<item>` Tag

A leaf element that places an AWS service icon inside a container.
Specify a positive signed 32-bit decimal ID from `service-catalog.csv` as the
`id` attribute (`1` through `2147483647`).
The icon is rendered to fit within the specified size (`item-size`).

The effective item size is resolved from the root `item-size` when present;
otherwise it comes from the render context. Native configuration and the
canonical embedded-asset profile default to `32` layout pixels. Callers that
provide a custom asset source may intentionally choose another value. For
cross-environment reproducibility, declare `item-size="32"` (or another fixed
positive value) on the document root.

```xml
<public-subnet id="public-subnet" title="Public Subnet">
  <item id="1178" />   <!-- with icon -->
  <item />             <!-- spacer: no icon, only a layout slot -->
  <item id="1189" />   <!-- with icon -->
</public-subnet>
```

| Attribute | Type | Required | Description |
|---|---|---|---|
| `id` | positive int32 | — | Decimal service ID `1..2147483647` from `service-catalog.csv`. Omitted or empty → treated as spacer; zero, signs, non-decimal syntax, and out-of-range values are invalid |
| `dx` / `dy` | float | — | Relative icon offset in pixels from the icon's normal layout `x,y` position. The moved icon rectangle must remain inside the parent frame/group border |

> If no icon is found for the given `id`, rendering skips the item and emits a warning rather than failing the document.

## `<spacer>` / `<blank>` Tags

Dedicated empty layout tags, usable as alternatives to `<item />`.
They occupy layout slots but render no icon, label, border, or text.

```xml
<public-subnet id="public-subnet" title="Public Subnet">
  <item id="1178" />
  <spacer />          <!-- empty slot: no icon -->
  <blank />           <!-- empty slot: no icon -->
  <item id="1189" />
</public-subnet>
```

No attributes (`id` is ignored if specified).

## `<connection>` Tag

Draws an **elbowed arrow** between `<item>` elements or group borders.
Must be written as a direct child of `<frame>` or inside a frame-level
`<connections>` tag, **outside** layout tags.
Use the same catalog IDs as `<item id="N">`, or assign `id`, `name`, or `ref`
to an AWS/group tag, for `src` / `dst`.

```xml
<frame width="1122" height="794" class="pa-4">
  <aws-cloud id="cloud" title="AWS Cloud">
    <public-subnet id="public" title="Public Subnet">
      <item id="1178" />
      <item id="1189" />
    </public-subnet>
  </aws-cloud>

  <!-- connections go last, as direct children of frame or inside <connections> -->
  <connections grid="8">
    <connection src="1178" dst="1189" />
    <connection src="public" dst="cloud" kind="route" />
  </connections>
</frame>
```

### `<connections>` Tag

`<connections>` is an optional direct child of `<frame>` that groups
`<connection>` elements and provides shared defaults. It does not render a
shape or occupy layout space. Any per-connection attribute overrides the parent
default.

Only these non-empty group attributes are inherited:
`arrowhead-size`, `kind`, `color`, `stroke-width`, `width`, `stroke-style`,
`start-arrowhead`, `end-arrowhead`, `arrowhead`, `scale`,
`coordinate-scale`, and `grid`. Endpoint identity and geometry are deliberately
not inherited: every child must supply its own `src` and `dst`, and
`src-side`, `dst-side`, `src-anchor`, `dst-anchor`, `src-frame-side`,
`dst-frame-side`, `src-frame-anchor`, `dst-frame-anchor`, bends, points, and
via data remain child-local. Defaults are applied to a connection snapshot
during scene construction; the parsed child node is not mutated.

`stroke-width`/`width`, `end-arrowhead`/`arrowhead`, and
`coordinate-scale`/`scale` are semantic alias pairs. If a child supplies either
name, neither name is inherited from the parent. When the child itself supplies
both, the first canonical name in each pair takes precedence.

`<connections>` may contain only `<connection>` children. A misspelled or
otherwise unknown child is a validation error rather than a silently skipped
connector.

```xml
<connections kind="traffic" color="#2563eb" grid="8" scale="1">
  <connection src="web" dst="app" />
  <connection src="app" dst="db" color="#059669" />
</connections>
```

| Attribute | Type | Required | Description |
|---|---|---|---|
| `src` | string | ✓ | Catalog ID, or `id`/`name`/`ref` of the arrow start item, AWS group, rectangle, port, or identified child frame |
| `dst` | string | ✓ | Catalog ID, or `id`/`name`/`ref` of the arrow end item, AWS group, rectangle, port, or identified child frame |
| `src-side` / `dst-side` | string | — | Optional endpoint side: `top`, `right`, `bottom`, or `left` |
| `src-anchor` / `dst-anchor` | string | — | Optional edge anchor. Each side has five inset positions (`top-1` through `top-5`, etc.) for 20 unique perimeter anchors |
| `src-frame-side` / `dst-frame-side` | string | — | Cross-frame-only logical page side, independent of the endpoint side; the drawable terminal uses that side's inward inset line |
| `src-frame-anchor` / `dst-frame-anchor` | string | — | Cross-frame-only logical page side and tangent slot. Uses `top|right|bottom|left-1..5`, or a side plus numeric/named slot; inward inset does not change the slot coordinate |
| `arrowhead-size` | string | — | V1 fixed arrowhead size: `"s"` (small). This is the default; `m` and `l` are not V1 values because V1 cannot preserve them across all render formats |
| `kind` | string | — | `connection` for the normal connector, `route` for a structural path without arrows, or `traffic` for directional flow drawn beside a matching route |
| `color` | `#RRGGBB` | — | Six-digit hexadecimal stroke color override. Named, short, and alpha colors are invalid in V1 so every format preserves the same color |
| `stroke-width` / `width` | float | — | Positive stroke width override; `width` is the compatibility alias |
| `stroke-style` | string | — | `solid`, `dashed`, or `dotted` |
| `start-arrowhead` / `end-arrowhead` | string | — | Independently set either end to `none`, `arrow`, `triangle`, `stealth`, `diamond`, or `oval`. An effective `kind="route"` permits only `none` |
| `arrowhead` | string | — | Backward-compatible alias for `end-arrowhead`; an effective route permits only `none` |
| `bends` / `points` / `via` | string | — | Backward-compatible inline coordinate list. Prefer child tags for multiple bend coordinates |
| `scale` / `coordinate-scale` | float | — | Positive multiplier applied to bend coordinates before routing. Default `1` |
| `grid` | float | — | Positive per-connection snap grid in layout pixels. Defaults to the router grid |

Default connections and `kind="traffic"` use a thin 1px line with
`start-arrowhead="none"` and a slender `stealth` end arrowhead. `kind="route"`
uses `start-arrowhead="none"` and `end-arrowhead="none"` by default. Default
colors are `#1e1e1e` for normal connections, `#64748b` for routes, and
`#2563eb` for traffic. A route is always headless in V1: after applying
`<connections>` defaults and child alias overrides, any effective non-`none`
`start-arrowhead`, `end-arrowhead`, or `arrowhead` is a source-positioned
validation error. Explicit `none` is accepted. Explicit `stroke-width`, color,
and stroke style are preserved for every kind; non-route arrowhead attributes
are also preserved.

For SVG, PPTX, PDF, and Excel Plan output, the render option `arrow-style`
supplies the global arrowhead (and, for `thin`/`standard`, width) only when the connection
does not explicitly set that semantic value. Explicit DSL or inherited group
values take precedence, and `kind="route"` remains headless. Excalidraw,
XYFlow, and Isoflow V1 output consume the resolved DSL scene rather than this
Plan-only option.

When a connection references endpoints in different frames, the shared scene
represents it as a page link instead of drawing one line across the inter-frame
gap. SVG, PPTX, PDF, Excel, and Excalidraw derive exactly two local stubs:

- the source stub runs from the source endpoint to the page-terminal inset line
  of its owning frame and has the exact label `to <destination frame ID>`; and
- the destination stub runs from the page-terminal inset line of its owning
  frame to the destination endpoint and has the exact label
  `from <source frame ID>`.

Angle brackets in those forms are rendered as literal punctuation. For a
connection from frame `overview` to frame `detail`, the visible strings are
therefore `to <detail>` and `from <overview>`.

Endpoint binding and logical frame-terminal geometry are separate. The
endpoint uses `src-anchor`/`dst-anchor`, then `src-side`/`dst-side`, then its
normal automatic binding. The logical page terminal uses
`src-frame-anchor`/`dst-frame-anchor`, then
`src-frame-side`/`dst-frame-side` as fixed choices. With neither frame-terminal
attribute, the legacy endpoint anchor, endpoint side, or normal nearest-border
result is only the preferred page side. The renderer keeps it when safe;
otherwise it chooses the nearest safe side from the endpoint's rendered visual
envelope. An item's envelope is the union of its icon and external label; other
endpoints use their rendered shape. Distance ties prefer a tied side facing the
remote frame, then `top`, `right`, `bottom`, `left`.

A side is safe when the resolved inset fits its normal frame dimension, it is
not the metadata edge, and an actual `top`/`bottom` terminal opposite metadata
does not enter the reservation strip. Validation of an automatic page terminal
checks only that this candidate set is non-empty. It must not infer the chosen
automatic side from layout `Box` geometry; final selection belongs to shared
scene construction after icon and label geometry is available. If the normal
preferred side is unsafe, scene construction remaps it to the nearest safe
candidate. No safe candidate is a source-positioned validation error at the
connection.

The endpoint- and frame-terminal-adjacent segments are perpendicular to their
own selected sides, so an endpoint may leave on `right` while its local stub
terminates at the page's bottom inset line. Frame-side and frame-anchor
attributes are valid only when the resolved endpoints belong to different
frames. Using any of them on a same-frame connection is a source-positioned
validation error.

Frame metadata reservation is a final safety constraint on that choice. For an
automatic page terminal, the metadata edge and any other unsafe side are
removed before the renderer's nearest-side choice. A terminal on a safe left
or right edge is clamped along that edge so it lies outside the top/bottom
reservation strip; any resulting coordinate difference is bridged
orthogonally. An explicit frame side or anchor that selects the reserved edge,
or an exact left/right anchor whose point lies inside the strip, is a
source-positioned validation error instead of being moved. Page-link paths and
labels remain outside the full strip.

When an explicit frame side is vertically opposite the metadata edge, its
actual terminal must remain outside the reservation strip. For bottom metadata
with explicit side `top`, the actual top terminal may not enter below the
strip's top boundary. For top metadata with explicit side `bottom`, the actual
bottom terminal may not enter above the strip's bottom boundary. A violation is
a source-positioned validation error at the connection. For an automatic page
terminal, the same conflict makes that candidate unsafe instead of immediately
rejecting the connection. A safe explicit `left` or `right` terminal remains
valid even if a hypothetical top/bottom inset line would enter the strip.

The parallel coordinate is resolved against the selected outer logical frame
edge before applying the normal inset. An explicit frame anchor keeps its exact
10/30/50/70/90-percent coordinate along the outer frame extent. An automatic
terminal's unconstrained parallel coordinate comes from the endpoint binding.
If it enters a 24-layout-px corner gutter, the parallel coordinate is clamped
and a two-bend orthogonal dogleg bridges the difference; a border shorter than
96 layout pixels uses one quarter of its length as an adaptive gutter. A
left/right terminal is also subject to the metadata reservation clamp described
above. Automatic left/right coincidence avoidance normally intersects that
corner-gutter range with an 8-layout-pixel clearance from the reservation. If a
very small non-reserved range cannot satisfy both preferences, it falls back to
the entire non-reserved interval, may touch its boundary, and never moves a
point outside the frame or inside the metadata strip.

The drawable terminal then lies on a page-terminal inset line parallel to that
outer edge. Let `i` be the resolved metadata `row-gap` when the frame has
metadata, or 4 layout pixels when it does not. The same `i` applies to every
terminal side regardless of metadata `position`; `i = 0` retains the outer
edge. An explicit `top`/`bottom` frame side requires `i < frame.height`; an
explicit `left`/`right` side requires `i < frame.width`. Failure is a
source-positioned validation error at that connection. For an automatic page
terminal, those inequalities classify candidates instead; only an empty safe
candidate set is an error. The resolved `i` is used exactly and is not reduced
to fit. With the resolved parallel coordinate represented by `u` for a
horizontal side or `v` for a vertical side, the terminal is:

```text
top:    (u, frame.y + i)
right:  (frame.x + frame.width - i, v)
bottom: (u, frame.y + frame.height - i)
left:   (frame.x + i, v)
```

The inset step changes only the normal coordinate. An explicit frame anchor
therefore retains its tangent slot and uses its local orthogonal stub for
visible separation. If an unconstrained final inset terminal would coincide
with the endpoint binding, its parallel coordinate moves by up to 24 layout
pixels within the available range so the stub remains visible. Manual bends do
not alter either local stub's geometry; bends remain logical routing metadata
for graph adapters.

There is one strict zero-inset case. When metadata is enabled with resolved
`row-gap="0"`, an endpoint resolves to its owning frame itself, and its explicit
frame anchor coincides with the resolved endpoint point, the connection is a
source-positioned validation error. An explicit endpoint anchor supplies that
point directly. An explicit endpoint side uses its center (`top` is `top-3`,
and likewise for the other sides); with neither endpoint attribute, the
automatically resolved endpoint side also uses its center. Fixed parallel
coordinates, perpendicular segments at both ends, and a visible local stub
cannot all be satisfied at that coincident point. The author must select a
different endpoint or frame anchor, or use a positive metadata `row-gap`.

Each `to <...>` / `from <...>` label is placed from the final inset terminal
with a 4-layout-pixel inward gap and a minimum 4-layout-pixel tangent gap.
Placement chooses the closest tangent position that avoids the endpoint
envelope and metadata reservation. Tiny pages clamp or shrink the label
fallback instead of moving it farther inward from that terminal.

Both scene stubs carry the same logical connector ID, original endpoint/frame
IDs, and V1 routing metadata. XYFlow and Isoflow use those fields to emit one
logical edge instead of two partial edges.

Default page-oriented export projects only the local stub belonging to each
frame: the source SVG/slide/page/worksheet contains `to <destination frame
ID>`, and the destination one contains `from <source frame ID>`.
`--combine-frames` places both local stubs on the compatibility canvas but
never reconnects them across the frame gap. Excalidraw also retains both stubs
in its one editable scene.

Output formats are projections of this resolved V1 meaning. A target schema
may not have fields for every V1 connector value; the upstream-compatible
Isoflow connector schema, for example, has no arbitrary metadata field. Such
adapters must use native constructs where available and must not add private,
schema-breaking fields. A V2 compatibility frontend consumes V1 directly and
must never use an output format as an intermediate representation.

When `src-side`, `dst-side`, `src-anchor`, and `dst-anchor` are omitted,
endpoint sides and anchor positions are calculated automatically from endpoint
geometry. Use `src-anchor` and `dst-anchor` to pin an endpoint to a specific
perimeter anchor. Cross-frame `src-frame-anchor` and `dst-frame-anchor` use the
same grammar to select the logical page side and tangent slot independently.
Each side has five positions at 10, 30, 50, 70, and 90 percent of the outer
frame extent, giving 20 unique tangent anchors. The drawable frame terminal
then moves only in the inward normal direction to the page-terminal inset line.
Corner anchors are not shared: `top-1` keeps a horizontal coordinate near the
left corner, while `left-1` keeps a vertical coordinate near the top corner.

```text
top:    top-1    top-2    top-3    top-4    top-5
right:  right-1  right-2  right-3  right-4  right-5
bottom: bottom-1 bottom-2 bottom-3 bottom-4 bottom-5
left:   left-1   left-2   left-3   left-4   left-5
```

Position numbers run left-to-right on `top` and `bottom`, and top-to-bottom on
`left` and `right`. Anchor positions are `1` through `5` from top/left to
bottom/right on the named side, inset from corners so each side owns its five
positions.
The aliases map one-to-one as `start=1`, `near=2`, `center=3`, `far=4`, and
`end=5`.

```xml
<connection src="web" dst="app"
            src-anchor="right-3"
            dst-anchor="left-3" />
<connection src="web" dst="app"
            src-side="right" src-anchor="3"
            dst-side="left" dst-anchor="3" />

<!-- The item and logical page terminal may use different sides. -->
<connection src="web" dst="detail.app"
            src-side="right" src-anchor="near"
            src-frame-side="bottom" src-frame-anchor="far"
            dst-side="left" dst-anchor="far"
            dst-frame-side="top" dst-frame-anchor="near" />
```

`src` and `dst` can also be expressed as child tags when the endpoint reference
and anchor should be declared together. The endpoint token can be tag text or
one of `id`, `ref`, `name`, or `target`.

```xml
<connection kind="traffic">
  <src anchor="right-3" frame-side="bottom" frame-anchor="far">web</src>
  <dst side="left" anchor="5" frame-anchor="top-2" ref="detail.app" />
</connection>
```

On `<src>` and `<dst>`, `frame-side` and `frame-anchor` map to the corresponding
source/destination connection attributes. A complete anchor such as
`bottom-4` supplies its side. With a separate side, slots accept `1..5` or the
aliases `start`, `near`, `center`, `far`, and `end`. Conflicting side and
complete-anchor values are validation errors for both endpoint and frame
anchors.

Excalidraw output always serializes arrowhead sizes as the smallest supported
size (`"s"`) to keep dense diagrams readable. The logical arrowhead type and
style metadata are still stored on the connector and used by SVG/PPTX export
and the SVG-based PDF/Excel projections.

Manual bend coordinates are expressed as child tags in the same Cartesian
layout coordinate space as the frame, with the origin at the upper-left of the
rendered frame and positive `x`/`y` extending right/down. SVG and PPTX route
calculations keep the connector orthogonal while forcing the route through each
listed bend in order. Excalidraw output stores the routing metadata on the
arrow; Excalidraw's own editor may still display its editable elbow connector
approximation.

```xml
<connection src="web" dst="db"
            scale="1" grid="8">
  <bend x="120" y="80" />
  <bend x="120" y="220" />
  <bend x="300" y="220" />
</connection>
```

`<point>`, `<via>`, and `<waypoint>` are accepted aliases for `<bend>`.
Coordinates can also be grouped inside `<bends>`, `<points>`, or `<path>`.

Items and group tags may define a connection reference with `id`, `name`, or
`ref`:

```xml
<item id="1178" name="web" />
<item id="1189" name="db" />
<vpc id="prod-vpc" />
web --- db
web ==> db
prod-vpc --- web
```

- `---` expands to `kind="route"`.
- `==>` expands to `kind="traffic"`.
- Operands may also be numeric item IDs or group IDs.
- Explicit `<connection src=... dst=...>` attributes resolve the same way.
- Shorthands must be direct text children of `<frame>`.
- References must be unique and must belong to an item or group with a
  non-empty ID.
- Use an explicit `<connection>` for color, width, or stroke overrides, and for
  arrowhead overrides on normal connections or traffic flows. Routes remain
  headless.

**Arrow spec:**
- `elbowed: true` — always right-angle connectors (Excalidraw "elbow connector")
- Arrowhead at end only by default. Excalidraw stores this as
  `endArrowhead: "arrow"` plus `endArrowheadSize: "s"`; xaligo metadata records
  the logical PPTX/SVG head as `stealth`.
- Stroke color `#1e1e1e`, stroke width `1px` for normal connections
- `kind="route"` defaults to `#64748b`, `1px`, lower route layer, no arrowheads
- `kind="traffic"` defaults to `#2563eb`, `1px`, higher traffic layer, directional end arrowhead
- A traffic line with the same endpoints as a route line is drawn beside that
  route in Excalidraw, SVG, and PPTX draw paths when possible.
- Start/end connect to the **edge midpoint** of the element
  - When direction is **downward**: label text element (`{id}-item-lbl`) bottom edge
  - Otherwise: icon image element (`{id}-item`) corresponding edge
- Edges are fixed with normalized coordinates via `fixedPoint`, so arrows snap correctly when the file is opened
- Arrow ID format: `conn-{src}-{dst}-{index}`
- Arrow ID is registered in `boundElements` of the bound elements
- Excalidraw item icons and labels are grouped with a 5x5 white anchor grid.
  Anchor grid cells are drawn above connectors and below the item content so
  lines do not cover icons/labels while endpoints remain visible.
- Excalidraw routing uses previously placed lines to offset exact or near-exact
  lane overlaps. Group header tags, item icons, and labels are treated as
  routing obstacles where possible.
- SVG/PPTX routing may additionally add automatic junction markers and line
  jump masks after the Excalidraw scene is built. PDF and Excel inherit the SVG
  projection. These are export-layer
  rendering features, not extra `.xal` tags.

**Edge selection logic:**

| Direction (dst relative to src) | Start edge | End edge |
|---|---|---|
| Right (dx ≥ dy) | right | left |
| Left | left | right |
| Down (dy > dx) | bottom (label) | top |
| Up | top | bottom (label) |

> If `src` / `dst` items are not rendered, a warning is emitted and the connection is skipped.

Connection endpoints must resolve to exactly one item, AWS group, rectangle,
port, or identified child frame. Numeric catalog IDs are valid only when that
ID appears once in the document; when the same service appears multiple times,
use unique `name` or `ref` values. Missing endpoints, ambiguous numeric IDs,
duplicate aliases, and `<connection>` tags nested below any tag other than
`<frame>` or its direct `<connections>` child are validation errors.

## AWS Group Tags

Like `container`, these stack children vertically, but are rendered with **AWS architecture diagram group border styles**.
Templates are in `etc/resources/aws/templates/excalidraw/` (`.excalidraw`) and `etc/resources/aws/templates/xal/` (`.xal`).
Icon SVGs are sourced from `etc/resources/aws/svg/Architecture-Group-Icons/`.

```xml
<aws-cloud id="production" title="Production Environment">
  <vpc id="vpc-main" title="vpc-0a1b2c3d">
    <private-subnet id="private-a" title="Private Subnet A">
      <card title="App Server" />
    </private-subnet>
  </vpc>
</aws-cloud>
```

| Tag | Display Name | Border Color | Style | Icon |
|---|---|---|---|---|
| `<aws-cloud>` | AWS Cloud | `#000000` | solid | AWS-Cloud-logo_32.svg |
| `<aws-cloud-alt>` | AWS Cloud | `#000000` | solid | AWS-Cloud_32.svg |
| `<region>` | Region | `#00A1C9` | dashed | Region_32.svg |
| `<availability-zone>` | Availability Zone | `#00A1C9` | dashed | — |
| `<security-group>` | Security group | `#CC0000` | dashed | — |
| `<auto-scaling-group>` | Auto Scaling group | `#E7601B` | dashed | Auto-Scaling-group_32.svg |
| `<vpc>` | Virtual private cloud (VPC) | `#8C4FFF` | solid | Virtual-private-cloud-VPC_32.svg |
| `<private-subnet>` | Private subnet | `#00A1C9` | solid | Private-subnet_32.svg |
| `<public-subnet>` | Public subnet | `#3F8624` | solid | Public-subnet_32.svg |
| `<server-contents>` | Server contents | `#7A7C7F` | solid | Server-contents_32.svg |
| `<corporate-data-center>` | Corporate data center | `#7A7C7F` | solid | Corporate-data-center_32.svg |
| `<ec2-instance-contents>` | EC2 instance contents | `#E7601B` | solid | EC2-instance-contents_32.svg |
| `<spot-fleet>` | Spot Fleet | `#E7601B` | solid | Spot-Fleet_32.svg |
| `<aws-account>` | AWS account | `#E7008A` | solid | AWS-Account_32.svg |
| `<aws-iot-greengrass-deployment>` | AWS IoT Greengrass Deployment | `#3F8624` | solid | AWS-IoT-Greengrass-Deployment_32.svg |
| `<aws-iot-greengrass>` | AWS IoT Greengrass | `#3F8624` | solid | — |
| `<elastic-beanstalk-container>` | Elastic Beanstalk container | `#E7601B` | solid | — |
| `<aws-step-functions-workflow>` | AWS Step Functions workflow | `#E7008A` | solid | — |
| `<generic-group>` | Generic group | `#AAB7B8` | dashed | Configurable with `icon-id` |
| `<capture>` | Capture | `#F5A623` | dashed | — |

All AWS group tags require a non-empty `id`. IDs for group tags, `<rectangle>`,
and `<port>` must be unique among frame-like components. Group tags otherwise
accept the same attributes as `container` (`title`, `class`, `gap`, etc.).

`<capture>` is a lightweight structural annotation container rather than an
AWS/architectural boundary. It participates in normal nested layout: its
children are allocated within its bordered content box, including the same
padding and optional title band used by other group tags. The border uses the
same title/text/tag-name fallback as every other group tag without implying
cloud/network semantics. Like every group tag, a `<capture>` is connectable by
`id`/`name`/`ref` from `<connection>`, including the `frameId.id` qualified
form, so a connection that starts or ends on a `<capture>` in another frame
renders as the same "to `<frame>`" / "from `<frame>`" cross-frame page-link
stubs used for any other connectable endpoint — no separate cross-boundary
arrow mechanism exists for captures.

`generic-group` additionally accepts `icon-id`, a positive signed 32-bit
decimal ID (`1..2147483647`) from `service-catalog.csv`. Zero, signs,
non-decimal syntax, and out-of-range values are invalid. It uses the same
embedded AWS, Tabler, and Yamaha icon catalog as `<item>` and renders a 32px
icon to the left of the title.
This matches the built-in group icon size. Every group header receives an
opaque mask matching its local background behind the icon and label, preventing
solid or dashed border strokes from crossing the header content.
Group header tag labels use the shared single-line text policy. The tag
background and label box use a conservative width estimate so no-wrap text
remains inside the tag in SVG and PowerPoint. Keep group tag text concise; if
changing group tag font, padding, or geometry, update the shared text-layout
policy, renderer width estimate, and regression tests together.
East Asian full-width characters, including Japanese labels, count as
double-width in group header and item label width estimates.

```xml
<generic-group id="network-topology" title="Network Topology" icon-id="104635">
  <item id="200036" />
</generic-group>
```

### Layout Control Attributes (shared by all containers)

Available on `frame` / `container` / `col`, all AWS group tags, and unknown
child-bearing container tags where noted.

| Attribute | Value | Description |
|---|---|---|
| `layout` | `"horizontal"` | Arrange children **horizontally** with proportional widths (use the `col` attribute for ratio) |
| `layout` | `"staggered"` | Stack children with a depth offset (AWS group tags and unknown child-bearing containers) |
| `gap` | float | Child spacing (px). Default `16` |
| `align` | `"{vertical}-{horizontal}"` | Position of content area and `<item>` icons. Item grids also support `spread`. Default item-grid alignment is `"middle-center"` |
| `content-width` / `content-height` | float | Shrink usable inner layout area, leaving whitespace |
| `width` / `height` | float | Fixed child size (root frame dimensions remain the paper/content frame) |
| `overflow` | `"error"` \| `"visible"` | Child containment policy. Default `error` |

**`align` values** — combine a vertical part and a horizontal part with `-`:

| Part | Values |
|---|---|
| vertical | `top` \| `middle` \| `bottom` |
| horizontal | `left` \| `center` \| `right` \| `spread` |

All 12 combinations are valid: `top-left`, `top-center`, `top-right`, `top-spread`, `middle-left`, `middle-center`, `middle-right`, `middle-spread`, `bottom-left`, `bottom-center`, `bottom-right`, `bottom-spread`.

> **`center` (default):** icons are packed together and the group is centred within the available area
> (equivalent to CSS `justify-content: center`).
>
> **`spread`:** icons are distributed with equal gaps between each icon and the container edges
> (equivalent to CSS `justify-content: space-evenly`).
>
> **`left` / `right`:** icons are packed at the respective edge with a fixed `8 px` gap between icons.

```xml
<!-- Icons centred vertically and horizontally inside the group (default) -->
<private-subnet id="app-tier" title="App Tier" align="middle-center">
  <item id="27" />
  <item id="547" />
</private-subnet>

<!-- Icons spread evenly across the full width -->
<generic-group id="global-services" title="Global Services" align="middle-spread">
  <item id="1179" />
  <item id="1178" />
  <item id="216" />
  <item id="227" />
</generic-group>

<!-- Icons pinned to the top-left -->
<generic-group id="security-services" title="Security" align="top-left">
  <item id="216" />
  <item id="227" />
</generic-group>
```

### Child Size Ratio Attributes

| Attribute | Direction | Description |
|---|---|---|
| `row` | Vertical (`layoutStack`) | **Height ratio** among children without explicit `height`. Default `1.0` |
| `col` | Horizontal (`layout="horizontal"`) | **Width ratio** among children without explicit `width`. Default `1.0` |

```xml
<!-- Horizontal: left 2 : right 1 width ratio -->
<vpc id="vpc-main" title="VPC" layout="horizontal">
  <public-subnet id="public-subnet" title="Public" col="2" />
  <private-subnet id="private-subnet" title="Private" col="1" />
</vpc>

<!-- Vertical: top 1 : bottom 2 height ratio -->
<region id="region-main" title="Region">
  <vpc id="vpc-a" title="VPC A" row="1" />
  <vpc id="vpc-b" title="VPC B" row="2" />
</region>
```

## Spacing Classes (`class` attribute)

Vuetify-style notation. **Unit: `spacingUnit = 8px`**.

### All-sides shorthand

| Class | Meaning |
|---|---|
| `pa-{n}` | padding all sides = n × 8px |
| `ma-{n}` | margin all sides = n × 8px |

### Axis shorthand

| Class | Meaning |
|---|---|
| `px-{n}` | padding left + right = n × 8px |
| `py-{n}` | padding top + bottom = n × 8px |
| `mx-{n}` | margin left + right = n × 8px |
| `my-{n}` | margin top + bottom = n × 8px |

### Per-side

| Class | Meaning |
|---|---|
| `pt-{n}` | padding-top |
| `pr-{n}` | padding-right |
| `pb-{n}` | padding-bottom |
| `pl-{n}` | padding-left |
| `mt-{n}` | margin-top |
| `mr-{n}` | margin-right |
| `mb-{n}` | margin-bottom |
| `ml-{n}` | margin-left |

Multiple classes are space-separated: `class="pa-4 mt-2"`

### Semantics

| Kind | Target tag | Behavior |
|---|---|---|
| `padding` | frame / container / col | Inner whitespace. Child layout starts `pad` pixels inward |
| `padding` | AWS group tags / unknown containers | **Added to** `defaultGroupTopInset(44)` / `defaultGroupSideInset(12)`. `pa-2` adds +16px below the header |
| `margin` | any child element | Read by the parent layout (`layoutStack` / `layoutRow`) and used as inter-sibling spacing (equivalent to CSS flex margin) |

## Layout Calculation Rules

1. Normalize and validate numeric attributes and enum values.
2. Resolve each parent's border box and content box after margin and padding.
3. `frame` / `container` / `col` → **vertical stack**: reserve fixed child
   heights, gaps, and margins, then divide the remainder by `row` weights.
4. `layout="horizontal"` → reserve fixed child widths, gaps, and margins, then
   divide the remainder by `col` weights.
5. `row` → **12-column grid** after validating each `span` and their total.
6. Leaf elements use the resolved `(x, y, w, h)` received from their parent;
   they do not replace the allocation after sibling placement.
7. Verify finite positive geometry and parent-content containment before scene
   construction. Respect only an explicit `overflow="visible"` exception.
8. Resolve item grids against the same occupied content area before encoding.

## Example

```xml
<frame width="1440" height="900" class="pa-4">
  <container class="pa-4">
    <row gap="20" class="mb-2">
      <col span="8" class="pa-2">
        <card title="Dashboard" />
      </col>
      <col span="4" class="pa-2">
        <card title="Summary" />
      </col>
    </row>

    <row gap="20">
      <col span="4" class="pa-2">
        <panel title="Filters" />
      </col>
      <col span="8" class="pa-2">
        <panel title="Main Chart" />
      </col>
    </row>
  </container>
</frame>
```

## Constraints and Notes

- The canonical root is `<xaligo version="1">`. Legacy `<frame>` and
  `<frames>` roots are accepted with a warning. Direct children of `<frames>`
  must be identified `<frame>` tags. V2 uses `<scene version="2">`, which is
  intentionally rejected by V1.
- Both self-closing (`<card title="..." />`) and regular (`<card title="..."></card>`) forms are supported.
- The sum of `span` values in direct children of `<row>` must not exceed 12.
  Excess is a validation error rather than implicit overflow to the right.
- `.xal` files must be saved in UTF-8.

---

# 08 Architecture


This document defines the implementation boundaries of xaligo. Product
direction lives in `roadmap.instructions.md`; DSL behavior lives in
`xal-spec.instructions.md`.

## Core pipeline

```text
.xal source
   -> internal/usecase orchestration
   -> internal/usecase/v1/engine parser functions
   -> validated numeric and enum attributes
   -> internal/usecase/v1/engine layout functions
   -> resolved canonical scene
   -> internal/usecase/v1/engine plan calculations
   -> internal/repository output encoder
   -> SVG | Excalidraw | PPTX | PDF | Excel | XYFlow | Isoflow
```

The parent `internal/usecase` package is the shared rendering and orchestration
boundary. Its `v1/engine` subpackage contains synchronous V1 calculation stages.
Format-rendering adapters (CLI, preview server, and WASM) call a
constructor-injected `RenderUsecase` instead of assembling a parallel
parser/layout/render pipeline. Utility
commands such as `generate xal` and `add service` may use their focused internal
builders and repositories directly.

## Language-version boundary

`<xaligo version="1">` selects canonical V1. Root `<frame>` and `<frames>` are
legacy V1 compatibility inputs and emit a migration warning. Native V2
uses the reject-safe `<scene version="2">` root. The V1 parser is not extended
to recognize `<scene>` and does not import or call V2 code.

The parent use-case boundary owns one lightweight root/version dispatch before
engine selection. It must inspect the first XML start element once, reject
contradictory root/version pairs, and pass the original bytes to exactly one
frontend. It must not select a version by retrying another parser after an
error.

V2 provides two frontends: its native `<scene version="2">` frontend and a V1
compatibility frontend that implements the frozen V1 behavior. Both lower
directly to one typed, version-neutral model consumed by V2 layout, routing,
and format encoders. The V1 compatibility path must not rewrite XML, parse a
document twice, serialize through an intermediate V1 scene, or invoke a full
V1 renderer and then reverse-engineer its output. This one-way relationship
allows V2 to render V1 while V1 remains unaware of V2.

## Package responsibilities

| Path | Responsibility |
|---|---|
| `internal/entity` | Independent entity layer containing cross-layer structures |
| `internal/usecase` | Render orchestration, context checks, repository port adaptation, and future parallel scheduling |
| `internal/usecase/v1/engine` | Synchronous V1 parser, validation, layout, scene, routing, and plan calculations; no repository or scheduling ownership |
| `internal/repository` | Filesystem, catalog, HTTP preview, and output-format encoding/export adapters |
| `internal/command.go` | Root Cobra command assembly |
| `internal/controller` | Cobra CLI argument and file-I/O adapters |
| `cmd/wasm` | JavaScript-global adapter over shared use cases and embedded assets |
| `external` | TypeScript external adapter layer mirroring `internal`: `command.ts`, `controller`, `entity`, `repository`, `usecase` |
| `test/unit` | Unit tests mirroring the source tree they cover |
| `test/integration` | Black-box tests of exported APIs and adapters |
| `etc/resources/aws` | Catalogs, templates, embedded assets, and attribution |

## Invariants

1. `.xal` is the only source DSL. Do not add adapter-specific parsers.
2. Mode selects visual semantics; format selects output serialization.
3. Format-rendering production paths call parser and layout through
   `internal/usecase`. Adapters use an injected `usecase.RenderUsecase`;
   controllers use separate narrow use cases for diagnostics, scene I/O,
   catalog access, and persisted export.
4. Routing and connector behavior belongs in shared scene/plan layers, not in
   individual output adapters.
5. Filesystem-less environments provide an `AssetSource`; they do not fork the
   render pipeline.
6. Native configuration remains the default when `RenderOptions.Assets` is nil.
7. New formats require a `Format` value, shared render function, CLI wiring,
   tests, and adapter documentation.
8. Errors are returned and wrapped with context. Core packages do not panic.
9. Native CLI dependency construction belongs in `NewRootCmd`; the WASM entry
   point is its own composition root. Controllers depend on use cases, never on
   other controllers.
10. Input/output destination dependencies belong in `internal/repository` and
    must not appear as use-case filenames.
11. Validation and rendering use the same parse, normalization, and geometry
    checks. An input accepted by `validate` must not fail during rendering
    because of malformed or non-finite geometry.
12. Resolved geometry contains only finite coordinates and strictly positive
    drawable sizes. A child stays inside its parent's content box unless the
    source explicitly selects a non-containing overflow policy.
13. Parent layout owns child allocation. Fixed main-axis sizes are reserved
    before flexible weights divide the remaining space; a child cannot silently
    replace its allocation after sibling positions have been calculated.
14. Text wrapping, fitting, clipping, line height, padding, and semantic role
    are part of the renderer-neutral draw contract. Encoders translate that
    contract and do not invent format-specific text behavior.
15. Geometry and typography use one effective layout-pixel transform. Changing
    PPI or paper fitting must not scale text independently from its shapes.
16. Item-grid placement is resolved before scene encoding and participates in
    the same occupancy and overflow checks as rectangles and other children.
17. Format dispatch has one use-case owner. Commands and controllers collect
    inputs and persist outputs; they do not maintain a second format switch or
    call an external output repository directly.
18. Shared types and functions use format-neutral names. Format-specific names
    may remain only as compatibility aliases at public boundaries. Neutral
    schemas must not require a renderer's JSON fields or generated-ID patterns;
    renaming an adapter-shaped type alone is not completion. Semantic element
    kind and parentage are explicit data; adapters must not reconstruct current
    scene hierarchy from rectangle containment.
19. Interfaces and constructors/factories live in the responsibility file that
    contains their concrete implementation. Do not create declaration-only
    `interface.go`, `interfaces.go`, `constructor.go`, `constructors.go`, or
    equivalent TypeScript files. Place an interface beside the concrete type
    and methods that implement it, and place `NewX`, `createX`, or another
    factory beside the type and behavior it constructs.
20. Layer components do not depend on peer components in the same layer. A
    repository must not construct, retain, or call another repository; a use
    case must not call another independently constructed use case; and a
    controller must not call another controller. Coordination between multiple
    repositories belongs to a use case, and coordination between multiple use
    cases belongs to a controller or composition/public-API boundary.
21. `internal/usecase/v1/engine` exposes synchronous calculation stages. It
    does not import `internal/repository`, interpret `context.Context`, create
    goroutines/channels/worker pools, or choose concurrency limits. The parent
    use case owns I/O, cancellation checks, job partitioning, result ordering,
    and any future parallel execution. Order-dependent routing within one plan
    remains sequential.
22. Language versions are selected by a root/version pair, never by parser
    fallback. `<frame>`/`<frames>` are V1 and `<scene version="2">` is V2, so a
    V1 reader rejects V2 before interpreting nested syntax.
23. V2 renders V1 through a frozen V1 compatibility frontend that lowers once
    into the typed neutral model. V1 has no V2 dependency, and neither XML
    rewriting, double parsing, nor renderer-output round-tripping is allowed.
24. Output schemas are capability projections of the shared semantic model.
    An encoder may omit a value its target schema cannot represent, but it must
    not invent private schema extensions or become an intermediate semantic
    model. Lossy capabilities are documented and tested explicitly.
25. Structural diff compares parsed `.xal` trees, never source lines or
    positional scene/plan IDs. The old side highlights removed and previous
    modified/moved nodes in pale red; the new side highlights added and current
    modified/moved nodes in pale green. Highlight overlays are added after
    layout and route resolution and must not become routing obstacles.
26. An identified child `<frame>` is one physical page by default. SVG emits
    one artifact per frame, PPTX maps one frame to one slide, PDF maps one
    frame to one page, and Excel maps one frame to one worksheet containing
    that frame's SVG image. `CombineFrames` is the explicit compatibility
    policy that preserves the former single-canvas, single-slide, single-page,
    or single-sheet result. Excalidraw, XYFlow, and Isoflow remain single
    logical documents and do not split by frame. Page-oriented encoders omit
    the page-frame outline in default and combined output; the frame remains a
    logical crop/page-link boundary rather than a visible rectangle.
    Excalidraw retains page-frame objects with transparent strokes.
    A default page-local SVG uses the exact logical frame rectangle as its
    canvas and clip boundary, without adding stroke/marker safety padding;
    PDF pages and Excel page images inherit that strict crop. Combined SVG
    compatibility output retains marker-safe canvas expansion.
27. Page projection happens only after the complete document scene, connector
    routing, and cross-frame link semantics are resolved. A per-frame encoder
    consumes an ordered `DocumentPlan` projection; it must not parse, lay out,
    route, or infer crop geometry independently. A one-frame SVG render returns
    the exact requested output path, while a multi-frame SVG render uses stable
    frame-derived artifact IDs and rejects filename collisions.
28. Native PDF and Excel encoders remain behind `!js` build constraints. The
    browser adapter uses lightweight `js` repository stubs because those
    formats are not exposed there; native canvas, font, and spreadsheet
    dependencies must not enter the browser-WASM dependency graph.
29. A frame metadata tag band is resolved once in the V1 shared layout and
    presentation scene as page-owned decoration. The resolved metadata
    `row-gap`, which defaults to 4 layout pixels, is both the inter-row spacing
    and the metadata page-edge inset. The selected top/bottom band edge and
    both horizontal row bounds are inset by that value, and every row wraps and
    aligns within `frame width - 2 * row-gap`; frame padding, content margins,
    and content-box offsets do not replace or add to this inset. The full-width
    reservation strip still runs from the outer logical frame edge to the final
    content-box boundary and is at least
    `row-gap + complete band height + 8` layout pixels deep.
    The inset is measured from the logical frame edge before any common PPTX
    slide centering; it is not an export `--paper-margin`. Normal items, text,
    local/UML connector paths and labels, and cross-frame page links cannot
    enter it. Legacy/automatic page-link side selection remaps a reserved edge
    to the nearest safe edge and clamps left/right terminals outside the strip.
    An explicit cross-frame `src-frame-side`, `dst-frame-side`,
    `src-frame-anchor`, or
    `dst-frame-anchor` that selects the reservation is instead a validation
    error. When metadata is enabled, its resolved `row-gap` is also the inward
    normal inset for page-link terminals on all four sides; without metadata,
    the terminal inset is 4 layout pixels. A resolved zero `row-gap` retains
    the outer logical frame edge. An explicit frame side/anchor requires the
    inset to fit that side's normal dimension and its actual terminal to avoid
    the reservation; failures are reported at the connection source position.
    Without an explicit frame terminal, validation only requires a non-empty
    set of sides satisfying those rules. Shared scene construction uses actual
    endpoint visual geometry to retain a safe preference or choose the nearest
    safe side; it does not use validator `Box` geometry to predict that side. A
    safe selected `left`/`right` terminal is not rejected for an unused
    top/bottom inset line. The inset is never implicitly clamped. Page-link
    labels stay adjacent to the final inset terminal with a 4-layout-pixel gap
    while avoiding metadata and endpoint geometry. These rules, text metrics,
    layer order, and per-page ownership are
    encoder-independent. SVG, PPTX, PDF, Excel, and Excalidraw consume that
    shared result; XYFlow and Isoflow may omit the decoration but must not
    reinterpret it as graph nodes or endpoints.
30. Cross-frame connector geometry distinguishes the item endpoint from the
    logical page terminal. `src-side`/`dst-side` and
    `src-anchor`/`dst-anchor` bind the endpoint; `src-frame-side`/
    `dst-frame-side` and `src-frame-anchor`/`dst-frame-anchor` select the
    owning frame side independently. The outer logical frame edge supplies the
    side and 10/30/50/70/90-percent tangent coordinate, but the drawable frame
    terminal is on the parallel page-terminal inset line. Applying the inset
    changes only the normal coordinate; an explicit frame anchor retains its
    tangent coordinate. The endpoint- and frame-terminal-adjacent route
    segments are perpendicular to their respective sides. Frame-terminal
    attributes are invalid on same-frame connections. At zero inset, an owning
    frame endpoint that coincides with an explicit frame anchor is a
    source-positioned validation error; explicit endpoint anchors keep their
    slot, while explicit endpoint sides and automatic endpoint sides resolve to
    their center slot for this check. Automatic left/right coincidence
    avoidance uses the corner gutter and metadata clearance when possible, but
    a tiny safe range falls back to the full non-reserved interval without
    leaving the frame.

## File organization

Files are divided by cohesive implementation responsibility, not by declaration
kind. A responsibility file may contain its private types, interface, concrete
implementation, constructor/factory, constants, and methods when those
declarations exist to support that implementation.

The package directory already identifies the architectural layer, so Go file
names never repeat `controller`, `usecase`, or `repository`. Use the component
responsibility as the filename prefix:

- `<component>.go` contains the component's public interface, constructor, and
  principal concrete implementation.
- `<component>_<detail>.go` contains a cohesive private implementation slice of
  the same component.
- The public interface is `<Component>Controller`, `<Component>Usecase`, or
  `<Component>Repository`, according to its package.
- The constructor is `New<Component>Controller`, `New<Component>Usecase`, or
  `New<Component>Repository` and returns that interface.
- The concrete implementation type is unexported.

Current component prefixes are `add`, `diff`, `generate`, `init`, `render`,
`serve`, `validate`, and `version` in `internal/controller`; `render`, `diff`, `diagnostics`,
`scene_io`, `catalog`, `export`, `parser`, `layout`, `element`, `pagination`,
`plan`, `scene`, and `theme` in `internal/usecase`; and `powerpoint`, `preview`,
`isoflow`, `svg`, `pdf`, `spreadsheet`, `xyflow`, `excalidraw`, and `xaligo` in
`internal/repository`. Repository supporting files retain the same prefix, such
as `powerpoint_export.go` and `isoflow_assets.go`. Every direct
`internal/usecase/*.go` file is a complete component as specified in
`coding.instructions.md`.

Calculation files in `internal/usecase/v1/engine` use functional prefixes
such as `parse_*`, `layout_*`, `scene_*`, `route_*`, and `plan_*`. They contain
cohesive algorithm slices and do not repeat the package or architectural layer
name in filenames.

- Keep a Go interface in the file containing the corresponding concrete
  implementation and its principal methods.
- Keep a Go constructor in the file containing the concrete type it returns or
  initializes.
- Keep a TypeScript interface and factory with the implementation that consumes
  or realizes that contract when the interface is implementation-specific.
- When several implementations satisfy one interface, keep the interface with
  the package's primary responsibility/implementation and keep each additional
  implementation with its own methods; do not introduce a declaration-only
  file merely to appear neutral.
- Cross-layer entity DTOs and renderer-neutral value contracts remain in
  `internal/entity` or `external/entity`; this rule does not move shared data
  models into implementation packages.
- File splitting must move complete responsibility slices. Do not split an
  interface, its constructor, and its concrete behavior into separate files.
- Place the interface, unexported concrete type, and constructor at the start of
  their responsibility file, after imports and any shared constants/log codes,
  and before implementation methods. Do not recreate a package-wide facade that
  lists methods implemented by unrelated responsibility components.
- Private files and functions may divide one concrete component's implementation
  (for example PPTX image, legend, and package helpers). Such helpers must not
  expose a peer-layer interface or constructor and do not constitute another
  repository/use-case/controller component.

## Geometry contract

Layout is a constraint-resolution stage, not a best-effort drawing stage. It
must establish these postconditions before a scene or plan is constructed:

- every coordinate, length, weight, gap, margin, padding, and scale is finite;
- drawable width and height are greater than zero;
- row and column weights are greater than zero, and grid spans are in range;
- each content box is derived once from its allocated border box;
- fixed-size children consume space before flexible children are distributed;
- gaps are subtracted exactly once and cursors advance by the resolved size;
- containment or the selected overflow policy is recorded explicitly; and
- invalid geometry is returned as a source-positioned diagnostic, not dropped
  later by scene construction or exposed to an output encoder.

With `overflow="visible"`, fixed children still consume their resolved sizes
and advance the cursor. If they leave no positive remainder while flexible
children exist, the parent's original usable main-axis extent becomes the flex
pool. Children remain in source order, and every sibling cursor advances by the
resolved size plus its declared gap and margins, so the resulting overflow is
explicit. The default `overflow="error"` rejects the same input.

`Validate` and `Render` must both call this same stage. Encoders may reject an
I/O or serialization failure, but they must never be the first component to
discover `NaN`, `Inf`, a negative drawable size, or an impossible grid ratio.

## Renderer-neutral text contract

Every text draw operation carries its resolved box plus a text-layout policy:

```text
wrap | no-wrap
fit: none | shrink
text overflow: visible | clip
line height
content padding
semantic role
```

The semantic role distinguishes ordinary labels, item labels, group headers,
ports, connector labels, and other future text without requiring encoders to
infer behavior from element IDs. Glyph overflow must either be included in
bounds/obstacle calculations or removed by the declared fit/clip policy.

Layout and canonical-scene values are expressed in layout pixels. The current
shared presentation plan stores geometry and padding in inches and font sizes in
points. For effective PPI `p`, conversion is `inch = px / p` and
`pt = px * 72 / p`; paper fitting changes `p` once and both formulas use that
same value. Fixed physical sizes, such as an explicitly specified PPTX label
size, must be represented as an intentional semantic policy rather than an
incidental conversion constant.

## Dependency direction

```text
internal command / controller / cmd/wasm
                  |
                  v
   internal/usecase orchestration
          /                 \
         v                   v
internal/usecase/       internal/repository
  v1/engine               interfaces and
     |                   implementations
     v                         |
 internal/entity <-------------+

external/command.ts
        |
        v
external/controller
        |
        v
external/usecase
        |
        v
external/repository
```

Entity and use-case packages must not depend on CLI, preview, WASM, or
TypeScript adapters. Encoders consume entity structures and must not depend on
use-case implementations merely to access types.

## Verification

Run after structural changes:

```bash
go test ./...
go build ./...
npm ci --ignore-scripts
npm run build --workspace=@xaligo/xaligo-external
npm --prefix external run build:pptx-exporter-wasm
```

Generated binaries, `node_modules`, `output`, and package `dist` directories are
ignored and must not be committed.

---

# 09 Coding


Read this file before planning, editing, reviewing, generating, or renaming Go
or TypeScript code. These rules are merge preconditions, not optional style
preferences.

## Responsibility-based files

- Organize files by cohesive implementation responsibility, not declaration
  kind.
- Use `<component>.go` for a layer component's principal implementation and
  `<component>_<detail>.go` for a cohesive implementation slice.
- Do not repeat `controller`, `usecase`, or `repository` in a filename; the
  package already identifies the layer.
- Keep interfaces, concrete implementations, constructors/factories, and their
  principal methods together. Do not create declaration-only interface or
  constructor files.
- Interface names are `<Component>Controller`, `<Component>Usecase`, or
  `<Component>Repository`. Constructors are `New<Component>Controller`,
  `New<Component>Usecase`, or `New<Component>Repository`.

### `internal/usecase` root contract

Every non-test Go file directly below `internal/usecase` is one complete
use-case component. Its declarations follow this order:

1. `XxxUsecase` interface;
2. unexported `xxxUsecase` concrete type;
3. `NewXxxUsecase` constructor returning the interface; and
4. receiver methods containing that component's orchestration.

Repository dependencies are constructor-injected fields on the concrete type.
Do not leave declaration-free algorithm/wrapper/helper files such as a file
containing only package functions. Calculation helpers belong in the versioned
engine. A shared-source compatibility function may follow the receiver methods
only when it is a thin, deprecated delegation with no independent logic.

Private orchestration helpers for a component stay in that component's file;
do not create `render_options.go`, `render_scene.go`, or another root-level
implementation fragment that lacks its own interface, concrete type, and
constructor.

## V1 engine identifiers

All package-scope identifiers declared below `internal/usecase/v1/engine` must
carry a suffix that identifies both the engine version and their responsibility
file:

```text
<base identifier>V1Engine<FileBaseCamelCase>
```

Examples:

```go
ParseV1EngineParseDocument
routeOneV1EngineRoutePath
SceneDependenciesV1EngineSceneTypes
defaultPxPerInchV1EnginePlanBuild
```

The rule applies to:

- exported and unexported functions;
- methods;
- named types;
- package-level constants; and
- package-level variables.

Derive `<FileBaseCamelCase>` from the filename without `.go`. For example,
`parse_document.go` becomes `ParseDocument`, and `plan_connector_style.go`
becomes `PlanConnectorStyle`. When moving a declaration to another file, update
its suffix and every reference in the same change.

Do not add this suffix to local variables, parameters, result names, struct
fields, imported identifiers, package names, or Go's special `init` function.
Stable public compatibility wrappers in the parent `internal/usecase` package
retain their existing names and delegate to the versioned engine names.

## Engine execution boundary

- `internal/usecase/v1/engine` contains synchronous calculation stages and
  explicitly supplied synchronous dependency ports.
- It must not import concrete repositories, interpret `context.Context`, start
  goroutines, own channels or worker pools, or select concurrency limits.
- The parent `internal/usecase` package owns repository adaptation, I/O,
  cancellation checks, stage ordering, job partitioning, result ordering, and
  future parallel-process control.
- Order-dependent routing within one document or plan remains sequential.

## Dependency direction

- A repository must not construct, retain, or call another repository.
- A use case must not call another independently constructed use case.
- A controller must not call another controller.
- Multi-repository coordination belongs to a use case. Multi-use-case
  coordination belongs to a controller or composition boundary.

## Verification

After structural or naming changes, run at minimum:

```bash
gofmt -w <changed-go-files>
go test ./...
go build ./...
git diff --check
```

The V1 engine naming regression test must pass; do not bypass it with aliases
that leave nonconforming package-scope declarations in the engine package.

---

# 10 PPTX and routing


This file is the current source of truth for PPTX export geometry.

## Brainstorm Reference

- ChatGPT share: https://chatgpt.com/share/6a35c5b9-4528-83e8-aff9-bc37907a4d80
- The share page may not be accessible from automated tooling. Keep the concrete
  decisions below authoritative for implementation.

## Confirmed Decisions

- PPTX export is an A3-landscape-first workflow for the current AWS sample.
- The PPTX export implementation should be compiled to WASM and invoked from
  the Go repository layer.
- Do not use `goja` or V8 for PPTX export execution.
- Avoid a long-term Node.js subprocess dependency for repository-layer PPTX
  export. Node may remain a development/build tool only while the WASM exporter
  is being prepared.
- All PPTX geometry and routing decisions are computed by the Go use-case
  pipeline before the exporter boundary.
- Each identified child frame becomes one diagram slide in source order by
  default. `--combine-frames` is the explicit compatibility path for the
  former single-slide canvas.
- A presentation has one common slide size. Multi-frame PPTX uses the largest
  resolved page width and height and centers smaller frame pages without
  scaling them independently.
- The PPTX drawing/export layer must not make independent layout/routing
  decisions.
- Lines must not visually cover icons or labels.
- If any obstacle-free route exists, obstacle-hitting routes must be rejected.
- Item labels are 8pt at the default 96 PPI and scale with the same effective
  PPI/paper-fit transform as item icons.
- Item icons should remain visually consistent with their labels; avoid shrinking
  icons merely to satisfy a cramped row when layout whitespace controls can be used.
- Legend belongs on separate PPTX slide(s), not outside the diagram page.
- Legend slide layout is fixed to 4 columns and contains icon, abbreviation, and
  official service name.
- DSL must support empty grid cells and both inner/outer whitespace controls.

## Current Pipeline

```text
.xal DSL
  -> Go parse and numeric-domain validation (typed normalization is the target)
  -> resolved layout and canonical scene
  -> ordered page-oriented Go document plan (neutral-schema migration remains)
  -> internal repository encoder (SVG), or
  -> Go repository -> WASM command -> external controller -> use case -> repository
  -> SVG | .pptx
```

Geometry belongs on the Go side. The WASM export module should only translate
the resolved plan into PPTX bytes. Excalidraw-compatible JSON may be one scene
serialization, but it is not the target architecture name or ownership boundary
for the shared plan.

## Go / WASM Boundary

The adopted integration style is Go invoking a WASM-compiled PPTX exporter from
the repository layer.

Implementation preconditions:

- Go owns CLI/controller/repository orchestration.
- WASM must be called from `internal/repository/powerpoint.go`, not directly from
  controller or command packages.
- The exporter must be compiled to WASM before repository-layer execution.
- Go forwards user-facing PPTX options to the WASM exporter through a typed
  options structure or JSON bridge.
- The WASM exporter consumes the resolved shared Go plan and returns PPTX
  bytes or writes them through a repository-controlled output path.
- The WASM exporter must not perform independent geometry, layout, or routing.
- The external WASI command calls its controller, the controller calls the
  external use case, and only the use case calls the external PPTX repository.
  Command/controller code must not bypass this path.
- Go repository/controller code must not implement PPTX/OOXML drawing or zip
  writing directly. Keep Go as the adapter that builds the plan, invokes the
  WASM exporter, and persists the returned bytes.
- If existing TypeScript/PptxGenJS code cannot be compiled into a practical WASM
  exporter, replace that drawing layer with a WASM-compatible PPTX writer rather
  than introducing `goja` or V8.

Other integration styles are not the current implementation target:

| Style | Status |
|---|---|
| stdin/stdout JSON-RPC | Candidate for long-running/high-volume workflows |
| HTTP API | Candidate for service/BFF separation |
| gRPC | Candidate for high-performance typed service boundaries |
| Node.js subprocess | Temporary fallback only; not the target architecture |
| Embedded JS engine (`goja`, V8) | Not a target for PPTX export |

Do not spend implementation time replacing the repository-layer exporter with
`goja` or V8 unless that architecture is explicitly re-approved.

## Ownership

| Area | Owner |
|---|---|
| DSL parse/layout | `internal/usecase/v1/engine/parse_*`, `internal/usecase/v1/engine/layout_*` |
| Canonical scene and item metadata | `internal/usecase/v1/engine/scene_*` |
| Plan geometry, text layout, paper scaling, routing, legend data | `internal/usecase/v1/engine/plan_*`, `internal/usecase/v1/engine/route_*` |
| WASM exporter invocation from Go | `internal/repository/powerpoint.go` |
| WASM-compatible PPTX drawing/export | `external` TypeScript package and implementation |
| PPTX WASI command entry | `external/command.ts` |
| Public browser/JavaScript API bridge | `cmd/wasm/main.go` |

## Paper / Scaling

- PPTX export supports `--paper`, `--orientation`, and paper-margin fitting
  flags.
- A3 landscape is generated with:

```bash
.bin/xaligo render docs/src/examples/samples/sample.xal \
  --format pptx \
  --services docs/src/examples/samples/services.csv \
  -o out.pptx \
  --paper A3 \
  --orientation landscape \
  --paper-margin-top 0.75 \
  --paper-margin-bottom 0.75
```

- The shared Go plan resolves paper size and computes one layout-pixel transform.
- The page-oriented plan is built after the full scene and cross-frame page
  links are resolved. Its frame projections preserve source order.
- Shape coordinates, font sizes, strokes, padding, and routing geometry use
  that same transform. `--px-per-inch 144` must not scale text independently
  from its containing shape.
- `--paper-margin N` applies an inch-based margin to every side before fitting
  the diagram to the selected paper.
- `--paper-margin-top`, `--paper-margin-right`, `--paper-margin-bottom`, and
  `--paper-margin-left` override the all-side value for individual sides.
- Paper margins do not change the slide size; they reduce the available fit
  area and centre the diagram within that inset area.
- The `paper-frame` element remains the content frame for scaling.
- Root `<frame margin="N">` or `class="ma-N"` is content outer whitespace: it
  insets diagram content without shrinking the paper frame itself.

## Routing Rules

- Route calculation is in `internal/usecase/v1/engine/route_*`.
- Obstacles include image and text rectangles from the Excalidraw scene.
- Start/end rectangles are excluded from obstacle checks for that connection.
- Binding `gap` from Excalidraw arrows must be honored in PPTX routing.
- If any obstacle-free candidate exists, obstacle-hitting candidates must not be
  selected.
- Lines on an obstacle boundary count as collision.
- Existing routed paths are included in scoring so later lines avoid overlap and
  near-parallel crowding.
- Excalidraw output also feeds previously routed lines back into the shared
  router so matching X/Y lanes are offset before export.
- Visible container borders are reserved routing paths. Connectors may cross a
  frame boundary, but parallel paths prefer the configured line margin.
- Previously placed line lanes are used as candidate offsets, so `--arrow-margin`
  affects routes that would otherwise share the same position.
- Final PPTX drawing order is:
  1. anchor backgrounds and containers/shapes
  2. route lines, traffic lines, and line-jump masks
  3. automatic junction markers
  4. icons and labels

This order prevents lines from visually covering icons even at endpoints.

Excalidraw output mirrors this readability rule with editable JSON elements:
each item image and label is grouped with a small 5x5 white anchor grid. The
grid is drawn above connector lines and below the icon/label so labels remain
readable without hiding the connector endpoint. Excalidraw routing treats group
header tags, item icons, and labels as obstacles where possible, and serializes
arrowhead sizes as `"s"` for dense diagrams.

### Cross-frame page links

A connection between different frames is a page link in page-oriented output;
it is never one line crossing the inter-frame canvas. The shared scene emits
two axis-aligned local stubs for Excalidraw, SVG, PPTX, PDF, and Excel:

- source endpoint to the source frame's page-terminal inset line, with the
  exact label `to <destination frame ID>`; and
- destination frame's page-terminal inset line to the destination endpoint,
  with the exact label `from <source frame ID>`.

Angle brackets are literal punctuation, so a link from `overview` to `detail`
renders `to <detail>` and `from <overview>`. The shared scene, not the PPTX
exporter, selects endpoint binding and logical frame terminal geometry. The
endpoint uses `src-side`/`dst-side` and `src-anchor`/`dst-anchor`. A cross-frame
connection may independently select its logical terminal with
`src-frame-side`/`dst-frame-side` or the more specific
`src-frame-anchor`/`dst-frame-anchor`. Every side has five anchors at
10/30/50/70/90 percent along the outer frame extent. The endpoint- and
frame-terminal-adjacent route segments remain perpendicular to their
respective sides even when the two sides differ.

Frame-terminal precedence is explicit frame anchor, explicit frame side,
legacy endpoint anchor, endpoint side, then automatic nearest-side selection.
The first two are fixed choices. Without either one, the legacy/automatic result
is a preferred side: the renderer keeps it when safe or selects the nearest
safe side from the endpoint's rendered visual envelope when it is unsafe. Ties
prefer a tied side facing the remote frame, then `top`, `right`, `bottom`,
`left`. Validation checks that at least one safe candidate exists but does not
predict the automatic side from layout `Box` geometry. Frame-terminal
attributes are cross-frame-only; using them on a same-frame connection is a
validation error.

A frame metadata reservation strip is a final safety constraint. The visible
metadata rows are inset from their selected vertical edge and both horizontal
edges by the resolved `row-gap`, while the reservation itself remains
full-width from the outer logical frame edge to the content boundary. Without
an explicit frame-terminal attribute, the renderer filters unsafe sides before
its visual nearest-side choice; an unsafe normal preference is remapped to the
nearest safe candidate. A left/right terminal is clamped outside the full-width
strip before the orthogonal dogleg is built. An explicit frame side or anchor
that selects the metadata edge, or an exact left/right anchor inside the strip,
is a validation error instead of being moved. Neither the local path nor its
label may enter the reservation strip.
If an explicit frame side is vertically opposite the metadata edge, its actual
terminal must remain outside the same strip: `top` against bottom metadata, or
`bottom` against top metadata. Moving that explicit terminal into the strip is
a source-positioned validation error. For an automatic page terminal, the same
conflict excludes that candidate rather than immediately rejecting the
connection. A safe explicit `left` or `right` terminal is allowed even if the
unused top/bottom inset line would intersect the reservation.

An automatic terminal first uses the endpoint binding's coordinate parallel to
the outer logical frame edge. If that coordinate enters a 24-layout-px corner
gutter, it is clamped and a two-bend orthogonal dogleg bridges the difference.
A border shorter than 96 layout pixels uses one quarter of its length as an
adaptive gutter. An explicit frame anchor instead retains its exact
10/30/50/70/90-percent tangent slot.

After side, tangent, and reservation handling, the drawable page terminal is
shifted inward from the outer logical frame edge along that side's normal. A
frame with metadata uses the resolved metadata `row-gap` on all four terminal
sides, regardless of whether the band is at the top or bottom. A frame without
metadata uses 4 layout pixels. A metadata `row-gap` of zero leaves the terminal
on the outer edge. This step does not change the resolved tangent coordinate.
For an explicit frame side or anchor, the inset must be strictly smaller than
that side's normal frame dimension: height for `top`/`bottom`, width for
`left`/`right`. A reservation conflict on that actual side is likewise an
error. Both failures are source-positioned at the connection. For an automatic
page terminal, the same tests classify each candidate side; an unsafe preferred
side is remapped and only an empty safe candidate set is an error. The inset is
applied exactly, without an implicit clamp.

The segments at both the endpoint and page-terminal inset line remain
perpendicular to their selected side; an explicit frame anchor's orthogonal
local stub supplies visible separation. If an unconstrained final inset
terminal and the endpoint coincide, the terminal shifts by up to 24 layout
pixels along the parallel axis within the available gutter range so the line
remains visible. On a left/right side with metadata, that preferred range also
keeps 8 layout pixels from the reservation. If a very small safe region cannot
retain both the corner gutter and that clearance, the shift falls back to the
entire non-reserved interval, may touch its boundary, and never leaves the
frame or enters the strip.

If an owning frame has metadata with resolved `row-gap="0"`, the endpoint
resolves to that frame itself, and an explicit frame anchor coincides with the
resolved endpoint point, that connection is a source-positioned validation
error. An explicit endpoint anchor supplies its side and slot; an explicit
endpoint side or automatically resolved side uses the center slot (`side-3`).
The author must use a different endpoint/frame anchor or a positive `row-gap`;
xaligo must not move the fixed frame-anchor tangent coordinate or emit an
invisible zero-length stub.

Manual bends remain connector metadata and do not steer page-local stubs. Both
stubs retain one logical connector ID; XYFlow and Isoflow reconstruct one graph
edge from that metadata rather than exporting the two page projections.

The `to <...>` / `from <...>` label is placed from the final inset terminal
with a 4-layout-pixel inward gap and a minimum 4-layout-pixel tangent gap.
Candidate placement chooses the closest tangent position that avoids
the endpoint envelope and metadata reservation; tiny pages use a clamped
fallback rather than increasing the normal label distance.

The outer logical page edge and the parallel terminal inset line are geometric,
not visible rectangles: SVG, PPTX, PDF, and Excel omit page-frame outlines in
both default and combined output.

Default PPTX output places the source and destination stubs on their respective
frame slides. `--combine-frames` places both stubs on the compatibility slide
but never draws a replacement line across the frame gap.

## Advanced Routing Features

### Line Jumps

Excalidraw does not provide reliable built-in line jumps/bridges for this
workflow. The shared draw plan therefore implements them for SVG/PPTX.

Current approach:

- Detect line segment intersections after routing.
- Determine which line is visually above the other by layer/kind/order.
- Render jumps as a 6px background-colored mask below the upper line in
  SVG/PPTX output. The mask uses the uppermost opaque container background at
  the crossing. A curved arc may replace the rectangular mask later.
- For Excalidraw output, approximate with normal lines or supported shape
  primitives when necessary.

SVG preview and PPTX can support line jumps more accurately than Excalidraw JSON.

### Route / Traffic Separation

Network diagrams distinguish structural route lines from traffic-flow lines.

Implemented model:

| Kind | Meaning | Visual Direction |
|---|---|---|
| `route` | Physical/logical connection path | Thin, lower layer, no arrowheads, shortest orthogonal route |
| `traffic` | Communication flow over a route | Offset beside a matching route, higher layer, directional arrow/style |

Potential DSL forms:

```xml
<connection src="A" dst="B" kind="route" />
<connection src="A" dst="B" kind="traffic" />
```

or future shorthand:

```text
A -> B
A => B
```

Routing orders routes below normal connections and traffic. When a traffic line
shares the same endpoints as a route line, the traffic line follows a nearby
parallel lane instead of drawing directly on top of the route.

### Route Connectors

Frozen V1 routes are headless in every format. Their effective
`start-arrowhead` and `end-arrowhead` must both resolve to `none` after
`<connections>` defaults and child aliases are merged. A non-`none` value is a
validation error rather than a renderer-specific circular endpoint.

Small circular route connector nodes remain a future versioned feature; they
must use a renderer-neutral connector-node concept instead of overloading V1
arrowheads.

Conceptual shape:

```text
[EC2] -- o -------- o -- [RDS]
```

Future behavior may render explicit connector nodes in SVG/PPTX and equivalent
editable shapes in Excalidraw. It is not part of the V1 compatibility profile.

## Connector Style Options

`xaligo render --format pptx` forwards all PPTX routing options:

| Flag | Meaning |
|---|---|
| `--arrow-style` | `thin`, `standard`, `triangle`, `stealth`, `arrow`, `diamond`, `oval`, `none` |
| `--arrow-stub` | Pixel stub before the first/last bend |
| `--arrow-margin` | Pixel margin reserved around existing line lanes |
| `--px-per-inch` | Layout scaling base, default 96 |
| `--paper` | Named slide paper size: `A5`, `A4`, `A3`, `A2`, `A1`, `Letter`, `Legal`, `Tabloid` |
| `--orientation` | `portrait` or `landscape`; auto-fit when omitted |
| `--paper-margin` | Inch margin applied to all sides before paper fitting |
| `--paper-margin-top/right/bottom/left` | Inch margin override for one side |

`--arrow-style` is a Plan-level default. A connection's explicit or inherited
DSL arrowhead and stroke width take precedence; `kind="route"` remains
headless. The `thin` and `standard` presets may supply a default line width only
when the DSL did not supply `stroke-width` or its `width` alias.

Every numeric render option must be finite. `--px-per-inch`, arrow stub and
margin values, and paper margins reject negative values; the internal zero
value selects the documented default. Validation happens before scene/plan
construction so `NaN` or infinity cannot first fail during JSON encoding.
Paper size, orientation, and arrow style are closed enums. Paper margins require
a named paper size, and their effective left/right and top/bottom sums must
leave a strictly positive content area in the selected (or at least one
automatic) orientation.

## Group Header Tags

- Group header tag labels are single-line in every output whose text engine can
  represent that policy. The shared draw plan marks their semantic role,
  wrapping, fitting, clipping, line height, and padding; the TS drawing layer
  consumes those values rather than inferring behavior from an element ID.
- Excalidraw scene generation must reserve conservative tag label width before
  PPTX export. `groupLabelCharW` is intentionally larger than the average
  Excalidraw text metric so PowerPoint no-wrap text stays inside the tag
  background.
- When changing group tag font size, font family, padding, or tag geometry,
  update both the scene width estimate and the group-header regression tests.

## Item Labels

- Item icon size defaults to 32px in native CLI config.
- Item label font is 8pt at the default 96 PPI. At another effective PPI it is
  `10.666...px * 72 / effectivePPI` points so its ratio to the icon is stable.
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

Fixed children are reserved before flexible `row`/`col` allocation. The
resolved size advances the sibling cursor and must remain inside the parent's
content box unless the source explicitly uses `overflow="visible"`. Layout
overflow is diagnosed before plan construction; SVG or PPTX clipping is not a
substitute for a valid layout.

For item grids, horizontal `spread` is also supported.

## Legend Pages

PPTX export adds legend slides after all frame/diagram slides when `--services`
is provided.

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
npm run build --workspace @xaligo/xaligo-external
.bin/xaligo render docs/src/examples/samples/sample.xal --format pptx --services docs/src/examples/samples/services.csv -o out.pptx --paper A3 --orientation landscape --arrow-style thin
unzip -t out.pptx
```

For icon-overlap regressions, inspect the resolved PPTX XML and ensure routed
custom geometry does not intersect target icon/label rectangles.

---

# 11 Diagram creation


Standard workflow for creating Excalidraw and PPTX diagrams.

---

## Step 1 — Find Service IDs

`etc/resources/aws/service-index.csv` maps service IDs to service names.
Use `grep` to search for the services you need.

```bash
# Format: id,service
grep -i "ec2"          etc/resources/aws/service-index.csv
grep -i "rds\|aurora"  etc/resources/aws/service-index.csv
grep -i "cloudfront"   etc/resources/aws/service-index.csv
```

Example output:
```
27,Amazon EC2
117,Amazon RDS
1178,Amazon CloudFront
```

---

## Step 2 — Create services.csv

`services.csv` lists the services to include in the diagram.

**Format:** `id,OfficialName,Abbreviation,Summary,Usage,Notes`

- Column 1 (`id`) as a number → icon is fetched from service-catalog.csv.
- Lines starting with `#` are treated as comments and ignored.
- For `xaligo render --services`, every non-comment row must have a positive
  numeric `id` and a non-empty `OfficialName`; duplicate IDs are rejected before
  rendering.
- `Abbreviation`, when set, is used as the **icon label inside the diagram** and in the standalone legend icon below the frame.
  - Takes priority over the built-in abbreviation table in
    `internal/entity/service.go`.
  - When empty, the built-in table is used as fallback, then the official name.
- `OfficialName` is displayed as the full-name text in legends.

```csv
# 3-tier Architecture service list — IDs must match <item> tags in the .xal file
# Format: id,OfficialName,Abbreviation,Summary,Usage,Notes
1179,Amazon Route 53,R53,DNS web service,Domain name resolution and health checks,
1581,Amazon VPC Internet Gateway,IGW,Internet connectivity,Inbound/outbound internet traffic,
1182,Elastic Load Balancing,ELB,Load balancing service,Distribute traffic across EC2 instances,
27,Amazon EC2,EC2,Virtual server,Application tier,
1582,Amazon VPC NAT Gateway,NATGW,NAT gateway,Outbound internet for private subnets,
110,Amazon Aurora,Aurora,Relational database,High-performance managed DB,
113,Amazon ElastiCache,EC,In-memory caching,Session and query cache,
```

> **Note:** `render --format excalidraw` warns to stderr when an `<item id="N">` in the .xal
> is not listed in services.csv, or when a services.csv entry has no corresponding
> `<item>` in the diagram.  Keep both files in sync to suppress these warnings.

Reference: [docs/src/examples/samples/services.csv](../../docs/src/examples/samples/services.csv)

---

## Step 3 — Create a .xal file

Use `<item id="N" />` to place service icons in the layout.
`N` is the service ID from the first column of service-index.csv.

### Choosing the right group tag

Use AWS-specific group tags only when the content matches the tag's meaning.
For logical groupings that do not correspond to a specific AWS construct, use `<generic-group>`.

| Tag | When to use |
|---|---|
| `<public-subnet>` | Items that belong to a public (internet-routable) subnet |
| `<private-subnet>` | Items that belong to a private subnet |
| `<security-group>` | Resources sharing an EC2 security group |
| `<auto-scaling-group>` | An EC2 Auto Scaling group |
| `<generic-group>` | Any logical grouping that does not fit the above (security services, storage tiers, CI/CD, etc.) |
| `<capture>` | A border-only structural annotation group (e.g. highlighting a "hot path") that participates in normal nested layout without conveying AWS/architectural semantics |

> **Incorrect:** using `<public-subnet title="Security &amp; Identity">` for IAM / WAF — these are not subnet resources.
> **Correct:** use `<generic-group title="Security &amp; Identity">` instead.

### Service Scope Validation

Before finalizing the `.xal`, verify that each service is placed at the correct scope level.
Placing a global or regional service inside an `<availability-zone>` is misleading.

| Scope | Placement in .xal | Typical services |
|---|---|---|
| **Global** | Direct child of `<aws-cloud>`, inside `<generic-group>` | Route 53, CloudFront, IAM, WAF |
| **Regional** | Inside `<region>`, outside `<vpc>`, inside `<generic-group>` | Lambda, S3, CloudWatch, SQS, SNS, EventBridge, Step Functions, CodePipeline, Macie |
| **VPC-level** | Inside `<vpc>`, outside `<availability-zone>`, inside `<generic-group>` | Internet Gateway, ELB/ALB, Secrets Manager |
| **AZ-specific** | Inside `<availability-zone>`, in `<public-subnet>` / `<private-subnet>` | EC2, NAT Gateway, RDS instance, Aurora replica, ElastiCache node, ECS task, EKS node |

> **Incorrect:** placing Route 53 or IAM inside `<availability-zone>` — these services are not AZ-bound.
> **Correct:** group them under `<generic-group title="Global Services">` as a direct child of `<aws-cloud>`.

Quick checklist:
- [ ] Global services (Route 53, CloudFront, IAM, WAF) → outside `<region>`
- [ ] Regional managed services (Lambda, S3, SQS, etc.) → inside `<region>`, outside `<vpc>`
- [ ] Network edge (IGW, ELB) → inside `<vpc>`, outside `<availability-zone>`
- [ ] Compute/DB instances → inside `<availability-zone>`
- [ ] Services not tied to a VPC → never inside `<vpc>` or `<availability-zone>`

```xml
<frame version="1" width="1440" height="900" class="pa-4">
  <aws-cloud id="aws-cloud" title="AWS Cloud">

    <!-- ✅ Global: outside <region> — not bound to any specific region -->
    <generic-group id="global-services" title="Global Services">
      <item id="1179" />  <!-- Route 53 -->
      <item id="216"  />  <!-- IAM -->
    </generic-group>

    <region id="region-apne1" title="ap-northeast-1" row="8">

      <!-- ✅ Regional: inside <region>, outside <vpc> — no VPC required -->
      <generic-group id="managed-serverless" title="Managed &amp; Serverless">
        <item id="13"   />  <!-- Lambda -->
        <item id="1020" />  <!-- S3 -->
      </generic-group>

      <vpc id="vpc-main" title="VPC (10.0.0.0/16)" row="6">

        <!-- ✅ VPC-edge: inside <vpc>, outside <availability-zone> -->
        <generic-group id="vpc-edge" title="VPC Edge">
          <item id="1581" />  <!-- Internet Gateway -->
          <item id="1182" />  <!-- ELB -->
        </generic-group>

        <row gap="8" row="5">
          <col span="6">
            <availability-zone id="az-apne1a" title="AZ: ap-northeast-1a">
              <!-- ✅ AZ-specific: public-subnet for NAT Gateway -->
              <public-subnet id="public-subnet-a" title="Public Subnet">
                <item id="1582" />  <!-- NAT Gateway -->
              </public-subnet>
              <!-- ✅ AZ-specific: compute instances in private subnet -->
              <private-subnet id="app-tier-a" title="Application Tier" row="3">
                <item id="27"  />   <!-- EC2 -->
                <item id="547" />   <!-- ECS -->
              </private-subnet>
            </availability-zone>
          </col>
          <col span="6">
            <availability-zone id="az-apne1b" title="AZ: ap-northeast-1b">
              <!-- ✅ AZ-specific: DB instances in private subnet -->
              <private-subnet id="data-tier-b" title="Data Tier">
                <item id="117" />   <!-- RDS -->
                <item id="110" />   <!-- Aurora -->
              </private-subnet>
            </availability-zone>
          </col>
        </row>

      </vpc>
    </region>
  </aws-cloud>

  <connection src="1182" dst="27" />
  <connection src="27"   dst="117" />
</frame>
```

Every `<connection>` must be a direct child of `<frame>` or a direct child of a
frame-level `<connections>` group. Each `src` / `dst` value must match exactly
one item, AWS group, rectangle, port, or identified child frame by catalog ID,
`id`, `name`, or `ref`. If the same service icon appears multiple times, give
the connected item a unique `name` or `ref` and use that value as the endpoint.

For network diagrams, define structural paths and communication flows
separately:

```xml
<connection src="client" dst="router" kind="route" />
<connection src="client" dst="router" kind="traffic" color="#2563eb" />
```

Routes have no arrowheads. V1 rejects a route whose effective start or end
arrowhead is non-`none`, including values inherited from `<connections>`.
Traffic lines are directional and, when they share the same endpoints as a
route, render beside the route lane. See
[docs/src/examples/samples/route-traffic.xal](../../docs/src/examples/samples/route-traffic.xal) for a compact
route/traffic example.

Excalidraw output uses the same orthogonal routing metadata and adds small
editable anchor grids behind item icons. These anchors keep lines from covering
icons/labels while preserving visible endpoints. When several lines would share
the same X or Y lane, the renderer offsets later lines where possible. Group
header tags are treated as route obstacles so tag labels stay readable.

Reference: [docs/src/examples/samples/sample.xal](../../docs/src/examples/samples/sample.xal)
DSL specification: [xal-spec.instructions.md](xal-spec.instructions.md)

---

## Step 4 — Render the Excalidraw file

```bash
xaligo render docs/src/examples/samples/sample.xal \
  --format excalidraw \
  -o output/sample.excalidraw \
  --services docs/src/examples/samples/services.csv
```

`--services` is strongly recommended for this workflow. The CSV provides
icon label overrides and service metadata. SVG output also uses it to draw a
service legend; place that legend with
`--svg-legend-position top|right|bottom|left` (default `bottom`).

> **Note:** Create the output directory if it does not already exist.
> ```bash
> mkdir -p output
> ```

---

## Command Reference

| Command | Description |
|---|---|
| `grep -i "<name>" etc/resources/aws/service-index.csv` | Search for a service ID |
| `xaligo render <xal> --format excalidraw -o <out> --services <csv>` | Convert .xal → .excalidraw with legend |
| `xaligo render <xal> --format svg -o <out.svg> --services <csv> --svg-legend-position right` | Convert .xal → SVG with a service legend |
| `xaligo render <xal> --format pptx -o <out.pptx> --services <csv> --paper A3 --orientation landscape` | Convert .xal → PPTX when the WASI exporter is configured |
| `xaligo render <xal> --format pdf -o <out.pdf>` | Convert .xal → PDF with one frame per page by default |
| `xaligo render <xal> --format excel -o <out.xlsx>` | Convert .xal → Excel with one frame SVG per worksheet by default |
| `xaligo add service --list <csv> --file <excalidraw>` | Add service icons to an existing file |
| `xaligo render <xal> -o <excalidraw>` | Convert .xal → .excalidraw without legend |

## PPTX Notes

- Native CLI export requires `xaligo.wasm`; the npm/WASM API currently
  exports through PptxGenJS.
- Diagram and legend icons are native SVG media. SVG-capable PowerPoint or a
  compatible viewer is required; legacy raster-only viewers do not display the
  icons because independently rasterized copies are not embedded.
- PPTX export adds separate legend slide(s) after all frame/diagram slides.
- Legend pages use 4 columns and show icon, abbreviation, and official name.
- Use `--paper A3 --orientation landscape --paper-margin-top 0.75 --paper-margin-bottom 0.75`
  for the current large AWS sample.
- Connector routing is resolved in Go/WASM and avoids icon/label obstacles.
- Group header tag labels are intentionally single-line in PPTX output; keep
  tag background width and label width in sync when adjusting tag text metrics.
- Group header and item label width estimates count East Asian full-width
  characters as double-width, so Japanese and other full-width labels keep their
  text boxes aligned across Excalidraw, SVG, and PPTX.
- Keep `docs/src/examples/samples/sample.xal` and `docs/src/examples/samples/services.csv` in sync so the legend
  includes every diagram service.
