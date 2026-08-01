package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

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
	repository := &powerpointRepository{export: exportPptxWithWASM}
	for _, option := range options {
		if option != nil {
			option(repository)
		}
	}
	return repository
}

// WithPowerpointExportFunc replaces the external byte exporter. It is intended
// for adapter tests; production construction uses the WASM implementation.
func WithPowerpointExportFunc(export func(context.Context, string, []byte) ([]byte, []byte, error)) PowerpointRepositoryOption {
	return func(repository *powerpointRepository) {
		if export != nil {
			repository.export = export
		}
	}
}

type pptxWasmRequest struct {
	Plan    json.RawMessage `json:"plan"`
	Options pptxWasmOptions `json:"options,omitempty"`
}

type pptxWasmOptions struct {
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
	req := pptxWasmRequest{
		Plan: json.RawMessage(opts.PlanJSON),
		Options: pptxWasmOptions{
			Title:       opts.Title,
			Author:      opts.Author,
			Company:     opts.Company,
			Subject:     opts.Subject,
			Compression: opts.Compression,
		},
	}
	reqJSON, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("encode PPTX WASM request: %w", err)
	}

	pptxBytes, _, err := rcvr.export(ctx, opts.ExporterWASM, reqJSON)
	if err != nil {
		return nil, err
	}
	if len(pptxBytes) == 0 {
		return nil, fmt.Errorf("PPTX WASM exporter produced no output")
	}
	logger.INFO(IRPEPWX002, "generated")
	return pptxBytes, nil
}
