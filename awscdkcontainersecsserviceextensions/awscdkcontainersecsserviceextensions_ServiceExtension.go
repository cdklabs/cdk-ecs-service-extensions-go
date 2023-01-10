// The CDK Construct Library that helps you build ECS services using simple extensions
package awscdkcontainersecsserviceextensions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-ecs-service-extensions-go/awscdkcontainersecsserviceextensions/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/constructs-go/constructs/v10"
)

// The shape of a service extension.
//
// This abstract class is implemented
// by other extensions that extend the hooks to implement any custom
// logic that they want to run during each step of preparing the service.
// Experimental.
type ServiceExtension interface {
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
	Prehook(parent Service, scope constructs.Construct)
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

// The jsii proxy struct for ServiceExtension
type jsiiProxy_ServiceExtension struct {
	_ byte // padding
}

func (j *jsiiProxy_ServiceExtension) Container() awsecs.ContainerDefinition {
	var returns awsecs.ContainerDefinition
	_jsii_.Get(
		j,
		"container",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceExtension) ContainerMutatingHooks() *[]ContainerMutatingHook {
	var returns *[]ContainerMutatingHook
	_jsii_.Get(
		j,
		"containerMutatingHooks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceExtension) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceExtension) ParentService() Service {
	var returns Service
	_jsii_.Get(
		j,
		"parentService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceExtension) Scope() constructs.Construct {
	var returns constructs.Construct
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}


// Experimental.
func NewServiceExtension_Override(s ServiceExtension, name *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk-containers/ecs-service-extensions.ServiceExtension",
		[]interface{}{name},
		s,
	)
}

func (j *jsiiProxy_ServiceExtension)SetContainer(val awsecs.ContainerDefinition) {
	_jsii_.Set(
		j,
		"container",
		val,
	)
}

func (j *jsiiProxy_ServiceExtension)SetContainerMutatingHooks(val *[]ContainerMutatingHook) {
	if err := j.validateSetContainerMutatingHooksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerMutatingHooks",
		val,
	)
}

func (j *jsiiProxy_ServiceExtension)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ServiceExtension)SetParentService(val Service) {
	if err := j.validateSetParentServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parentService",
		val,
	)
}

func (j *jsiiProxy_ServiceExtension)SetScope(val constructs.Construct) {
	if err := j.validateSetScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (s *jsiiProxy_ServiceExtension) AddContainerMutatingHook(hook ContainerMutatingHook) {
	if err := s.validateAddContainerMutatingHookParameters(hook); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addContainerMutatingHook",
		[]interface{}{hook},
	)
}

func (s *jsiiProxy_ServiceExtension) AddHooks() {
	_jsii_.InvokeVoid(
		s,
		"addHooks",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceExtension) ConnectToService(service Service, connectToProps *ConnectToProps) {
	if err := s.validateConnectToServiceParameters(service, connectToProps); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"connectToService",
		[]interface{}{service, connectToProps},
	)
}

func (s *jsiiProxy_ServiceExtension) ModifyServiceProps(props *ServiceBuild) *ServiceBuild {
	if err := s.validateModifyServicePropsParameters(props); err != nil {
		panic(err)
	}
	var returns *ServiceBuild

	_jsii_.Invoke(
		s,
		"modifyServiceProps",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceExtension) ModifyTaskDefinitionProps(props *awsecs.TaskDefinitionProps) *awsecs.TaskDefinitionProps {
	if err := s.validateModifyTaskDefinitionPropsParameters(props); err != nil {
		panic(err)
	}
	var returns *awsecs.TaskDefinitionProps

	_jsii_.Invoke(
		s,
		"modifyTaskDefinitionProps",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceExtension) Prehook(parent Service, scope constructs.Construct) {
	if err := s.validatePrehookParameters(parent, scope); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"prehook",
		[]interface{}{parent, scope},
	)
}

func (s *jsiiProxy_ServiceExtension) ResolveContainerDependencies() {
	_jsii_.InvokeVoid(
		s,
		"resolveContainerDependencies",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceExtension) UseService(service interface{}) {
	if err := s.validateUseServiceParameters(service); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"useService",
		[]interface{}{service},
	)
}

func (s *jsiiProxy_ServiceExtension) UseTaskDefinition(taskDefinition awsecs.TaskDefinition) {
	if err := s.validateUseTaskDefinitionParameters(taskDefinition); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"useTaskDefinition",
		[]interface{}{taskDefinition},
	)
}

