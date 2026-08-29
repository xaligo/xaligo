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
│   ├── controller/              CLI/LSP/RAG flags and transport/file-I/O adapters
│   ├── core/profiles/           declarative builtin and domain profile boundaries
│   ├── entity/                  internal structures; independent entity layer
│   ├── lsp/                     LSP 3.18 session and protocol adapter
│   ├── usecase/
│   │   ├── render.go            RenderUsecase, dispatch, and orchestration
│   │   ├── diff.go              DiffUsecase, structural comparison, and paired SVG orchestration
│   │   ├── diagnostics.go       DiagnosticsUsecase and shared validation
│   │   ├── project.go           docs-only RAG indexing and explicit .xal concept analysis
│   │   ├── parser.go            ParserUsecase over the V1 engine
│   │   ├── layout.go            LayoutUsecase over the V1 engine
│   │   ├── pagination.go        PaginationUsecase over the V1 engine
│   │   ├── plan.go              PlanUsecase over the V1 engine
│   │   ├── scene.go             SceneUsecase and repository-port adapter
│   │   ├── v1/engine/           synchronous V1 parser/layout/scene/plan logic
│   │   └── v2/                  V2 Go orchestration over the Rust engine ABI
│   ├── repository/              filesystem/output adapters plus SQLite icon/project stores
│   └── config/                  project configuration
├── test/
│   ├── unit/                    unit tests mirroring the project tree
│   │   └── external/            Rust engine/exporter unit tests
│   └── integration/             black-box use-case/adapter tests
├── external/
│   ├── engine/                  Go/cgo adapter and one Rust staticlib crate
│   │   ├── abi/fields.csv       authoritative Go/Rust ABI field index schema
│   │   ├── src/base.rs          Rust engine composition root
│   │   ├── src/cnf/             ABI constants, limits, and defaults
│   │   ├── src/ent/             model, request, and response entities
│   │   ├── src/usc/             flat dispatch, layout_*, routing, and SVG calculations
│   │   ├── src/rep.rs           reserved future external-representation layer
│   │   ├── src/ctl/             C ABI controller
│   │   └── src/util/            explicit codecs/traits, message codes, logging, and errors
│   └── exporter/                Rust `pptx` adapter with a C ABI
│       ├── Cargo.toml
│       └── src/                 same cnf/ent/usc/rep/ctl/util layers as engine
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
