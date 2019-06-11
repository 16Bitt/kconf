package kconf

import (
	"gopkg.in/yaml.v2"

	"github.com/16bitt/kconf/internal/dockercompose"
)

// GenerateDockerCompose and write the file to path
func (kc *KConfig) GenerateDockerCompose(env string) (string, error) {
	root := dockercompose.New()
	for _, app := range kc.Apps {
		svc := dockercompose.Service{
			Ports:       app.Ports,
			Environment: make(map[string]string),
		}

		if app.Build != "" {
			svc.Build = app.Build
		} else {
			svc.Image = app.Image
		}

    for _, section := range app.Configs {
      configs := kc.Configs[section]
      for _, config := range configs {
        svc.Environment[config.Name] = config.GetDefaultForEnvironment(env, ConfTypeDockerCompose)
      }
    }

		root.Services[app.Name] = svc
	}

	bytes, err := yaml.Marshal(root)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
