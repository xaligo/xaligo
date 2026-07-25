package entity

import "time"

// PreviewKind selects the content produced by a live preview server.
type PreviewKind string

const (
	// PreviewKindSVG serves a single combined SVG canvas (the .xal preview).
	PreviewKindSVG PreviewKind = "svg"
	// PreviewKindHTML serves a full HTML document with rendered diagrams
	// embedded inline (the Markdown preview).
	PreviewKindHTML PreviewKind = "html"
)

type PreviewOptions struct {
	Render       RenderOptions
	PollInterval time.Duration
	Kind         PreviewKind
}

type PreviewStatus struct {
	Version     uint64       `json:"version"`
	Error       string       `json:"error,omitempty"`
	Diagnostics []Diagnostic `json:"diagnostics,omitempty"`
}
