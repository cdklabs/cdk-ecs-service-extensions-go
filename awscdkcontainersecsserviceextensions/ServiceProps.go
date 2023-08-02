package awscdkcontainersecsserviceextensions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// The settings for an ECS Service.
// Experimental.
type ServiceProps struct {
	// The environment to launch the service in.
	// Experimental.
	Environment IEnvironment `field:"required" json:"environment" yaml:"environment"`
	// The ServiceDescription used to build the service.
	// Experimental.
	ServiceDescription ServiceDescription `field:"required" json:"serviceDescription" yaml:"serviceDescription"`
	// The options for configuring the auto scaling target.
	// Default: none.
	//
	// Experimental.
	AutoScaleTaskCount *AutoScalingOptions `field:"optional" json:"autoScaleTaskCount" yaml:"autoScaleTaskCount"`
	// The desired number of instantiations of the task definition to keep running on the service.
	// Default: - When creating the service, default is 1; when updating the service, default uses
	// the current task number.
	//
	// Experimental.
	DesiredCount *float64 `field:"optional" json:"desiredCount" yaml:"desiredCount"`
	// The name of the IAM role that grants containers in the task permission to call AWS APIs on your behalf.
	// Default: - A task role is automatically created for you.
	//
	// Experimental.
	TaskRole awsiam.IRole `field:"optional" json:"taskRole" yaml:"taskRole"`
}

