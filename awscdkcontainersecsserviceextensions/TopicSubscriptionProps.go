package awscdkcontainersecsserviceextensions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

// The topic-specific settings for creating the queue subscriptions.
// Experimental.
type TopicSubscriptionProps struct {
	// The SNS Topic to subscribe to.
	// Experimental.
	Topic awssns.ITopic `field:"required" json:"topic" yaml:"topic"`
	// The user-provided queue to subscribe to the given topic.
	// Deprecated: use `topicSubscriptionQueue`.
	Queue awssqs.IQueue `field:"optional" json:"queue" yaml:"queue"`
	// The object representing topic-specific queue and corresponding queue delay fields to configure auto scaling.
	//
	// If not provided, the default `eventsQueue` will subscribe to the given topic.
	// Experimental.
	TopicSubscriptionQueue *SubscriptionQueue `field:"optional" json:"topicSubscriptionQueue" yaml:"topicSubscriptionQueue"`
}

