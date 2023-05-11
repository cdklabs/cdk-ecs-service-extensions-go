//go:build !no_runtime_type_checking

package awscdkcontainersecsserviceextensions

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/constructs-go/constructs/v10"
)

func (s *jsiiProxy_Service) validateAddURLParameters(urlName *string, url *string) error {
	if urlName == nil {
		return fmt.Errorf("parameter urlName is required, but nil was provided")
	}

	if url == nil {
		return fmt.Errorf("parameter url is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_Service) validateConnectToParameters(service Service, connectToProps *ConnectToProps) error {
	if service == nil {
		return fmt.Errorf("parameter service is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(connectToProps, func() string { return "parameter connectToProps" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_Service) validateGetURLParameters(urlName *string) error {
	if urlName == nil {
		return fmt.Errorf("parameter urlName is required, but nil was provided")
	}

	return nil
}

func validateService_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Service) validateSetEcsServiceParameters(val interface{}) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	switch val.(type) {
	case awsecs.Ec2Service:
		// ok
	case awsecs.FargateService:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: awsecs.Ec2Service, awsecs.FargateService; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_Service) validateSetTaskDefinitionParameters(val awsecs.TaskDefinition) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewServiceParameters(scope constructs.Construct, id *string, props *ServiceProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

