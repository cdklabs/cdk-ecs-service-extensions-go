package awscdkcontainersecsserviceextensions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
)

// Experimental.
type AliasedPortMutatingHookProps struct {
	// The port on the container which receives traffic.
	//
	// This is the same as the `containerPort` property of port mapping.
	// Experimental.
	AliasPort *float64 `field:"required" json:"aliasPort" yaml:"aliasPort"`
	// The name by which to refer to this port mapping.
	// Experimental.
	PortMappingName *string `field:"required" json:"portMappingName" yaml:"portMappingName"`
	// The protocol which this port mapping expects to receive.
	// Default: - none.
	//
	// Experimental.
	Protocol awsecs.AppProtocol `field:"optional" json:"protocol" yaml:"protocol"`
}

