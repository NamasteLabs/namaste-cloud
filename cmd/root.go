package cmd

import (
	"fmt"
	"namaste-cloud/cmd/instances"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd is the root command for the CLI.
var RootCmd = &cobra.Command{
	Use:   "namaste-cloud",
	Short: "Namaste Cloud CLI for managing cloud resources",
	Long:  "A CLI tool to manage cloud resources across AWS, GCP, and Azure.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Namaste Cloud CLI!")
	},
}

// Execute adds all child commands to the root command and runs the CLI.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.AddCommand(ConfigureCommand())
	RootCmd.AddCommand(UseCloudCommand())
	RootCmd.AddCommand(StatusCommand())
	RootCmd.AddCommand(instances.ListInstancesCommand())
	RootCmd.AddCommand(instances.CreateInstanceCommand())

}
