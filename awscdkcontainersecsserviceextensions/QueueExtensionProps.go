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
	// Default: none.
	//
	// Experimental.
	EventsQueue awssqs.IQueue `field:"optional" json:"eventsQueue" yaml:"eventsQueue"`
	// The user-provided queue delay fields to configure auto scaling for the default queue.
	// Default: none.
	//
	// Experimental.
	ScaleOnLatency *QueueAutoScalingOptions `field:"optional" json:"scaleOnLatency" yaml:"scaleOnLatency"`
	// The list of subscriptions for this service.
	// Default: none.
	//
	// Experimental.
	Subscriptions *[]ISubscribable `field:"optional" json:"subscriptions" yaml:"subscriptions"`
}

