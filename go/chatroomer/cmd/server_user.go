package cmd

import (
	"github.com/spf13/cobra"

	"github.com/maslias/chatroomer/internal/wire"
)

var serveruserCmd = &cobra.Command{
	Use:   "user",
	Short: "initialize user grpc server",
	Run:   serveruserTask,
}

func init() {
	serverCmd.AddCommand(serveruserCmd)
}

func serveruserTask(cmd *cobra.Command, args []string) {
	userserver, err := wire.InitializeUser()
	if err != nil {
		Logger.Errorw("error - userserver wire init", "error", err)
	}

	if err := userserver.Run(); err != nil {
		Logger.Errorw("error - userserver crashed", "error", err)
	}
}
