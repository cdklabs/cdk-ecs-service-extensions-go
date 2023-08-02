package awscdkcontainersecsserviceextensions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// The autoscaling settings.
// Deprecated: use the `minTaskCount` and `maxTaskCount` properties of `autoScaleTaskCount` in the `Service` construct
// to configure the auto scaling target for the service. For more information, please refer
// https://github.com/aws/aws-cdk/blob/master/packages/%40aws-cdk-containers/ecs-service-extensions/README.md#task-auto-scaling .
type CpuScalingProps struct {
	// How many tasks to launch initially.
	// Default: - 2.
	//
	// Deprecated: use the `minTaskCount` and `maxTaskCount` properties of `autoScaleTaskCount` in the `Service` construct
	// to configure the auto scaling target for the service. For more information, please refer
	// https://github.com/aws/aws-cdk/blob/master/packages/%40aws-cdk-containers/ecs-service-extensions/README.md#task-auto-scaling .
	InitialTaskCount *float64 `field:"optional" json:"initialTaskCount" yaml:"initialTaskCount"`
	// The maximum number of tasks when scaling out.
	// Default: - 8.
	//
	// Deprecated: use the `minTaskCount` and `maxTaskCount` properties of `autoScaleTaskCount` in the `Service` construct
	// to configure the auto scaling target for the service. For more information, please refer
	// https://github.com/aws/aws-cdk/blob/master/packages/%40aws-cdk-containers/ecs-service-extensions/README.md#task-auto-scaling .
	MaxTaskCount *float64 `field:"optional" json:"maxTaskCount" yaml:"maxTaskCount"`
	// The minimum number of tasks when scaling in.
	// Default: - 2.
	//
	// Deprecated: use the `minTaskCount` and `maxTaskCount` properties of `autoScaleTaskCount` in the `Service` construct
	// to configure the auto scaling target for the service. For more information, please refer
	// https://github.com/aws/aws-cdk/blob/master/packages/%40aws-cdk-containers/ecs-service-extensions/README.md#task-auto-scaling .
	MinTaskCount *float64 `field:"optional" json:"minTaskCount" yaml:"minTaskCount"`
	// How long to wait between scale in actions.
	// Default: - 60 seconds.
	//
	// Deprecated: use the `minTaskCount` and `maxTaskCount` properties of `autoScaleTaskCount` in the `Service` construct
	// to configure the auto scaling target for the service. For more information, please refer
	// https://github.com/aws/aws-cdk/blob/master/packages/%40aws-cdk-containers/ecs-service-extensions/README.md#task-auto-scaling .
	ScaleInCooldown awscdk.Duration `field:"optional" json:"scaleInCooldown" yaml:"scaleInCooldown"`
	// How long to wait between scale out actions.
	// Default: - 60 seconds.
	//
	// Deprecated: use the `minTaskCount` and `maxTaskCount` properties of `autoScaleTaskCount` in the `Service` construct
	// to configure the auto scaling target for the service. For more information, please refer
	// https://github.com/aws/aws-cdk/blob/master/packages/%40aws-cdk-containers/ecs-service-extensions/README.md#task-auto-scaling .
	ScaleOutCooldown awscdk.Duration `field:"optional" json:"scaleOutCooldown" yaml:"scaleOutCooldown"`
	// The CPU utilization to try ot maintain.
	// Default: - 50%.
	//
	// Deprecated: use the `minTaskCount` and `maxTaskCount` properties of `autoScaleTaskCount` in the `Service` construct
	// to configure the auto scaling target for the service. For more information, please refer
	// https://github.com/aws/aws-cdk/blob/master/packages/%40aws-cdk-containers/ecs-service-extensions/README.md#task-auto-scaling .
	TargetCpuUtilization *float64 `field:"optional" json:"targetCpuUtilization" yaml:"targetCpuUtilization"`
}

