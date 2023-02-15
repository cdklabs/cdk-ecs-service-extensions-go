// The CDK Construct Library that helps you build ECS services using simple extensions
package awscdkcontainersecsserviceextensions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Settings for the hook which mutates the application container to route logs through FireLens.
// Experimental.
type FirelensProps struct {
	// The log group into which logs should be routed.
	// Experimental.
	LogGroup awslogs.LogGroup `field:"required" json:"logGroup" yaml:"logGroup"`
	// The parent service that is being mutated.
	// Experimental.
	ParentService Service `field:"required" json:"parentService" yaml:"parentService"`
}

