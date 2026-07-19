package usecase

import (
	"context"

	"github.com/xaligo/xaligo/internal/entity"
	v1engine "github.com/xaligo/xaligo/internal/usecase/v1/engine"
)

type DiagnosticsUsecase interface {
	Validate(context.Context, []byte) error
	ValidateWithImports(context.Context, []byte, *entity.ImportSource) error
	Diagnose(context.Context, []byte) ([]entity.Diagnostic, error)
	DiagnoseWithImports(context.Context, []byte, *entity.ImportSource) ([]entity.Diagnostic, error)
}

type diagnosticsUsecase struct{}

func NewDiagnosticsUsecase() DiagnosticsUsecase {
	return &diagnosticsUsecase{}
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
	diagnostics := v1engine.DiagnoseWithImportsV1EngineDiagnoseDocument(input, imports)
	if err := checkContext(ctx); err != nil {
		return nil, err
	}
	return diagnostics, nil
}

// Validate is kept as a source-compatible package boundary.
// Deprecated: construct DiagnosticsUsecase and call Validate.
func Validate(ctx context.Context, input []byte) error {
	return (&diagnosticsUsecase{}).Validate(ctx, input)
}

// Diagnose is kept as a source-compatible package boundary.
// Deprecated: construct DiagnosticsUsecase and call Diagnose.
func Diagnose(ctx context.Context, input []byte) ([]entity.Diagnostic, error) {
	return (&diagnosticsUsecase{}).Diagnose(ctx, input)
}

func checkContext(ctx context.Context) error {
	if ctx == nil {
		return nil
	}
	return ctx.Err()
}
