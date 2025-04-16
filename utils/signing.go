package utils

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type SigningInfo struct {
	KeyAlias      string
	KeyPassword   string
	StorePassword string
	Org           string
	City          string
	State         string
	Country       string
	Validity      string
}

func CreateKeystoreAndProperties(info SigningInfo) error {
	androidAppPath := filepath.Join("android", "app")
	if _, err := os.Stat(androidAppPath); errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("'%s' directory not found. Make sure you are in the root of a Flutter project", androidAppPath)
	}

	if _, err := exec.LookPath("keytool"); err != nil {
		return errors.New("'keytool' not found. Make sure Java is installed and keytool is in PATH")
	}

	keystorePath := filepath.Join(androidAppPath, "my-release-key.jks")
	cmdArgs := []string{
		"-genkey",
		"-v",
		"-keystore", keystorePath,
		"-alias", info.KeyAlias,
		"-keyalg", "RSA",
		"-keysize", "2048",
		"-validity", info.Validity,
		"-storepass", info.StorePassword,
		"-keypass", info.KeyPassword,
		"-dname", fmt.Sprintf("CN=%s, OU=%s, O=%s, L=%s, S=%s, C=%s",
			info.KeyAlias, info.Org, info.Org, info.City, info.State, info.Country),
	}

	fmt.Println("\n🔐 Generating keystore file...")
	if err := exec.Command("keytool", cmdArgs...).Run(); err != nil {
		return fmt.Errorf("failed to generate keystore: %w", err)
	}
	fmt.Println("✅ Keystore generated successfully at", keystorePath)

	keyProps := fmt.Sprintf(`storePassword=%s
keyPassword=%s
keyAlias=%s
storeFile=my-release-key.jks
`, info.StorePassword, info.KeyPassword, info.KeyAlias)

	if err := os.WriteFile("key.properties", []byte(keyProps), 0644); err != nil {
		return fmt.Errorf("failed to write key.properties: %w", err)
	}
	fmt.Println("✅ key.properties file created.")
	fmt.Println("ℹ️ Please update your android/app/build.gradle with signingConfigs and release buildTypes.")
	return nil
}
