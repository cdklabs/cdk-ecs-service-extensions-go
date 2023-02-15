// The CDK Construct Library that helps you build ECS services using simple extensions
package awscdkcontainersecsserviceextensions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
)

// An environment into which to deploy a service.
// Experimental.
type IEnvironment interface {
	// Add a default cloudmap namespace to the environment's cluster.
	// Experimental.
	AddDefaultCloudMapNamespace(options *awsecs.CloudMapNamespaceOptions)
	// The capacity type used by the service's cluster.
	// Experimental.
	CapacityType() EnvironmentCapacityType
	// The cluster that is providing capacity for this service.
	// Experimental.
	Cluster() awsecs.ICluster
	// The name of this environment.
	// Experimental.
	Id() *string
	// The VPC into which environment services should be placed.
	// Experimental.
	Vpc() awsec2.IVpc
}

// The jsii proxy for IEnvironment
type jsiiProxy_IEnvironment struct {
	_ byte // padding
}

func (i *jsiiProxy_IEnvironment) AddDefaultCloudMapNamespace(options *awsecs.CloudMapNamespaceOptions) {
	if err := i.validateAddDefaultCloudMapNamespaceParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addDefaultCloudMapNamespace",
		[]interface{}{options},
	)
}

func (j *jsiiProxy_IEnvironment) CapacityType() EnvironmentCapacityType {
	var returns EnvironmentCapacityType
	_jsii_.Get(
		j,
		"capacityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironment) Cluster() awsecs.ICluster {
	var returns awsecs.ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironment) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}

