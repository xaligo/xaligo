//go:build js

package repository

import (
	"context"
	"errors"

	"github.com/xaligo/xaligo/internal/entity"
)

// PDFRepository is unavailable in the JavaScript runtime because the native
// implementation depends on the Canvas PDF encoder.
type PDFRepository interface {
	Export(context.Context, []entity.RenderPage) ([]byte, error)
}

type pdfRepository struct{}

func NewPDFRepository() PDFRepository {
	return &pdfRepository{}
}

func (rcvr *pdfRepository) Export(context.Context, []entity.RenderPage) ([]byte, error) {
	return nil, errors.New("export PDF: unavailable in JavaScript runtime")
}
