---
applyTo: ".github/instructions/manual/**"
---

# 09.01 Coding: Responsibility-based files

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
