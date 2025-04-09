package cmd

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

var ideCmd = &cobra.Command{
	Use:   "ide [ide_name] [project_path]",
	Short: "Open the Flutter project in the specified IDE (e.g., vscode, androidstudio)",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		ide := args[0]
		path := args[1]

		var openCmd *exec.Cmd

		switch ide {
		case "vscode":
			openCmd = exec.Command("code", path)
		case "androidstudio":
			if runtime.GOOS == "darwin" {
				openCmd = exec.Command("open", "-a", "Android Studio", path)
			} else {
				openCmd = exec.Command("studio", path)
			}
		default:
			fmt.Println("Unsupported IDE. Use 'vscode' or 'androidstudio'")
			return
		}

		if err := openCmd.Start(); err != nil {
			fmt.Printf("Failed to open IDE: %v\n", err)
			return
		}

		fmt.Println("✅ Project opened in", ide)
	},
}

func init() {
	rootCmd.AddCommand(ideCmd)
}
