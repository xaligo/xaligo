---
applyTo: ".github/instructions/manual/**"
---

# 01.02 General: Directory structure

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
├── external/
│   ├── engine/                  Rust layout/SVG engine workspace
│   └── pptx-exporter/           TypeScript/PptxGenJS PPTX adapter
│       ├── index.ts             package API composition boundary
│       ├── command.ts           TypeScript/WASI entry point
│       ├── controller/          request and byte-I/O adapter
│       ├── entity/              TypeScript PPTX plan types
│       ├── repository/          PptxGenJS and package adapters
│       └── usecase/             independent PPTX application use cases
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
