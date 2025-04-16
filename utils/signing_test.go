package utils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCreateKeystoreAndPropertiesMock(t *testing.T) {
	// Setup dummy android/app structure
	appPath := filepath.Join("android", "app")
	os.MkdirAll(appPath, os.ModePerm)
	defer os.RemoveAll("android")

	info := SigningInfo{
		KeyAlias:      "testkey",
		KeyPassword:   "123456",
		StorePassword: "654321",
		Org:           "TestOrg",
		City:          "City",
		State:         "State",
		Country:       "IN",
		Validity:      "10000",
	}

	// Mock keytool
	fakeKeytool := filepath.Join(os.TempDir(), "keytool")
	_ = os.WriteFile(fakeKeytool, []byte("#!/bin/bash\necho 'Fake keytool'\n"), 0755)
	defer os.Remove(fakeKeytool)
	os.Setenv("PATH", filepath.Dir(fakeKeytool)+":"+os.Getenv("PATH"))

	err := CreateKeystoreAndProperties(info)
	if err != nil {
		t.Fatalf("Expected CreateKeystoreAndProperties to succeed, got error: %v", err)
	}

	// Validate key.properties written
	_, err = os.Stat("key.properties")
	if err != nil {
		t.Errorf("Expected key.properties to exist, got error: %v", err)
	}
	os.Remove("key.properties")
}
