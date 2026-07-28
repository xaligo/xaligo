---
applyTo: ".github/instructions/manual/**"
---

# 05.04.02 Issues and quality: Q02 Data Registry and Imports

### Q02 Data Registry and Imports

| Task | Status | Scope | Verification target |
|---|---|---|---|
| Q02.1 | not-started | Inventory supported `<data>` children and their reference syntax. | spec/docs consistency check |
| Q02.2 | not-started | Verify relative import resolution, missing files, unsupported extensions, and provenance. | diagnostics tests with temp files |
| Q02.3 | not-started | Verify duplicate data IDs and unresolved data references across frames. | parser/validation unit tests |
| Q02.4 | not-started | Confirm imports are deterministic and do not execute commands or read outside the allowed context. | security-oriented unit tests |
| Q02.5 | not-started | Confirm imported data renders identically to equivalent inline data where applicable. | golden or structural scene comparison |
