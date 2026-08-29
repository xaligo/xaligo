// Package exporter contains the Rust PPTX exporter and its narrow Go/C ABI adapter.
package exporter

import "errors"

var ErrUnavailable = errors.New("Rust PPTX exporter is unavailable; build with make build-exporter or make build")
