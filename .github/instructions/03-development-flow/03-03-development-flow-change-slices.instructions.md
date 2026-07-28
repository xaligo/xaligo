---
applyTo: ".github/instructions/manual/**"
---

# 03.03 Development flow: Change Slices

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
