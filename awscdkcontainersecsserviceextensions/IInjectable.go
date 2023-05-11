package awscdkcontainersecsserviceextensions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// An interface that will be implemented by all the resources that can be published events or written data to.
// Experimental.
type IInjectable interface {
	// Experimental.
	EnvironmentVariables() *map[string]*string
}

// The jsii proxy for IInjectable
type jsiiProxy_IInjectable struct {
	_ byte // padding
}

func (i *jsiiProxy_IInjectable) EnvironmentVariables() *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"environmentVariables",
		nil, // no parameters
		&returns,
	)

	return returns
}

