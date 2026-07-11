---
applyTo: "**/*.{go,ts}"
---

# Coding Preconditions

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
