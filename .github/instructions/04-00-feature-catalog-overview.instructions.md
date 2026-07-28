---
applyTo: ".github/instructions/manual/**"
---

# 04.00 Feature catalog


This file is the authoritative, ID-addressable catalog of xaligo's supported
and planned features. Read it as a precondition alongside
`02-00-agent-guide-overview.instructions.md` to understand what the product already does (and
what it has already committed to doing) before proposing new work, filing a
roadmap entry, or judging whether a request is a bug fix, an extension of an
existing feature, genuinely new scope, or an already-tracked planned item.

Each row is a stable 7-digit feature ID (`XAL-GNNNNNN`), so implementation
notes, commit messages, roadmap entries, and issue reports can reference a
feature without repeating its full description. IDs are never reused or
renumbered once assigned; a removed feature's ID is retired, not recycled.

Every row carries a `Status` column:

- `Implemented` — shipped and available today.
- `Planned` — not yet implemented; tracked in `06-00-roadmap-overview.instructions.md` and/or
  the `05-00-issues-and-quality-overview.instructions.md` Q05 backlog, which is the
  authoritative source for its exact sequencing and scope.
- `Excluded unless justified` — a considered but deliberately unsupported
  capability that stays out of scope until a non-substitutable use case is
  identified.

Do not remove or renumber a row when its status changes; update its `Status`
and Summary in place instead (e.g., `Planned` -> `Implemented` once a feature
ships).

- The leading digit `G` is the group (major capability area); see the section
  headers below.
- The remaining 6 digits are a per-group sequence number in steps of 10
  (`000010`, `000020`, ...), leaving room to insert a new fine-grained feature
  between two existing ones without renumbering the group.
- Add a new feature at the next free step within its group, or open a new
  group (next leading digit) for a capability area not covered below.
