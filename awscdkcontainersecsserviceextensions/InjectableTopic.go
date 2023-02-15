// The CDK Construct Library that helps you build ECS services using simple extensions
package awscdkcontainersecsserviceextensions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-ecs-service-extensions-go/awscdkcontainersecsserviceextensions/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
)

// The `InjectableTopic` class represents SNS Topic resource that can be published events to by the parent service.
// Experimental.
type InjectableTopic interface {
	IGrantInjectable
	// Experimental.
	Topic() awssns.ITopic
	// Experimental.
	EnvironmentVariables() *map[string]*string
	// Experimental.
	Grant(taskDefinition awsecs.TaskDefinition)
}

// The jsii proxy struct for InjectableTopic
type jsiiProxy_InjectableTopic struct {
	jsiiProxy_IGrantInjectable
}

func (j *jsiiProxy_InjectableTopic) Topic() awssns.ITopic {
	var returns awssns.ITopic
	_jsii_.Get(
		j,
		"topic",
		&returns,
	)
	return returns
}


// Experimental.
func NewInjectableTopic(props *InjectableTopicProps) InjectableTopic {
	_init_.Initialize()

	if err := validateNewInjectableTopicParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_InjectableTopic{}

	_jsii_.Create(
		"@aws-cdk-containers/ecs-service-extensions.InjectableTopic",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewInjectableTopic_Override(i InjectableTopic, props *InjectableTopicProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk-containers/ecs-service-extensions.InjectableTopic",
		[]interface{}{props},
		i,
	)
}

func (i *jsiiProxy_InjectableTopic) EnvironmentVariables() *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"environmentVariables",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InjectableTopic) Grant(taskDefinition awsecs.TaskDefinition) {
	if err := i.validateGrantParameters(taskDefinition); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"grant",
		[]interface{}{taskDefinition},
	)
}

