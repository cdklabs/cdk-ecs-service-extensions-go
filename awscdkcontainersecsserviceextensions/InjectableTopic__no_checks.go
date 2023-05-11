//go:build no_runtime_type_checking

package awscdkcontainersecsserviceextensions

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InjectableTopic) validateGrantParameters(taskDefinition awsecs.TaskDefinition) error {
	return nil
}

func validateNewInjectableTopicParameters(props *InjectableTopicProps) error {
	return nil
}

