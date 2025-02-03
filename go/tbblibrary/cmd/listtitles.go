package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listtitleCmd = &cobra.Command{
	Use:   "title",
	Short: "list of titles",
	Long:  `list of all available titles from videos`,
	Run:   listtitleTask,
}

func init() {
	listCmd.AddCommand(listtitleCmd)
}

func listtitleTask(cmd *cobra.Command, args []string) {
	// isF, _ := cmd.Flags().GetBool("finished")

	out := []string{}

	for _, vn := range jd.VideoNodes {
		out = append(out, string(vn.Title))
	}

	out = RemoveDuplicateStr(out)

	for _, c := range out {
		fmt.Printf("%s\n", c)
	}
}
