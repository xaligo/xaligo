---
applyTo: ".github/instructions/manual/**"
---

# 05.04.11 Issues and quality: Q11 CI, Release, and Tooling

### Q11 CI, Release, and Tooling

| Task | Status | Scope | Verification target |
|---|---|---|---|
| Q11.1 | not-started | Audit GitHub Actions for Go, TypeScript, docs, version, release, and pages jobs. | workflow review and CI run results |
| Q11.2 | not-started | Confirm `VERSION`, npm package metadata, and release tags stay coherent. | version-gate reproduction |
| Q11.3 | not-started | Confirm npm lockfile policy is reflected in workflows, Makefile, and docs. | TypeScript install/test command |
| Q11.4 | not-started | Verify RTK and security-check preconditions are current and actionable. | `make security-check` and instruction review |
| Q11.5 | not-started | Confirm generated artifacts, vendored dependencies, binaries, caches, and docs images follow repository policy. | `git status`, ignore rules, and release/package review |
