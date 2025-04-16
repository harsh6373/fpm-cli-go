package helpers

import (
	"errors"
	"os"
	"path/filepath"
)

// ResolvePath returns the absolute path or current working directory if empty.
func ResolvePath(pathFlag string) string {
	if pathFlag == "" {
		cwd, _ := os.Getwd()
		return cwd
	}
	absPath, err := filepath.Abs(pathFlag)
	if err != nil {
		return pathFlag
	}
	return absPath
}

// FindFlutterProjectRoot traverses up to find the Flutter project root (where pubspec.yaml exists).
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
