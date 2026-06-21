package share_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ryo-arima/xaligo/internal/share"
)

func TestMCodePaddingUsesRegisteredMaxLength(t *testing.T) {
	short := share.NewMCode("T1", "short")
	long := share.NewMCode("TEST-LONG-1", "long")

	if share.GetMaxCodeLength() < len(long.Code) {
		t.Fatalf("max code length = %d, want at least %d", share.GetMaxCodeLength(), len(long.Code))
	}
	if got := short.PaddedCode(); len(got) != share.GetMaxCodeLength() || !strings.HasPrefix(got, short.Code) {
		t.Fatalf("padded code = %q, max = %d", got, share.GetMaxCodeLength())
	}
}

func TestLoggerWritesStructuredEntry(t *testing.T) {
	path := filepath.Join(t.TempDir(), "xaligo.log")
	logger := share.NewLogger(share.LoggerConfig{
		Component:  "unit",
		Service:    "share",
		Level:      "debug",
		Structured: true,
		Output:     path,
	})

	logger.DEBUG(share.NewMCode("TLOG1", "debug message"), "detail", map[string]any{"answer": 42, "error": "boom"})

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	var entry share.LogEntry
	if err := json.Unmarshal([]byte(strings.TrimSpace(string(data))), &entry); err != nil {
		t.Fatalf("log entry is not JSON: %s: %v", data, err)
	}
	if entry.Level != "DEBUG" || entry.Code != "TLOG1" || entry.Component != "unit" || entry.Service != "share" {
		t.Fatalf("entry metadata = %#v", entry)
	}
	if entry.Message != "debug message: detail" || entry.Error != "boom" {
		t.Fatalf("entry message/error = %#v", entry)
	}
	if entry.Fields["answer"] != float64(42) {
		t.Fatalf("entry fields = %#v", entry.Fields)
	}
}

func TestLoggerFiltersBelowConfiguredLevel(t *testing.T) {
	path := filepath.Join(t.TempDir(), "xaligo.log")
	logger := share.NewLogger(share.LoggerConfig{Level: "error", Output: path})

	logger.INFO(share.NewMCode("TFILTER1", "ignore"), "")
	logger.ERROR(share.NewMCode("TFILTER2", "write"), "")

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	text := string(data)
	if strings.Contains(text, "TFILTER1") || !strings.Contains(text, "TFILTER2") {
		t.Fatalf("filtered output = %q", text)
	}
}
