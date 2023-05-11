package awscdkcontainersecsserviceextensions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
)

// Experimental.
type AssignPublicIpDnsOptions struct {
	// Name of the record to add to the zone and in which to add the task IP addresses to.
	//
	// Example:
	//   'myservice'
	//
	// Experimental.
	RecordName *string `field:"required" json:"recordName" yaml:"recordName"`
	// A DNS Zone to expose task IPs in.
	// Experimental.
	Zone awsroute53.IHostedZone `field:"required" json:"zone" yaml:"zone"`
}

