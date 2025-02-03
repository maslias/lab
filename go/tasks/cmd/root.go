package cmd

import (
	"os"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/maslias/tasks/db"
	"github.com/maslias/tasks/db/models"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tasks",
	Short: "a minimal task manager",
	Long:  `a minimal task mananger with -list -add and -remove`,
}


var appDb = &db.AppDb{}

func CreateTable(tasks []models.Task) *table.Table{

	rows := [][]string{}

	for _, t := range tasks {
		id := strconv.Itoa(t.Id)
		title := t.Title
		details := ""
		if t.Details.Valid {
			details = t.Details.String
		}
		daysLeft := int(t.TerminatedAt.Sub(t.CreatedAt).Hours() / 24)
		doneAt := "today"

		if daysLeft > 0 {
			doneAt = strconv.Itoa(daysLeft) + ` days left`
		} else if daysLeft < 0 {
			doneAt = strconv.Itoa(daysLeft) + ` days expired`
		}

		if t.DoneAt.Valid {
			doneAt = `done`
		}

		rows = append(rows, []string{id, title, details, doneAt})

	}

	return table.New().
		StyleFunc(func(row, col int) lipgloss.Style {
			if row == 0 {
				return lipgloss.NewStyle().Align(lipgloss.Center)
			} else {
				switch {
				case strings.Contains(rows[row-1][3], "today"):
                    return lipgloss.NewStyle().Padding(0, 1).Width(14).Foreground(lipgloss.Color("#f1ff5e"))
				case strings.Contains(rows[row-1][3], "left"):
                    return lipgloss.NewStyle().Padding(0, 1).Width(14)
				case strings.Contains(rows[row-1][3], "expired"):
					return lipgloss.NewStyle().Padding(0, 1).Width(14).Foreground(lipgloss.Color("#ff6e5e"))
				default:
                    return lipgloss.NewStyle().Padding(0, 1).Width(14).Foreground(lipgloss.Color("#5eff6c"))
				}
			}
		}).
        Width(100).
		Headers("id", "title", "details", "timer").
		Rows(rows...)


}

func Execute() {

    appDb.Execute()

    err := rootCmd.Execute()
    if err != nil {
        os.Exit(1)
    }
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
