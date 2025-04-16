package utils

import (
	"os"
	"strings"
	"testing"
)

func TestGenerateReadme(t *testing.T) {
	projectName := "TestProject"
	description := "This is a test project"

	GenerateReadme(projectName, description)

	data, err := os.ReadFile("README.md")
	if err != nil {
		t.Fatalf("Expected README.md to be created, got error: %v", err)
	}

	content := string(data)
	if !strings.Contains(content, projectName) {
		t.Errorf("README.md content does not contain project name: %s", projectName)
	}
	if !strings.Contains(content, description) {
		t.Errorf("README.md content does not contain description: %s", description)
	}

	_ = os.Remove("README.md")
}
