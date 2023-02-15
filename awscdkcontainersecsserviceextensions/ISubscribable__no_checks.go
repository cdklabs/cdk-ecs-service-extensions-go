//go:build no_runtime_type_checking

// The CDK Construct Library that helps you build ECS services using simple extensions
package awscdkcontainersecsserviceextensions

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ISubscribable) validateSubscribeParameters(extension QueueExtension) error {
	return nil
}

