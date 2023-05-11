//go:build !no_runtime_type_checking

package awscdkcontainersecsserviceextensions

import (
	"fmt"
)

func (s *jsiiProxy_ServiceDescription) validateAddParameters(extension ServiceExtension) error {
	if extension == nil {
		return fmt.Errorf("parameter extension is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_ServiceDescription) validateGetParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ServiceDescription) validateSetExtensionsParameters(val *map[string]ServiceExtension) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

