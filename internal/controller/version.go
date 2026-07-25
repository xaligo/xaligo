package controller

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/xaligo/xaligo/internal/share"
)

var version string

var (
	ICVERSIONIVC001 = share.NewMCode("ICVERSIONIVC-001", "Init version command start")
	ICVERSIONIVC002 = share.NewMCode("ICVERSIONIVC-002", "Init version command output version")
)

type VersionController interface {
	Command() *cobra.Command
}

type versionController struct{}

func NewVersionController() VersionController { return &versionController{} }

func (rcvr *versionController) Command() *cobra.Command {
	logger.DEBUG(ICVERSIONIVC001, "start")
	return &cobra.Command{
		Use:   "version",
		Short: "Print xaligo version",
		Long: `Print the resolved xaligo build/release version.

The version is resolved in this order:
  1. the embedded build-time value;
  2. the explicit file named by XALIGO_VERSION_FILE, when set;
  3. $XALIGO_HOME/VERSION;
  4. the nearest VERSION file in the current directory or its parents; and
  5. "dev" when no version is available.

An unreadable XALIGO_VERSION_FILE is an explicit override and resolves to
"dev" instead of consulting lower-priority VERSION files.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			resolved := resolvedVersion()
			logger.DEBUG(ICVERSIONIVC002, "version", map[string]any{"version": resolved})
			if _, err := fmt.Fprintln(cmd.OutOrStdout(), resolved); err != nil {
				return fmt.Errorf("write version output: %w", err)
			}
			return nil
		},
	}
}

func resolvedVersion() string {
	if version != "" {
		return version
	}
	if version, ok := readVersionFile(); ok {
		return version
	}
	return "dev"
}

func readVersionFile() (string, bool) {
	if path := os.Getenv("XALIGO_VERSION_FILE"); path != "" {
		return readVersionPath(path)
	}
	if home := os.Getenv("XALIGO_HOME"); home != "" {
		if version, ok := readVersionPath(filepath.Join(home, "VERSION")); ok {
			return version, true
		}
	}
	dir, err := os.Getwd()
	if err != nil {
		return "", false
	}
	for {
		if version, ok := readVersionPath(filepath.Join(dir, "VERSION")); ok {
			return version, true
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", false
		}
		dir = parent
	}
}

func readVersionPath(path string) (string, bool) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", false
	}
	version := strings.TrimSpace(string(data))
	return version, version != ""
}
