//go:build no_runtime_type_checking

// The CDK Construct Library that helps you build ECS services using simple extensions
package awscdkcontainersecsserviceextensions

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Service) validateAddURLParameters(urlName *string, url *string) error {
	return nil
}

func (s *jsiiProxy_Service) validateConnectToParameters(service Service, connectToProps *ConnectToProps) error {
	return nil
}

func (s *jsiiProxy_Service) validateGetURLParameters(urlName *string) error {
	return nil
}

func validateService_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Service) validateSetEcsServiceParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Service) validateSetTaskDefinitionParameters(val awsecs.TaskDefinition) error {
	return nil
}

func validateNewServiceParameters(scope constructs.Construct, id *string, props *ServiceProps) error {
	return nil
}

