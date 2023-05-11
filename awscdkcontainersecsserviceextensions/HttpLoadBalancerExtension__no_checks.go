//go:build no_runtime_type_checking

package awscdkcontainersecsserviceextensions

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HttpLoadBalancerExtension) validateAddContainerMutatingHookParameters(hook ContainerMutatingHook) error {
	return nil
}

func (h *jsiiProxy_HttpLoadBalancerExtension) validateConnectToServiceParameters(service Service, connectToProps *ConnectToProps) error {
	return nil
}

func (h *jsiiProxy_HttpLoadBalancerExtension) validateModifyServicePropsParameters(props *ServiceBuild) error {
	return nil
}

func (h *jsiiProxy_HttpLoadBalancerExtension) validateModifyTaskDefinitionPropsParameters(props *awsecs.TaskDefinitionProps) error {
	return nil
}

func (h *jsiiProxy_HttpLoadBalancerExtension) validatePrehookParameters(service Service, scope constructs.Construct) error {
	return nil
}

func (h *jsiiProxy_HttpLoadBalancerExtension) validateUseServiceParameters(service interface{}) error {
	return nil
}

func (h *jsiiProxy_HttpLoadBalancerExtension) validateUseTaskDefinitionParameters(taskDefinition awsecs.TaskDefinition) error {
	return nil
}

func (j *jsiiProxy_HttpLoadBalancerExtension) validateSetContainerMutatingHooksParameters(val *[]ContainerMutatingHook) error {
	return nil
}

func (j *jsiiProxy_HttpLoadBalancerExtension) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_HttpLoadBalancerExtension) validateSetParentServiceParameters(val Service) error {
	return nil
}

func (j *jsiiProxy_HttpLoadBalancerExtension) validateSetScopeParameters(val constructs.Construct) error {
	return nil
}

func validateNewHttpLoadBalancerExtensionParameters(props *HttpLoadBalancerProps) error {
	return nil
}

