//go:build no_runtime_type_checking

// The CDK Construct Library that helps you build ECS services using simple extensions
package awscdkcontainersecsserviceextensions

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServiceDescription) validateAddParameters(extension ServiceExtension) error {
	return nil
}

func (s *jsiiProxy_ServiceDescription) validateGetParameters(name *string) error {
	return nil
}

func (j *jsiiProxy_ServiceDescription) validateSetExtensionsParameters(val *map[string]ServiceExtension) error {
	return nil
}

