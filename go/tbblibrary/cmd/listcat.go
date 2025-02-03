package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listcatCmd = &cobra.Command{
	Use:   "cat",
	Short: "list of categories",
	Long:  `list of all available categories from videos`,
	Run:   listcatTask,
}

func init() {
	listCmd.AddCommand(listcatCmd)
}

func listcatTask(cmd *cobra.Command, args []string) {
	// isF, _ := cmd.Flags().GetBool("finished")

	out := []string{}

	for _, vn := range jd.VideoNodes {
		out = append(out, string(vn.Module))
	}

	out = RemoveDuplicateStr(out)

	for _, c := range out {
		fmt.Printf("%s\n", c)
	}
}
