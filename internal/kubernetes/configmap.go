package kubernetes

import (
	"gopkg.in/yaml.v2"
)

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

// SafeString returns the object as a string, returning an error on failure
func (cfg Configmap) SafeString() (string, error) {
	return serialize(cfg)
}

// Secret represents a confidential configuration for a deployment
type Secret struct {
	Metadata   Metadata
	Kind       string
	Type       string
	ApiVersion string
	Data       map[string]string
}

// NewSecret returns a default empty secret
func NewSecret() Secret {
	return Secret{
		Kind:       "Secret",
		Type: 			"Opaque",
		ApiVersion: ConfigsVersion,
		Data:       make(map[string]string),
		Metadata:   NewMetadata(),
	}
}

// SafeString returns the object as a string, returning an error on failure
func (secret Secret) SafeString() (string, error) {
	return serialize(secret)
}

func serialize(cfg interface{}) (string, error) {
	bytes, err := yaml.Marshal(cfg)
	if err != nil {
		return "", err
	}

	return string(bytes), err
}