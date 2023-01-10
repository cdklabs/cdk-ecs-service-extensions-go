// The CDK Construct Library that helps you build ECS services using simple extensions
package awscdkcontainersecsserviceextensions


// Experimental.
type AutoScalingOptions struct {
	// The maximum number of tasks when scaling out.
	// Experimental.
	MaxTaskCount *float64 `field:"required" json:"maxTaskCount" yaml:"maxTaskCount"`
	// The minimum number of tasks when scaling in.
	// Experimental.
	MinTaskCount *float64 `field:"optional" json:"minTaskCount" yaml:"minTaskCount"`
	// The target value for CPU utilization across all tasks in the service.
	// Experimental.
	TargetCpuUtilization *float64 `field:"optional" json:"targetCpuUtilization" yaml:"targetCpuUtilization"`
	// The target value for memory utilization across all tasks in the service.
	// Experimental.
	TargetMemoryUtilization *float64 `field:"optional" json:"targetMemoryUtilization" yaml:"targetMemoryUtilization"`
}

