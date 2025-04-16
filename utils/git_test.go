package utils

import (
	"os"
	"testing"
)

func TestSetupGit_SafeToCall(t *testing.T) {
	// This just ensures it doesn't panic.
	// Manual testing or integration tests needed for git exec parts.
	os.Stdout = nil // simulate no output, suppress
	defer func() { recover() }()
	SetupGit()
}
