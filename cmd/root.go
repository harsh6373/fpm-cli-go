package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command
var rootCmd = &cobra.Command{
	Use:   "flutter-cli-manager",
	Short: "A CLI tool to manage Flutter projects",
	Long:  `flutter-cli-manager is a CLI application written in Go to streamline Flutter project management tasks.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
