package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list of your tasks",
	Long:  `list all your task. filter unfinished -u or finished -f task`,
	Run:   listTask,
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP("unfinished", "u", false, "show only unfinished tasks")
	listCmd.Flags().BoolP("finished", "f", false, "show only finished tasks")
}


func listTask(cmd *cobra.Command, args []string) {
	 isU, _ := cmd.Flags().GetBool("unfinished")
	 isF, _ := cmd.Flags().GetBool("finished")

	 tasks, err := appDb.Tasks.All(isU, isF)
	 if err != nil {
	 	log.Fatal(err)
	 }
	

    fmt.Println(CreateTable(tasks))

}
