// Package engine contains the Rust V2 engine workspace and its narrow Go/C
// ABI adapter. Application callers should depend on v2.EngineUsecase.
package engine

import "errors"

// ErrUnavailable reports a build that did not link the Rust engine static
// library. The canonical native build enables it with the xaligo_engine tag.
var ErrUnavailable = errors.New("Rust engine is unavailable; build with make build-engine or make build")
