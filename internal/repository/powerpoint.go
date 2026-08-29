package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/xaligo/xaligo/external/exporter"
	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/share"
)

var (
	logger     = share.DefaultLogger()
	IRPEPWX002 = share.NewMCode("IRPEPWX-002", "Export PPTX bytes generated output")
)

type PowerpointRepository interface {
	ExportPptxBytes(ctx context.Context, opts entity.PptxExportOptions) ([]byte, error)
}

type powerpointExportFunc func(context.Context, string, []byte) ([]byte, []byte, error)

type PowerpointRepositoryOption func(*powerpointRepository)

type powerpointRepository struct {
	export powerpointExportFunc
}

func NewPowerpointRepository(options ...PowerpointRepositoryOption) PowerpointRepository {
	repository := &powerpointRepository{export: exportPptxNative}
	for _, option := range options {
		if option != nil {
			option(repository)
		}
	}
	return repository
}

// WithPowerpointExportFunc replaces the external byte exporter. It is intended
// for adapter tests; production construction uses the statically linked Rust implementation.
func WithPowerpointExportFunc(export func(context.Context, string, []byte) ([]byte, []byte, error)) PowerpointRepositoryOption {
	return func(repository *powerpointRepository) {
		if export != nil {
			repository.export = export
		}
	}
}

type pptxExporterRequest struct {
	Plan    json.RawMessage     `json:"plan"`
	Options pptxExporterOptions `json:"options,omitempty"`
}

type pptxExporterOptions struct {
	Title       string `json:"title,omitempty"`
	Author      string `json:"author,omitempty"`
	Company     string `json:"company,omitempty"`
	Subject     string `json:"subject,omitempty"`
	Compression *bool  `json:"compression,omitempty"`
}

func (rcvr *powerpointRepository) ExportPptxBytes(ctx context.Context, opts entity.PptxExportOptions) ([]byte, error) {
	if len(bytes.TrimSpace(opts.PlanJSON)) == 0 {
		return nil, fmt.Errorf("PPTX plan JSON is required")
	}
	req := pptxExporterRequest{
		Plan: json.RawMessage(opts.PlanJSON),
		Options: pptxExporterOptions{
			Title:       opts.Title,
			Author:      opts.Author,
			Company:     opts.Company,
			Subject:     opts.Subject,
			Compression: opts.Compression,
		},
	}
	reqJSON, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("encode PPTX exporter request: %w", err)
	}

	pptxBytes, _, err := rcvr.export(ctx, "", reqJSON)
	if err != nil {
		return nil, err
	}
	if len(pptxBytes) == 0 {
		return nil, fmt.Errorf("Rust PPTX exporter produced no output")
	}
	logger.INFO(IRPEPWX002, "generated")
	return pptxBytes, nil
}

func exportPptxNative(ctx context.Context, _ string, requestJSON []byte) ([]byte, []byte, error) {
	data, err := exporter.ProcessContext(ctx, requestJSON)
	return data, nil, err
}
