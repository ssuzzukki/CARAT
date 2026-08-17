/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)


const version = "0.1.0"


// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "carat",
	Short: "Compile Any Ruby Application Tool",
	Long: "CARAT converts Ruby applications into standalone executables.",
	Version: version,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}


/*
//まだ使わない
func init() {

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
*/