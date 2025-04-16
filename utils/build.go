package helpers

import (
	"os"
	"os/exec"
)

// GenerateFlutterBuildArgs returns the build arguments and output extension based on the build type.
func GenerateFlutterBuildArgs(buildType string, env string) ([]string, string) {
	var buildArgs []string
	var outExt string

	switch buildType {
	case "APK":
		buildArgs = []string{"build", "apk", "--release", "--no-tree-shake-icons", "--dart-define=ENVIRONMENT=" + env}
		outExt = ".apk"
	case "App Bundle":
		buildArgs = []string{"build", "appbundle", "--release", "--no-tree-shake-icons", "--dart-define=ENVIRONMENT=" + env}
		outExt = ".aab"
	default:
		buildArgs = []string{}
		outExt = ""
	}
	return buildArgs, outExt
}

// RunFlutterBuild executes the flutter build command.
func RunFlutterBuild(buildArgs []string) error {
	cmd := exec.Command("flutter", buildArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// GetBuildOutputPath returns the default path where Flutter outputs the built APK or AAB.
func GetBuildOutputPath(buildType string) string {
	switch buildType {
	case "APK":
		return "build/app/outputs/flutter-apk/app-release.apk"
	case "App Bundle":
		return "build/app/outputs/bundle/release/app-release.aab"
	default:
		return ""
	}
}
