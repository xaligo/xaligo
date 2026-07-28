---
applyTo: "**/*.{go,ts,xal}"
---

# PPTX output

Go owns renderer-neutral scene/plan geometry; configured WASM/PptxGenJS writes
PPTX. Do not add another OOXML writer or PPTX-only routing.

Explicit DSL connector style overrides plan defaults. Geometry and text share
one effective PPI/paper-fit transform; validate finite options and positive
paper content area before plan construction. Legends follow diagram slides and
contain only referenced services.

Flags, routing precedence, label metrics, page fitting, and checks:
`reference.md` section `10`.
