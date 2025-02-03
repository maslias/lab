package cmd

import (
	"github.com/spf13/cobra"

	"github.com/maslias/chatroomer/internal/wire"
)

var serverwebCmd = &cobra.Command{
	Use:   "web",
	Short: "initialize frontend webserver for chatrommer",
	Run:   serverwebTask,
}

func init() {
	serverCmd.AddCommand(serverwebCmd)
}

func serverwebTask(cmd *cobra.Command, args []string) {
	webserver, err := wire.InitializeWeb()
	if err != nil {
		Logger.Errorw("error - webserver wire init","error", err)
	}

	if err := webserver.Run(); err != nil {
		Logger.Errorw("error - webserver crashed", "error", err)
	}
}
