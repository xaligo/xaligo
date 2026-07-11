package engine

import (
	"bytes"
	"errors"

	"github.com/xaligo/xaligo/internal/entity"
)

// DiagnoseV1EngineDiagnoseDocument parses and lays out one document without performing I/O or
// interpreting cancellation. The parent usecase owns those concerns.
func DiagnoseV1EngineDiagnoseDocument(input []byte) []entity.Diagnostic {
	doc, err := ParseV1EngineParseDocument(bytes.NewReader(input))
	if err != nil {
		return []entity.Diagnostic{diagnosticFromErrorV1EngineDiagnoseDocument(err)}
	}
	if _, err := BuildV1EngineLayoutBuild(doc); err != nil {
		return []entity.Diagnostic{diagnosticFromErrorV1EngineDiagnoseDocument(err)}
	}
	if _, specified := doc.Root.Attrs["version"]; !specified {
		return []entity.Diagnostic{{
			Severity: SeverityWarningV1EngineOptionRender,
			Message:  implicitV1VersionWarningV1EngineParseNode(doc.Root),
			Offset:   doc.Root.Position.Offset,
			Line:     doc.Root.Position.Line,
			Column:   doc.Root.Position.Column,
		}}
	}
	return nil
}

func diagnosticFromErrorV1EngineDiagnoseDocument(err error) entity.Diagnostic {
	diagnostic := entity.Diagnostic{Severity: SeverityErrorV1EngineOptionRender, Message: err.Error()}
	var positioned *entity.ParseError
	if errors.As(err, &positioned) {
		diagnostic.Message = positioned.Err.Error()
		diagnostic.Offset = positioned.Position.Offset
		diagnostic.Line = positioned.Position.Line
		diagnostic.Column = positioned.Position.Column
	}
	return diagnostic
}
