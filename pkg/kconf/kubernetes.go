package kconf

import (
	"fmt"

	"github.com/16bitt/kconf/internal/kubernetes"
	"gopkg.in/yaml.v2"
)

// GenerateConfigmaps parses the kubernetes config and writes a configmap to path
func (kc *KConfig) GenerateConfigmaps(path string) error {
	configs := []kubernetes.Configmap{}
	for key, section := range kc.Configs {
		config := kubernetes.NewConfigmap()
		config.Metadata.Name = key
		config.Metadata.Namespace = kc.Project.Kubernetes.Namespace
		config.Metadata.Labels["environment"] = "staging"
		for _, option := range section {
			if option.Secret {
				continue
			}

			config.Data[option.Name] = option.GetDefaultForEnvironment("staging", ConfTypeKubernetes)
		}
		configs = append(configs, config)
	}

	for _, cfg := range configs {
		bytes, err := yaml.Marshal(cfg)
		if err != nil {
			return err
		}

		fmt.Println(string(bytes))
	}
	return nil
}
