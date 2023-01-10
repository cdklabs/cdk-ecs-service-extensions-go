// The CDK Construct Library that helps you build ECS services using simple extensions
package awscdkcontainersecsserviceextensions


// The types of capacity that are supported.
//
// These capacity types may change the
// behavior of an extension.
// Experimental.
type EnvironmentCapacityType string

const (
	// Specify that the environment should use AWS Fargate for hosting containers.
	// Experimental.
	EnvironmentCapacityType_FARGATE EnvironmentCapacityType = "FARGATE"
	// Specify that the environment should launch containers onto EC2 instances.
	// Experimental.
	EnvironmentCapacityType_EC2 EnvironmentCapacityType = "EC2"
)

