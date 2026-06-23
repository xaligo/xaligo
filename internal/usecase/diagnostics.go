package usecase

import (
	"bytes"
	"context"
	"errors"

	"github.com/ryo-arima/xaligo/internal/entity"
)

// Validate runs the same parser and layout validation used by Render.
func Validate(ctx context.Context, input []byte) error {
	diagnostics, err := Diagnose(ctx, input)
	if err != nil {
		return err
	}
	if len(diagnostics) > 0 {
		return &entity.DiagnosticsError{Diagnostics: diagnostics}
	}
	return nil
}

// Diagnose validates a document and returns editor-friendly source positions.
func Diagnose(ctx context.Context, input []byte) ([]entity.Diagnostic, error) {
	if err := checkContext(ctx); err != nil {
		return nil, err
	}
	doc, err := Parse(bytes.NewReader(input))
	if err != nil {
		return []entity.Diagnostic{diagnosticFromError(err)}, nil
	}
	if _, err := Build(doc); err != nil {
		return []entity.Diagnostic{{Severity: SeverityError, Message: err.Error()}}, nil
	}
	if err := checkContext(ctx); err != nil {
		return nil, err
	}
	return nil, nil
}

func diagnosticFromError(err error) entity.Diagnostic {
	diagnostic := entity.Diagnostic{Severity: SeverityError, Message: err.Error()}
	var positioned *entity.ParseError
	if errors.As(err, &positioned) {
		diagnostic.Message = positioned.Err.Error()
		diagnostic.Offset = positioned.Position.Offset
		diagnostic.Line = positioned.Position.Line
		diagnostic.Column = positioned.Position.Column
	}
	return diagnostic
}

func checkContext(ctx context.Context) error {
	if ctx == nil {
		return nil
	}
	return ctx.Err()
}
