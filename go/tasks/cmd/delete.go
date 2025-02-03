
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete a task",
	Long:  `delete a task with -id`,
	Run:   deleteTask,
}

func deleteTask(cmd *cobra.Command, args []string) {
	id, _ := cmd.Flags().GetInt("id")
	tasks, err := appDb.Tasks.Delete(id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(CreateTable(tasks))
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().IntP("id", "i", -1, "task id to delete the task")
	deleteCmd.MarkFlagRequired("id")
}
