package helpers

import (
	"fmt"
	"os"
	"os/exec"
)

// CreateFlutterProject initializes a Flutter project in the specified directory
func CreateFlutterProject(projectPath string) error {
	// Check if the project path exists
	_, err := os.Stat(projectPath)
	if err != nil && os.IsNotExist(err) {
		// If the directory doesn't exist, create it
		if err := os.MkdirAll(projectPath, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create directory %s: %v", projectPath, err)
		}
	}

	// Initialize a new Flutter project
	cmd := exec.Command("flutter", "create", ".")
	cmd.Dir = projectPath
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to create Flutter project in %s: %v", projectPath, err)
	}

	fmt.Println("✅ Flutter project created successfully at:", projectPath)
	return nil
}
