package awscdkcontainersecsserviceextensions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-ecs-service-extensions-go/awscdkcontainersecsserviceextensions/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/constructs-go/constructs/v10"
)

// Experimental.
type AliasedPortExtension interface {
	ServiceExtension
	// Experimental.
	Alias() *string
	// Experimental.
	SetAlias(val *string)
	// Experimental.
	AliasPort() *float64
	// Experimental.
	SetAliasPort(val *float64)
	// Experimental.
	AppProtocol() awsecs.AppProtocol
	// Experimental.
	SetAppProtocol(val awsecs.AppProtocol)
	// The container for this extension.
	//
	// Most extensions have a container, but not
	// every extension is required to have a container. Some extensions may just
	// modify the properties of the service, or create external resources
	// connected to the service.
	// Experimental.
	Container() awsecs.ContainerDefinition
	// Experimental.
	SetContainer(val awsecs.ContainerDefinition)
	// Experimental.
	ContainerMutatingHooks() *[]ContainerMutatingHook
	// Experimental.
	SetContainerMutatingHooks(val *[]ContainerMutatingHook)
	// The name of the extension.
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	Namespace() *string
	// Experimental.
	SetNamespace(val *string)
	// The service which this extension is being added to.
	//
	// Initially, extensions are collected into a ServiceDescription, but no service
	// exists yet. Later, when the ServiceDescription is used to create a service,
	// the extension is told what Service it is now working on.
	// Experimental.
	ParentService() Service
	// Experimental.
	SetParentService(val Service)
	// Experimental.
	Scope() constructs.Construct
	// Experimental.
	SetScope(val constructs.Construct)
	// This hook allows another service extension to register a mutating hook for changing the primary container of this extension.
	//
	// This is primarily used
	// for the application extension. For example, the Firelens extension wants to
	// be able to modify the settings of the application container to
	// route logs through Firelens.
	// Experimental.
	AddContainerMutatingHook(hook ContainerMutatingHook)
	// A hook that allows the extension to add hooks to other extensions that are registered.
	// Experimental.
	AddHooks()
	// This hook allows the extension to establish a connection to extensions from another service.
	//
	// Usually used for things like
	// allowing one service to talk to the load balancer or service mesh
	// proxy for another service.
	// Experimental.
	ConnectToService(service Service, connectToProps *ConnectToProps)
	// Prior to launching the task definition as a service, this hook is called on each extension to give it a chance to mutate the properties of the service to be created.
	// Experimental.
	ModifyServiceProps(props *ServiceBuild) *ServiceBuild
	// This is a hook which allows extensions to modify the settings of the task definition prior to it being created.
	//
	// For example, the App Mesh
	// extension needs to configure an Envoy proxy in the task definition,
	// or the Application extension wants to set the overall resource for
	// the task.
	// Experimental.
	ModifyTaskDefinitionProps(props *awsecs.TaskDefinitionProps) *awsecs.TaskDefinitionProps
	// A hook that is called for each extension ahead of time to allow for any initial setup, such as creating resources in advance.
	// Experimental.
	Prehook(service Service, scope constructs.Construct)
	// Once all containers are added to the task definition, this hook is called for each extension to give it a chance to resolve its dependency graph so that its container starts in the right order based on the other extensions that were enabled.
	// Experimental.
	ResolveContainerDependencies()
	// When this hook is implemented by extension, it allows the extension to use the service which has been created.
	//
	// It is generally used to
	// create any final resources which might depend on the service itself.
	// Experimental.
	UseService(service interface{})
	// Once the task definition is created, this hook is called for each extension to give it a chance to add containers to the task definition, change the task definition's role to add permissions, etc.
	// Experimental.
	UseTaskDefinition(taskDefinition awsecs.TaskDefinition)
}

// The jsii proxy struct for AliasedPortExtension
type jsiiProxy_AliasedPortExtension struct {
	jsiiProxy_ServiceExtension
}

func (j *jsiiProxy_AliasedPortExtension) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AliasedPortExtension) AliasPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"aliasPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AliasedPortExtension) AppProtocol() awsecs.AppProtocol {
	var returns awsecs.AppProtocol
	_jsii_.Get(
		j,
		"appProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AliasedPortExtension) Container() awsecs.ContainerDefinition {
	var returns awsecs.ContainerDefinition
	_jsii_.Get(
		j,
		"container",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AliasedPortExtension) ContainerMutatingHooks() *[]ContainerMutatingHook {
	var returns *[]ContainerMutatingHook
	_jsii_.Get(
		j,
		"containerMutatingHooks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AliasedPortExtension) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AliasedPortExtension) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AliasedPortExtension) ParentService() Service {
	var returns Service
	_jsii_.Get(
		j,
		"parentService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AliasedPortExtension) Scope() constructs.Construct {
	var returns constructs.Construct
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}


// Experimental.
func NewAliasedPortExtension(props *AliasedPortProps) AliasedPortExtension {
	_init_.Initialize()

	if err := validateNewAliasedPortExtensionParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AliasedPortExtension{}

	_jsii_.Create(
		"@aws-cdk-containers/ecs-service-extensions.AliasedPortExtension",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewAliasedPortExtension_Override(a AliasedPortExtension, props *AliasedPortProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk-containers/ecs-service-extensions.AliasedPortExtension",
		[]interface{}{props},
		a,
	)
}

func (j *jsiiProxy_AliasedPortExtension)SetAlias(val *string) {
	if err := j.validateSetAliasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_AliasedPortExtension)SetAliasPort(val *float64) {
	_jsii_.Set(
		j,
		"aliasPort",
		val,
	)
}

func (j *jsiiProxy_AliasedPortExtension)SetAppProtocol(val awsecs.AppProtocol) {
	_jsii_.Set(
		j,
		"appProtocol",
		val,
	)
}

func (j *jsiiProxy_AliasedPortExtension)SetContainer(val awsecs.ContainerDefinition) {
	_jsii_.Set(
		j,
		"container",
		val,
	)
}

func (j *jsiiProxy_AliasedPortExtension)SetContainerMutatingHooks(val *[]ContainerMutatingHook) {
	if err := j.validateSetContainerMutatingHooksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerMutatingHooks",
		val,
	)
}

func (j *jsiiProxy_AliasedPortExtension)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AliasedPortExtension)SetNamespace(val *string) {
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_AliasedPortExtension)SetParentService(val Service) {
	if err := j.validateSetParentServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parentService",
		val,
	)
}

func (j *jsiiProxy_AliasedPortExtension)SetScope(val constructs.Construct) {
	if err := j.validateSetScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (a *jsiiProxy_AliasedPortExtension) AddContainerMutatingHook(hook ContainerMutatingHook) {
	if err := a.validateAddContainerMutatingHookParameters(hook); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addContainerMutatingHook",
		[]interface{}{hook},
	)
}

func (a *jsiiProxy_AliasedPortExtension) AddHooks() {
	_jsii_.InvokeVoid(
		a,
		"addHooks",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AliasedPortExtension) ConnectToService(service Service, connectToProps *ConnectToProps) {
	if err := a.validateConnectToServiceParameters(service, connectToProps); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"connectToService",
		[]interface{}{service, connectToProps},
	)
}

func (a *jsiiProxy_AliasedPortExtension) ModifyServiceProps(props *ServiceBuild) *ServiceBuild {
	if err := a.validateModifyServicePropsParameters(props); err != nil {
		panic(err)
	}
	var returns *ServiceBuild

	_jsii_.Invoke(
		a,
		"modifyServiceProps",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AliasedPortExtension) ModifyTaskDefinitionProps(props *awsecs.TaskDefinitionProps) *awsecs.TaskDefinitionProps {
	if err := a.validateModifyTaskDefinitionPropsParameters(props); err != nil {
		panic(err)
	}
	var returns *awsecs.TaskDefinitionProps

	_jsii_.Invoke(
		a,
		"modifyTaskDefinitionProps",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AliasedPortExtension) Prehook(service Service, scope constructs.Construct) {
	if err := a.validatePrehookParameters(service, scope); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"prehook",
		[]interface{}{service, scope},
	)
}

func (a *jsiiProxy_AliasedPortExtension) ResolveContainerDependencies() {
	_jsii_.InvokeVoid(
		a,
		"resolveContainerDependencies",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AliasedPortExtension) UseService(service interface{}) {
	if err := a.validateUseServiceParameters(service); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"useService",
		[]interface{}{service},
	)
}

func (a *jsiiProxy_AliasedPortExtension) UseTaskDefinition(taskDefinition awsecs.TaskDefinition) {
	if err := a.validateUseTaskDefinitionParameters(taskDefinition); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"useTaskDefinition",
		[]interface{}{taskDefinition},
	)
}

