package awscdkcontainersecsserviceextensions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-ecs-service-extensions-go/awscdkcontainersecsserviceextensions/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/constructs-go/constructs/v10"
)

// This extension helps you scale your service according to CPU utilization.
// Deprecated: To enable target tracking based on CPU utilization, use the `targetCpuUtilization` property of `autoScaleTaskCount` in the `Service` construct.
// For more information, please refer https://github.com/aws/aws-cdk/blob/master/packages/%40aws-cdk-containers/ecs-service-extensions/README.md#task-auto-scaling .
type ScaleOnCpuUtilization interface {
	ServiceExtension
	// The container for this extension.
	//
	// Most extensions have a container, but not
	// every extension is required to have a container. Some extensions may just
	// modify the properties of the service, or create external resources
	// connected to the service.
	// Deprecated: To enable target tracking based on CPU utilization, use the `targetCpuUtilization` property of `autoScaleTaskCount` in the `Service` construct.
	// For more information, please refer https://github.com/aws/aws-cdk/blob/master/packages/%40aws-cdk-containers/ecs-service-extensions/README.md#task-auto-scaling .
	Container() awsecs.ContainerDefinition
	// Deprecated: To enable target tracking based on CPU utilization, use the `targetCpuUtilization` property of `autoScaleTaskCount` in the `Service` construct.
	// For more information, please refer https://github.com/aws/aws-cdk/blob/master/packages/%40aws-cdk-containers/ecs-service-extensions/README.md#task-auto-scaling .
	SetContainer(val awsecs.ContainerDefinition)
	// Deprecated: To enable target tracking based on CPU utilization, use the `targetCpuUtilization` property of `autoScaleTaskCount` in the `Service` construct.
	// For more information, please refer https://github.com/aws/aws-cdk/blob/master/packages/%40aws-cdk-containers/ecs-service-extensions/README.md#task-auto-scaling .
	ContainerMutatingHooks() *[]ContainerMutatingHook
	// Deprecated: To enable target tracking based on CPU utilization, use the `targetCpuUtilization` property of `autoScaleTaskCount` in the `Service` construct.
	// For more information, please refer https://github.com/aws/aws-cdk/blob/master/packages/%40aws-cdk-containers/ecs-service-extensions/README.md#task-auto-scaling .
	SetContainerMutatingHooks(val *[]ContainerMutatingHook)
	// How many tasks to launch initially.
	// Deprecated: To enable target tracking based on CPU utilization, use the `targetCpuUtilization` property of `autoScaleTaskCount` in the `Service` construct.
	// For more information, please refer https://github.com/aws/aws-cdk/blob/master/packages/%40aws-cdk-containers/ecs-service-extensions/README.md#task-auto-scaling .
	InitialTaskCount() *float64
	// The maximum number of tasks when scaling out.
	// Deprecated: To enable target tracking based on CPU utilization, use the `targetCpuUtilization` property of `autoScaleTaskCount` in the `Service` construct.
	// For more information, please refer https://github.com/aws/aws-cdk/blob/master/packages/%40aws-cdk-containers/ecs-service-extensions/README.md#task-auto-scaling .
	MaxTaskCount() *float64
	// The minimum number of tasks when scaling in.
	// Deprecated: To enable target tracking based on CPU utilization, use the `targetCpuUtilization` property of `autoScaleTaskCount` in the `Service` construct.
	// For more information, please refer https://github.com/aws/aws-cdk/blob/master/packages/%40aws-cdk-containers/ecs-service-extensions/README.md#task-auto-scaling .
	MinTaskCount() *float64
	// The name of the extension.
	// Deprecated: To enable target tracking based on CPU utilization, use the `targetCpuUtilization` property of `autoScaleTaskCount` in the `Service` construct.
	// For more information, please refer https://github.com/aws/aws-cdk/blob/master/packages/%40aws-cdk-containers/ecs-service-extensions/README.md#task-auto-scaling .
	Name() *string
	// Deprecated: To enable target tracking based on CPU utilization, use the `targetCpuUtilization` property of `autoScaleTaskCount` in the `Service` construct.
	// For more information, please refer https://github.com/aws/aws-cdk/blob/master/packages/%40aws-cdk-containers/ecs-service-extensions/README.md#task-auto-scaling .
	SetName(val *string)
	// The service which this extension is being added to.
	//
	// Initially, extensions are collected into a ServiceDescription, but no service
	// exists yet. Later, when the ServiceDescription is used to create a service,
	// the extension is told what Service it is now working on.
	// Deprecated: To enable target tracking based on CPU utilization, use the `targetCpuUtilization` property of `autoScaleTaskCount` in the `Service` construct.
	// For more information, please refer https://github.com/aws/aws-cdk/blob/master/packages/%40aws-cdk-containers/ecs-service-extensions/README.md#task-auto-scaling .
	ParentService() Service
	// Deprecated: To enable target tracking based on CPU utilization, use the `targetCpuUtilization` property of `autoScaleTaskCount` in the `Service` construct.
	// For more information, please refer https://github.com/aws/aws-cdk/blob/master/packages/%40aws-cdk-containers/ecs-service-extensions/README.md#task-auto-scaling .
	SetParentService(val Service)
	// How long to wait between scale in actions.
	// Deprecated: To enable target tracking based on CPU utilization, use the `targetCpuUtilization` property of `autoScaleTaskCount` in the `Service` construct.
	// For more information, please refer https://github.com/aws/aws-cdk/blob/master/packages/%40aws-cdk-containers/ecs-service-extensions/README.md#task-auto-scaling .
	ScaleInCooldown() awscdk.Duration
	// How long to wait between scale out actions.
	// Deprecated: To enable target tracking based on CPU utilization, use the `targetCpuUtilization` property of `autoScaleTaskCount` in the `Service` construct.
	// For more information, please refer https://github.com/aws/aws-cdk/blob/master/packages/%40aws-cdk-containers/ecs-service-extensions/README.md#task-auto-scaling .
	ScaleOutCooldown() awscdk.Duration
	// Deprecated: To enable target tracking based on CPU utilization, use the `targetCpuUtilization` property of `autoScaleTaskCount` in the `Service` construct.
	// For more information, please refer https://github.com/aws/aws-cdk/blob/master/packages/%40aws-cdk-containers/ecs-service-extensions/README.md#task-auto-scaling .
	Scope() constructs.Construct
	// Deprecated: To enable target tracking based on CPU utilization, use the `targetCpuUtilization` property of `autoScaleTaskCount` in the `Service` construct.
	// For more information, please refer https://github.com/aws/aws-cdk/blob/master/packages/%40aws-cdk-containers/ecs-service-extensions/README.md#task-auto-scaling .
	SetScope(val constructs.Construct)
	// The CPU utilization to try ot maintain.
	// Deprecated: To enable target tracking based on CPU utilization, use the `targetCpuUtilization` property of `autoScaleTaskCount` in the `Service` construct.
	// For more information, please refer https://github.com/aws/aws-cdk/blob/master/packages/%40aws-cdk-containers/ecs-service-extensions/README.md#task-auto-scaling .
	TargetCpuUtilization() *float64
	// This hook allows another service extension to register a mutating hook for changing the primary container of this extension.
	//
	// This is primarily used
	// for the application extension. For example, the Firelens extension wants to
	// be able to modify the settings of the application container to
	// route logs through Firelens.
	// Deprecated: To enable target tracking based on CPU utilization, use the `targetCpuUtilization` property of `autoScaleTaskCount` in the `Service` construct.
	// For more information, please refer https://github.com/aws/aws-cdk/blob/master/packages/%40aws-cdk-containers/ecs-service-extensions/README.md#task-auto-scaling .
	AddContainerMutatingHook(hook ContainerMutatingHook)
	// A hook that allows the extension to add hooks to other extensions that are registered.
	// Deprecated: To enable target tracking based on CPU utilization, use the `targetCpuUtilization` property of `autoScaleTaskCount` in the `Service` construct.
	// For more information, please refer https://github.com/aws/aws-cdk/blob/master/packages/%40aws-cdk-containers/ecs-service-extensions/README.md#task-auto-scaling .
	AddHooks()
	// This hook allows the extension to establish a connection to extensions from another service.
	//
	// Usually used for things like
	// allowing one service to talk to the load balancer or service mesh
	// proxy for another service.
	// Deprecated: To enable target tracking based on CPU utilization, use the `targetCpuUtilization` property of `autoScaleTaskCount` in the `Service` construct.
	// For more information, please refer https://github.com/aws/aws-cdk/blob/master/packages/%40aws-cdk-containers/ecs-service-extensions/README.md#task-auto-scaling .
	ConnectToService(service Service, connectToProps *ConnectToProps)
	// Prior to launching the task definition as a service, this hook is called on each extension to give it a chance to mutate the properties of the service to be created.
	// Deprecated: To enable target tracking based on CPU utilization, use the `targetCpuUtilization` property of `autoScaleTaskCount` in the `Service` construct.
	// For more information, please refer https://github.com/aws/aws-cdk/blob/master/packages/%40aws-cdk-containers/ecs-service-extensions/README.md#task-auto-scaling .
	ModifyServiceProps(props *ServiceBuild) *ServiceBuild
	// This is a hook which allows extensions to modify the settings of the task definition prior to it being created.
	//
	// For example, the App Mesh
	// extension needs to configure an Envoy proxy in the task definition,
	// or the Application extension wants to set the overall resource for
	// the task.
	// Deprecated: To enable target tracking based on CPU utilization, use the `targetCpuUtilization` property of `autoScaleTaskCount` in the `Service` construct.
	// For more information, please refer https://github.com/aws/aws-cdk/blob/master/packages/%40aws-cdk-containers/ecs-service-extensions/README.md#task-auto-scaling .
	ModifyTaskDefinitionProps(props *awsecs.TaskDefinitionProps) *awsecs.TaskDefinitionProps
	// A hook that is called for each extension ahead of time to allow for any initial setup, such as creating resources in advance.
	// Deprecated: To enable target tracking based on CPU utilization, use the `targetCpuUtilization` property of `autoScaleTaskCount` in the `Service` construct.
	// For more information, please refer https://github.com/aws/aws-cdk/blob/master/packages/%40aws-cdk-containers/ecs-service-extensions/README.md#task-auto-scaling .
	Prehook(parent Service, scope constructs.Construct)
	// Once all containers are added to the task definition, this hook is called for each extension to give it a chance to resolve its dependency graph so that its container starts in the right order based on the other extensions that were enabled.
	// Deprecated: To enable target tracking based on CPU utilization, use the `targetCpuUtilization` property of `autoScaleTaskCount` in the `Service` construct.
	// For more information, please refer https://github.com/aws/aws-cdk/blob/master/packages/%40aws-cdk-containers/ecs-service-extensions/README.md#task-auto-scaling .
	ResolveContainerDependencies()
	// When this hook is implemented by extension, it allows the extension to use the service which has been created.
	//
	// It is generally used to
	// create any final resources which might depend on the service itself.
	// Deprecated: To enable target tracking based on CPU utilization, use the `targetCpuUtilization` property of `autoScaleTaskCount` in the `Service` construct.
	// For more information, please refer https://github.com/aws/aws-cdk/blob/master/packages/%40aws-cdk-containers/ecs-service-extensions/README.md#task-auto-scaling .
	UseService(service interface{})
	// Once the task definition is created, this hook is called for each extension to give it a chance to add containers to the task definition, change the task definition's role to add permissions, etc.
	// Deprecated: To enable target tracking based on CPU utilization, use the `targetCpuUtilization` property of `autoScaleTaskCount` in the `Service` construct.
	// For more information, please refer https://github.com/aws/aws-cdk/blob/master/packages/%40aws-cdk-containers/ecs-service-extensions/README.md#task-auto-scaling .
	UseTaskDefinition(taskDefinition awsecs.TaskDefinition)
}

// The jsii proxy struct for ScaleOnCpuUtilization
type jsiiProxy_ScaleOnCpuUtilization struct {
	jsiiProxy_ServiceExtension
}

func (j *jsiiProxy_ScaleOnCpuUtilization) Container() awsecs.ContainerDefinition {
	var returns awsecs.ContainerDefinition
	_jsii_.Get(
		j,
		"container",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScaleOnCpuUtilization) ContainerMutatingHooks() *[]ContainerMutatingHook {
	var returns *[]ContainerMutatingHook
	_jsii_.Get(
		j,
		"containerMutatingHooks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScaleOnCpuUtilization) InitialTaskCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialTaskCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScaleOnCpuUtilization) MaxTaskCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTaskCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScaleOnCpuUtilization) MinTaskCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minTaskCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScaleOnCpuUtilization) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScaleOnCpuUtilization) ParentService() Service {
	var returns Service
	_jsii_.Get(
		j,
		"parentService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScaleOnCpuUtilization) ScaleInCooldown() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"scaleInCooldown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScaleOnCpuUtilization) ScaleOutCooldown() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"scaleOutCooldown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScaleOnCpuUtilization) Scope() constructs.Construct {
	var returns constructs.Construct
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScaleOnCpuUtilization) TargetCpuUtilization() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetCpuUtilization",
		&returns,
	)
	return returns
}


// Deprecated: To enable target tracking based on CPU utilization, use the `targetCpuUtilization` property of `autoScaleTaskCount` in the `Service` construct.
// For more information, please refer https://github.com/aws/aws-cdk/blob/master/packages/%40aws-cdk-containers/ecs-service-extensions/README.md#task-auto-scaling .
func NewScaleOnCpuUtilization(props *CpuScalingProps) ScaleOnCpuUtilization {
	_init_.Initialize()

	if err := validateNewScaleOnCpuUtilizationParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ScaleOnCpuUtilization{}

	_jsii_.Create(
		"@aws-cdk-containers/ecs-service-extensions.ScaleOnCpuUtilization",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Deprecated: To enable target tracking based on CPU utilization, use the `targetCpuUtilization` property of `autoScaleTaskCount` in the `Service` construct.
// For more information, please refer https://github.com/aws/aws-cdk/blob/master/packages/%40aws-cdk-containers/ecs-service-extensions/README.md#task-auto-scaling .
func NewScaleOnCpuUtilization_Override(s ScaleOnCpuUtilization, props *CpuScalingProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk-containers/ecs-service-extensions.ScaleOnCpuUtilization",
		[]interface{}{props},
		s,
	)
}

func (j *jsiiProxy_ScaleOnCpuUtilization)SetContainer(val awsecs.ContainerDefinition) {
	_jsii_.Set(
		j,
		"container",
		val,
	)
}

func (j *jsiiProxy_ScaleOnCpuUtilization)SetContainerMutatingHooks(val *[]ContainerMutatingHook) {
	if err := j.validateSetContainerMutatingHooksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerMutatingHooks",
		val,
	)
}

func (j *jsiiProxy_ScaleOnCpuUtilization)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ScaleOnCpuUtilization)SetParentService(val Service) {
	if err := j.validateSetParentServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parentService",
		val,
	)
}

func (j *jsiiProxy_ScaleOnCpuUtilization)SetScope(val constructs.Construct) {
	if err := j.validateSetScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (s *jsiiProxy_ScaleOnCpuUtilization) AddContainerMutatingHook(hook ContainerMutatingHook) {
	if err := s.validateAddContainerMutatingHookParameters(hook); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addContainerMutatingHook",
		[]interface{}{hook},
	)
}

func (s *jsiiProxy_ScaleOnCpuUtilization) AddHooks() {
	_jsii_.InvokeVoid(
		s,
		"addHooks",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ScaleOnCpuUtilization) ConnectToService(service Service, connectToProps *ConnectToProps) {
	if err := s.validateConnectToServiceParameters(service, connectToProps); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"connectToService",
		[]interface{}{service, connectToProps},
	)
}

func (s *jsiiProxy_ScaleOnCpuUtilization) ModifyServiceProps(props *ServiceBuild) *ServiceBuild {
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

func (s *jsiiProxy_ScaleOnCpuUtilization) ModifyTaskDefinitionProps(props *awsecs.TaskDefinitionProps) *awsecs.TaskDefinitionProps {
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

func (s *jsiiProxy_ScaleOnCpuUtilization) Prehook(parent Service, scope constructs.Construct) {
	if err := s.validatePrehookParameters(parent, scope); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"prehook",
		[]interface{}{parent, scope},
	)
}

func (s *jsiiProxy_ScaleOnCpuUtilization) ResolveContainerDependencies() {
	_jsii_.InvokeVoid(
		s,
		"resolveContainerDependencies",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ScaleOnCpuUtilization) UseService(service interface{}) {
	if err := s.validateUseServiceParameters(service); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"useService",
		[]interface{}{service},
	)
}

func (s *jsiiProxy_ScaleOnCpuUtilization) UseTaskDefinition(taskDefinition awsecs.TaskDefinition) {
	if err := s.validateUseTaskDefinitionParameters(taskDefinition); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"useTaskDefinition",
		[]interface{}{taskDefinition},
	)
}

