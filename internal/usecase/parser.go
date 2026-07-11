package usecase

import (
	"io"

	"github.com/xaligo/xaligo/internal/entity"
	v1engine "github.com/xaligo/xaligo/internal/usecase/v1/engine"
)

type ParserUsecase interface {
	Parse(io.Reader) (entity.Document, error)
}

type parserUsecase struct{}

func NewParserUsecase() ParserUsecase {
	return &parserUsecase{}
}

func (rcvr *parserUsecase) Parse(r io.Reader) (entity.Document, error) {
	return v1engine.ParseV1EngineParseDocument(r)
}

// Parse delegates DSL parsing to ParserUsecase.
// Deprecated: use NewParserUsecase().Parse instead.
func Parse(r io.Reader) (entity.Document, error) {
	return NewParserUsecase().Parse(r)
}
