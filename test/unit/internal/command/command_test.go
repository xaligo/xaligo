package command_test

import (
	"os"
	"os/exec"
	"testing"

	command "github.com/ryo-arima/xaligo/internal"
	"github.com/spf13/cobra"
)

func TestNewRootCmdAssemblesSubcommands(t *testing.T) {
	for _, cmd := range []*cobra.Command{command.NewRootCmd()} {
		if cmd.Use != "xaligo" || cmd.Short == "" || len(cmd.Commands()) < 7 {
			t.Fatalf("root command = use %q short %q subcommands %d", cmd.Use, cmd.Short, len(cmd.Commands()))
		}
		seen := map[string]bool{}
		for _, sub := range cmd.Commands() {
			seen[sub.Name()] = true
		}
		for _, name := range []string{"render", "validate", "serve", "init", "version", "add", "generate"} {
			if !seen[name] {
				t.Fatalf("subcommand %q missing from %#v", name, seen)
			}
		}
	}
}

func TestExecuteRunsRootCommand(t *testing.T) {
	if os.Getenv("XALIGO_TEST_EXECUTE_FAIL") == "1" {
		os.Args = []string{"xaligo", "unknown"}
		command.Execute()
		return
	}

	testBinary := os.Args[0]
	oldArgs := os.Args
	os.Args = []string{"xaligo", "version"}
	t.Cleanup(func() { os.Args = oldArgs })

	command.Execute()

	cmd := exec.Command(testBinary, "-test.run", "^TestExecuteRunsRootCommand$")
	cmd.Env = append(os.Environ(), "XALIGO_TEST_EXECUTE_FAIL=1")
	err := cmd.Run()
	if exitErr, ok := err.(*exec.ExitError); !ok || exitErr.ExitCode() != 1 {
		t.Fatalf("execute failure subprocess err = %v", err)
	}
}
