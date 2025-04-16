package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generate shell completion scripts",
	Long: `To load completions:

Bash:
  $ source <(flutter-cli completion bash)
  $ flutter-cli completion bash > /etc/bash_completion.d/flutter-cli

Zsh:
  $ source <(flutter-cli completion zsh)
  $ flutter-cli completion zsh > "${fpath[1]}/_flutter-cli"

Fish:
  $ flutter-cli completion fish | source
  $ flutter-cli completion fish > ~/.config/fish/completions/flutter-cli.fish

PowerShell:
  PS> flutter-cli completion powershell | Out-String | Invoke-Expression
  PS> flutter-cli completion powershell > flutter-cli.ps1`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactValidArgs(1),
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			rootCmd.GenBashCompletion(os.Stdout)
		case "zsh":
			rootCmd.GenZshCompletion(os.Stdout)
		case "fish":
			rootCmd.GenFishCompletion(os.Stdout, true)
		case "powershell":
			rootCmd.GenPowerShellCompletionWithDesc(os.Stdout)
		default:
			fmt.Fprintf(os.Stderr, "Unsupported shell: %s\n", args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
