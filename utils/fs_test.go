package utils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestPrepareProjectDirectory(t *testing.T) {
	tmpDir := filepath.Join(os.TempDir(), "fpm_test")
	defer os.RemoveAll(tmpDir)

	err := PrepareProjectDirectory(tmpDir)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	_, err = os.Stat(tmpDir)
	if os.IsNotExist(err) {
		t.Errorf("Expected directory to exist: %s", tmpDir)
	}
}
