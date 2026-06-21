package repository

import (
	"encoding/csv"
	"io"
)

func newCatalogCSVReader(r io.Reader) *csv.Reader {
	reader := csv.NewReader(r)
	reader.Comment = '#'
	reader.FieldsPerRecord = -1
	return reader
}
