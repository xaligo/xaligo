package usecase

import (
	"bytes"
	"context"
	"errors"
)

// Validate runs the same parser and layout validation used by Render.
func Validate(ctx context.Context, input []byte) error {
	diagnostics, err := Diagnose(ctx, input)
	if err != nil {
		return err
	}
	if len(diagnostics) > 0 {
		return &DiagnosticsError{Diagnostics: diagnostics}
	}
	return nil
}

// Diagnose validates a document and returns editor-friendly source positions.
func Diagnose(ctx context.Context, input []byte) ([]Diagnostic, error) {
	if err := checkContext(ctx); err != nil {
		return nil, err
	}
	doc, err := Parse(bytes.NewReader(input))
	if err != nil {
		return []Diagnostic{diagnosticFromError(err)}, nil
	}
	if _, err := Build(doc); err != nil {
		return []Diagnostic{{Severity: SeverityError, Message: err.Error()}}, nil
	}
	if err := checkContext(ctx); err != nil {
		return nil, err
	}
	return nil, nil
}

func diagnosticFromError(err error) Diagnostic {
	diagnostic := Diagnostic{Severity: SeverityError, Message: err.Error()}
	var positioned *Error
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
