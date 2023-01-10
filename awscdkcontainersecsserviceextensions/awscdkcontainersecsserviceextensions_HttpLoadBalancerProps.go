// The CDK Construct Library that helps you build ECS services using simple extensions
package awscdkcontainersecsserviceextensions


// Experimental.
type HttpLoadBalancerProps struct {
	// The number of ALB requests per target.
	// Experimental.
	RequestsPerTarget *float64 `field:"optional" json:"requestsPerTarget" yaml:"requestsPerTarget"`
}

