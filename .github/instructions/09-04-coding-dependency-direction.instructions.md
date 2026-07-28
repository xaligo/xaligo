---
applyTo: ".github/instructions/manual/**"
---

# 09.04 Coding: Dependency direction

## Dependency direction

- A repository must not construct, retain, or call another repository.
- A use case must not call another independently constructed use case.
- A controller must not call another controller.
- Multi-repository coordination belongs to a use case. Multi-use-case
  coordination belongs to a controller or composition boundary.
