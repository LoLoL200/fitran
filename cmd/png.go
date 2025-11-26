/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fitran/convert"
	"fmt"

	"github.com/spf13/cobra"
)

var pngCmd = &cobra.Command{
	Use:   "png",
	Short: "Convert with JPG in PNG",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("============List================")
		convert.FileInPNG()
	},
}

func init() {
	rootCmd.AddCommand(pngCmd)
}
