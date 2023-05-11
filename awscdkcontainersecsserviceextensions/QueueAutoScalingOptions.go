package awscdkcontainersecsserviceextensions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Options for configuring SQS Queue auto scaling.
// Experimental.
type QueueAutoScalingOptions struct {
	// Acceptable amount of time a message can sit in the queue (including the time required to process it).
	// Experimental.
	AcceptableLatency awscdk.Duration `field:"required" json:"acceptableLatency" yaml:"acceptableLatency"`
	// Average amount of time for processing a single message in the queue.
	// Experimental.
	MessageProcessingTime awscdk.Duration `field:"required" json:"messageProcessingTime" yaml:"messageProcessingTime"`
}

