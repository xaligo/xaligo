package repository_test

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ryo-arima/xaligo/internal/repository"
)

type failingReader struct{}

func (failingReader) Read([]byte) (int, error) {
	return 0, errors.New("read failed")
}

func TestReadServiceListFromReaderParsesSupportedForms(t *testing.T) {
	entries, err := repository.ReadServiceListFromReader(strings.NewReader(`
# comment

Amazon S3
27,Amazon EC2
Amazon RDS,Database
117,Amazon RDS,RDS,Relational database,Primary database,
Custom Service,CS,Custom summary
`))
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 5 {
		t.Fatalf("entries = %d, want 5: %#v", len(entries), entries)
	}
	checks := []struct {
		index        int
		catalogID    int
		officialName string
		abbreviation string
	}{
		{0, 0, "Amazon S3", ""},
		{1, 27, "Amazon EC2", ""},
		{2, 0, "Amazon RDS", ""},
		{3, 117, "Amazon RDS", "RDS"},
		{4, 0, "Custom Service", "CS"},
	}
	for _, check := range checks {
		entry := entries[check.index]
		if entry.CatalogID != check.catalogID || entry.OfficialName != check.officialName || entry.Abbreviation != check.abbreviation {
			t.Fatalf("entry[%d] = %#v", check.index, entry)
		}
	}
}

func TestReadServiceListUsesFilePath(t *testing.T) {
	path := filepath.Join(t.TempDir(), "services.csv")
	if err := osWriteFile(path, "1020,Amazon S3,S3\n"); err != nil {
		t.Fatal(err)
	}
	entries, err := repository.ReadServiceList(path)
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 1 || entries[0].CatalogID != 1020 || entries[0].Abbreviation != "S3" {
		t.Fatalf("entries = %#v", entries)
	}
}

func TestReadServiceListErrors(t *testing.T) {
	if _, err := repository.ReadServiceList(filepath.Join(t.TempDir(), "missing.csv")); err == nil {
		t.Fatal("ReadServiceList missing file error = nil")
	}
	if _, err := repository.ReadServiceListFromReader(failingReader{}); err == nil {
		t.Fatal("ReadServiceListFromReader scanner error = nil")
	}
}

func osWriteFile(path, data string) error {
	return os.WriteFile(path, []byte(data), 0644)
}