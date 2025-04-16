package utils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func PrepareProjectDirectory(path string) error {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create base path: %w", err)
	}
	testFile := filepath.Join(path, ".fpm_write_test")
	if err := os.WriteFile(testFile, []byte("test"), 0644); err != nil {
		return errors.New("cannot write to the specified path")
	}
	os.Remove(testFile)
	return nil
}

func FindFlutterProjectRoot(startDir string) (string, error) {
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
