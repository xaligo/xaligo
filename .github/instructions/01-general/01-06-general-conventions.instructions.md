---
applyTo: ".github/instructions/manual/**"
---

# 01.06 General: Conventions

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
