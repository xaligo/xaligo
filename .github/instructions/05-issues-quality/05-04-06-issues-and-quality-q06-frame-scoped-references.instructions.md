---
applyTo: ".github/instructions/manual/**"
---

# 05.04.06 Issues and quality: Q06 Frame-Scoped References

### Q06 Frame-Scoped References

| Task | Status | Scope | Verification target |
|---|---|---|---|
| Q06.1 | not-started | Confirm local endpoint references resolve inside the owning frame first. | reference-resolution unit tests |
| Q06.2 | not-started | Confirm `frame-id.endpoint-id` resolves only to the intended frame endpoint. | cross-frame tests |
| Q06.3 | not-started | Validate duplicate IDs across frames, ambiguous endpoints, and missing frames. | diagnostics tests |
| Q06.4 | not-started | Confirm reference diagnostics include actionable source positions. | error-message assertions |
| Q06.5 | not-started | Check docs and examples explain local vs qualified reference behavior. | docs review and `mdbook build docs` |
