package awscdkcontainersecsserviceextensions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-ecs-service-extensions-go/awscdkcontainersecsserviceextensions/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/cdk-ecs-service-extensions-go/awscdkcontainersecsserviceextensions/v2/internal"
)

// Experimental.
type ImportedEnvironment interface {
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
	// The VPC into which environment services should be placed.
	// Experimental.
	Vpc() awsec2.IVpc
	// Adding a default cloudmap namespace to the cluster will throw an error, as we don't own it.
	// Experimental.
	AddDefaultCloudMapNamespace(_options *awsecs.CloudMapNamespaceOptions)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ImportedEnvironment
type jsiiProxy_ImportedEnvironment struct {
	internal.Type__constructsConstruct
	jsiiProxy_IEnvironment
}

func (j *jsiiProxy_ImportedEnvironment) CapacityType() EnvironmentCapacityType {
	var returns EnvironmentCapacityType
	_jsii_.Get(
		j,
		"capacityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImportedEnvironment) Cluster() awsecs.ICluster {
	var returns awsecs.ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImportedEnvironment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImportedEnvironment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImportedEnvironment) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


// Experimental.
func NewImportedEnvironment(scope constructs.Construct, id *string, props *EnvironmentAttributes) ImportedEnvironment {
	_init_.Initialize()

	if err := validateNewImportedEnvironmentParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ImportedEnvironment{}

	_jsii_.Create(
		"@aws-cdk-containers/ecs-service-extensions.ImportedEnvironment",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewImportedEnvironment_Override(i ImportedEnvironment, scope constructs.Construct, id *string, props *EnvironmentAttributes) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk-containers/ecs-service-extensions.ImportedEnvironment",
		[]interface{}{scope, id, props},
		i,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func ImportedEnvironment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateImportedEnvironment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk-containers/ecs-service-extensions.ImportedEnvironment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImportedEnvironment) AddDefaultCloudMapNamespace(_options *awsecs.CloudMapNamespaceOptions) {
	if err := i.validateAddDefaultCloudMapNamespaceParameters(_options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addDefaultCloudMapNamespace",
		[]interface{}{_options},
	)
}

func (i *jsiiProxy_ImportedEnvironment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

