---
applyTo: ".github/instructions/manual/**"
---

# 06.01.01 Roadmap: V1 Structured-Diagram Profile

### V1 Structured-Diagram Profile

The table, relational-database, and UML design extends the V1 engine. Canonical
V1 uses `<xaligo version="1">`; historical root `<frame>` and `<frames>`
documents remain compatible but emit a migration warning. V2 uses the same
document envelope with `<xaligo version="2">`.

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
