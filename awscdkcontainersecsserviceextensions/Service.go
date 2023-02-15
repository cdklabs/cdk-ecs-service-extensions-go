// The CDK Construct Library that helps you build ECS services using simple extensions
package awscdkcontainersecsserviceextensions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-ecs-service-extensions-go/awscdkcontainersecsserviceextensions/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/cdk-ecs-service-extensions-go/awscdkcontainersecsserviceextensions/v2/internal"
)

// This Service construct serves as a Builder class for an ECS service.
//
// It
// supports various extensions and keeps track of any mutating state, allowing
// it to build up an ECS service progressively.
// Experimental.
type Service interface {
	constructs.Construct
	// The capacity type that this service will use.
	//
	// Valid values are EC2 or FARGATE.
	// Experimental.
	CapacityType() EnvironmentCapacityType
	// The cluster that is providing capacity for this service.
	//
	// [disable-awslint:ref-via-interface].
	// Experimental.
	Cluster() awsecs.ICluster
	// The underlying ECS service that was created.
	// Experimental.
	EcsService() interface{}
	// Experimental.
	SetEcsService(val interface{})
	// The environment where this service was launched.
	// Experimental.
	Environment() IEnvironment
	// The name of the service.
	// Experimental.
	Id() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// The scalable attribute representing task count.
	// Experimental.
	ScalableTaskCount() awsecs.ScalableTaskCount
	// The ServiceDescription used to build this service.
	// Experimental.
	ServiceDescription() ServiceDescription
	// The application target group if the service has an HTTPLoadBalancerExtension.
	// Experimental.
	TargetGroup() awselasticloadbalancingv2.ApplicationTargetGroup
	// Experimental.
	SetTargetGroup(val awselasticloadbalancingv2.ApplicationTargetGroup)
	// The generated task definition for this service.
	//
	// It is only
	// generated after .prepare() has been executed.
	// Experimental.
	TaskDefinition() awsecs.TaskDefinition
	// Experimental.
	SetTaskDefinition(val awsecs.TaskDefinition)
	// The VPC where this service should be placed.
	// Experimental.
	Vpc() awsec2.IVpc
	// This method adds a new URL for the service.
	//
	// This allows extensions to
	// submit a URL for the service. For example, a load balancer might add its
	// URL, or App Mesh can add its DNS name for the service.
	// Experimental.
	AddURL(urlName *string, url *string)
	// Tell extensions from one service to connect to extensions from another sevice if they have implemented a hook for it.
	// Experimental.
	ConnectTo(service Service, connectToProps *ConnectToProps)
	// This helper method is used to set the `autoScalingPoliciesEnabled` attribute whenever an auto scaling policy is configured for the service.
	// Experimental.
	EnableAutoScalingPolicy()
	// Retrieve a URL for the service.
	//
	// The URL must have previously been
	// stored by one of the URL providing extensions.
	// Experimental.
	GetURL(urlName *string) *string
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Service
type jsiiProxy_Service struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_Service) CapacityType() EnvironmentCapacityType {
	var returns EnvironmentCapacityType
	_jsii_.Get(
		j,
		"capacityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) Cluster() awsecs.ICluster {
	var returns awsecs.ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) EcsService() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecsService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) Environment() IEnvironment {
	var returns IEnvironment
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) ScalableTaskCount() awsecs.ScalableTaskCount {
	var returns awsecs.ScalableTaskCount
	_jsii_.Get(
		j,
		"scalableTaskCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) ServiceDescription() ServiceDescription {
	var returns ServiceDescription
	_jsii_.Get(
		j,
		"serviceDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) TargetGroup() awselasticloadbalancingv2.ApplicationTargetGroup {
	var returns awselasticloadbalancingv2.ApplicationTargetGroup
	_jsii_.Get(
		j,
		"targetGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) TaskDefinition() awsecs.TaskDefinition {
	var returns awsecs.TaskDefinition
	_jsii_.Get(
		j,
		"taskDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


// Experimental.
func NewService(scope constructs.Construct, id *string, props *ServiceProps) Service {
	_init_.Initialize()

	if err := validateNewServiceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Service{}

	_jsii_.Create(
		"@aws-cdk-containers/ecs-service-extensions.Service",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewService_Override(s Service, scope constructs.Construct, id *string, props *ServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk-containers/ecs-service-extensions.Service",
		[]interface{}{scope, id, props},
		s,
	)
}

func (j *jsiiProxy_Service)SetEcsService(val interface{}) {
	if err := j.validateSetEcsServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ecsService",
		val,
	)
}

func (j *jsiiProxy_Service)SetTargetGroup(val awselasticloadbalancingv2.ApplicationTargetGroup) {
	_jsii_.Set(
		j,
		"targetGroup",
		val,
	)
}

func (j *jsiiProxy_Service)SetTaskDefinition(val awsecs.TaskDefinition) {
	if err := j.validateSetTaskDefinitionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskDefinition",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func Service_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateService_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk-containers/ecs-service-extensions.Service",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Service) AddURL(urlName *string, url *string) {
	if err := s.validateAddURLParameters(urlName, url); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addURL",
		[]interface{}{urlName, url},
	)
}

func (s *jsiiProxy_Service) ConnectTo(service Service, connectToProps *ConnectToProps) {
	if err := s.validateConnectToParameters(service, connectToProps); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"connectTo",
		[]interface{}{service, connectToProps},
	)
}

func (s *jsiiProxy_Service) EnableAutoScalingPolicy() {
	_jsii_.InvokeVoid(
		s,
		"enableAutoScalingPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Service) GetURL(urlName *string) *string {
	if err := s.validateGetURLParameters(urlName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getURL",
		[]interface{}{urlName},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Service) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

