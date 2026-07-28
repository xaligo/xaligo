---
applyTo: "**/*.{go,ts}"
---

# File and component shape

Organize by responsibility, not declaration kind. Use `<component>.go` and
`<component>_<detail>.go`; do not repeat layer names in filenames or create
declaration-only interface/constructor files.

Keep interface, unexported implementation, constructor/factory, and principal
methods together. Root `internal/usecase` files are complete components;
calculations belong in the versioned engine. Layer names are
`<Component>{Controller|Usecase|Repository}` with matching `New<Component>...`.

Run `gofmt`, tests, build, and `git diff --check`. Exact contract:
`reference.md` section `09`.
