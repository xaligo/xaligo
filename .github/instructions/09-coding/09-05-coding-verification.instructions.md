---
applyTo: ".github/instructions/manual/**"
---

# 09.05 Coding: Verification

## Verification

After structural or naming changes, run at minimum:

```bash
gofmt -w <changed-go-files>
go test ./...
go build ./...
git diff --check
```

The V1 engine naming regression test must pass; do not bypass it with aliases
that leave nonconforming package-scope declarations in the engine package.
