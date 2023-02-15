// The CDK Construct Library that helps you build ECS services using simple extensions
package awscdkcontainersecsserviceextensions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

// `SubscriptionQueue` represents the subscription queue object which includes the topic-specific queue and its corresponding auto scaling fields.
// Experimental.
type SubscriptionQueue struct {
	// The user-provided queue to subscribe to the given topic.
	// Experimental.
	Queue awssqs.IQueue `field:"required" json:"queue" yaml:"queue"`
	// The user-provided queue delay fields to configure auto scaling for the topic-specific queue.
	// Experimental.
	ScaleOnLatency *QueueAutoScalingOptions `field:"optional" json:"scaleOnLatency" yaml:"scaleOnLatency"`
}

