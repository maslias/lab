package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/maslias/tbblibrary/jsondata"
)

var gettitleCmd = &cobra.Command{
	Use:   "title",
	Short: "list of titles",
	Long:  `list of all available titles from videos`,
	Run:   gettitleTask,
}

func init() {
	getCmd.AddCommand(gettitleCmd)
	gettitleCmd.Flags().StringP("titlename", "n", "", "list videos with in this category")
}

func gettitleTask(cmd *cobra.Command, args []string) {
	titleName, _ := cmd.Flags().GetString("titlename")

	vnModule := "nan"

	nodes := []jsondata.VideoNode{}

	for _, vn := range jd.VideoNodes {
		if vn.Title == titleName {
			vnModule = vn.Module
			nodes = append(nodes, vn)
		}
	}

	fmt.Printf("module: %s | %v\n\n", vnModule)
	fmt.Printf("title: %s | %v\n\n", titleName)

	for _, n := range nodes {
		fmt.Printf("- %s\n", n.Title)
		fmt.Printf("Link: %s\n\n", n.Wistia_link)
	}
}
