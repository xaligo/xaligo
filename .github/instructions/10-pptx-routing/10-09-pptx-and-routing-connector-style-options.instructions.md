---
applyTo: ".github/instructions/manual/**"
---

# 10.09 PPTX and routing: Connector Style Options

## Connector Style Options

`xaligo render --format pptx` forwards all PPTX routing options:

| Flag | Meaning |
|---|---|
| `--arrow-style` | `thin`, `standard`, `triangle`, `stealth`, `arrow`, `diamond`, `oval`, `none` |
| `--arrow-stub` | Pixel stub before the first/last bend |
| `--arrow-margin` | Pixel margin reserved around existing line lanes |
| `--px-per-inch` | Layout scaling base, default 96 |
| `--paper` | Named slide paper size: `A5`, `A4`, `A3`, `A2`, `A1`, `Letter`, `Legal`, `Tabloid` |
| `--orientation` | `portrait` or `landscape`; auto-fit when omitted |
| `--paper-margin` | Inch margin applied to all sides before paper fitting |
| `--paper-margin-top/right/bottom/left` | Inch margin override for one side |

`--arrow-style` is a Plan-level default. A connection's explicit or inherited
DSL arrowhead and stroke width take precedence; `kind="route"` remains
headless. The `thin` and `standard` presets may supply a default line width only
when the DSL did not supply `stroke-width` or its `width` alias.

Every numeric render option must be finite. `--px-per-inch`, arrow stub and
margin values, and paper margins reject negative values; the internal zero
value selects the documented default. Validation happens before scene/plan
construction so `NaN` or infinity cannot first fail during JSON encoding.
Paper size, orientation, and arrow style are closed enums. Paper margins require
a named paper size, and their effective left/right and top/bottom sums must
leave a strictly positive content area in the selected (or at least one
automatic) orientation.
