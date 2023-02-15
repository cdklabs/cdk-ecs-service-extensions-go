// The CDK Construct Library that helps you build ECS services using simple extensions
package awscdkcontainersecsserviceextensions


// Enum of supported AppMesh protocols.
// Experimental.
type Protocol string

const (
	// Experimental.
	Protocol_HTTP Protocol = "HTTP"
	// Experimental.
	Protocol_TCP Protocol = "TCP"
	// Experimental.
	Protocol_HTTP2 Protocol = "HTTP2"
	// Experimental.
	Protocol_GRPC Protocol = "GRPC"
)

