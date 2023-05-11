package awscdkcontainersecsserviceextensions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Setting for the main application container of a service.
// Experimental.
type ContainerExtensionProps struct {
	// How much CPU the container requires.
	// Experimental.
	Cpu *float64 `field:"required" json:"cpu" yaml:"cpu"`
	// The image to run.
	// Experimental.
	Image awsecs.ContainerImage `field:"required" json:"image" yaml:"image"`
	// How much memory in megabytes the container requires.
	// Experimental.
	MemoryMiB *float64 `field:"required" json:"memoryMiB" yaml:"memoryMiB"`
	// What port the image listen for traffic on.
	// Experimental.
	TrafficPort *float64 `field:"required" json:"trafficPort" yaml:"trafficPort"`
	// Environment variables to pass into the container.
	// Experimental.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// The environment files to pass to the container.
	// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/taskdef-envfiles.html
	//
	// Experimental.
	EnvironmentFiles *[]awsecs.EnvironmentFile `field:"optional" json:"environmentFiles" yaml:"environmentFiles"`
	// The log group into which application container logs should be routed.
	// Experimental.
	LogGroup awslogs.ILogGroup `field:"optional" json:"logGroup" yaml:"logGroup"`
	// The secret environment variables to pass to the container.
	// Experimental.
	Secrets *map[string]awsecs.Secret `field:"optional" json:"secrets" yaml:"secrets"`
}

