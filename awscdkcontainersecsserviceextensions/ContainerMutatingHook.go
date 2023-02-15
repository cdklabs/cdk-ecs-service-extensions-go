// The CDK Construct Library that helps you build ECS services using simple extensions
package awscdkcontainersecsserviceextensions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-ecs-service-extensions-go/awscdkcontainersecsserviceextensions/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
)

// This is an abstract class wrapper for a mutating hook.
//
// It is
// extended by any extension which wants to mutate other extension's containers.
// Experimental.
type ContainerMutatingHook interface {
	// This is a hook for modifying the container definition of any upstream containers.
	//
	// This is primarily used for the main application container.
	// For example, the Firelens extension wants to be able to modify the logging
	// settings of the application container.
	// Experimental.
	MutateContainerDefinition(props *awsecs.ContainerDefinitionOptions) *awsecs.ContainerDefinitionOptions
}

// The jsii proxy struct for ContainerMutatingHook
type jsiiProxy_ContainerMutatingHook struct {
	_ byte // padding
}

// Experimental.
func NewContainerMutatingHook_Override(c ContainerMutatingHook) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk-containers/ecs-service-extensions.ContainerMutatingHook",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_ContainerMutatingHook) MutateContainerDefinition(props *awsecs.ContainerDefinitionOptions) *awsecs.ContainerDefinitionOptions {
	if err := c.validateMutateContainerDefinitionParameters(props); err != nil {
		panic(err)
	}
	var returns *awsecs.ContainerDefinitionOptions

	_jsii_.Invoke(
		c,
		"mutateContainerDefinition",
		[]interface{}{props},
		&returns,
	)

	return returns
}

