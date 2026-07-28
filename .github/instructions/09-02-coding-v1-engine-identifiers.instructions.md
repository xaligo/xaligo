---
applyTo: ".github/instructions/manual/**"
---

# 09.02 Coding: V1 engine identifiers

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
