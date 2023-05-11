//go:build no_runtime_type_checking

package awscdkcontainersecsserviceextensions

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_Container) validateAddContainerMutatingHookParameters(hook ContainerMutatingHook) error {
	return nil
}

func (c *jsiiProxy_Container) validateConnectToServiceParameters(service Service, connectToProps *ConnectToProps) error {
	return nil
}

func (c *jsiiProxy_Container) validateModifyServicePropsParameters(props *ServiceBuild) error {
	return nil
}

func (c *jsiiProxy_Container) validateModifyTaskDefinitionPropsParameters(props *awsecs.TaskDefinitionProps) error {
	return nil
}

func (c *jsiiProxy_Container) validatePrehookParameters(service Service, scope constructs.Construct) error {
	return nil
}

func (c *jsiiProxy_Container) validateUseServiceParameters(service interface{}) error {
	return nil
}

func (c *jsiiProxy_Container) validateUseTaskDefinitionParameters(taskDefinition awsecs.TaskDefinition) error {
	return nil
}

func (j *jsiiProxy_Container) validateSetContainerMutatingHooksParameters(val *[]ContainerMutatingHook) error {
	return nil
}

func (j *jsiiProxy_Container) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Container) validateSetParentServiceParameters(val Service) error {
	return nil
}

func (j *jsiiProxy_Container) validateSetScopeParameters(val constructs.Construct) error {
	return nil
}

func validateNewContainerParameters(props *ContainerExtensionProps) error {
	return nil
}

