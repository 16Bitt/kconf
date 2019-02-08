package kubernetes

// ConfigsVersion is the default resource version for kubernetes
// configmaps and secrets
const ConfigsVersion = "api/v1"

// Configmap represents a general configuration
type Configmap struct {
	Metadata   Metadata
	Kind       string
	ApiVersion string
	Data       map[string]string
}

// NewConfigmap returns a default empty configmap
func NewConfigmap() Configmap {
	return Configmap{
		Kind:       "Configmap",
		ApiVersion: ConfigsVersion,
		Data:       make(map[string]string),
		Metadata:   NewMetadata(),
	}
}

// Secret represents a confidential configuration for a deployment
type Secret struct {
	Metadata   Metadata
	Kind       string
	Type       string
	ApiVersion string
	Data       map[string]string
}
