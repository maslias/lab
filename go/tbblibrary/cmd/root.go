/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/maslias/tbblibrary/jsondata"
)

var jd = &jsondata.JsonData{}

var rootCmd = &cobra.Command{
	Use:   "tbblibrary",
	Short: "video library of transforming basketball",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	jd = jsondata.NewJsonData()

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
