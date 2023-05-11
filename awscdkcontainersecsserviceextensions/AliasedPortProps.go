package awscdkcontainersecsserviceextensions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
)

// AliasedPortProps defines the properties of an aliased port extension.
// Experimental.
type AliasedPortProps struct {
	// The DNS alias to advertise for downstream clients.
	// Experimental.
	Alias *string `field:"required" json:"alias" yaml:"alias"`
	// The traffic port for clients to use to connect to the DNS alias.
	// Experimental.
	AliasPort *float64 `field:"optional" json:"aliasPort" yaml:"aliasPort"`
	// The protocol to use over the specified port.
	//
	// May be one of HTTP, HTTP2, or GRPC.
	// Experimental.
	AppProtocol awsecs.AppProtocol `field:"optional" json:"appProtocol" yaml:"appProtocol"`
}

