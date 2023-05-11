//go:build no_runtime_type_checking

package awscdkcontainersecsserviceextensions

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InjecterExtension) validateAddContainerMutatingHookParameters(hook ContainerMutatingHook) error {
	return nil
}

func (i *jsiiProxy_InjecterExtension) validateConnectToServiceParameters(service Service, connectToProps *ConnectToProps) error {
	return nil
}

func (i *jsiiProxy_InjecterExtension) validateModifyServicePropsParameters(props *ServiceBuild) error {
	return nil
}

func (i *jsiiProxy_InjecterExtension) validateModifyTaskDefinitionPropsParameters(props *awsecs.TaskDefinitionProps) error {
	return nil
}

func (i *jsiiProxy_InjecterExtension) validatePrehookParameters(service Service, scope constructs.Construct) error {
	return nil
}

func (i *jsiiProxy_InjecterExtension) validateUseServiceParameters(service interface{}) error {
	return nil
}

func (i *jsiiProxy_InjecterExtension) validateUseTaskDefinitionParameters(taskDefinition awsecs.TaskDefinition) error {
	return nil
}

func (j *jsiiProxy_InjecterExtension) validateSetContainerMutatingHooksParameters(val *[]ContainerMutatingHook) error {
	return nil
}

func (j *jsiiProxy_InjecterExtension) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_InjecterExtension) validateSetParentServiceParameters(val Service) error {
	return nil
}

func (j *jsiiProxy_InjecterExtension) validateSetScopeParameters(val constructs.Construct) error {
	return nil
}

func validateNewInjecterExtensionParameters(props *InjecterExtensionProps) error {
	return nil
}

