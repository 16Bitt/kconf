package kubernetes

// DeploymentAPIVersion is the fixed version for our deployments
const DeploymentAPIVersion = "api/v1"

// Deployment represents how an app is configured and run
type Deployment struct {
	ApiVersion string
	Metadata   Metadata
	Spec       DeploymentSpec
}

// DeploymentSpec is the main contents of a deployment
type DeploymentSpec struct {
	Replicas string
	Selector MatchSelector
	Template DeploymentTemplate
}

// MatchSelector provides a map a labels for querying the deployment
type MatchSelector struct {
	MatchLabels map[string]string
}

// DeploymentTemplate wraps the functionality of a deployment
type DeploymentTemplate struct {
	Metadata Metadata
	Spec     ContainerSpec
}

// ContainerSpec defines the details of a container within a deployment
type ContainerSpec struct {
	Name      string
	Image     string
	Ports     []PortSpec          `yaml:",omitempty"`
	Command   []string            `yaml:",omitempty"`
	Args      []string            `yaml:",omitempty"`
	Env       []EnvironmentConfig `yaml:",omitempty"`
	Resources ResourceConstraints
}

// PortSpec defines the open ports to communicate with a container in a pod
type PortSpec struct {
	ContainerPort string
}

// EnvironmentConfig defines a configuration and how it is mapped to the container
type EnvironmentConfig struct {
	Name      string
	Value     string              `yaml:",omitempty"`
	ValueFrom EnvironmentValueMap `yaml:",omitempty"`
}

// EnvironmentValueMap defines how the environment value is mapped to your container
type EnvironmentValueMap struct {
	ConfigMapKeyRef KeyRef `yaml:",omitempty"`
	SecretKeyRef    KeyRef `yaml:",omitempty"`
}

// KeyRef maps a secret or configmap value to an environment config
type KeyRef struct {
	Name string
	Key  string
}

// ResourceConstraints defines the limits for a pod running in the cluster
type ResourceConstraints struct {
	Limits   ResourceConstraint `yaml:",omitempty"`
	Requests ResourceConstraint `yaml:",omitempty"`
}

// ResourceConstraint is the basic way of configuring resources
type ResourceConstraint struct {
	Cpu    string
	Memory string
}
