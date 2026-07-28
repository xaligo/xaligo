---
applyTo: ".github/instructions/manual/**"
---

# 03.08 Development flow: Completion Audit

## Completion Audit

Before handing work back:

1. Confirm every task-owned tracked and untracked change is committed.
2. Confirm the worktree has no unexpected staged or unstaged changes.
3. Review the new commit sequence with `git log --oneline` for ordering,
   granularity, and clear subjects.
4. Report the commit count or range, verification results, any intentionally
   uncommitted user changes, and whether anything was not run.
