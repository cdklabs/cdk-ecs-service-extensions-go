//go:build no_runtime_type_checking

package awscdkcontainersecsserviceextensions

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ContainerMutatingHook) validateMutateContainerDefinitionParameters(props *awsecs.ContainerDefinitionOptions) error {
	return nil
}

