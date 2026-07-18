# UML

`<uml>` contains exactly one of these diagram kinds:

- `class-diagram`, `object-diagram`, `component-diagram`
- `deployment-diagram`, `package-diagram`, `composite-structure-diagram`
- `profile-diagram`, `use-case-diagram`
- `activity-diagram`, `state-machine-diagram`
- `sequence-diagram`, `communication-diagram`
- `interaction-overview-diagram`, `timing-diagram`

All kinds accept identified UML elements and relations. Common element tags
include `class`, `interface`, `object`, `component`, `node`, `artifact`,
`package`, `part`, `port`, `profile`, `stereotype`, `actor`, `use-case`,
`activity`, `action`, `decision`, `state`, `participant`, `lifeline`,
`interaction`, `time-state`, and the neutral `element` tag.

Common relation tags include `association`, `aggregation`, `composition`,
`generalization`, `realization`, `dependency`, `include`, `extend`,
`control-flow`, `object-flow`, `transition`, `message`, `link`, `occurrence`,
`duration`, and the neutral `relation` tag. Every relation requires `src` and
`dst` IDs from the same diagram.

Child tags inside an element become text compartments. Every diagram is
normalized to renderer-neutral shapes and semantic connections, so SVG,
Excalidraw, PPTX, XYFlow, and Isoflow use the same resolved geometry.
