package controller

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ryo-arima/xaligo/internal/share"
	"github.com/spf13/cobra"
)

var (
	ICIIIC001 = share.NewMCode("ICIIIC-001", "Init init command start")
	ICIRI001  = share.NewMCode("ICIRI-001", "Run init create output directory failed")
	ICIRI002  = share.NewMCode("ICIRI-002", "Run init write sample DSL failed")
	ICIRI003  = share.NewMCode("ICIRI-003", "Run init created sample DSL")
)

const sampleDSL = `<frame width="1440" height="900" class="pa-4">
  <container class="pa-4">
    <row gap="20" class="mb-2">
      <col span="8" class="pa-2">
        <card title="Dashboard" />
      </col>
      <col span="4" class="pa-2">
        <card title="Summary" />
      </col>
    </row>

    <row gap="20">
      <col span="4" class="pa-2">
        <panel title="Filters" />
      </col>
      <col span="8" class="pa-2">
        <panel title="Main Chart" />
      </col>
    </row>
  </container>
</frame>
`

func InitInitCmd() *cobra.Command {
	logger.DEBUG(ICIIIC001, "start")
	var outputDir string
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Create starter xaligo template",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunInit(outputDir)
		},
	}
	cmd.Flags().StringVarP(&outputDir, "output", "o", ".", "output directory")
	return cmd
}

func RunInit(outputDir string) error {
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		logger.ERROR(ICIRI001, "create output directory failed", map[string]any{"outputDir": outputDir, "error": err})
		return fmt.Errorf("create output dir: %w", err)
	}
	path := filepath.Join(outputDir, "sample.xal")
	if err := os.WriteFile(path, []byte(sampleDSL), 0644); err != nil {
		logger.ERROR(ICIRI002, "write sample DSL failed", map[string]any{"path": path, "error": err})
		return fmt.Errorf("write sample DSL: %w", err)
	}
	logger.INFO(ICIRI003, "created", map[string]any{"path": path})
	return nil
}
