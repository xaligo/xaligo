---
applyTo: ".github/instructions/manual/**"
---

# 03.07 Development flow: Documentation and Architecture Diagrams

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
