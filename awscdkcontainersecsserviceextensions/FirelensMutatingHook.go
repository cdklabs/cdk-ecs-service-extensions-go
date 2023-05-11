package awscdkcontainersecsserviceextensions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-ecs-service-extensions-go/awscdkcontainersecsserviceextensions/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
)

// This hook modifies the application container's settings so that it routes logs using FireLens.
// Experimental.
type FirelensMutatingHook interface {
	ContainerMutatingHook
	// This is a hook for modifying the container definition of any upstream containers.
	//
	// This is primarily used for the main application container.
	// For example, the Firelens extension wants to be able to modify the logging
	// settings of the application container.
	// Experimental.
	MutateContainerDefinition(props *awsecs.ContainerDefinitionOptions) *awsecs.ContainerDefinitionOptions
}

// The jsii proxy struct for FirelensMutatingHook
type jsiiProxy_FirelensMutatingHook struct {
	jsiiProxy_ContainerMutatingHook
}

// Experimental.
func NewFirelensMutatingHook(props *FirelensProps) FirelensMutatingHook {
	_init_.Initialize()

	if err := validateNewFirelensMutatingHookParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_FirelensMutatingHook{}

	_jsii_.Create(
		"@aws-cdk-containers/ecs-service-extensions.FirelensMutatingHook",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewFirelensMutatingHook_Override(f FirelensMutatingHook, props *FirelensProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk-containers/ecs-service-extensions.FirelensMutatingHook",
		[]interface{}{props},
		f,
	)
}

func (f *jsiiProxy_FirelensMutatingHook) MutateContainerDefinition(props *awsecs.ContainerDefinitionOptions) *awsecs.ContainerDefinitionOptions {
	if err := f.validateMutateContainerDefinitionParameters(props); err != nil {
		panic(err)
	}
	var returns *awsecs.ContainerDefinitionOptions

	_jsii_.Invoke(
		f,
		"mutateContainerDefinition",
		[]interface{}{props},
		&returns,
	)

	return returns
}

