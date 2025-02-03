package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add new task",
	Long:  `add a new task with -t <title>, optional with -d <details> and -f <terminate date>, default target date is 7+ days from created date`,
	Run:   addTask,
}

func addTask(cmd *cobra.Command, args []string) {
	title, _ := cmd.Flags().GetString("title")
	details, _ := cmd.Flags().GetString("details")
	terminated,_ := cmd.Flags().GetInt("terminated")


    tasks, err := appDb.Tasks.Add(title, details, terminated)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(CreateTable(tasks))


}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("title", "t", "", "title of task")
	addCmd.Flags().StringP("details", "d", "", "details of task")
	addCmd.Flags().IntP("terminated", "f", 7, "terminate task in days")
	addCmd.MarkFlagRequired("title")
}
