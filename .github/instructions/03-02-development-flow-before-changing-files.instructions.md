---
applyTo: ".github/instructions/manual/**"
---

# 03.02 Development flow: Before Changing Files

## Before Changing Files

1. Read `01-00-general-overview.instructions.md` and every instruction whose `applyTo` matches
   the current scope. Re-evaluate matching instructions whenever the scope
   expands.
2. Use RTK for repository operations when an RTK wrapper exists, especially
  noisy commands such as `git`, `rg`, `go`, `npm`, tests, diffs, logs, and
  dependency inspection. Prefer `rtk <tool>` or another RTK-filtered command
  shape to reduce token-heavy output while preserving enough diagnostic
  signal. If RTK is missing in a local environment, install it with
  `curl -fsSL https://raw.githubusercontent.com/rtk-ai/rtk/refs/heads/master/install.sh | sh`
  before continuing unless the task is urgent or the environment disallows
  network installs.
3. Inspect `git status --short`, the unstaged diff, and the staged diff before
   editing. Existing changes belong to the user unless the current task clearly
   created them.
4. Identify the affected contracts, implementation layers, tests,
   documentation, generated source-controlled assets, and verification
   commands before choosing commit boundaries.
5. Reproduce a reported defect or establish an observable invariant before
   changing code. Prefer a regression test that fails for the diagnosed cause.
6. Fix the earliest shared layer that owns the information. Do not hide an
   engine or contract defect by changing only one sample, diagram, controller,
   or output encoder.

When one occurrence reveals a structural pattern, audit the corresponding
controller, use-case, repository, entity, V1-engine, and external TypeScript
layers as applicable. Examples include peer-layer dependencies, misplaced
interfaces or constructors, duplicate dispatch, stale format-specific names,
and inconsistent validation/render paths.
