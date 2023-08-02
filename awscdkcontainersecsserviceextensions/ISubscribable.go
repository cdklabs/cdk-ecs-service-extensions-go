package awscdkcontainersecsserviceextensions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

// An interface that will be implemented by all the resources that can be subscribed to.
// Experimental.
type ISubscribable interface {
	// All classes implementing this interface must also implement the `subscribe()` method.
	// Experimental.
	Subscribe(extension QueueExtension) awssqs.IQueue
	// The `SubscriptionQueue` object for the `ISubscribable` object.
	// Default: none.
	//
	// Experimental.
	SubscriptionQueue() *SubscriptionQueue
}

// The jsii proxy for ISubscribable
type jsiiProxy_ISubscribable struct {
	_ byte // padding
}

func (i *jsiiProxy_ISubscribable) Subscribe(extension QueueExtension) awssqs.IQueue {
	if err := i.validateSubscribeParameters(extension); err != nil {
		panic(err)
	}
	var returns awssqs.IQueue

	_jsii_.Invoke(
		i,
		"subscribe",
		[]interface{}{extension},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_ISubscribable) SubscriptionQueue() *SubscriptionQueue {
	var returns *SubscriptionQueue
	_jsii_.Get(
		j,
		"subscriptionQueue",
		&returns,
	)
	return returns
}

