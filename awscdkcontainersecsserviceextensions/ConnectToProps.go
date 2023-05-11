package awscdkcontainersecsserviceextensions


// connectToProps will have all the extra parameters which are required for connecting services.
// Experimental.
type ConnectToProps struct {
	// localBindPort is the local port that this application should use when calling the upstream service in ECS Consul Mesh Extension Currently, this parameter will only be used in the ECSConsulMeshExtension https://github.com/aws-ia/ecs-consul-mesh-extension.
	// Experimental.
	LocalBindPort *float64 `field:"optional" json:"localBindPort" yaml:"localBindPort"`
}

