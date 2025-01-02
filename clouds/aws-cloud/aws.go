package awscloud

import (
	"context"
	"fmt"
	"log"
	"namaste-cloud/internal"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	ec2types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

// ExecuteAWSCommand performs AWS operations based on the given command.
func ExecuteAWSCommand(command string, args ...string) {
	// Load AWS credentials from storage
	cred, err := internal.GetCredential("aws")
	if err != nil {
		log.Fatalf("Failed to load AWS credentials: %v", err)
	}

	// Load AWS configuration with credentials
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
			return aws.Credentials{
				AccessKeyID:     cred.AccessKey,
				SecretAccessKey: cred.SecretKey,
			}, nil
		})),
	)
	if err != nil {
		log.Fatalf("Failed to load AWS configuration: %v", err)
	}

	// Create an EC2 client
	ec2Client := ec2.NewFromConfig(cfg)

	// Dispatch command based on the input
	switch command {
	case "list-instances":
		listInstances(ec2Client)
	case "create-instance":
		if len(args) < 1 {
			fmt.Println("Please specify the AMI ID for the instance.")
			return
		}
		createInstance(ec2Client, args[0])
	case "stop-instance":
		if len(args) < 1 {
			fmt.Println("Please specify the Instance ID to stop.")
			return
		}
		stopInstance(ec2Client, args[0])
	case "start-instance":
		if len(args) < 1 {
			fmt.Println("Please specify the Instance ID to start.")
			return
		}
		startInstance(ec2Client, args[0])
	case "terminate-instance":
		if len(args) < 1 {
			fmt.Println("Please specify the Instance ID to terminate.")
			return
		}
		terminateInstance(ec2Client, args[0])
	case "describe-instance":
		if len(args) < 1 {
			fmt.Println("Please specify the Instance ID to describe.")
			return
		}
		describeInstance(ec2Client, args[0])
	case "list-regions":
		listRegions(ec2Client)
	case "create-key-pair":
		if len(args) < 1 {
			fmt.Println("Please specify a name for the key pair.")
			return
		}
		createKeyPair(ec2Client, args[0])
	case "create-security-group":
		if len(args) < 2 {
			fmt.Println("Please specify a name and description for the security group.")
			return
		}
		createSecurityGroup(ec2Client, args[0], args[1])
	case "authorize-security-group":
		if len(args) < 3 {
			fmt.Println("Please specify the Security Group ID, protocol, and port range.")
			return
		}
		authorizeSecurityGroup(ec2Client, args[0], args[1], args[2])
	default:
		fmt.Println("Unsupported AWS command.")
	}
}

// listInstances lists all EC2 instances.
func listInstances(client *ec2.Client) {
	resp, err := client.DescribeInstances(context.TODO(), &ec2.DescribeInstancesInput{})
	if err != nil {
		log.Fatalf("Failed to list instances: %v", err)
	}

	for _, reservation := range resp.Reservations {
		for _, instance := range reservation.Instances {
			fmt.Printf("Instance ID: %s, State: %s\n", *instance.InstanceId, instance.State.Name)
		}
	}
}

// createInstance creates an EC2 instance with the specified AMI ID.
func createInstance(client *ec2.Client, amiID string) {
	resp, err := client.RunInstances(context.TODO(), &ec2.RunInstancesInput{
		ImageId:      aws.String(amiID),
		InstanceType: "t2.micro",
		MinCount:     aws.Int32(1),
		MaxCount:     aws.Int32(1),
	})
	if err != nil {
		log.Fatalf("Failed to create instance: %v", err)
	}

	for _, instance := range resp.Instances {
		fmt.Printf("Created instance with ID: %s\n", *instance.InstanceId)
	}
}

// stopInstance stops an EC2 instance.
func stopInstance(client *ec2.Client, instanceID string) {
	_, err := client.StopInstances(context.TODO(), &ec2.StopInstancesInput{
		InstanceIds: []string{instanceID},
	})
	if err != nil {
		log.Fatalf("Failed to stop instance: %v", err)
	}

	fmt.Printf("Stopping instance with ID: %s\n", instanceID)
}

// startInstance starts an EC2 instance.
func startInstance(client *ec2.Client, instanceID string) {
	_, err := client.StartInstances(context.TODO(), &ec2.StartInstancesInput{
		InstanceIds: []string{instanceID},
	})
	if err != nil {
		log.Fatalf("Failed to start instance: %v", err)
	}

	fmt.Printf("Starting instance with ID: %s\n", instanceID)
}

// terminateInstance terminates an EC2 instance.
func terminateInstance(client *ec2.Client, instanceID string) {
	_, err := client.TerminateInstances(context.TODO(), &ec2.TerminateInstancesInput{
		InstanceIds: []string{instanceID},
	})
	if err != nil {
		log.Fatalf("Failed to terminate instance: %v", err)
	}

	fmt.Printf("Terminating instance with ID: %s\n", instanceID)
}

// describeInstance provides details of an EC2 instance.
func describeInstance(client *ec2.Client, instanceID string) {
	resp, err := client.DescribeInstances(context.TODO(), &ec2.DescribeInstancesInput{
		InstanceIds: []string{instanceID},
	})
	if err != nil {
		log.Fatalf("Failed to describe instance: %v", err)
	}

	for _, reservation := range resp.Reservations {
		for _, instance := range reservation.Instances {
			fmt.Printf("Instance ID: %s, State: %s, Public IP: %s\n",
				*instance.InstanceId, instance.State.Name, aws.ToString(instance.PublicIpAddress))
		}
	}
}

// listRegions lists all available AWS regions.
func listRegions(client *ec2.Client) {
	resp, err := client.DescribeRegions(context.TODO(), &ec2.DescribeRegionsInput{})
	if err != nil {
		log.Fatalf("Failed to list regions: %v", err)
	}

	for _, region := range resp.Regions {
		fmt.Printf("Region Name: %s\n", *region.RegionName)
	}
}

// createKeyPair creates a new key pair.
func createKeyPair(client *ec2.Client, keyName string) {
	resp, err := client.CreateKeyPair(context.TODO(), &ec2.CreateKeyPairInput{
		KeyName: aws.String(keyName),
	})
	if err != nil {
		log.Fatalf("Failed to create key pair: %v", err)
	}

	fmt.Printf("Key Pair Created: %s\nPrivate Key:\n%s\n", *resp.KeyName, *resp.KeyMaterial)
}

// createSecurityGroup creates a new security group.
func createSecurityGroup(client *ec2.Client, name, description string) {
	resp, err := client.CreateSecurityGroup(context.TODO(), &ec2.CreateSecurityGroupInput{
		GroupName:   aws.String(name),
		Description: aws.String(description),
	})
	if err != nil {
		log.Fatalf("Failed to create security group: %v", err)
	}

	fmt.Printf("Security Group Created: %s\n", *resp.GroupId)
}

// authorizeSecurityGroup adds a rule to a security group.
func authorizeSecurityGroup(client *ec2.Client, groupID, protocol, portRange string) {
	// Parse the port range
	var fromPort, toPort int32
	fmt.Sscanf(portRange, "%d-%d", &fromPort, &toPort)

	_, err := client.AuthorizeSecurityGroupIngress(context.TODO(), &ec2.AuthorizeSecurityGroupIngressInput{
		GroupId: aws.String(groupID),
		IpPermissions: []ec2types.IpPermission{
			{
				IpProtocol: aws.String(protocol),
				FromPort:   aws.Int32(fromPort),
				ToPort:     aws.Int32(toPort),
				IpRanges: []ec2types.IpRange{
					{CidrIp: aws.String("0.0.0.0/0")},
				},
			},
		},
	})
	if err != nil {
		log.Fatalf("Failed to authorize security group ingress: %v", err)
	}

	fmt.Printf("Ingress rule added to Security Group: %s\n", groupID)
}
