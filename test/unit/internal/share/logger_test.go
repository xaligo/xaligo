package share_test

import (
	"encoding/json"
	"os"
	"os/exec"
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

func TestLoggerEnvCallerAndOutputBranches(t *testing.T) {
	t.Setenv("XALIGO_LOG_LEVEL", "warning")
	t.Setenv("XALIGO_LOG_STRUCTURED", "yes")
	t.Setenv("XALIGO_LOG_CALLER", "on")
	path := filepath.Join(t.TempDir(), "env.log")
	t.Setenv("XALIGO_LOG_OUTPUT", path)

	logger := share.NewEnvLogger("component", "service")
	logger.INFO(share.NewMCode("TENV1", "skip"), "")
	logger.WARN(share.NewMCode("TENV2", "warn"), "", map[string]any{"error": os.ErrNotExist})

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(data), "TENV1") {
		t.Fatalf("info log was not filtered: %s", data)
	}
	var entry share.LogEntry
	if err := json.Unmarshal([]byte(strings.TrimSpace(string(data))), &entry); err != nil {
		t.Fatalf("env log entry is not JSON: %s: %v", data, err)
	}
	if entry.Level != "WARN" || entry.Component != "component" || entry.Service != "service" || entry.File == "" || entry.Function == "" || entry.Error == "" {
		t.Fatalf("env log entry = %#v", entry)
	}

	share.NewLogger(share.LoggerConfig{Level: "fatal", Output: filepath.Join(t.TempDir(), "missing", "xaligo.log")}).ERROR(share.NewMCode("TENV3", "fallback"), "")
	share.NewLogger(share.LoggerConfig{Level: "mystery", Output: "stderr"}).WARN(share.NewMCode("TENV4", "stderr"), "")
}

func TestLoggerFatalExits(t *testing.T) {
	if os.Getenv("XALIGO_TEST_FATAL") == "1" {
		share.FATAL(share.NewMCode("TFATAL", "fatal"), "detail")
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run", "^TestLoggerFatalExits$")
	cmd.Env = append(os.Environ(), "XALIGO_TEST_FATAL=1")
	err := cmd.Run()
	if exitErr, ok := err.(*exec.ExitError); !ok || exitErr.ExitCode() != 1 {
		t.Fatalf("fatal subprocess err = %v", err)
	}
}

func TestPackageLevelLoggerHelpers(t *testing.T) {
	share.DEBUG(share.NewMCode("TPKGD", "debug"), "detail")
	share.INFO(share.NewMCode("TPKGI", "info"), "detail")
	share.WARN(share.NewMCode("TPKGW", "warn"), "detail")
	share.ERROR(share.NewMCode("TPKGE", "error"), "detail")
	if got := share.Mcode(share.NewMCode("TPKGM", "message")); got.Code != "TPKGM" || got.Message != "message" {
		t.Fatalf("Mcode = %#v", got)
	}
}
