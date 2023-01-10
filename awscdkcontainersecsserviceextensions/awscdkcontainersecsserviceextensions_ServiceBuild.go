// The CDK Construct Library that helps you build ECS services using simple extensions
package awscdkcontainersecsserviceextensions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
)

// A set of mutable service props in the process of being assembled using a builder pattern.
//
// They will eventually to be translated into an
// ecs.Ec2ServiceProps or ecs.FargateServiceProps interface, depending on the
// environment's capacity type.
// Experimental.
type ServiceBuild struct {
	// The cluster in which to launch the service.
	// Experimental.
	Cluster awsecs.ICluster `field:"required" json:"cluster" yaml:"cluster"`
	// The task definition registered to this service.
	// Experimental.
	TaskDefinition awsecs.TaskDefinition `field:"required" json:"taskDefinition" yaml:"taskDefinition"`
	// Specifies whether the task's elastic network interface receives a public IP address.
	//
	// If true, each task will receive a public IP address.
	// Experimental.
	AssignPublicIp *bool `field:"optional" json:"assignPublicIp" yaml:"assignPublicIp"`
	// Configuration for how to register the service in service discovery.
	// Experimental.
	CloudMapOptions *awsecs.CloudMapOptions `field:"optional" json:"cloudMapOptions" yaml:"cloudMapOptions"`
	// How many tasks to run.
	// Experimental.
	DesiredCount *float64 `field:"optional" json:"desiredCount" yaml:"desiredCount"`
	// How long the healthcheck can fail during initial task startup before the task is considered unhealthy.
	//
	// This is used to give the task more
	// time to start passing healthchecks.
	// Experimental.
	HealthCheckGracePeriod awscdk.Duration `field:"optional" json:"healthCheckGracePeriod" yaml:"healthCheckGracePeriod"`
	// Maximum percentage of tasks that can be launched.
	// Experimental.
	MaxHealthyPercent *float64 `field:"optional" json:"maxHealthyPercent" yaml:"maxHealthyPercent"`
	// Minimum healthy task percentage.
	// Experimental.
	MinHealthyPercent *float64 `field:"optional" json:"minHealthyPercent" yaml:"minHealthyPercent"`
	// Configuration for service connect for this service.
	// Experimental.
	ServiceConnectConfiguration *awsecs.ServiceConnectProps `field:"optional" json:"serviceConnectConfiguration" yaml:"serviceConnectConfiguration"`
}

