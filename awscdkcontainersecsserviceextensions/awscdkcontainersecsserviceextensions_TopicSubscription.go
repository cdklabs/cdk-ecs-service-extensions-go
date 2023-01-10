// The CDK Construct Library that helps you build ECS services using simple extensions
package awscdkcontainersecsserviceextensions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-ecs-service-extensions-go/awscdkcontainersecsserviceextensions/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

// The `TopicSubscription` class represents an SNS Topic resource that can be subscribed to by the service queues.
// Experimental.
type TopicSubscription interface {
	ISubscribable
	// The queue that subscribes to the given topic.
	// Deprecated: use `subscriptionQueue`.
	Queue() awssqs.IQueue
	// The subscription queue object for this subscription.
	// Experimental.
	SubscriptionQueue() *SubscriptionQueue
	// Experimental.
	Topic() awssns.ITopic
	// This method sets up SNS Topic subscriptions for the SQS queue provided by the user.
	//
	// If a `queue` is not provided,
	// the default `eventsQueue` subscribes to the given topic.
	//
	// Returns: the queue subscribed to the given topic.
	// Experimental.
	Subscribe(extension QueueExtension) awssqs.IQueue
}

// The jsii proxy struct for TopicSubscription
type jsiiProxy_TopicSubscription struct {
	jsiiProxy_ISubscribable
}

func (j *jsiiProxy_TopicSubscription) Queue() awssqs.IQueue {
	var returns awssqs.IQueue
	_jsii_.Get(
		j,
		"queue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TopicSubscription) SubscriptionQueue() *SubscriptionQueue {
	var returns *SubscriptionQueue
	_jsii_.Get(
		j,
		"subscriptionQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TopicSubscription) Topic() awssns.ITopic {
	var returns awssns.ITopic
	_jsii_.Get(
		j,
		"topic",
		&returns,
	)
	return returns
}


// Experimental.
func NewTopicSubscription(props *TopicSubscriptionProps) TopicSubscription {
	_init_.Initialize()

	if err := validateNewTopicSubscriptionParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_TopicSubscription{}

	_jsii_.Create(
		"@aws-cdk-containers/ecs-service-extensions.TopicSubscription",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewTopicSubscription_Override(t TopicSubscription, props *TopicSubscriptionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk-containers/ecs-service-extensions.TopicSubscription",
		[]interface{}{props},
		t,
	)
}

func (t *jsiiProxy_TopicSubscription) Subscribe(extension QueueExtension) awssqs.IQueue {
	if err := t.validateSubscribeParameters(extension); err != nil {
		panic(err)
	}
	var returns awssqs.IQueue

	_jsii_.Invoke(
		t,
		"subscribe",
		[]interface{}{extension},
		&returns,
	)

	return returns
}

