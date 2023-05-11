package awscdkcontainersecsserviceextensions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
)

// Experimental.
type EnvironmentAttributes struct {
	// The capacity type used by the service's cluster.
	// Experimental.
	CapacityType EnvironmentCapacityType `field:"required" json:"capacityType" yaml:"capacityType"`
	// The cluster that is providing capacity for this service.
	// Experimental.
	Cluster awsecs.ICluster `field:"required" json:"cluster" yaml:"cluster"`
}

