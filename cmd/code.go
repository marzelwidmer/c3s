/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/marzelwidmer/uzo/util"
)

// codeCmd represents the code command
var codeCmd = &cobra.Command{
	Use:   "code",
	Short: "open with vscode (code)",
	Long: `open with vscode (code)`,
	Run: func(cmd *cobra.Command, args []string) {
		var fileName string
		var err error
		var argument string

		argument = args[0]
		fileExists , err := util.FileExists(argument)
		
	},
}

func init() {
	rootCmd.AddCommand(codeCmd)
}
