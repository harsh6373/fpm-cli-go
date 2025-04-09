package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create [project_name]",
	Short: "Create a new Flutter project with Git, README and .gitignore",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		project := args[0]

		// Step 1: flutter create
		createCmd := exec.Command("flutter", "create", project)
		createCmd.Stdout = os.Stdout
		createCmd.Stderr = os.Stderr
		if err := createCmd.Run(); err != nil {
			fmt.Printf("❌ Failed to create Flutter project: %v\n", err)
			return
		}
		fmt.Println("✅ Flutter project created.")

		// Change working dir to new project
		projectPath := filepath.Join(".", project)
		if err := os.Chdir(projectPath); err != nil {
			fmt.Printf("❌ Failed to change to project directory: %v\n", err)
			return
		}

		// Step 2: Create README.md
		readmeContent := fmt.Sprintf("# %s\n\nCreated with fpm-cli-go 🚀", project)
		if err := os.WriteFile("README.md", []byte(readmeContent), 0644); err != nil {
			fmt.Println("⚠️ Failed to write README.md:", err)
		} else {
			fmt.Println("✅ README.md created.")
		}

		// Step 3: Create .gitignore
		gitignoreContent := `# Flutter/Dart/Pub related
.dart_tool/
.flutter-plugins
.packages
.pub-cache/
.pub/
build/
ios/Flutter/Flutter.framework
ios/Flutter/Flutter.podspec
coverage/
*.log
*.tmp
*.swp
*.lock
.idea/
*.iml
*.ipr
*.iws
.vscode/
`
		if err := os.WriteFile(".gitignore", []byte(gitignoreContent), 0644); err != nil {
			fmt.Println("⚠️ Failed to write .gitignore:", err)
		} else {
			fmt.Println("✅ .gitignore created.")
		}

		// Step 4: git init
		gitInitCmd := exec.Command("git", "init")
		gitInitCmd.Stdout = os.Stdout
		gitInitCmd.Stderr = os.Stderr
		if err := gitInitCmd.Run(); err != nil {
			fmt.Println("⚠️ Failed to initialize Git repo:", err)
		} else {
			fmt.Println("✅ Git repository initialized.")
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
