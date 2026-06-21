package entity

import "fmt"

type Diagnostic struct {
	Severity DiagnosticSeverity `json:"severity"`
	Message  string             `json:"message"`
	Offset   int                `json:"offset,omitempty"`
	Line     int                `json:"line,omitempty"`
	Column   int                `json:"column,omitempty"`
}

type DiagnosticsError struct {
	Diagnostics []Diagnostic
}

func (e *DiagnosticsError) Error() string {
	if len(e.Diagnostics) == 0 {
		return "validation failed"
	}
	d := e.Diagnostics[0]
	if d.Line > 0 {
		return fmt.Sprintf("line %d, column %d: %s", d.Line, d.Column, d.Message)
	}
	return d.Message
}
