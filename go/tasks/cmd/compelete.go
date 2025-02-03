package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "complete a task",
	Long:  `complete a task with -id`,
	Run:   completeTask,
}

func completeTask(cmd *cobra.Command, args []string) {
	id, _ := cmd.Flags().GetInt("id")
	tasks, err := appDb.Tasks.Complete(id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(CreateTable(tasks))
}

func init() {
	rootCmd.AddCommand(completeCmd)
	completeCmd.Flags().IntP("id", "i", -1, "task id to complete the task")
	completeCmd.MarkFlagRequired("id")
}
