---
applyTo: ".github/instructions/manual/**"
---

# 05.04.14 Issues and quality: Q14 CLI, API, and Preview Contracts

### Q14 CLI, API, and Preview Contracts

| Task | Status | Scope | Verification target |
|---|---|---|---|
| Q14.1 | not-started | Audit CLI flags, defaults, aliases, validation, exit codes, stdout/stderr, and output path behavior. | controller/command integration tests |
| Q14.2 | not-started | Confirm constructor-injected use-case APIs and convenience methods return equivalent contracts. | use-case API tests |
| Q14.3 | not-started | Verify native and embedded asset sources produce equivalent output at matching settings. | native/embedded parity tests |
| Q14.4 | not-started | Audit live preview initial render, reload events, diagnostics, browser refresh, and file-change recovery. | preview tests and local serve check |
| Q14.5 | not-started | Confirm SVG is the default, SVG/PPTX are the only accepted formats, and retired format aliases are rejected. | compatibility tests and docs review |
