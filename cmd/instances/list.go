// instances/listInstances.go
package instances

import (
	"fmt"
	awscloud "namaste-cloud/clouds/aws-cloud"
	gcpcloud "namaste-cloud/clouds/gcp-cloud"
	"namaste-cloud/internal" // Import the internal package to load the config

	"github.com/spf13/cobra"
)

// ListInstancesCommand returns the `list-instances` command.
func ListInstancesCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "list-instances",
		Short: "List instances/VMs in the selected cloud provider",
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

			// Perform the cloud-specific action based on the active cloud provider
			switch cfg.ActiveCloud {
			case "aws":
				awscloud.ExecuteAWSCommand("list-instances")
			case "gcp":
				gcpcloud.ExecuteGCPCommand("list-vms")
			case "azure":
				// azurecloud.ExecuteGCPCommand("list-compute")
			default:
				fmt.Println("Unsupported cloud provider. Please use `namaste-cloud use-cloud` to select a supported cloud provider.")
			}
		},
	}
}
