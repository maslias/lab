package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/maslias/tbblibrary/jsondata"
)

var getcatCmd = &cobra.Command{
	Use:   "cat",
	Short: "list of categories",
	Long:  `list of all available categories from videos`,
	Run:   getcatTask,
}

func init() {
	getCmd.AddCommand(getcatCmd)
	getcatCmd.Flags().StringP("modulename", "n", "", "list videos with in this category")
}

func getcatTask(cmd *cobra.Command, args []string) {
	moduleName, _ := cmd.Flags().GetString("modulename")
    vnModule := "nan"

	nodes := []jsondata.VideoNode{}

	for _, vn := range jd.VideoNodes {
		if vn.Module == moduleName {
            vnModule = vn.Module
			nodes = append(nodes, vn)
		}
	}

	fmt.Printf("module: %s | %v\n\n", vnModule, len(nodes))

	for _, n := range nodes {
		fmt.Printf("- %s\n", n.Title)
		fmt.Printf("Link: %s\n\n", n.Wistia_link)
	}
}
