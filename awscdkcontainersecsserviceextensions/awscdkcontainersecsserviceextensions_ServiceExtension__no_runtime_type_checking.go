//go:build no_runtime_type_checking

// The CDK Construct Library that helps you build ECS services using simple extensions
package awscdkcontainersecsserviceextensions

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServiceExtension) validateAddContainerMutatingHookParameters(hook ContainerMutatingHook) error {
	return nil
}

func (s *jsiiProxy_ServiceExtension) validateConnectToServiceParameters(service Service, connectToProps *ConnectToProps) error {
	return nil
}

func (s *jsiiProxy_ServiceExtension) validateModifyServicePropsParameters(props *ServiceBuild) error {
	return nil
}

func (s *jsiiProxy_ServiceExtension) validateModifyTaskDefinitionPropsParameters(props *awsecs.TaskDefinitionProps) error {
	return nil
}

func (s *jsiiProxy_ServiceExtension) validatePrehookParameters(parent Service, scope constructs.Construct) error {
	return nil
}

func (s *jsiiProxy_ServiceExtension) validateUseServiceParameters(service interface{}) error {
	return nil
}

func (s *jsiiProxy_ServiceExtension) validateUseTaskDefinitionParameters(taskDefinition awsecs.TaskDefinition) error {
	return nil
}

func (j *jsiiProxy_ServiceExtension) validateSetContainerMutatingHooksParameters(val *[]ContainerMutatingHook) error {
	return nil
}

func (j *jsiiProxy_ServiceExtension) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ServiceExtension) validateSetParentServiceParameters(val Service) error {
	return nil
}

func (j *jsiiProxy_ServiceExtension) validateSetScopeParameters(val constructs.Construct) error {
	return nil
}

func validateNewServiceExtensionParameters(name *string) error {
	return nil
}

