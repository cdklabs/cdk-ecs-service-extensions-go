// The CDK Construct Library that helps you build ECS services using simple extensions
package awscdkcontainersecsserviceextensions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-ecs-service-extensions-go/awscdkcontainersecsserviceextensions/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappmesh"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/constructs-go/constructs/v10"
)

// This extension adds an Envoy sidecar to the task definition and creates the App Mesh resources required to route network traffic to the container in a service mesh.
//
// The service will then be available to other App Mesh services at the
// address `<service name>.<environment name>`. For example, a service called
// `orders` deploying in an environment called `production` would be accessible
// to other App Mesh enabled services at the address `http://orders.production`.
// Experimental.
type AppMeshExtension interface {
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
	// The protocol used for AppMesh routing.
	//
	// default - Protocol.HTTP
	// Experimental.
	Protocol() Protocol
	// Experimental.
	Route() awsappmesh.Route
	// Experimental.
	SetRoute(val awsappmesh.Route)
	// Experimental.
	Scope() constructs.Construct
	// Experimental.
	SetScope(val constructs.Construct)
	// Experimental.
	VirtualNode() awsappmesh.VirtualNode
	// Experimental.
	SetVirtualNode(val awsappmesh.VirtualNode)
	// Experimental.
	VirtualRouter() awsappmesh.VirtualRouter
	// Experimental.
	SetVirtualRouter(val awsappmesh.VirtualRouter)
	// Experimental.
	VirtualService() awsappmesh.VirtualService
	// Experimental.
	SetVirtualService(val awsappmesh.VirtualService)
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
	ConnectToService(otherService Service, _connectToProps *ConnectToProps)
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

// The jsii proxy struct for AppMeshExtension
type jsiiProxy_AppMeshExtension struct {
	jsiiProxy_ServiceExtension
}

func (j *jsiiProxy_AppMeshExtension) Container() awsecs.ContainerDefinition {
	var returns awsecs.ContainerDefinition
	_jsii_.Get(
		j,
		"container",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppMeshExtension) ContainerMutatingHooks() *[]ContainerMutatingHook {
	var returns *[]ContainerMutatingHook
	_jsii_.Get(
		j,
		"containerMutatingHooks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppMeshExtension) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppMeshExtension) ParentService() Service {
	var returns Service
	_jsii_.Get(
		j,
		"parentService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppMeshExtension) Protocol() Protocol {
	var returns Protocol
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppMeshExtension) Route() awsappmesh.Route {
	var returns awsappmesh.Route
	_jsii_.Get(
		j,
		"route",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppMeshExtension) Scope() constructs.Construct {
	var returns constructs.Construct
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppMeshExtension) VirtualNode() awsappmesh.VirtualNode {
	var returns awsappmesh.VirtualNode
	_jsii_.Get(
		j,
		"virtualNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppMeshExtension) VirtualRouter() awsappmesh.VirtualRouter {
	var returns awsappmesh.VirtualRouter
	_jsii_.Get(
		j,
		"virtualRouter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppMeshExtension) VirtualService() awsappmesh.VirtualService {
	var returns awsappmesh.VirtualService
	_jsii_.Get(
		j,
		"virtualService",
		&returns,
	)
	return returns
}


// Experimental.
func NewAppMeshExtension(props *MeshProps) AppMeshExtension {
	_init_.Initialize()

	if err := validateNewAppMeshExtensionParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppMeshExtension{}

	_jsii_.Create(
		"@aws-cdk-containers/ecs-service-extensions.AppMeshExtension",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewAppMeshExtension_Override(a AppMeshExtension, props *MeshProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk-containers/ecs-service-extensions.AppMeshExtension",
		[]interface{}{props},
		a,
	)
}

func (j *jsiiProxy_AppMeshExtension)SetContainer(val awsecs.ContainerDefinition) {
	_jsii_.Set(
		j,
		"container",
		val,
	)
}

func (j *jsiiProxy_AppMeshExtension)SetContainerMutatingHooks(val *[]ContainerMutatingHook) {
	if err := j.validateSetContainerMutatingHooksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerMutatingHooks",
		val,
	)
}

func (j *jsiiProxy_AppMeshExtension)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AppMeshExtension)SetParentService(val Service) {
	if err := j.validateSetParentServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parentService",
		val,
	)
}

func (j *jsiiProxy_AppMeshExtension)SetRoute(val awsappmesh.Route) {
	if err := j.validateSetRouteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"route",
		val,
	)
}

func (j *jsiiProxy_AppMeshExtension)SetScope(val constructs.Construct) {
	if err := j.validateSetScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_AppMeshExtension)SetVirtualNode(val awsappmesh.VirtualNode) {
	if err := j.validateSetVirtualNodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualNode",
		val,
	)
}

func (j *jsiiProxy_AppMeshExtension)SetVirtualRouter(val awsappmesh.VirtualRouter) {
	if err := j.validateSetVirtualRouterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualRouter",
		val,
	)
}

func (j *jsiiProxy_AppMeshExtension)SetVirtualService(val awsappmesh.VirtualService) {
	if err := j.validateSetVirtualServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualService",
		val,
	)
}

func (a *jsiiProxy_AppMeshExtension) AddContainerMutatingHook(hook ContainerMutatingHook) {
	if err := a.validateAddContainerMutatingHookParameters(hook); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addContainerMutatingHook",
		[]interface{}{hook},
	)
}

func (a *jsiiProxy_AppMeshExtension) AddHooks() {
	_jsii_.InvokeVoid(
		a,
		"addHooks",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppMeshExtension) ConnectToService(otherService Service, _connectToProps *ConnectToProps) {
	if err := a.validateConnectToServiceParameters(otherService, _connectToProps); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"connectToService",
		[]interface{}{otherService, _connectToProps},
	)
}

func (a *jsiiProxy_AppMeshExtension) ModifyServiceProps(props *ServiceBuild) *ServiceBuild {
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

func (a *jsiiProxy_AppMeshExtension) ModifyTaskDefinitionProps(props *awsecs.TaskDefinitionProps) *awsecs.TaskDefinitionProps {
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

func (a *jsiiProxy_AppMeshExtension) Prehook(service Service, scope constructs.Construct) {
	if err := a.validatePrehookParameters(service, scope); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"prehook",
		[]interface{}{service, scope},
	)
}

func (a *jsiiProxy_AppMeshExtension) ResolveContainerDependencies() {
	_jsii_.InvokeVoid(
		a,
		"resolveContainerDependencies",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppMeshExtension) UseService(service interface{}) {
	if err := a.validateUseServiceParameters(service); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"useService",
		[]interface{}{service},
	)
}

func (a *jsiiProxy_AppMeshExtension) UseTaskDefinition(taskDefinition awsecs.TaskDefinition) {
	if err := a.validateUseTaskDefinitionParameters(taskDefinition); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"useTaskDefinition",
		[]interface{}{taskDefinition},
	)
}

