//go:build !no_runtime_type_checking

package awscdkcontainersecsserviceextensions

import (
	"fmt"
)

func (i *jsiiProxy_ISubscribable) validateSubscribeParameters(extension QueueExtension) error {
	if extension == nil {
		return fmt.Errorf("parameter extension is required, but nil was provided")
	}

	return nil
}

