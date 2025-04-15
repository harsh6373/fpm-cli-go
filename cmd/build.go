package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
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

		flutterRoot, err := findFlutterProjectRoot(projectPath)
		if err != nil {
			fmt.Println("❌", err)
			return
		}

		// Change to Flutter project root
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
			{
				Name: "Env",
				Prompt: &survey.Select{
					Message: "Select environment:",
					Options: []string{"DEVELOPMENT", "STAGING", "PRODUCTION"},
				},
			},
			{
				Name: "Type",
				Prompt: &survey.Select{
					Message: "Build type:",
					Options: []string{"APK", "App Bundle"},
				},
			},
			{
				Name:     "Version",
				Prompt:   &survey.Input{Message: "Enter version (e.g. 1.0.0):"},
				Validate: survey.Required,
			},
			{
				Name:     "Build",
				Prompt:   &survey.Input{Message: "Enter build number (e.g. 5):"},
				Validate: survey.Required,
			},
		}, &answers)

		timestamp := time.Now().Format("060102") // YYMMDD
		outputName := fmt.Sprintf("%s_%s_v%s+%s_%s", strings.ToLower(projectName), strings.ToLower(answers.Env), answers.Version, answers.Build, timestamp)
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

		// Move final build to builds/ directory
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

func findFlutterProjectRoot(startDir string) (string, error) {
	absStart, err := filepath.Abs(startDir)
	if err != nil {
		return "", err
	}
	for dir := absStart; dir != "/" && dir != "."; dir = filepath.Dir(dir) {
		if _, err := os.Stat(filepath.Join(dir, "pubspec.yaml")); err == nil {
			return dir, nil
		}
	}
	return "", errors.New("flutter project root not found (missing pubspec.yaml)")
}

func init() {
	rootCmd.AddCommand(buildCmd)
	buildCmd.Flags().String("path", "", "Path to the Flutter project (optional)")
}
