---
applyTo: ".github/instructions/manual/**"
---

# 05.07 Issues and quality: First UML Quality Slice

## First UML Quality Slice

After the current Q01 canonical-envelope follow-ups are closed or explicitly
paused, start the UML quality pass with
`docs/src/examples/samples/uml-activity.xal`. It is the smallest UML slice that
exercises both semantic precision and visible design quality, and should
verify:

- `initial`, `action`, `object-node`, `decision`, `fork`, `join`, `merge`, and
  `final` nodes.
- `control-flow` and `object-flow` rendering and visual distinction.
- `guard`, `title`, nested `responsibility`, and nested `constraint` behavior.
- Left-to-right activity readability.
- Decision diamonds, fork/join bars, merge nodes, final nodes, and connector
  labels avoiding collisions.
- SVG output quality sufficient for documentation, with additional format
  checks added when shared contracts or encoder behavior changes.
