// The CDK Construct Library that helps you build ECS services using simple extensions
package awscdkcontainersecsserviceextensions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-ecs-service-extensions-go/awscdkcontainersecsserviceextensions/v2/jsii"
)

// A ServiceDescription is a wrapper for all of the extensions that a user wants to add to an ECS Service.
//
// It collects all of the extensions that are added
// to a service, allowing each extension to query the full list of extensions
// added to a service to determine information about how to self-configure.
// Experimental.
type ServiceDescription interface {
	// The list of extensions that have been registered to run when preparing this service.
	// Experimental.
	Extensions() *map[string]ServiceExtension
	// Experimental.
	SetExtensions(val *map[string]ServiceExtension)
	// Adds a new extension to the service.
	//
	// The extensions mutate a service
	// to add resources to or configure properties for the service.
	// Experimental.
	Add(extension ServiceExtension) ServiceDescription
	// Get the extension with a specific name.
	//
	// This is generally used by
	// extensions in order to discover each other.
	// Experimental.
	Get(name *string) ServiceExtension
}

// The jsii proxy struct for ServiceDescription
type jsiiProxy_ServiceDescription struct {
	_ byte // padding
}

func (j *jsiiProxy_ServiceDescription) Extensions() *map[string]ServiceExtension {
	var returns *map[string]ServiceExtension
	_jsii_.Get(
		j,
		"extensions",
		&returns,
	)
	return returns
}


// Experimental.
func NewServiceDescription() ServiceDescription {
	_init_.Initialize()

	j := jsiiProxy_ServiceDescription{}

	_jsii_.Create(
		"@aws-cdk-containers/ecs-service-extensions.ServiceDescription",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewServiceDescription_Override(s ServiceDescription) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk-containers/ecs-service-extensions.ServiceDescription",
		nil, // no parameters
		s,
	)
}

func (j *jsiiProxy_ServiceDescription)SetExtensions(val *map[string]ServiceExtension) {
	if err := j.validateSetExtensionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extensions",
		val,
	)
}

func (s *jsiiProxy_ServiceDescription) Add(extension ServiceExtension) ServiceDescription {
	if err := s.validateAddParameters(extension); err != nil {
		panic(err)
	}
	var returns ServiceDescription

	_jsii_.Invoke(
		s,
		"add",
		[]interface{}{extension},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceDescription) Get(name *string) ServiceExtension {
	if err := s.validateGetParameters(name); err != nil {
		panic(err)
	}
	var returns ServiceExtension

	_jsii_.Invoke(
		s,
		"get",
		[]interface{}{name},
		&returns,
	)

	return returns
}

