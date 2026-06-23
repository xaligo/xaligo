package entity

import "io"

// PptxExportOptions contains a resolved plan and adapter write options.
type PptxExportOptions struct {
	PlanJSON     []byte
	Output       string
	Title        string
	Author       string
	Company      string
	Subject      string
	Compression  *bool
	ExporterWASM string
	Stdout       io.Writer
	Stderr       io.Writer
}
