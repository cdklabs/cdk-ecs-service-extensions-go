//go:build !no_runtime_type_checking

package awscdkcontainersecsserviceextensions

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
)

func (i *jsiiProxy_IGrantInjectable) validateGrantParameters(taskDefinition awsecs.TaskDefinition) error {
	if taskDefinition == nil {
		return fmt.Errorf("parameter taskDefinition is required, but nil was provided")
	}

	return nil
}

