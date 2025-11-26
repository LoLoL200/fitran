/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fitran/convert"
	"fmt"

	"github.com/spf13/cobra"
)

var jpgCmd = &cobra.Command{
	Use:   "jpg",
	Short: `Convert with PNG in JPG`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("============List================")
		convert.FileInJPG()
	},
}

func init() {
	rootCmd.AddCommand(jpgCmd)
}
