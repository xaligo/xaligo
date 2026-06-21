// Package render exposes the stable renderer-independent use case used by the
// CLI, preview server, WebAssembly adapter, editors, and integration tests.
package usecase

import (
	"errors"

	"github.com/ryo-arima/xaligo/internal/entity"
)

type Mode = entity.Mode
type Format = entity.Format
type DiagnosticSeverity = entity.DiagnosticSeverity

const (
	ModeStandard Mode = "standard"
	ModeNetwork  Mode = "network"
	ModeAWS      Mode = "aws"

	FormatExcalidraw Format = "excalidraw"
	FormatSVG        Format = "svg"
	FormatPPTX       Format = "pptx"
	FormatXYFlow     Format = "xyflow"
	FormatIsoflow    Format = "isoflow"

	SeverityError DiagnosticSeverity = "error"
)

var ErrNotImplemented = errors.New("renderer not implemented")

// AssetSource describes an embedded or virtual asset tree. Leave Assets nil
// to use the native project configuration and filesystem paths.
type AssetSource = entity.AssetSource

// RenderOptions contains renderer-independent presentation, routing, and
// output options. Assets is intended for filesystem-less adapters such as
// WebAssembly; native callers normally leave it nil.
type RenderOptions = entity.RenderOptions
type Diagnostic = entity.Diagnostic
type DiagnosticsError = entity.DiagnosticsError
