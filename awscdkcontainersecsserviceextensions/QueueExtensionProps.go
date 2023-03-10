// The CDK Construct Library that helps you build ECS services using simple extensions
package awscdkcontainersecsserviceextensions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

// The settings for the Queue extension.
// Experimental.
type QueueExtensionProps struct {
	// The user-provided default queue for this service.
	//
	// If the `eventsQueue` is not provided, a default SQS Queue is created for the service.
	// Experimental.
	EventsQueue awssqs.IQueue `field:"optional" json:"eventsQueue" yaml:"eventsQueue"`
	// The user-provided queue delay fields to configure auto scaling for the default queue.
	// Experimental.
	ScaleOnLatency *QueueAutoScalingOptions `field:"optional" json:"scaleOnLatency" yaml:"scaleOnLatency"`
	// The list of subscriptions for this service.
	// Experimental.
	Subscriptions *[]ISubscribable `field:"optional" json:"subscriptions" yaml:"subscriptions"`
}

