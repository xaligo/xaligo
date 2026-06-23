package controller

import (
	"github.com/ryo-arima/xaligo/internal/share"
	"github.com/spf13/cobra"
)

var version = "0.1.0"

var (
	ICVERSIONIVC001 = share.NewMCode("ICVERSIONIVC-001", "Init version command start")
	ICVERSIONIVC002 = share.NewMCode("ICVERSIONIVC-002", "Init version command output version")
)

type VersionController struct{}

func NewVersionController() *VersionController { return &VersionController{} }

func (rcvr *VersionController) Command() *cobra.Command {
	logger.DEBUG(ICVERSIONIVC001, "start")
	return &cobra.Command{
		Use:   "version",
		Short: "Print xaligo version",
		Run: func(cmd *cobra.Command, args []string) {
			logger.INFO(ICVERSIONIVC002, "version", map[string]any{"version": version})
		},
	}
}
