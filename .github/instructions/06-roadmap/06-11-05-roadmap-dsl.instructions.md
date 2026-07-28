---
applyTo: ".github/instructions/manual/**"
---

# 06.11.05 Roadmap: DSL

### DSL

```text
web --- db
web ==> db
```

Status: implemented using `<item name="...">`, `<item ref="...">`, or numeric
item IDs. Shorthands expand into the shared connection model during parsing.

---
