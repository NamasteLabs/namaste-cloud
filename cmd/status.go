package cmd

import (
	"fmt"
	"namaste-cloud/internal"

	"github.com/spf13/cobra"
)

// StatusCommand returns the `status` command.
func StatusCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Display the current cloud provider and credential status",
		Run: func(cmd *cobra.Command, args []string) {
			// Load the active cloud provider from the file
			activeCloud, err := internal.LoadActiveCloudProvider()
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Printf("Active Cloud Provider: %s\n", activeCloud)

			// Validate credentials
			cred, err := internal.GetCredential(activeCloud)
			if err != nil {
				fmt.Println("Invalid or missing credentials. Use `namaste-cloud configure` to update.")
				return
			}

			fmt.Printf("Credentials for %s are valid.\n", cred.Cloud)
		},
	}
}
