package kubernetes

// Deployment represents how an app is configured and run
type Deployment struct {
	ApiVersion string
	Metadata   Metadata
}
