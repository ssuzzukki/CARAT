/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build [file]",
	Short: "Build Ruby application",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Build command received: %s\n", args[0])
	},
}

func init() {
	buildCmd.Flags().StringP(
		"output", 
		"o", 
		"", 
		"Output executable", 
	)

	buildCmd.Flags().StringP(
		"icon", 
		"i", 
		"", 
		"Application icon", 
	)
	
	rootCmd.AddCommand(buildCmd)

}
