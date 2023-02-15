// The CDK Construct Library that helps you build ECS services using simple extensions
package awscdkcontainersecsserviceextensions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
)

// The settings for the `InjectableTopic` class.
// Experimental.
type InjectableTopicProps struct {
	// The SNS Topic to publish events to.
	// Experimental.
	Topic awssns.ITopic `field:"required" json:"topic" yaml:"topic"`
}

