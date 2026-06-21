package repository

import (
	"encoding/csv"
	"io"
	"strings"
)

const svgDataURLPrefix = "data:image/svg+xml;base64,"

func newCatalogCSVReader(r io.Reader) *csv.Reader {
	reader := csv.NewReader(r)
	reader.Comment = '#'
	reader.FieldsPerRecord = -1
	return reader
}

func svgDataURL(raw string) string {
	raw = strings.TrimSpace(raw)
	if raw == "" || strings.HasPrefix(raw, "data:") {
		return raw
	}
	return svgDataURLPrefix + raw
}
