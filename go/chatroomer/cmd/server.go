package cmd

import "github.com/spf13/cobra"

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "initialze a server",
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
