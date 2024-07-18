/*
Copyright © 2024 thaidmfinnick
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sw",
	Short: "Starwars API inside command line",
	Long: `Starwars API helps you can interact with Starwars data easily with command line.

It stills in development.
	`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Log all data")
	rootCmd.PersistentFlags().IntP("limit", "l", -1, "Limit data show in terminal")
}
