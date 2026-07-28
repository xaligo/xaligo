---
applyTo: ".github/instructions/manual/**"
---

# 05.04.18 Issues and quality: Q18 Configuration, Logging, and Observability

### Q18 Configuration, Logging, and Observability

| Task | Status | Scope | Verification target |
|---|---|---|---|
| Q18.1 | not-started | Audit configuration defaults, file/environment/flag precedence, and invalid-value diagnostics. | config/controller tests |
| Q18.2 | not-started | Confirm logs and wrapped errors identify the failed stage without duplicating or obscuring the root cause. | error-chain assertions and command checks |
| Q18.3 | not-started | Verify normal commands keep stdout machine-usable and send diagnostics/progress to the intended stream. | CLI output tests |
| Q18.4 | not-started | Confirm logs and diagnostics do not expose imported source contents, credentials, tokens, or sensitive paths unnecessarily. | redaction/security tests |
| Q18.5 | not-started | Define optional timing/count observability for parser, layout, routing, scene, and encoder stages without changing default output. | design decision and focused tests if implemented |
