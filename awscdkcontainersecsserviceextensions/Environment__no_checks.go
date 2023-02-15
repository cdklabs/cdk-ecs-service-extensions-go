//go:build no_runtime_type_checking

// The CDK Construct Library that helps you build ECS services using simple extensions
package awscdkcontainersecsserviceextensions

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Environment) validateAddDefaultCloudMapNamespaceParameters(options *awsecs.CloudMapNamespaceOptions) error {
	return nil
}

func validateEnvironment_FromEnvironmentAttributesParameters(scope constructs.Construct, id *string, attrs *EnvironmentAttributes) error {
	return nil
}

func validateEnvironment_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewEnvironmentParameters(scope constructs.Construct, id *string, props *EnvironmentProps) error {
	return nil
}

