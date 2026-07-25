package controller

import (
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

The version is taken from an embedded build-time value first, then the file
path named by the XALIGO_VERSION_FILE environment variable, then falls back
to the repository VERSION file, and finally to "dev" if none is available.`,
		Run: func(cmd *cobra.Command, args []string) {
			logger.INFO(ICVERSIONIVC002, "version", map[string]any{"version": resolvedVersion()})
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
