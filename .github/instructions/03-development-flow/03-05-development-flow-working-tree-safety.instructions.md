---
applyTo: ".github/instructions/manual/**"
---

# 03.05 Development flow: Working-Tree Safety

## Working-Tree Safety

- Preserve unrelated tracked and untracked changes. Never discard, reset, or
  stash them merely to simplify the current task.
- Do not stage another task's changes unless the user explicitly asks to commit
  all current changes.
- When the user explicitly asks to commit all current changes, inventory every
  modified, deleted, and untracked path, classify the complete set by cohesive
  responsibility, and commit those groups in dependency order. Do not turn
  that instruction into one catch-all stage or commit.
- If the current task overlaps an existing edit, inspect the combined diff and
  preserve both intents. Stop for direction only when they cannot be reconciled
  safely.
- When work is delegated in parallel, give workers non-overlapping edit scopes.
  One coordinating owner performs staging and commits serially so concurrent
  index changes cannot mix responsibilities.
- Do not commit dependencies, caches, binaries, package `dist`, `output`, or
  WASM build artifacts. Source-controlled documentation images are allowed
  only when the repository intentionally tracks them and their editable source
  is committed in the same slice.
