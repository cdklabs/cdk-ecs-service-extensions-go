// The CDK Construct Library that helps you build ECS services using simple extensions
package awscdkcontainersecsserviceextensions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-ecs-service-extensions-go/awscdkcontainersecsserviceextensions/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
)

// This extension creates a default `eventsQueue` for the service (if not provided) and accepts a list of objects of type `ISubscribable` that the `eventsQueue` subscribes to.
//
// It creates the subscriptions and sets up permissions
// for the service to consume messages from the SQS Queues.
//
// It also configures a target tracking scaling policy for the service to maintain an acceptable queue latency by tracking
// the backlog per task. For more information, please refer: https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-using-sqs-queue.html .
//
// The default queue for this service can be accessed using the getter `<extension>.eventsQueue`.
// Experimental.
type QueueExtension interface {
	ServiceExtension
	// Experimental.
	AutoscalingOptions() *QueueAutoScalingOptions
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
	// Experimental.
	EventsQueue() awssqs.IQueue
	// The log group created by the extension where the AWS Lambda function logs are stored.
	// Experimental.
	LogGroup() awslogs.ILogGroup
	// Experimental.
	SetLogGroup(val awslogs.ILogGroup)
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
	// Add hooks to the main application extension so that it is modified to add the events queue URL to the container environment.
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
	// This hook creates (if required) and sets the default queue `eventsQueue`.
	//
	// It also sets up the subscriptions for
	// the provided `ISubscribable` objects.
	// Experimental.
	Prehook(service Service, scope constructs.Construct)
	// Once all containers are added to the task definition, this hook is called for each extension to give it a chance to resolve its dependency graph so that its container starts in the right order based on the other extensions that were enabled.
	// Experimental.
	ResolveContainerDependencies()
	// When this hook is implemented by extension, it allows the extension to use the service which has been created.
	//
	// It is used to add target tracking
	// scaling policies for the SQS Queues of the service. It also creates an AWS Lambda
	// Function for calculating the backlog per task metric.
	// Experimental.
	UseService(service interface{})
	// After the task definition has been created, this hook grants SQS permissions to the task role.
	// Experimental.
	UseTaskDefinition(taskDefinition awsecs.TaskDefinition)
}

// The jsii proxy struct for QueueExtension
type jsiiProxy_QueueExtension struct {
	jsiiProxy_ServiceExtension
}

func (j *jsiiProxy_QueueExtension) AutoscalingOptions() *QueueAutoScalingOptions {
	var returns *QueueAutoScalingOptions
	_jsii_.Get(
		j,
		"autoscalingOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueExtension) Container() awsecs.ContainerDefinition {
	var returns awsecs.ContainerDefinition
	_jsii_.Get(
		j,
		"container",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueExtension) ContainerMutatingHooks() *[]ContainerMutatingHook {
	var returns *[]ContainerMutatingHook
	_jsii_.Get(
		j,
		"containerMutatingHooks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueExtension) EventsQueue() awssqs.IQueue {
	var returns awssqs.IQueue
	_jsii_.Get(
		j,
		"eventsQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueExtension) LogGroup() awslogs.ILogGroup {
	var returns awslogs.ILogGroup
	_jsii_.Get(
		j,
		"logGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueExtension) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueExtension) ParentService() Service {
	var returns Service
	_jsii_.Get(
		j,
		"parentService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueExtension) Scope() constructs.Construct {
	var returns constructs.Construct
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}


// Experimental.
func NewQueueExtension(props *QueueExtensionProps) QueueExtension {
	_init_.Initialize()

	if err := validateNewQueueExtensionParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_QueueExtension{}

	_jsii_.Create(
		"@aws-cdk-containers/ecs-service-extensions.QueueExtension",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewQueueExtension_Override(q QueueExtension, props *QueueExtensionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk-containers/ecs-service-extensions.QueueExtension",
		[]interface{}{props},
		q,
	)
}

func (j *jsiiProxy_QueueExtension)SetContainer(val awsecs.ContainerDefinition) {
	_jsii_.Set(
		j,
		"container",
		val,
	)
}

func (j *jsiiProxy_QueueExtension)SetContainerMutatingHooks(val *[]ContainerMutatingHook) {
	if err := j.validateSetContainerMutatingHooksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerMutatingHooks",
		val,
	)
}

func (j *jsiiProxy_QueueExtension)SetLogGroup(val awslogs.ILogGroup) {
	_jsii_.Set(
		j,
		"logGroup",
		val,
	)
}

func (j *jsiiProxy_QueueExtension)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_QueueExtension)SetParentService(val Service) {
	if err := j.validateSetParentServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parentService",
		val,
	)
}

func (j *jsiiProxy_QueueExtension)SetScope(val constructs.Construct) {
	if err := j.validateSetScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (q *jsiiProxy_QueueExtension) AddContainerMutatingHook(hook ContainerMutatingHook) {
	if err := q.validateAddContainerMutatingHookParameters(hook); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"addContainerMutatingHook",
		[]interface{}{hook},
	)
}

func (q *jsiiProxy_QueueExtension) AddHooks() {
	_jsii_.InvokeVoid(
		q,
		"addHooks",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QueueExtension) ConnectToService(service Service, connectToProps *ConnectToProps) {
	if err := q.validateConnectToServiceParameters(service, connectToProps); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"connectToService",
		[]interface{}{service, connectToProps},
	)
}

func (q *jsiiProxy_QueueExtension) ModifyServiceProps(props *ServiceBuild) *ServiceBuild {
	if err := q.validateModifyServicePropsParameters(props); err != nil {
		panic(err)
	}
	var returns *ServiceBuild

	_jsii_.Invoke(
		q,
		"modifyServiceProps",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueueExtension) ModifyTaskDefinitionProps(props *awsecs.TaskDefinitionProps) *awsecs.TaskDefinitionProps {
	if err := q.validateModifyTaskDefinitionPropsParameters(props); err != nil {
		panic(err)
	}
	var returns *awsecs.TaskDefinitionProps

	_jsii_.Invoke(
		q,
		"modifyTaskDefinitionProps",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueueExtension) Prehook(service Service, scope constructs.Construct) {
	if err := q.validatePrehookParameters(service, scope); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"prehook",
		[]interface{}{service, scope},
	)
}

func (q *jsiiProxy_QueueExtension) ResolveContainerDependencies() {
	_jsii_.InvokeVoid(
		q,
		"resolveContainerDependencies",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QueueExtension) UseService(service interface{}) {
	if err := q.validateUseServiceParameters(service); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"useService",
		[]interface{}{service},
	)
}

func (q *jsiiProxy_QueueExtension) UseTaskDefinition(taskDefinition awsecs.TaskDefinition) {
	if err := q.validateUseTaskDefinitionParameters(taskDefinition); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"useTaskDefinition",
		[]interface{}{taskDefinition},
	)
}

