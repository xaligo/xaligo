---
applyTo: ".github/instructions/manual/**"
---

# 03.01 Development flow: Authorization and Scope

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
