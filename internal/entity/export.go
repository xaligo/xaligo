package entity

// PptxExportOptions contains a resolved plan and exporter options.
type PptxExportOptions struct {
	PlanJSON    []byte
	Title       string
	Author      string
	Company     string
	Subject     string
	Compression *bool
}
