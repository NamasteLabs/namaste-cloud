package instances

import (
	"fmt"
	"namaste-cloud/internal" // Import the internal package to load the config

	"github.com/spf13/cobra"
)

// CreateInstanceCommand returns the `create-instance` command.
func CreateInstanceCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "create-instance",
		Short: "Create an instance in the selected cloud provider",
		Run: func(cmd *cobra.Command, args []string) {
			// Load the active cloud provider from the configuration file
			cfg, err := internal.LoadConfig()
			if err != nil {
				fmt.Println("Error loading configuration:", err)
				return
			}

			if cfg.ActiveCloud == "" {
				fmt.Println("No active cloud provider set. Please use `namaste-cloud use-cloud` to select one.")
				return
			}

			// Assuming we have an amiID for AWS or instance config for other clouds
			switch cfg.ActiveCloud {
			case "aws":
				// awscloud.CreateInstance("ami-12345")
			case "gcp":
				// gcpcloud.CreateVM("instance-config")
			case "azure":
				// azurecloud.CreateVM("instance-config")
			default:
				fmt.Println("Unsupported cloud provider. Please use `namaste-cloud use-cloud` to select a supported cloud provider.")
			}
		},
	}
}
