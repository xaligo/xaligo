# Instructions index

Read `00-core-general.instructions.md` first, then only matching files. Numeric
prefixes group concerns; the suffix names the exact content. `reference.md`
contains all detail in numbered reading order.

| Scope | Read |
|---|---|
| Any change | `00-core-general`, `01-workflow-changes` |
| Go/TypeScript structure | `10-code-layers`, `11-code-files`, `12-code-v1-engine` |
| `.xal` syntax/layout | `20-dsl-document`, `21-dsl-layout` |
| Connections/PPTX | `22-dsl-connections`, `30-output-pptx` |
| Diagram authoring | `40-author-diagrams` |
| Scope/status/planning | `50-product-features`, `51-product-quality`, `52-product-roadmap` |

All names above end in `.instructions.md`. Search detailed material narrowly:

```bash
rg -n '^# (0[1-9]|1[01]) |<term|feature-id>' .github/instructions/reference.md
```

Reference authority: `07` DSL behavior; `08` architecture; `09` code shape;
`03` workflow; `04–06` product status/direction.
