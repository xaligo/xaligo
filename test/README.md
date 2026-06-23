# Tests

`unit/` contains unit tests arranged by layer. Each layer directory is flat;
use filenames such as `parser_test.go`, `excalidraw_connection_test.go`, and
`pptxplan_test.go` to mirror source responsibilities.

`integration/` contains black-box tests for shared use cases and reusable
adapters. These tests import production packages through the same boundaries as
the CLI, preview, and WASM adapters.
