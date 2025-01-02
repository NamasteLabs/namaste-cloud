package cmd

import (
	"fmt"
	"strings"

	"namaste-cloud/internal"

	"github.com/spf13/cobra"
)

// UseCloudCommand returns the `use-cloud` command.
func UseCloudCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "use-cloud [cloud-provider]",
		Short: "Set the active cloud provider",
		Args:  cobra.ExactArgs(1), // Ensure exactly one argument is provided
		Run: func(cmd *cobra.Command, args []string) {
			cloud := strings.ToLower(args[0])
			if cloud == "aws" || cloud == "gcp" || cloud == "azure" {
				// Check if credentials exist for the cloud
				_, err := internal.GetCredential(cloud)
				if err != nil {
					fmt.Printf("No credentials found for %s. Use `namaste-cloud configure` to set them.\n", cloud)
					return
				}

				// Load existing configuration
				cfg, err := internal.LoadConfig()
				if err != nil {
					fmt.Printf("Failed to load configuration: %v\n", err)
					return
				}

				// Set the active cloud provider in the configuration
				cfg.ActiveCloud = cloud
				if err := internal.SaveConfig(cfg); err != nil {
					fmt.Printf("Failed to save active cloud provider: %v\n", err)
					return
				}

				fmt.Printf("Active cloud provider set to: %s\n", cloud)
			} else {
				fmt.Println("Unsupported cloud provider. Please choose from aws, gcp, or azure.")
			}
		},
	}
}
