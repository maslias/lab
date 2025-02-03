package cmd

import (
	"github.com/spf13/cobra"

	"github.com/maslias/chatroomer/internal/wire"
)

var serverhealthCmd = &cobra.Command{
	Use:   "health",
	Short: "initialize health http server",
	Run:   serverhealthTask,
}

func init() {
	serverCmd.AddCommand(serverhealthCmd)
}

func serverhealthTask(cmd *cobra.Command, agrs []string) {
	healthserver, err := wire.InitliazeHealth()
	if err != nil {
		Logger.Errorw("error - healthserver wire init", "error", err)
	}

	if err := healthserver.Run(); err != nil {
		Logger.Errorw("error - healthserver crashed", "error", err)
	}
}
