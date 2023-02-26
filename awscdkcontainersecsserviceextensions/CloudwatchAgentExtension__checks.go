//go:build !no_runtime_type_checking

// The CDK Construct Library that helps you build ECS services using simple extensions
package awscdkcontainersecsserviceextensions

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/constructs-go/constructs/v10"
)

func (c *jsiiProxy_CloudwatchAgentExtension) validateAddContainerMutatingHookParameters(hook ContainerMutatingHook) error {
	if hook == nil {
		return fmt.Errorf("parameter hook is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudwatchAgentExtension) validateConnectToServiceParameters(service Service, connectToProps *ConnectToProps) error {
	if service == nil {
		return fmt.Errorf("parameter service is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(connectToProps, func() string { return "parameter connectToProps" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CloudwatchAgentExtension) validateModifyServicePropsParameters(props *ServiceBuild) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CloudwatchAgentExtension) validateModifyTaskDefinitionPropsParameters(props *awsecs.TaskDefinitionProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CloudwatchAgentExtension) validatePrehookParameters(service Service, scope constructs.Construct) error {
	if service == nil {
		return fmt.Errorf("parameter service is required, but nil was provided")
	}

	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudwatchAgentExtension) validateUseServiceParameters(service interface{}) error {
	if service == nil {
		return fmt.Errorf("parameter service is required, but nil was provided")
	}
	switch service.(type) {
	case awsecs.Ec2Service:
		// ok
	case awsecs.FargateService:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(service) {
			return fmt.Errorf("parameter service must be one of the allowed types: awsecs.Ec2Service, awsecs.FargateService; received %#v (a %T)", service, service)
		}
	}

	return nil
}

func (c *jsiiProxy_CloudwatchAgentExtension) validateUseTaskDefinitionParameters(taskDefinition awsecs.TaskDefinition) error {
	if taskDefinition == nil {
		return fmt.Errorf("parameter taskDefinition is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CloudwatchAgentExtension) validateSetContainerMutatingHooksParameters(val *[]ContainerMutatingHook) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CloudwatchAgentExtension) validateSetNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CloudwatchAgentExtension) validateSetParentServiceParameters(val Service) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CloudwatchAgentExtension) validateSetScopeParameters(val constructs.Construct) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}
