---
applyTo: ".github/instructions/manual/**"
---

# 09.01.01 Coding: `internal/usecase` root contract

### `internal/usecase` root contract

Every non-test Go file directly below `internal/usecase` is one complete
use-case component. Its declarations follow this order:

1. `XxxUsecase` interface;
2. unexported `xxxUsecase` concrete type;
3. `NewXxxUsecase` constructor returning the interface; and
4. receiver methods containing that component's orchestration.

Repository dependencies are constructor-injected fields on the concrete type.
Do not leave declaration-free algorithm/wrapper/helper files such as a file
containing only package functions. Calculation helpers belong in the versioned
engine. A shared-source compatibility function may follow the receiver methods
only when it is a thin, deprecated delegation with no independent logic.

Private orchestration helpers for a component stay in that component's file;
do not create `render_options.go`, `render_scene.go`, or another root-level
implementation fragment that lacks its own interface, concrete type, and
constructor.
