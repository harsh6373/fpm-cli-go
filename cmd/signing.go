package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/harsh6373/fpm-cli-go/utils"
	"github.com/spf13/cobra"
)

var signingCmd = &cobra.Command{
	Use:   "signing",
	Short: "Setup Android signing configuration",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)

		input := func(prompt string, validate func(string) bool) string {
			for {
				fmt.Print(prompt)
				val, _ := reader.ReadString('\n')
				val = strings.TrimSpace(val)
				if validate == nil || validate(val) {
					return val
				}
				fmt.Println("❌ Invalid input. Try again.")
			}
		}

		keyAlias := input("Enter key alias: ", nil)
		keyPassword := input("Enter key password: ", nil)
		storePassword := input("Enter store password: ", nil)
		orgName := input("Enter organization/name: ", nil)
		city := input("Enter city/locality: ", nil)
		state := input("Enter state/province: ", nil)
		country := input("Enter country code (2-letter): ", func(s string) bool { return len(s) == 2 })
		validity := input("Enter validity in days (e.g. 10000): ", func(s string) bool { return s != "" })

		confirmOverwrite := false
		if _, err := os.Stat("key.properties"); err == nil {
			fmt.Print("\n'key.properties' already exists. Overwrite? (y/n): ")
			confirm, _ := reader.ReadString('\n')
			confirm = strings.ToLower(strings.TrimSpace(confirm))
			confirmOverwrite = (confirm == "y")
		} else {
			confirmOverwrite = true
		}

		if !confirmOverwrite {
			fmt.Println("Aborting without modifying 'key.properties'.")
			return
		}

		if err := utils.CreateKeystoreAndProperties(utils.SigningInfo{
			KeyAlias:      keyAlias,
			KeyPassword:   keyPassword,
			StorePassword: storePassword,
			Org:           orgName,
			City:          city,
			State:         state,
			Country:       country,
			Validity:      validity,
		}); err != nil {
			log.Fatalf("❌ Error: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(signingCmd)
}
