package kconf

import (
	"strings"

	"github.com/16bitt/kconf/internal/kubernetes"
)

// GenerateConfigmaps parses the kubernetes config and returns the configmap as a yaml string
func (kc *KConfig) GenerateConfigmaps(environment string) (string, error) {
	configs := []kubernetes.Configmap{}
	for key, section := range kc.Configs {
		config := kubernetes.NewConfigmap()
		config.Metadata.Name = key
		config.Metadata.Namespace = kc.Project.Kubernetes.Namespace
		config.Metadata.Labels["environment"] = environment
		for _, option := range section {
			if option.Secret {
				continue
			}

			config.Data[option.Name] = option.GetDefaultForEnvironment(environment, ConfTypeKubernetes)
		}
		configs = append(configs, config)
	}

	var builder strings.Builder
	for _, config := range configs {
		yaml, err := config.SafeString()
		if err != nil {
			return "", err
		}
		builder.WriteString("---\n")
		builder.WriteString(yaml)
		builder.WriteString("\n")
	}

	return builder.String(), nil
}

// GenerateSecrets parses the kubernetes config and returns the secret as a yaml string
func (kc *KConfig) GenerateSecrets(environment string) (string, error) {
	secrets := []kubernetes.Secret{}
	for key, section := range kc.Configs {
		secret := kubernetes.NewSecret()
		secret.Metadata.Name = key
		secret.Metadata.Namespace = kc.Project.Kubernetes.Namespace
		secret.Metadata.Labels["environment"] = environment
		for _, option := range section {
			if !option.Secret {
				continue
			}

			secret.Data[option.Name] = option.GetDefaultForEnvironment(environment, ConfTypeKubernetes)
		}
		secrets = append(secrets, secret)
	}

	var builder strings.Builder
	for _, secret := range secrets {
		yaml, err := secret.SafeString()
		if err != nil {
			return "", err
		}
		builder.WriteString("---\n")
		builder.WriteString(yaml)
		builder.WriteString("\n")
	}

	return builder.String(), nil
}

func (kc *KConfig) GenerateDeployments(environment string) (string, error) {
	deps := []kubernetes.Deployment{}
	for _, app := range kc.Apps {
		dep := kubernetes.Deployment{
			ApiVersion: kubernetes.DeploymentAPIVersion,
			Metadata: kubernetes.Metadata{
				Name:      app.Name,
				Namespace: kc.Project.Kubernetes.Namespace,
				Labels:    make(map[string]string),
			},
			Spec: DeploymentSpec{},
		}

		dep.Metadata.Labels["environment"] = environment
	}
	return "", nil
}
