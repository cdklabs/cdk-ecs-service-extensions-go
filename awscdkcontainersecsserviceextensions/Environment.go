package awscdkcontainersecsserviceextensions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-ecs-service-extensions-go/awscdkcontainersecsserviceextensions/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/cdk-ecs-service-extensions-go/awscdkcontainersecsserviceextensions/v2/internal"
)

// An environment into which to deploy a service.
//
// This environment
// can either be instantiated with a pre-existing AWS VPC and ECS cluster,
// or it can create its own VPC and cluster. By default, it will create
// a cluster with Fargate capacity.
// Experimental.
type Environment interface {
	constructs.Construct
	IEnvironment
	// The capacity type used by the service's cluster.
	// Experimental.
	CapacityType() EnvironmentCapacityType
	// The cluster that is providing capacity for this service.
	// Experimental.
	Cluster() awsecs.ICluster
	// The name of this environment.
	// Experimental.
	Id() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// The VPC where environment services should be placed.
	// Experimental.
	Vpc() awsec2.IVpc
	// Add a default cloudmap namespace to the environment's cluster.
	//
	// The environment's cluster must not be imported.
	// Experimental.
	AddDefaultCloudMapNamespace(options *awsecs.CloudMapNamespaceOptions)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Environment
type jsiiProxy_Environment struct {
	internal.Type__constructsConstruct
	jsiiProxy_IEnvironment
}

func (j *jsiiProxy_Environment) CapacityType() EnvironmentCapacityType {
	var returns EnvironmentCapacityType
	_jsii_.Get(
		j,
		"capacityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Environment) Cluster() awsecs.ICluster {
	var returns awsecs.ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Environment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Environment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Environment) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


// Experimental.
func NewEnvironment(scope constructs.Construct, id *string, props *EnvironmentProps) Environment {
	_init_.Initialize()

	if err := validateNewEnvironmentParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Environment{}

	_jsii_.Create(
		"@aws-cdk-containers/ecs-service-extensions.Environment",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewEnvironment_Override(e Environment, scope constructs.Construct, id *string, props *EnvironmentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk-containers/ecs-service-extensions.Environment",
		[]interface{}{scope, id, props},
		e,
	)
}

// Import an existing environment from its attributes.
// Experimental.
func Environment_FromEnvironmentAttributes(scope constructs.Construct, id *string, attrs *EnvironmentAttributes) IEnvironment {
	_init_.Initialize()

	if err := validateEnvironment_FromEnvironmentAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IEnvironment

	_jsii_.StaticInvoke(
		"@aws-cdk-containers/ecs-service-extensions.Environment",
		"fromEnvironmentAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func Environment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEnvironment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk-containers/ecs-service-extensions.Environment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Environment) AddDefaultCloudMapNamespace(options *awsecs.CloudMapNamespaceOptions) {
	if err := e.validateAddDefaultCloudMapNamespaceParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addDefaultCloudMapNamespace",
		[]interface{}{options},
	)
}

func (e *jsiiProxy_Environment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

