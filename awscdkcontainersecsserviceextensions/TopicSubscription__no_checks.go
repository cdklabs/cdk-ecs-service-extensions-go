//go:build no_runtime_type_checking

package awscdkcontainersecsserviceextensions

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TopicSubscription) validateSubscribeParameters(extension QueueExtension) error {
	return nil
}

func validateNewTopicSubscriptionParameters(props *TopicSubscriptionProps) error {
	return nil
}

