/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/maslias/chatroomer/pkg/common"
)

var rootCmd = &cobra.Command{
	Use:   "chatroomer",
	Short: "golang websocket chatsystem",
}

var Logger *common.HttpLog

func Execute() {
	Logger = common.NewHttpLog()
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
