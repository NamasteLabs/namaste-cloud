package azurecloud

import (
	// "context"
	// "fmt"
	// "log"
	// "namaste-cloud/internal"

	// // "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	// "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

// // ExecuteAzureCommand performs Azure operations based on the given command.
// func ExecuteAzureCommand(command string, args ...string) {
// 	// Load Azure credentials from storage
// 	cred, err := internal.GetCredential("azure")
// 	if err != nil {
// 		log.Fatalf("Failed to load Azure credentials: %v", err)
// 	}

// 	// Use Azure credentials to create a client
// 	// credClient, err := azidentity.NewClientSecretCredential(cred.TenantID, cred.AccessKey, cred.SecretKey, nil)
// 	if err != nil {
// 		log.Fatalf("Failed to create Azure client: %v", err)
// 	}

// 	// // Create an Azure compute client
// 	// client := armcompute.NewVirtualMachinesClient(cred.SubscriptionID, credClient)

// 	// Dispatch command based on the input
// 	switch command {
// 	case "list-instances":
// 		// listInstances(client)
// 	case "create-instance":
// 		if len(args) < 1 {
// 			fmt.Println("Please specify the VM size for the instance.")
// 			return
// 		}
// 		// createInstance(client, args[0])
// 	case "stop-instance":
// 		if len(args) < 1 {
// 			fmt.Println("Please specify the Instance ID to stop.")
// 			return
// 		}
// 		// stopInstance(client, args[0])
// 	case "start-instance":
// 		if len(args) < 1 {
// 			fmt.Println("Please specify the Instance ID to start.")
// 			return
// 		}
// 		// startInstance(client, args[0])
// 	case "terminate-instance":
// 		if len(args) < 1 {
// 			fmt.Println("Please specify the Instance ID to terminate.")
// 			return
// 		}
// 		// terminateInstance(client, args[0])
// 	default:
// 		fmt.Println("Unsupported Azure command.")
// 	}
// }

// // listInstances lists all Azure virtual machines.
// func listInstances(client *armcompute.VirtualMachinesClient) {
// 	ctx := context.Background()
// 	pager := client.NewListPager("", nil)

// 	for pager.More() {
// 		page, err := pager.NextPage(ctx)
// 		if err != nil {
// 			log.Fatalf("Failed to list instances: %v", err)
// 		}

// 		for _, vm := range page.Value {
// 			fmt.Printf("VM ID: %s, State: %s\n", *vm.Name, *vm.Status)
// 		}
// 	}
// }

// // createInstance creates an Azure virtual machine with the specified size.
// func createInstance(client *armcompute.VirtualMachinesClient, vmSize string) {
// 	// Example of creating a VM (you would need to customize the config)
// 	vm := armcompute.VirtualMachine{
// 		Location: "eastus",
// 		Properties: &armcompute.VirtualMachineProperties{
// 			HardwareProfile: &armcompute.HardwareProfile{
// 				VMSize: armcompute.VirtualMachineSizeTypes(vmSize),
// 			},
// 		},
// 	}

// 	// Create VM
// 	ctx := context.Background()
// 	_, err := client.BeginCreateOrUpdate(ctx, "your-resource-group", "new-vm", vm, nil)
// 	if err != nil {
// 		log.Fatalf("Failed to create instance: %v", err)
// 	}

// 	fmt.Println("VM created successfully.")
// }

// // stopInstance stops an Azure virtual machine.
// func stopInstance(client *armcompute.VirtualMachinesClient, instanceID string) {
// 	ctx := context.Background()
// 	_, err := client.BeginDeallocate(ctx, "your-resource-group", instanceID, nil)
// 	if err != nil {
// 		log.Fatalf("Failed to stop instance: %v", err)
// 	}

// 	fmt.Printf("Stopping instance with ID: %s\n", instanceID)
// }

// // startInstance starts an Azure virtual machine.
// func startInstance(client *armcompute.VirtualMachinesClient, instanceID string) {
// 	ctx := context.Background()
// 	_, err := client.BeginStart(ctx, "your-resource-group", instanceID, nil)
// 	if err != nil {
// 		log.Fatalf("Failed to start instance: %v", err)
// 	}

// 	fmt.Printf("Starting instance with ID: %s\n", instanceID)
// }

// // terminateInstance terminates an Azure virtual machine.
// func terminateInstance(client *armcompute.VirtualMachinesClient, instanceID string) {
// 	ctx := context.Background()
// 	_, err := client.BeginDelete(ctx, "your-resource-group", instanceID, nil)
// 	if err != nil {
// 		log.Fatalf("Failed to terminate instance: %v", err)
// 	}

// 	fmt.Printf("Terminating instance with ID: %s\n", instanceID)
// }
