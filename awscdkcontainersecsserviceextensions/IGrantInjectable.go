package awscdkcontainersecsserviceextensions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
)

// An interface that will be implemented by all the injectable resources that need to grant permissions to the task role.
// Experimental.
type IGrantInjectable interface {
	IInjectable
	// Experimental.
	Grant(taskDefinition awsecs.TaskDefinition)
}

// The jsii proxy for IGrantInjectable
type jsiiProxy_IGrantInjectable struct {
	jsiiProxy_IInjectable
}

func (i *jsiiProxy_IGrantInjectable) Grant(taskDefinition awsecs.TaskDefinition) {
	if err := i.validateGrantParameters(taskDefinition); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"grant",
		[]interface{}{taskDefinition},
	)
}

