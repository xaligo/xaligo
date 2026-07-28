---
applyTo: ".github/instructions/manual/**"
---

# 03.04 Development flow: Commit Cadence

## Commit Cadence

Repository-changing work is committed incrementally unless the user explicitly
requests no commits. Commit each completed slice as soon as its focused
verification succeeds; do not wait until the entire request is finished.
Long-running work should therefore produce several small, ordered commits
rather than one end-of-task snapshot.

Before each commit:

1. Stage explicit paths or hunks. Prefer `git add -p` for mixed files and avoid
   broad staging in a dirty worktree.
2. Inspect `git diff --cached --name-status`, `git diff --cached --stat`, and
   `git diff --cached` to confirm that the index contains one responsibility
   and no unrelated user changes.
3. Run `git diff --cached --check`.
4. Run `make security-check`. Security scanning is a mandatory commit
   precondition for every change, including documentation-only changes. Run
   `make security-setup` once after cloning or whenever the pinned scanner
   version changes.
5. Run the narrowest relevant test, build, validation, or render command.
6. Use the existing concise subject style, normally
   `<type>: <imperative summary>`, and describe the outcome rather than the
   edited filenames.
7. Re-run `git status --short` after the commit and confirm that no unintended
   path remains staged.

Do not create empty checkpoint or `WIP` commits. Do not amend, squash, rebase,
reset, or otherwise rewrite existing history unless the user explicitly asks.
Do not push commits or publish artifacts unless the user explicitly asks.
