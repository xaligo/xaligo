---
applyTo: ".github/instructions/manual/**"
---

# 08.11 Architecture: Single binary, shared service, and external runtimes

## Status and purpose

This file is a mandatory V2 and platform precondition. The repository migrates
incrementally toward one native executable and one constructor-injected
application service. Existing controllers and use cases remain valid migration
steps, but new work must not introduce a daemon, sidecar service, duplicate
command-specific core pipeline, or separately distributed engine executable.

## Distribution and process model

The target native distribution is one cross-platform executable:

```text
xaligo | xaligo.exe
├── lsp
├── mcp
├── serve
├── rag
├── icon
├── model
├── validate
├── render
├── format
└── license
```

Long-running capabilities are modes of that executable:

```text
xaligo lsp
xaligo mcp --stdio
xaligo mcp --http
xaligo serve
xaligo rag watch
```

Single-operation commands terminate after producing their result. Do not add:

- a `xaligo-daemon` or separate local client;
- automatic daemon startup, reconnect, or lifecycle management;
- Unix sockets, named pipes, or an internal RPC protocol for local features;
- `RemoteService` versus `LocalService` selection for in-process work; or
- separate LSP, MCP, layout, SVG, RAG, icon, or model executables.

## Shared application service

The command dispatcher, LSP adapter, MCP adapter, preview server, and ordinary
CLI controllers call one constructor-injected application service. They do not
assemble independent parser, validator, layout, formatter, renderer, registry,
search, or persistence pipelines.

The target responsibility is:

```go
type Service interface {
	OpenProject(ctx context.Context, root string) error
	Parse(ctx context.Context, uri string) (*Document, error)
	Validate(ctx context.Context, uri string) ([]Diagnostic, error)
	Format(ctx context.Context, uri string) ([]byte, error)
	Render(ctx context.Context, uri string, opts RenderOptions) ([]byte, error)
	Apply(ctx context.Context, uri string, op Operation) error
	Search(ctx context.Context, query SearchQuery) ([]SearchResult, error)
}
```

This declaration describes target ownership, not a requirement to add a
declaration-only facade. During migration, complete constructor-injected
components in `internal/usecase` are the shared service boundary.

```text
LSP | MCP | serve | validate | render | format | icon
                         |
                         v
                 shared Xaligo service
                         |
          parser / validator / registry / renderer
                         |
                         v
                in-process Rust engine
```

## Target command and source layout

```text
xaligo/
├── cmd/xaligo/main.go
├── internal/
│   ├── command/
│   ├── service/
│   ├── core/
│   ├── parser/
│   ├── ast/
│   ├── ir/
│   ├── validator/
│   ├── formatter/
│   ├── renderer/
│   ├── project/
│   ├── lsp/
│   ├── mcp/
│   ├── rag/
│   ├── ai/
│   ├── icon/
│   └── storage/
├── external/
│   ├── engine/
│   │   ├── api.go
│   │   ├── bridge_cgo.go
│   │   ├── bridge_stub.go
│   │   ├── Cargo.toml
│   │   ├── include/xaligo_engine.h
│   │   ├── lib/               generated and ignored
│   │   └── src/
│   │       ├── base.rs
│   │       ├── cnf/
│   │       ├── ent/model/
│   │       ├── ent/request/
│   │       ├── ent/response/
│   │       ├── usc/
│   │       ├── ctl/
│   │       └── util/
│   └── pptx-exporter/
├── assets/
├── Makefile
└── LICENSES/
```

The physical Go package migration follows behavior and dependency movement; do
not perform a directory-only rewrite. `external/engine` and
`external/pptx-exporter` are source workspaces, not separately launched runtime
programs.

## Rust engine integration

Rust owns generic layout and SVG calculation. Native builds compile the single
layered Rust crate to a platform-native static library, copy the generated
archive to the ignored `external/engine/lib` cgo link location, and build the
Go executable with cgo and the engine build tag. The final executable contains
the Rust archive and calls it through a versioned C ABI in process.

```text
Go EngineUsecase
      |
      | cgo + versioned C ABI
      v
linked Rust static library
      |
      +-- ctl -> usc -> ctl
      +-- rep (reserved for future Rust-owned encoders such as PPTX)
      +-- cnf + ent + util
```

The Rust source layout follows `ryo-arima/vem/src`, adapted from a CLI binary
to a static library: `base.rs` composes the request pipeline; `cnf` owns ABI
configuration; `ent` owns model, request, and response values; `usc` coordinates
operations and contains cohesive layout, routing, and SVG calculation files;
`ctl` exposes C ABI ingress and egress; `rep` remains empty until Rust owns an
external representation encoder; and `util` contains explicit binary serialization/deserialization,
standard-trait implementations, message codes, logging, and errors. Entity
behavior is implemented there without derive or serde annotations. The engine
is one crate. Do not split these responsibilities into independently versioned
layout, SVG, or FFI crates.

The generated static libraries, Cargo `target`, npm `dist`, and `node_modules`
content are never committed. Each target platform builds its own Rust archive;
do not cross-link a host archive into another target. The C ABI keeps variable
length request/response data behind an explicit Rust-owned buffer contract and
requires the matching Rust free function. The engine remains part of the one
native executable, not a separately distributed dynamic library.

The TypeScript PPTX exporter remains under `external/pptx-exporter` during its
migration. All build, test, configuration, packaging, and documentation paths
must use that boundary. It must not gain layout or routing ownership.

## Persistent data and concurrency

The executable may open two durable SQLite databases directly:

```text
xaligo-assets.db       bundled/imported specifications and SVG icons
.xaligo/project.db     project Markdown, FTS, embeddings, symbols, AST indexes,
                       diagram structure, SVG metadata, and model profiles
```

SQLite implementations use WAL mode, a bounded `busy_timeout`, short
transactions, a single write queue per process, and explicit project-file
locking where migrations or bulk imports require it. Processes share durable
indexes, not in-progress ASTs, render intermediates, session state, or temporary
caches.

## Integrated serve mode

`serve` is the explicit multi-capability process:

```text
xaligo serve --preview --mcp --rag-watch --port 8080
├── web preview
├── file watcher
├── RAG watcher
├── MCP HTTP
├── shared AST and render caches
└── shared SQLite connections
```

Separately launched `xaligo lsp` and `xaligo mcp` have independent in-memory
caches and may share the same durable databases. That simplicity tradeoff is
intentional and is not a reason to add a daemon.

## Migration order

1. keep one `xaligo` native entry point and establish command dispatch;
2. retain complete `internal/usecase` components as the shared boundary;
3. move the generic layout/SVG engine to embedded Rust;
4. route `validate`, `format`, and `render` through one service composition;
5. add the SQLite SVG registry and `icon` commands;
6. add RAG, LSP, MCP, and model command families;
7. add integrated `serve` capabilities without a hidden daemon; and
8. package the same single-binary topology for every supported platform.
