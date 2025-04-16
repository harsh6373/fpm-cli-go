package utils

import (
	"strings"
	"testing"
)

func TestGenerateArtifactName(t *testing.T) {
	name := GenerateArtifactName("MyApp", "STAGING", "1.2.3", "7")
	if !strings.Contains(name, "myapp_staging_v1.2.3+7_") {
		t.Errorf("Unexpected artifact name: %s", name)
	}
}
