package awscdkcontainersecsserviceextensions


// The settings for the Injecter extension.
// Experimental.
type InjecterExtensionProps struct {
	// The list of injectable resources for this service.
	// Experimental.
	Injectables *[]IInjectable `field:"required" json:"injectables" yaml:"injectables"`
}

