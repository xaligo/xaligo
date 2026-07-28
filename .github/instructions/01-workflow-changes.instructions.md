---
applyTo: "**"
---

# Change workflow

Implementation requests authorize small local commits unless the user says no;
reviews and diagnoses are read-only. Never push, publish, tag, open a PR, or
rewrite history without explicit permission.

Before editing: inspect status and both diffs, preserve user work, identify
contracts/tests/docs, reproduce behavior, and fix the earliest shared owner.
Slice in dependency order: contract → engine → use case → repository →
controller → docs. Keep tests with behavior.

Before each commit: stage explicit paths/hunks; inspect the staged diff; run
`git diff --cached --check`, `make security-check`, and focused checks. Use
`<type>: <imperative outcome>`; never WIP commits.

Before handoff run broad relevant checks and report commits, checks, omissions,
and preserved changes. Detail: `reference.md` section `03`.
