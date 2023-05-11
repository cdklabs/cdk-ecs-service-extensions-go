package awscdkcontainersecsserviceextensions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-ecs-service-extensions-go/awscdkcontainersecsserviceextensions/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/constructs-go/constructs/v10"
)

// This extension accepts a list of `Injectable` resources that the parent service can publish events or write data to.
//
// It sets up the corresponding permissions for the task role of the parent service.
// Experimental.
type InjecterExtension interface {
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
	// Add hooks to the main application extension so that it is modified to add the injectable resource environment variables to the container environment.
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
	// After the task definition has been created, this hook grants the required permissions to the task role for the parent service.
	// Experimental.
	UseTaskDefinition(taskDefinition awsecs.TaskDefinition)
}

// The jsii proxy struct for InjecterExtension
type jsiiProxy_InjecterExtension struct {
	jsiiProxy_ServiceExtension
}

func (j *jsiiProxy_InjecterExtension) Container() awsecs.ContainerDefinition {
	var returns awsecs.ContainerDefinition
	_jsii_.Get(
		j,
		"container",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InjecterExtension) ContainerMutatingHooks() *[]ContainerMutatingHook {
	var returns *[]ContainerMutatingHook
	_jsii_.Get(
		j,
		"containerMutatingHooks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InjecterExtension) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InjecterExtension) ParentService() Service {
	var returns Service
	_jsii_.Get(
		j,
		"parentService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InjecterExtension) Scope() constructs.Construct {
	var returns constructs.Construct
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}


// Experimental.
func NewInjecterExtension(props *InjecterExtensionProps) InjecterExtension {
	_init_.Initialize()

	if err := validateNewInjecterExtensionParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_InjecterExtension{}

	_jsii_.Create(
		"@aws-cdk-containers/ecs-service-extensions.InjecterExtension",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewInjecterExtension_Override(i InjecterExtension, props *InjecterExtensionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk-containers/ecs-service-extensions.InjecterExtension",
		[]interface{}{props},
		i,
	)
}

func (j *jsiiProxy_InjecterExtension)SetContainer(val awsecs.ContainerDefinition) {
	_jsii_.Set(
		j,
		"container",
		val,
	)
}

func (j *jsiiProxy_InjecterExtension)SetContainerMutatingHooks(val *[]ContainerMutatingHook) {
	if err := j.validateSetContainerMutatingHooksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerMutatingHooks",
		val,
	)
}

func (j *jsiiProxy_InjecterExtension)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_InjecterExtension)SetParentService(val Service) {
	if err := j.validateSetParentServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parentService",
		val,
	)
}

func (j *jsiiProxy_InjecterExtension)SetScope(val constructs.Construct) {
	if err := j.validateSetScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (i *jsiiProxy_InjecterExtension) AddContainerMutatingHook(hook ContainerMutatingHook) {
	if err := i.validateAddContainerMutatingHookParameters(hook); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addContainerMutatingHook",
		[]interface{}{hook},
	)
}

func (i *jsiiProxy_InjecterExtension) AddHooks() {
	_jsii_.InvokeVoid(
		i,
		"addHooks",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InjecterExtension) ConnectToService(service Service, connectToProps *ConnectToProps) {
	if err := i.validateConnectToServiceParameters(service, connectToProps); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"connectToService",
		[]interface{}{service, connectToProps},
	)
}

func (i *jsiiProxy_InjecterExtension) ModifyServiceProps(props *ServiceBuild) *ServiceBuild {
	if err := i.validateModifyServicePropsParameters(props); err != nil {
		panic(err)
	}
	var returns *ServiceBuild

	_jsii_.Invoke(
		i,
		"modifyServiceProps",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InjecterExtension) ModifyTaskDefinitionProps(props *awsecs.TaskDefinitionProps) *awsecs.TaskDefinitionProps {
	if err := i.validateModifyTaskDefinitionPropsParameters(props); err != nil {
		panic(err)
	}
	var returns *awsecs.TaskDefinitionProps

	_jsii_.Invoke(
		i,
		"modifyTaskDefinitionProps",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InjecterExtension) Prehook(service Service, scope constructs.Construct) {
	if err := i.validatePrehookParameters(service, scope); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"prehook",
		[]interface{}{service, scope},
	)
}

func (i *jsiiProxy_InjecterExtension) ResolveContainerDependencies() {
	_jsii_.InvokeVoid(
		i,
		"resolveContainerDependencies",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InjecterExtension) UseService(service interface{}) {
	if err := i.validateUseServiceParameters(service); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"useService",
		[]interface{}{service},
	)
}

func (i *jsiiProxy_InjecterExtension) UseTaskDefinition(taskDefinition awsecs.TaskDefinition) {
	if err := i.validateUseTaskDefinitionParameters(taskDefinition); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"useTaskDefinition",
		[]interface{}{taskDefinition},
	)
}

