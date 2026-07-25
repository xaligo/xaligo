package controller_test

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/xaligo/xaligo/internal/controller"
)

const versionCommandHelperEnv = "XALIGO_TEST_VERSION_COMMAND_HELPER"

func TestVersionCommandWritesStableStdoutAcrossLogFormats(t *testing.T) {
	versionFile := filepath.Join(t.TempDir(), "VERSION")
	if err := os.WriteFile(versionFile, []byte("1.2.3-test\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	for _, structured := range []string{"", "1"} {
		t.Run("structured="+structured, func(t *testing.T) {
			cmd := exec.Command(os.Args[0], "-test.run=^TestVersionCommandProcessHelper$")
			cmd.Env = append(os.Environ(),
				versionCommandHelperEnv+"=1",
				"XALIGO_VERSION_FILE="+versionFile,
				"XALIGO_HOME=",
				"XALIGO_LOG_LEVEL=",
				"XALIGO_LOG_OUTPUT=stdout",
				"XALIGO_LOG_STRUCTURED="+structured,
			)

			output, err := cmd.Output()
			if err != nil {
				t.Fatalf("version command: %v", err)
			}
			if got, want := string(output), "1.2.3-test\n"; got != want {
				t.Fatalf("stdout = %q, want %q", got, want)
			}
		})
	}
}

func TestVersionCommandProcessHelper(t *testing.T) {
	if os.Getenv(versionCommandHelperEnv) != "1" {
		return
	}
	if err := controller.NewVersionController().Command().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	os.Exit(0)
}

func TestVersionCommandFallbackOrder(t *testing.T) {
	explicitDir := t.TempDir()
	explicitPath := filepath.Join(explicitDir, "release-version")
	if err := os.WriteFile(explicitPath, []byte("explicit\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	homeDir := t.TempDir()
	if err := os.WriteFile(filepath.Join(homeDir, "VERSION"), []byte("home\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	workingDir := t.TempDir()
	if err := os.WriteFile(filepath.Join(workingDir, "VERSION"), []byte("working\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	childDir := filepath.Join(workingDir, "nested")
	if err := os.Mkdir(childDir, 0o755); err != nil {
		t.Fatal(err)
	}
	emptyDir := t.TempDir()

	tests := []struct {
		name        string
		versionFile string
		home        string
		dir         string
		want        string
	}{
		{
			name:        "explicit file wins",
			versionFile: explicitPath,
			home:        homeDir,
			dir:         childDir,
			want:        "explicit",
		},
		{
			name: "home wins over working tree",
			home: homeDir,
			dir:  childDir,
			want: "home",
		},
		{
			name: "working tree ancestors",
			dir:  childDir,
			want: "working",
		},
		{
			name:        "unreadable explicit file resolves to dev",
			versionFile: filepath.Join(explicitDir, "missing"),
			home:        homeDir,
			dir:         childDir,
			want:        "dev",
		},
		{
			name: "no version available",
			dir:  emptyDir,
			want: "dev",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cmd := exec.Command(os.Args[0], "-test.run=^TestVersionCommandProcessHelper$")
			cmd.Dir = test.dir
			cmd.Env = append(os.Environ(),
				versionCommandHelperEnv+"=1",
				"XALIGO_VERSION_FILE="+test.versionFile,
				"XALIGO_HOME="+test.home,
				"XALIGO_LOG_LEVEL=",
				"XALIGO_LOG_OUTPUT=stdout",
				"XALIGO_LOG_STRUCTURED=",
			)

			output, err := cmd.Output()
			if err != nil {
				t.Fatalf("version command: %v", err)
			}
			if got, want := strings.TrimSpace(string(output)), test.want; got != want {
				t.Fatalf("stdout = %q, want %q", got, want)
			}
		})
	}
}

func TestVersionCommandHelpDocumentsFallbackOrder(t *testing.T) {
	var output bytes.Buffer
	cmd := controller.NewVersionController().Command()
	cmd.SetOut(&output)
	cmd.SetArgs([]string{"--help"})

	if err := cmd.Execute(); err != nil {
		t.Fatal(err)
	}

	help := output.String()
	ordered := []string{
		"embedded build-time value",
		"XALIGO_VERSION_FILE",
		"$XALIGO_HOME/VERSION",
		"current directory or its parents",
		`"dev"`,
	}
	previous := -1
	for _, text := range ordered {
		index := strings.Index(help, text)
		if index < 0 {
			t.Fatalf("help does not contain %q:\n%s", text, help)
		}
		if index <= previous {
			t.Fatalf("help does not document fallback order at %q:\n%s", text, help)
		}
		previous = index
	}
}
