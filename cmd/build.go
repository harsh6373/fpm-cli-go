package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/harsh6373/fpm-cli-go/utils"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build a Flutter project with environment, version, and build type",
	Run: func(cmd *cobra.Command, args []string) {
		pathFlag, _ := cmd.Flags().GetString("path")
		projectPath := pathFlag
		if projectPath == "" {
			cwd, _ := os.Getwd()
			projectPath = cwd
		}

		flutterRoot, err := utils.FindFlutterProjectRoot(projectPath)
		if err != nil {
			fmt.Println("❌", err)
			return
		}

		if err := os.Chdir(flutterRoot); err != nil {
			fmt.Println("❌ Failed to switch to project directory:", err)
			return
		}

		projectName := filepath.Base(flutterRoot)

		var answers struct {
			Env     string
			Type    string
			Version string
			Build   string
		}

		survey.Ask([]*survey.Question{
			{Name: "Env", Prompt: &survey.Select{Message: "Select environment:", Options: []string{"DEVELOPMENT", "STAGING", "PRODUCTION"}}},
			{Name: "Type", Prompt: &survey.Select{Message: "Build type:", Options: []string{"APK", "App Bundle"}}},
			{Name: "Version", Prompt: &survey.Input{Message: "Enter version (e.g. 1.0.0):"}, Validate: survey.Required},
			{Name: "Build", Prompt: &survey.Input{Message: "Enter build number (e.g. 5):"}, Validate: survey.Required},
		}, &answers)

		outputName := utils.GenerateArtifactName(projectName, answers.Env, answers.Version, answers.Build)
		buildsDir := filepath.Join(flutterRoot, "builds")
		_ = os.MkdirAll(buildsDir, os.ModePerm)

		var outExt string
		var buildArgs []string
		if answers.Type == "APK" {
			outExt = ".apk"
			buildArgs = []string{"build", "apk", "--release", "--no-tree-shake-icons", fmt.Sprintf("--dart-define=ENVIRONMENT=%s", answers.Env)}
		} else {
			outExt = ".aab"
			buildArgs = []string{"build", "appbundle", "--release", "--no-tree-shake-icons", fmt.Sprintf("--dart-define=ENVIRONMENT=%s", answers.Env)}
		}

		fmt.Println("🚀 Running Flutter build...")
		cmdBuild := exec.Command("flutter", buildArgs...)
		cmdBuild.Stdout = os.Stdout
		cmdBuild.Stderr = os.Stderr
		if err := cmdBuild.Run(); err != nil {
			fmt.Println("❌ Build failed:", err)
			return
		}

		var buildOutputPath string
		if answers.Type == "APK" {
			buildOutputPath = filepath.Join("build", "app", "outputs", "flutter-apk", "app-release.apk")
		} else {
			buildOutputPath = filepath.Join("build", "app", "outputs", "bundle", "release", "app-release.aab")
		}

		finalOutputPath := filepath.Join(buildsDir, outputName+outExt)
		if err := os.Rename(buildOutputPath, finalOutputPath); err != nil {
			fmt.Println("❌ Failed to move build artifact:", err)
			return
		}

		fmt.Println("✅ Build completed: ", finalOutputPath)
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
	buildCmd.Flags().String("path", "", "Path to the Flutter project (optional)")
}
