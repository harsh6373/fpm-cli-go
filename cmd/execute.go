package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// executeCmd represents the execute command
var executeCmd = &cobra.Command{
	Use:   "exec [command]",
	Short: "Execute common Flutter commands like clean, build, pub get",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		flutterCmd := exec.Command("flutter", args...)
		flutterCmd.Stdout = os.Stdout
		flutterCmd.Stderr = os.Stderr
		if err := flutterCmd.Run(); err != nil {
			fmt.Printf("Error executing Flutter command: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(executeCmd)
}
