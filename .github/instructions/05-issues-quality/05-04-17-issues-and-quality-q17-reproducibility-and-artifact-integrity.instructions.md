---
applyTo: ".github/instructions/manual/**"
---

# 05.04.17 Issues and quality: Q17 Reproducibility and Artifact Integrity

### Q17 Reproducibility and Artifact Integrity

| Task | Status | Scope | Verification target |
|---|---|---|---|
| Q17.1 | not-started | Confirm checked-in documentation SVGs reproduce byte-for-byte or through a documented normalization step. | regenerate-and-compare command |
| Q17.2 | not-started | Detect stale generated images, missing source `.xal` files, and orphaned documentation assets. | source/asset inventory script or test |
| Q17.3 | not-started | Audit npm, release, and documentation package contents for required and forbidden files. | package dry run and archive listing |
| Q17.4 | not-started | Confirm generators are idempotent and do not reorder or rewrite unrelated files. | double-run clean-worktree check |
| Q17.5 | not-started | Verify timestamps, random IDs, map iteration, and platform paths do not make supported outputs unstable. | cross-run and cross-platform determinism checks |
