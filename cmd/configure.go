package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"namaste-cloud/internal"

	"github.com/spf13/cobra"
)

// ConfigureCommand returns the `configure` command.
func ConfigureCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "configure",
		Short: "Configure credentials for a cloud provider",
		Run: func(cmd *cobra.Command, args []string) {
			reader := bufio.NewReader(os.Stdin)

			// Prompt for cloud provider
			fmt.Println("Select a cloud provider (aws, gcp, azure):")
			cloud, _ := reader.ReadString('\n')
			cloud = strings.TrimSpace(strings.ToLower(cloud))

			// Validate cloud provider
			if cloud != "aws" && cloud != "gcp" && cloud != "azure" {
				fmt.Println("Unsupported cloud provider. Please choose from aws, gcp, or azure.")
				return
			}

			// Prompt for credentials
			fmt.Printf("Enter credentials for %s:\n", cloud)
			fmt.Print("Access Key: ")
			accessKey, _ := reader.ReadString('\n')
			fmt.Print("Secret Key: ")
			secretKey, _ := reader.ReadString('\n')

			// Validate input
			accessKey = strings.TrimSpace(accessKey)
			secretKey = strings.TrimSpace(secretKey)

			if accessKey == "" || secretKey == "" {
				fmt.Println("Invalid credentials. Access Key and Secret Key cannot be empty. Please try again.")
				return
			}

			// Save credentials securely
			cred := internal.Credential{
				Cloud:     cloud,
				AccessKey: accessKey,
				SecretKey: secretKey,
			}
			if err := internal.SaveCredential(cred); err != nil {
				fmt.Printf("Failed to save credentials: %v\n", err)
				return
			}

			fmt.Printf("Credentials for %s saved successfully.\n", cloud)
		},
	}
}
