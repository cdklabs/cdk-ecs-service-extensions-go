package awscdkcontainersecsserviceextensions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-ecs-service-extensions-go/awscdkcontainersecsserviceextensions/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/constructs-go/constructs/v10"
)

// This extension adds a FluentBit log router to the task definition and does all the configuration necessarily to enable log routing for the task using FireLens.
// Experimental.
type FireLensExtension interface {
	ServiceExtension
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

// The jsii proxy struct for FireLensExtension
type jsiiProxy_FireLensExtension struct {
	jsiiProxy_ServiceExtension
}

func (j *jsiiProxy_FireLensExtension) Container() awsecs.ContainerDefinition {
	var returns awsecs.ContainerDefinition
	_jsii_.Get(
		j,
		"container",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FireLensExtension) ContainerMutatingHooks() *[]ContainerMutatingHook {
	var returns *[]ContainerMutatingHook
	_jsii_.Get(
		j,
		"containerMutatingHooks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FireLensExtension) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FireLensExtension) ParentService() Service {
	var returns Service
	_jsii_.Get(
		j,
		"parentService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FireLensExtension) Scope() constructs.Construct {
	var returns constructs.Construct
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}


// Experimental.
func NewFireLensExtension() FireLensExtension {
	_init_.Initialize()

	j := jsiiProxy_FireLensExtension{}

	_jsii_.Create(
		"@aws-cdk-containers/ecs-service-extensions.FireLensExtension",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewFireLensExtension_Override(f FireLensExtension) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk-containers/ecs-service-extensions.FireLensExtension",
		nil, // no parameters
		f,
	)
}

func (j *jsiiProxy_FireLensExtension)SetContainer(val awsecs.ContainerDefinition) {
	_jsii_.Set(
		j,
		"container",
		val,
	)
}

func (j *jsiiProxy_FireLensExtension)SetContainerMutatingHooks(val *[]ContainerMutatingHook) {
	if err := j.validateSetContainerMutatingHooksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerMutatingHooks",
		val,
	)
}

func (j *jsiiProxy_FireLensExtension)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_FireLensExtension)SetParentService(val Service) {
	if err := j.validateSetParentServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parentService",
		val,
	)
}

func (j *jsiiProxy_FireLensExtension)SetScope(val constructs.Construct) {
	if err := j.validateSetScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (f *jsiiProxy_FireLensExtension) AddContainerMutatingHook(hook ContainerMutatingHook) {
	if err := f.validateAddContainerMutatingHookParameters(hook); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addContainerMutatingHook",
		[]interface{}{hook},
	)
}

func (f *jsiiProxy_FireLensExtension) AddHooks() {
	_jsii_.InvokeVoid(
		f,
		"addHooks",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FireLensExtension) ConnectToService(service Service, connectToProps *ConnectToProps) {
	if err := f.validateConnectToServiceParameters(service, connectToProps); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"connectToService",
		[]interface{}{service, connectToProps},
	)
}

func (f *jsiiProxy_FireLensExtension) ModifyServiceProps(props *ServiceBuild) *ServiceBuild {
	if err := f.validateModifyServicePropsParameters(props); err != nil {
		panic(err)
	}
	var returns *ServiceBuild

	_jsii_.Invoke(
		f,
		"modifyServiceProps",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FireLensExtension) ModifyTaskDefinitionProps(props *awsecs.TaskDefinitionProps) *awsecs.TaskDefinitionProps {
	if err := f.validateModifyTaskDefinitionPropsParameters(props); err != nil {
		panic(err)
	}
	var returns *awsecs.TaskDefinitionProps

	_jsii_.Invoke(
		f,
		"modifyTaskDefinitionProps",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FireLensExtension) Prehook(service Service, scope constructs.Construct) {
	if err := f.validatePrehookParameters(service, scope); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"prehook",
		[]interface{}{service, scope},
	)
}

func (f *jsiiProxy_FireLensExtension) ResolveContainerDependencies() {
	_jsii_.InvokeVoid(
		f,
		"resolveContainerDependencies",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FireLensExtension) UseService(service interface{}) {
	if err := f.validateUseServiceParameters(service); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"useService",
		[]interface{}{service},
	)
}

func (f *jsiiProxy_FireLensExtension) UseTaskDefinition(taskDefinition awsecs.TaskDefinition) {
	if err := f.validateUseTaskDefinitionParameters(taskDefinition); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"useTaskDefinition",
		[]interface{}{taskDefinition},
	)
}

