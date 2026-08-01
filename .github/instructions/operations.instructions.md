---
applyTo: "**"
---

# AI development operations

Use this minimal loop; open other instructions only when index search selects
them.

```bash
# 1. Route: search index, then read only matching files.
rg -i '<task terms>' .github/instructions/index.instructions.md
sed -n '1,220p' .github/instructions/<matched-file>

# 2. Establish state; never overwrite unrelated changes.
git status --short
git diff -- <in-scope-paths>
git diff --cached -- <in-scope-paths>

# 3. Locate before reading; narrow output by path and symbol.
rg -n '<symbol|error|feature-id>' <likely-paths>
sed -n '<start>,<end>p' <matched-file>

# 4. Edit the smallest coherent slice, then run focused checks.
gofmt -w <changed-go-files>
go test ./<affected-package>/... -count=1
git diff --check

# 5. Commit only task-owned paths after mandatory security checks.
git add <explicit-paths>
git diff --cached --name-status
git diff --cached --check
make security-check
git commit -m '<type>: <imperative outcome>'

# 6. Handoff audit.
git status --short
git log -1 --oneline
```

Use `npm --prefix external/pptx-exporter test`, `go build ./...`,
`mdbook build docs`, or
format-specific render checks only when the changed scope requires them.
Diagnose/read-only requests stop before editing, staging, or committing.
Never push, publish, tag, open a PR, or rewrite history without explicit user
authorization.
