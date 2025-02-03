package cmd

import (
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "list of categories",
	Long:  `list of all available categories from videos`,
}

func init() {
	rootCmd.AddCommand(getCmd)
}
