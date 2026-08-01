---
applyTo: ".github/instructions/manual/**"
---

# 03.06 Development flow: Verification Cadence

## Verification Cadence

Use layered verification rather than running every expensive command after
every hunk:

- run focused unit or adapter tests before the commit that introduces the
  behavior;
- run package-level or integration tests after a structural cutover;
- run the broad relevant suite at implementation milestones and before final
  handoff; and
- rerun verification after generators or formatters in case they changed the
  worktree.

Typical final checks are selected by scope:

```bash
make security-check
go test ./... -count=1
go build ./...
npm --prefix external/pptx-exporter test
mdbook build docs
git diff --check
git status --short
```

For `.xal` changes, validate and render the affected source with the same V1
pipeline used in production. For PPTX changes, also exercise the external
exporter tests and package validation when the required WASM/tooling is
available. A sandbox, missing optional tool, or pre-existing failure must be
reported explicitly; it must not be represented as a successful check.

When a shared DSL, layout, routing, scene, or renderer contract changes, verify
the relevant format matrix and representative documentation `.xal` corpus, not
only the one source that originally exposed the defect.
