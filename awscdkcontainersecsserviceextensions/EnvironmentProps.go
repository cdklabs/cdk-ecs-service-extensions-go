package awscdkcontainersecsserviceextensions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
)

// Settings for the environment where you want to deploy your services.
// Experimental.
type EnvironmentProps struct {
	// The type of capacity to use for this environment.
	// Default: - EnvironmentCapacityType.FARGATE
	//
	// Experimental.
	CapacityType EnvironmentCapacityType `field:"optional" json:"capacityType" yaml:"capacityType"`
	// The ECS cluster which provides compute capacity to this service.
	//
	// [disable-awslint:ref-via-interface].
	// Default: - Create a new cluster.
	//
	// Experimental.
	Cluster awsecs.Cluster `field:"optional" json:"cluster" yaml:"cluster"`
	// The VPC used by the service for networking.
	// Default: - Create a new VPC.
	//
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

