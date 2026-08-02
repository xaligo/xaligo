package usecase

import (
	"context"
	"errors"

	"github.com/xaligo/xaligo/internal/entity"
	v1engine "github.com/xaligo/xaligo/internal/usecase/v1/engine"
	v2usecase "github.com/xaligo/xaligo/internal/usecase/v2"
)

type DiagnosticsUsecase interface {
	Validate(context.Context, []byte) error
	ValidateWithImports(context.Context, []byte, *entity.ImportSource) error
	Diagnose(context.Context, []byte) ([]entity.Diagnostic, error)
	DiagnoseWithImports(context.Context, []byte, *entity.ImportSource) ([]entity.Diagnostic, error)
}

type diagnosticsUsecase struct {
	frontend v2usecase.FrontendUsecase
	engine   v2usecase.EngineUsecase
}

func NewDiagnosticsUsecase() DiagnosticsUsecase {
	return &diagnosticsUsecase{frontend: v2usecase.NewFrontendUsecase(), engine: v2usecase.NewEngineUsecase()}
}

func (rcvr *diagnosticsUsecase) Validate(ctx context.Context, input []byte) error {
	return rcvr.ValidateWithImports(ctx, input, nil)
}

func (rcvr *diagnosticsUsecase) ValidateWithImports(ctx context.Context, input []byte, imports *entity.ImportSource) error {
	diagnostics, err := rcvr.DiagnoseWithImports(ctx, input, imports)
	if err != nil {
		return err
	}
	errors := make([]entity.Diagnostic, 0, len(diagnostics))
	for _, diagnostic := range diagnostics {
		if diagnostic.Severity == SeverityError {
			errors = append(errors, diagnostic)
		}
	}
	if len(errors) > 0 {
		return &entity.DiagnosticsError{Diagnostics: errors}
	}
	return nil
}

func (rcvr *diagnosticsUsecase) Diagnose(ctx context.Context, input []byte) ([]entity.Diagnostic, error) {
	return rcvr.DiagnoseWithImports(ctx, input, nil)
}

func (rcvr *diagnosticsUsecase) DiagnoseWithImports(ctx context.Context, input []byte, imports *entity.ImportSource) ([]entity.Diagnostic, error) {
	if err := checkContext(ctx); err != nil {
		return nil, err
	}
	if renderDocumentVersion(input) == "2" {
		spec, _, err := rcvr.frontend.Lower(input)
		if err != nil {
			return []entity.Diagnostic{{Code: "XAL-E1001", Severity: SeverityError, Stage: "parse", Message: err.Error()}}, nil
		}
		_, err = rcvr.engine.Resolve(ctx, spec)
		if err == nil {
			return nil, nil
		}
		var engineErr *entity.EngineDiagnosticError
		if !errors.As(err, &engineErr) {
			return nil, err
		}
		diagnostic := entity.Diagnostic{
			Code: engineErr.Diagnostic.Code, Severity: SeverityError, Stage: engineErr.Diagnostic.Stage,
			Element: engineErr.Diagnostic.ElementID, Parameter: engineErr.Diagnostic.Parameter,
			Message: engineErr.Diagnostic.Message,
		}
		for _, span := range spec.Spans {
			if span.ID == engineErr.Diagnostic.SpanID {
				diagnostic.Offset, diagnostic.Line, diagnostic.Column = span.Offset, span.Line, span.Column
				break
			}
		}
		return []entity.Diagnostic{diagnostic}, nil
	}
	diagnostics := v1engine.DiagnoseWithImportsV1EngineDiagnoseDocument(input, imports)
	if err := checkContext(ctx); err != nil {
		return nil, err
	}
	return diagnostics, nil
}

// Validate is kept as a source-compatible package boundary.
// Deprecated: construct DiagnosticsUsecase and call Validate.
func Validate(ctx context.Context, input []byte) error {
	return NewDiagnosticsUsecase().Validate(ctx, input)
}

// Diagnose is kept as a source-compatible package boundary.
// Deprecated: construct DiagnosticsUsecase and call Diagnose.
func Diagnose(ctx context.Context, input []byte) ([]entity.Diagnostic, error) {
	return NewDiagnosticsUsecase().Diagnose(ctx, input)
}

func checkContext(ctx context.Context) error {
	if ctx == nil {
		return nil
	}
	return ctx.Err()
}
