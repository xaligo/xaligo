//go:build js

package repository

import (
	"context"
	"errors"

	"github.com/xaligo/xaligo/internal/entity"
)

// SpreadsheetRepository is unavailable in the JavaScript runtime because the
// native implementation depends on the Excelize and Canvas encoders.
type SpreadsheetRepository interface {
	Export(context.Context, []entity.RenderPage) ([]byte, error)
}

type spreadsheetRepository struct{}

func NewSpreadsheetRepository() SpreadsheetRepository {
	return &spreadsheetRepository{}
}

func (rcvr *spreadsheetRepository) Export(context.Context, []entity.RenderPage) ([]byte, error) {
	return nil, errors.New("export spreadsheet: unavailable in JavaScript runtime")
}
