package entity

import "time"

type UseCaseDependencies struct{}

type PreviewOptions struct {
	Render       RenderOptions
	PollInterval time.Duration
}

type PreviewStatus struct {
	Version     uint64       `json:"version"`
	Error       string       `json:"error,omitempty"`
	Diagnostics []Diagnostic `json:"diagnostics,omitempty"`
}
