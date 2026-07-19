package usecase_test

import (
	"bytes"
	"compress/zlib"
	"context"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/usecase"
)

func TestRenderSamplePDFRetainsTextAndAWSVectorIcons(t *testing.T) {
	exampleDir := filepath.Join(repoRoot(t), "docs", "src", "examples", "samples")
	source, err := os.ReadFile(filepath.Join(exampleDir, "sample.xal"))
	if err != nil {
		t.Fatal(err)
	}
	services, err := os.ReadFile(filepath.Join(exampleDir, "services.csv"))
	if err != nil {
		t.Fatal(err)
	}

	document, err := newUsecase().RenderPDF(context.Background(), source, entity.RenderOptions{
		Format:        usecase.FormatPDF,
		Theme:         "light",
		PxPerInch:     96,
		ServicesCSV:   services,
		Mode:          usecase.ModeNetwork,
		CombineFrames: true,
	})
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.HasPrefix(document, []byte("%PDF-")) || !bytes.Contains(document, []byte("/Font")) {
		t.Fatal("sample PDF is missing its PDF header or embedded text font")
	}

	inflated := inflateSamplePDFStreams(t, document)
	if count := strings.Count(inflated, " BT"); count < 100 {
		t.Fatalf("sample PDF text objects = %d, want at least 100", count)
	}
	// AWS data:image/svg+xml assets contain the curved service-icon paths.
	// The surrounding Xaligo diagram is predominantly rectangles and straight
	// connectors, so this threshold detects the prior image-dropping behavior.
	if count := strings.Count(inflated, " c"); count < 1000 {
		t.Fatalf("sample PDF cubic path commands = %d, want at least 1000 from AWS vector icons", count)
	}
}

func inflateSamplePDFStreams(t *testing.T, document []byte) string {
	t.Helper()
	pattern := regexp.MustCompile(`(?s)stream\r?\n(.*?)\r?\nendstream`)
	var output strings.Builder
	for _, match := range pattern.FindAllSubmatch(document, -1) {
		reader, err := zlib.NewReader(bytes.NewReader(match[1]))
		if err != nil {
			continue
		}
		data, readErr := io.ReadAll(reader)
		closeErr := reader.Close()
		if readErr != nil {
			t.Fatalf("read PDF stream: %v", readErr)
		}
		if closeErr != nil {
			t.Fatalf("close PDF stream: %v", closeErr)
		}
		output.Write(data)
	}
	return output.String()
}
