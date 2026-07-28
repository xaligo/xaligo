---
applyTo: ".github/instructions/manual/**"
---

# 01.04 General: Testing rules

## Testing rules

- Put unit tests under `test/unit`, mirroring the source tree they cover.
- Put black-box tests of exported APIs and adapters in `test/integration`.
- Prefer externally observable behavior over package-private helper assertions
  when moving tests outside implementation packages.
- Add focused coverage for behavior changes and preserve regression tests.
