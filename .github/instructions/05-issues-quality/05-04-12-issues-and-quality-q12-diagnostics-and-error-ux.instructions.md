---
applyTo: ".github/instructions/manual/**"
---

# 05.04.12 Issues and quality: Q12 Diagnostics and Error UX

### Q12 Diagnostics and Error UX

| Task | Status | Scope | Verification target |
|---|---|---|---|
| Q12.1 | not-started | Audit diagnostic severity, line/column positions, element context, and stable wording. | diagnostics unit tests |
| Q12.2 | not-started | Confirm multiple independent input errors are reported without hiding the first actionable cause. | aggregate-diagnostics tests |
| Q12.3 | not-started | Confirm warnings are non-blocking while errors consistently fail validate and render. | validate/diagnose agreement tests |
| Q12.4 | not-started | Verify canceled contexts stop diagnostics and rendering with wrapped context errors. | cancellation tests |
| Q12.5 | not-started | Review CLI and preview presentation of diagnostics for concise, user-correctable output. | controller/preview tests and manual command check |
