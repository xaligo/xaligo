# Terminal Output

Terminal output projects a resolved V2 document directly into text for TUIs,
SSH sessions, CI logs, and command-line inspection. It does not rasterize SVG
and does not pass through the V1 compatibility scene.

```bash
xaligo render diagram-v2.xal --format terminal
```

Output is written to standard output by default. Use `-o` to create a file:

```bash
xaligo render diagram-v2.xal --format terminal -o diagram.txt
```

V1 input returns an explicit error. SVG and PPTX remain the cross-version
visual outputs.

## Layouts

`diagram` preserves resolved spatial relationships on a character grid:

```bash
xaligo render diagram-v2.xal --format terminal \
  --terminal-layout diagram --terminal-width 120 --terminal-height 40
```

`semantic` prioritizes hierarchy and connections, making it suitable for
narrow terminals and machine logs:

```bash
xaligo render diagram-v2.xal --format terminal --terminal-layout semantic
```

`hybrid` places a compact diagram above a component, flow, or focused-element
detail pane:

```bash
xaligo render diagram-v2.xal --format terminal \
  --terminal-layout hybrid --terminal-focus api
```

## Character and color modes

Unicode box drawing is the default. Strict ASCII output contains only 7-bit
ASCII characters:

```bash
xaligo render diagram-v2.xal --format terminal --terminal-style ascii
```

`--color auto` emits ANSI color only when standard output is an interactive
terminal. Redirected and file output is deterministic and uncolored unless
`--color always` is specified. `--color never` is useful when embedding output
in documentation.

## Size and detail

For an interactive terminal, omitted dimensions use `COLUMNS` and `LINES` when
available. Redirected output defaults to 100 columns and 40 rows. Explicit
dimensions are clamped to safe bounds of 20–500 columns and 8–200 rows.

`--terminal-detail compact|normal|full` controls semantic metadata.
`--terminal-icons label|symbol|none` selects whether icon-backed elements use
their label, a terminal-safe symbol plus label, or no fallback icon label.

Terminal output intentionally simplifies visual capabilities. Images become
labels or symbols, fills become optional ANSI foreground accents, arbitrary
curves become orthogonal character-cell paths, and exact typography is not
preserved. The resolved hierarchy, element labels, relative placement, and
connection paths remain available without adding a second parser or layout
pipeline.
