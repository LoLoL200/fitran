/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fitran/config"

	"github.com/spf13/cobra"
)

// historyCmd represents the history command
var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "Display all file actions",
	Run:   history,
}

func init() {
	rootCmd.AddCommand(historyCmd)
}
func history(cmd *cobra.Command, args []string) {
	rootDir := "./"
	config.HistoryFiles(rootDir)
}
