package cmd

import (
	"github.com/spf13/cobra"

	"github.com/maslias/chatroomer/internal/wire"
)

var serverforwarderCmd = &cobra.Command{
	Use:   "forwarder",
	Short: "initialize forwarder http and grpc server",
	Run:   serverforwarderTask,
}

func init() {
	serverCmd.AddCommand(serverforwarderCmd)
}

func serverforwarderTask(cmd *cobra.Command, args []string) {
	forwarderserver, err := wire.InitializeForwarder()
	if err != nil {
		Logger.Errorw("error - forwarder wire init", "error", err)
	}

	if err := forwarderserver.Run(); err != nil {
		Logger.Errorw("error - forwarder crashed", "error", err)
	}



}
