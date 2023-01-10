// The CDK Construct Library that helps you build ECS services using simple extensions
package awscdkcontainersecsserviceextensions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-ecs-service-extensions-go/awscdkcontainersecsserviceextensions/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
)

// This hook modifies the application container's settings so that its primary port mapping has a name.
// Experimental.
type AliasedPortMutatingHook interface {
	ContainerMutatingHook
	// This is a hook for modifying the container definition of any upstream containers.
	//
	// This is primarily used for the main application container.
	// For example, the Firelens extension wants to be able to modify the logging
	// settings of the application container.
	// Experimental.
	MutateContainerDefinition(props *awsecs.ContainerDefinitionOptions) *awsecs.ContainerDefinitionOptions
}

// The jsii proxy struct for AliasedPortMutatingHook
type jsiiProxy_AliasedPortMutatingHook struct {
	jsiiProxy_ContainerMutatingHook
}

// Experimental.
func NewAliasedPortMutatingHook(props *AliasedPortMutatingHookProps) AliasedPortMutatingHook {
	_init_.Initialize()

	if err := validateNewAliasedPortMutatingHookParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AliasedPortMutatingHook{}

	_jsii_.Create(
		"@aws-cdk-containers/ecs-service-extensions.AliasedPortMutatingHook",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewAliasedPortMutatingHook_Override(a AliasedPortMutatingHook, props *AliasedPortMutatingHookProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk-containers/ecs-service-extensions.AliasedPortMutatingHook",
		[]interface{}{props},
		a,
	)
}

func (a *jsiiProxy_AliasedPortMutatingHook) MutateContainerDefinition(props *awsecs.ContainerDefinitionOptions) *awsecs.ContainerDefinitionOptions {
	if err := a.validateMutateContainerDefinitionParameters(props); err != nil {
		panic(err)
	}
	var returns *awsecs.ContainerDefinitionOptions

	_jsii_.Invoke(
		a,
		"mutateContainerDefinition",
		[]interface{}{props},
		&returns,
	)

	return returns
}

