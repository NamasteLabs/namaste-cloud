package gcpcloud

import (
	"context"
	"fmt"
	"log"
	"namaste-cloud/internal"

	compute "cloud.google.com/go/compute/apiv1"
	"google.golang.org/api/option"
)

// ExecuteGCPCommand performs GCP operations based on the given command.
func ExecuteGCPCommand(command string, args ...string) {
	// Load GCP credentials from storage
	cred, err := internal.GetCredential("gcp")
	if err != nil {
		log.Fatalf("Failed to load GCP credentials: %v", err)
	}

	// Load GCP configuration with credentials
	ctx := context.Background()
	client, err := compute.NewInstancesRESTClient(ctx, option.WithCredentialsJSON([]byte(cred.AccessKey)))
	if err != nil {
		log.Fatalf("Failed to create GCP client: %v", err)
	}
	defer client.Close()

	// Dispatch command based on the input
	switch command {
	case "list-instances":
		// listInstances(client)
	case "create-instance":
		if len(args) < 1 {
			fmt.Println("Please specify the machine type for the instance.")
			return
		}
		// createInstance(client, args[0])
	case "stop-instance":
		if len(args) < 1 {
			fmt.Println("Please specify the Instance ID to stop.")
			return
		}
		// stopInstance(client, args[0])
	case "start-instance":
		if len(args) < 1 {
			fmt.Println("Please specify the Instance ID to start.")
			return
		}
		// startInstance(client, args[0])
	case "terminate-instance":
		if len(args) < 1 {
			fmt.Println("Please specify the Instance ID to terminate.")
			return
		}
		// terminateInstance(client, args[0])
	default:
		fmt.Println("Unsupported GCP command.")
	}
}

// // listInstances lists all GCP instances.
// func listInstances(client *compute.InstancesClient) {
// 	ctx := context.Background()
// 	req := &compute.AggregatedListInstancesRequest{}
// 	it := client.AggregatedList(ctx, req)

// 	for {
// 		resp, err := it.Next()
// 		if err != nil {
// 			if err == iterator.Done {
// 				break
// 			}
// 			log.Fatalf("Failed to list instances: %v", err)
// 		}

// 		for _, instances := range resp.Items {
// 			for _, instance := range instances.Instances {
// 				fmt.Printf("Instance ID: %s, State: %s\n", instance.Name, instance.Status)
// 			}
// 		}
// 	}
// }

// // createInstance creates a GCP instance with the specified machine type.
// func createInstance(client *compute.InstancesClient, machineType string) {
// 	// Example of creating an instance (you would need to customize the config)
// 	instance := &compute.Instance{
// 		Name:        "new-instance",
// 		MachineType: fmt.Sprintf("zones/us-central1-a/machineTypes/%s", machineType),
// 	}

// 	// Set other parameters like disk, image, network, etc.

// 	// Create instance
// 	ctx := context.Background()
// 	_, err := client.Insert(ctx, &compute.InsertInstanceRequest{
// 		Project:  "your-project-id",
// 		Zone:     "us-central1-a",
// 		Instance: instance,
// 	})
// 	if err != nil {
// 		log.Fatalf("Failed to create instance: %v", err)
// 	}

// 	fmt.Println("Instance created successfully.")
// }

// // stopInstance stops a GCP instance.
// func stopInstance(client *compute.InstancesClient, instanceID string) {
// 	ctx := context.Background()
// 	_, err := client.Stop(ctx, "your-project-id", "us-central1-a", instanceID)
// 	if err != nil {
// 		log.Fatalf("Failed to stop instance: %v", err)
// 	}

// 	fmt.Printf("Stopping instance with ID: %s\n", instanceID)
// }

// // startInstance starts a GCP instance.
// func startInstance(client *compute.InstancesClient, instanceID string) {
// 	ctx := context.Background()
// 	_, err := client.Start(ctx, "your-project-id", "us-central1-a", instanceID)
// 	if err != nil {
// 		log.Fatalf("Failed to start instance: %v", err)
// 	}

// 	fmt.Printf("Starting instance with ID: %s\n", instanceID)
// }

// // terminateInstance terminates a GCP instance.
// func terminateInstance(client *compute.InstancesClient, instanceID string) {
// 	ctx := context.Background()
// 	_, err := client.Delete(ctx, "your-project-id", "us-central1-a", instanceID)
// 	if err != nil {
// 		log.Fatalf("Failed to terminate instance: %v", err)
// 	}

// 	fmt.Printf("Terminating instance with ID: %s\n", instanceID)
// }
