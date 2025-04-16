package utils

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/AlecAivazis/survey/v2"
)

func SetupGit() {
	var initGit bool
	survey.AskOne(&survey.Confirm{Message: "Initialize with Git?", Default: true}, &initGit)
	if !initGit {
		return
	}
	cmds := [][]string{
		{"git", "init"},
		{"git", "add", "."},
		{"git", "commit", "-m", "initial commit"},
		{"git", "branch", "-M", "main"},
	}
	for _, c := range cmds {
		cmd := exec.Command(c[0], c[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println("❌ Git error:", err)
			return
		}
	}
	var remote string
	survey.AskOne(&survey.Input{Message: "Remote URL (blank to skip):"}, &remote)
	if remote != "" {
		exec.Command("git", "remote", "add", "origin", remote).Run()
		exec.Command("git", "push", "-u", "origin", "main").Run()
	}
}
