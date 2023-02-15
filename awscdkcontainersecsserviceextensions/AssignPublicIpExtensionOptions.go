// The CDK Construct Library that helps you build ECS services using simple extensions
package awscdkcontainersecsserviceextensions


// Experimental.
type AssignPublicIpExtensionOptions struct {
	// Enable publishing task public IPs to a recordset in a Route 53 hosted zone.
	//
	// Note: If you want to change the DNS zone or record name, you will need to
	// remove this extension completely and then re-add it.
	// Experimental.
	Dns *AssignPublicIpDnsOptions `field:"optional" json:"dns" yaml:"dns"`
}

