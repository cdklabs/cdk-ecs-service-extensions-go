// The CDK Construct Library that helps you build ECS services using simple extensions
package awscdkcontainersecsserviceextensions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappmesh"
)

// The settings for the App Mesh extension.
// Experimental.
type MeshProps struct {
	// The service mesh into which to register the service.
	// Experimental.
	Mesh awsappmesh.Mesh `field:"required" json:"mesh" yaml:"mesh"`
	// The protocol of the service.
	//
	// Valid values are Protocol.HTTP, Protocol.HTTP2, Protocol.TCP, Protocol.GRPC
	// Experimental.
	Protocol Protocol `field:"optional" json:"protocol" yaml:"protocol"`
}

