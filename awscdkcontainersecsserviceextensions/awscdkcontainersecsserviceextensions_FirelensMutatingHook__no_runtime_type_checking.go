//go:build no_runtime_type_checking

// The CDK Construct Library that helps you build ECS services using simple extensions
package awscdkcontainersecsserviceextensions

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FirelensMutatingHook) validateMutateContainerDefinitionParameters(props *awsecs.ContainerDefinitionOptions) error {
	return nil
}

func validateNewFirelensMutatingHookParameters(props *FirelensProps) error {
	return nil
}

