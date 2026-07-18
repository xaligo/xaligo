---
applyTo: "**"
---

# Development Flow Preconditions

Read this file before planning, editing, reviewing, generating, testing, or
committing repository changes. The goal is a traceable history made of the
smallest coherent changes that remain safe to review, test, revert, and build.

## Authorization and Scope

- An implementation or change request authorizes incremental local commits for
  its in-scope changes unless the user explicitly requests no commits. A review,
  diagnosis, explanation, or status request remains read-only and does not
  authorize a commit.
- A local commit does not authorize pushing, force-pushing, publishing,
  tagging, opening a pull request, or rewriting existing history. Those actions
  require an explicit request.
- Treat durable user decisions as repository preconditions. Record them in the
  applicable instruction file during the same task instead of leaving them
  only in conversation history.

## Before Changing Files

1. Read `general.instructions.md` and every instruction whose `applyTo` matches
   the current scope. Re-evaluate matching instructions whenever the scope
   expands.
2. Inspect `git status --short`, the unstaged diff, and the staged diff before
   editing. Existing changes belong to the user unless the current task clearly
   created them.
3. Identify the affected contracts, implementation layers, tests,
   documentation, generated source-controlled assets, and verification
   commands before choosing commit boundaries.
4. Reproduce a reported defect or establish an observable invariant before
   changing code. Prefer a regression test that fails for the diagnosed cause.
5. Fix the earliest shared layer that owns the information. Do not hide an
   engine or contract defect by changing only one sample, diagram, controller,
   or output encoder.

When one occurrence reveals a structural pattern, audit the corresponding
controller, use-case, repository, entity, V1-engine, and external TypeScript
layers as applicable. Examples include peer-layer dependencies, misplaced
interfaces or constructors, duplicate dispatch, stale format-specific names,
and inconsistent validation/render paths.

## Change Slices

Split work into the smallest cohesive slices that provide an independently
understandable outcome. Good boundaries include:

- one precondition or coding-rule change;
- one entity or cross-layer contract change;
- one synchronous V1-engine responsibility with its focused tests;
- one root use-case component or orchestration boundary;
- one repository or output-format behavior;
- one controller, command, or composition-root migration;
- one external TypeScript command/controller/use-case/repository boundary;
- one user-facing specification or documentation topic; and
- one architecture diagram source together with its rendered documentation
  asset.

Separate a mechanical rename or move from a behavior change when each is valid
on its own. Keep a rename, its required reference updates, and any deletion of
the superseded file together when splitting them would leave broken imports,
links, or duplicate definitions.

Use dependency order where possible:

```text
contract/entity
  -> synchronous engine
  -> use-case orchestration
  -> repository/format adapter
  -> controller/command/composition root
  -> documentation and source-controlled derived assets
```

Keep implementation and its focused regression tests in the same commit.
Keep a checked-in derived documentation asset with the source that generates
it. A contract-changing implementation and its authoritative specification
must be committed together when separating them would make either commit
misleading; otherwise documentation may be a separate immediately following
commit.

Prefer every commit to build and pass its focused tests. If an interface or
constructor signature creates an inherently coupled cross-layer cutover, keep
that cutover in one atomic commit. Do not create a knowingly broken
intermediate commit, and do not add temporary aliases or compatibility facades
solely to increase the number of commits.

Do not split changes mechanically by line count or filename. A large cohesive
engine move may be one commit, while unrelated hunks in one file must be split
with patch staging.

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
npm --prefix external test
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

## Documentation and Architecture Diagrams

- Documentation about internals links to the actual implementation files or
  functions. Re-check links after moves, renames, or line-sensitive edits.
- Design-first documentation distinguishes current behavior from a planned
  target. It must not describe an unimplemented target as the current system.
- Distinguish dependency construction from runtime data flow, and include every
  real format path relevant to the documented behavior.
- Source-controlled architecture diagrams are authored as `.xal`; do not edit
  their generated SVG as the source of truth.
- Commit a diagram's `.xal` source and rendered SVG together, then build the
  documentation.
- In the internal pipeline diagram, preserve the hierarchy
  `xaligo -> internal/external -> Main/Other`. `Main` follows
  `command -> controller -> usecase -> repository` from top to bottom; command
  is an entry point. Packages are nested groups, functions are rectangles, and
  conceptual data is represented by ports. `Other` contains entities,
  configuration, shared utilities, tools, and generated artifacts.
- Check generated diagrams for group containment, sibling overlap, port-label
  overflow, and stale or missing package paths. A local size adjustment is safe
  only when the shared layout invariant has already been verified or separately
  recorded as a structural issue.

## Completion Audit

Before handing work back:

1. Confirm every task-owned tracked and untracked change is committed.
2. Confirm the worktree has no unexpected staged or unstaged changes.
3. Review the new commit sequence with `git log --oneline` for ordering,
   granularity, and clear subjects.
4. Report the commit count or range, verification results, any intentionally
   uncommitted user changes, and whether anything was not run.
