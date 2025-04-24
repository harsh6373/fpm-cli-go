package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/AlecAivazis/survey/v2"
	"github.com/harsh6373/fpm-cli-go/boilerplate"
	"github.com/harsh6373/fpm-cli-go/utils"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new Flutter project with state management boilerplate",
	Run: func(cmd *cobra.Command, args []string) {
		var answers struct {
			ProjectName  string
			PackageName  string
			Description  string
			StateManager string
			ProjectPath  string
		}

		qs := []*survey.Question{
			{Name: "ProjectName", Prompt: &survey.Input{Message: "Enter project name:"}, Validate: survey.Required},
			{Name: "PackageName", Prompt: &survey.Input{Message: "Enter package name (e.g. com.example.app):"}, Validate: survey.Required},
			{Name: "Description", Prompt: &survey.Input{Message: "Project description:"}},
			{Name: "ProjectPath", Prompt: &survey.Input{Message: "Enter the full path where the project should be created:", Default: "./"}, Validate: survey.Required},
			{Name: "StateManager", Prompt: &survey.Select{Message: "Select a state management solution:", Options: []string{"GetX", "BLoC", "Provider", "Riverpod"}}},
		}

		if err := survey.Ask(qs, &answers); err != nil {
			fmt.Println("Prompt failed:", err)
			return
		}

		projectBase := answers.ProjectPath
		projectName := answers.ProjectName

		if err := utils.PrepareProjectDirectory(projectBase); err != nil {
			fmt.Println("❌", err)
			return
		}
		os.Chdir(projectBase)

		fmt.Println("🚀 Creating Flutter project...")

		flutterPath, err := exec.LookPath("flutter")
		if err != nil {
			fmt.Println("❌ 'flutter' command not found in PATH.")
			fmt.Println("💡 Add Flutter to your PATH or use the full path to the flutter binary.")
			return
		}
		fmt.Println("✅ Found flutter at:", flutterPath)

		createFlutterCmd := exec.Command(flutterPath, "create",
			"--org", answers.PackageName,
			"--project-name", projectName,
			"--description", answers.Description,
			projectName,
		)
		createFlutterCmd.Stdout = os.Stdout
		createFlutterCmd.Stderr = os.Stderr

		if err := createFlutterCmd.Run(); err != nil {
			fmt.Println("❌ Failed to create Flutter project:", err)
			return
		}

		if err := os.Chdir(projectName); err != nil {
			fmt.Println("❌ Failed to switch into project directory:", err)
			return
		}

		switch answers.StateManager {
		case "GetX":
			boilerplate.AddGetXBoilerplate()
		case "BLoC":
			boilerplate.AddBlocBoilerplate()
		case "Provider":
			boilerplate.AddProviderBoilerplate()
		case "Riverpod":
			boilerplate.AddRiverpodBoilerplate()
		default:
			fmt.Println("Invalid selection.")
		}

		fmt.Println("✅ Project setup complete.")
		utils.GenerateReadme(projectName, answers.Description)
		utils.SetupGit()
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
