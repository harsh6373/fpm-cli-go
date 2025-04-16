package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var signingCmd = &cobra.Command{
	Use:   "signing",
	Short: "Setup Android signing configuration",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)

		// Check if android/app folder exists
		androidAppPath := filepath.Join("android", "app")
		if _, err := os.Stat(androidAppPath); errors.Is(err, os.ErrNotExist) {
			log.Fatalf("Error: '%s' directory not found. Make sure you are in the root of a Flutter project.", androidAppPath)
		}

		// Check if keytool is installed
		if _, err := exec.LookPath("keytool"); err != nil {
			log.Fatalf("Error: 'keytool' not found. Make sure Java is installed and keytool is available in PATH.")
		}

		fmt.Print("Enter key alias: ")
		keyAlias, _ := reader.ReadString('\n')
		keyAlias = strings.TrimSpace(keyAlias)

		fmt.Print("Enter key password: ")
		keyPassword, _ := reader.ReadString('\n')
		keyPassword = strings.TrimSpace(keyPassword)

		fmt.Print("Enter store password: ")
		storePassword, _ := reader.ReadString('\n')
		storePassword = strings.TrimSpace(storePassword)

		fmt.Print("Enter organization/name: ")
		orgName, _ := reader.ReadString('\n')
		orgName = strings.TrimSpace(orgName)

		fmt.Print("Enter city/locality: ")
		city, _ := reader.ReadString('\n')
		city = strings.TrimSpace(city)

		fmt.Print("Enter state/province: ")
		state, _ := reader.ReadString('\n')
		state = strings.TrimSpace(state)

		fmt.Print("Enter country code (2-letter): ")
		country, _ := reader.ReadString('\n')
		country = strings.TrimSpace(country)
		if len(country) != 2 {
			log.Fatal("Error: Country code must be exactly 2 letters.")
		}

		fmt.Print("Enter validity in days (e.g. 10000): ")
		validity, _ := reader.ReadString('\n')
		validity = strings.TrimSpace(validity)
		if validity == "" {
			log.Fatal("Error: Validity period is required.")
		}

		keystorePath := filepath.Join("android", "app", "my-release-key.jks")
		keyPropsPath := "key.properties"

		cmdArgs := []string{
			"-genkey",
			"-v",
			"-keystore", keystorePath,
			"-alias", keyAlias,
			"-keyalg", "RSA",
			"-keysize", "2048",
			"-validity", validity,
			"-storepass", storePassword,
			"-keypass", keyPassword,
			"-dname", fmt.Sprintf("CN=%s, OU=%s, O=%s, L=%s, S=%s, C=%s", keyAlias, orgName, orgName, city, state, country),
		}

		fmt.Println("\nGenerating keystore file...")
		err := exec.Command("keytool", cmdArgs...).Run()
		if err != nil {
			log.Fatalf("Failed to generate keystore: %v", err)
		}
		fmt.Println("Keystore generated successfully at", keystorePath)

		keyProps := fmt.Sprintf(`storePassword=%s
keyPassword=%s
keyAlias=%s
storeFile=my-release-key.jks
`, storePassword, keyPassword, keyAlias)

		// Check if key.properties already exists
		if _, err := os.Stat(keyPropsPath); err == nil {
			fmt.Print("\n'key.properties' already exists. Overwrite? (y/n): ")
			confirm, _ := reader.ReadString('\n')
			confirm = strings.ToLower(strings.TrimSpace(confirm))
			if confirm != "y" {
				fmt.Println("Aborting without modifying 'key.properties'.")
				return
			}
		}

		err = os.WriteFile(keyPropsPath, []byte(keyProps), 0644)
		if err != nil {
			log.Fatalf("Failed to write key.properties: %v", err)
		}
		fmt.Println("key.properties file created.")

		fmt.Println("\nPlease ensure you update your android/app/build.gradle file with signingConfigs and buildTypes release section.")
	},
}

func init() {
	rootCmd.AddCommand(signingCmd)
}
