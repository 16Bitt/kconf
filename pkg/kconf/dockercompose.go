package kconf

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"github.com/16bitt/kconf/internal/dockercompose"
)

// GenerateDockerCompose and write the file to path
func (kc *KConfig) GenerateDockerCompose(path string) error {
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

		vars, err := kc.AllVars(app)
		if err != nil {
			return err
		}

		for _, name := range vars {
			svc.Environment[name] = ""
		}

		root.Services[app.Name] = svc
	}

	bytes, err := yaml.Marshal(root)
	if err != nil {
		return err
	}
	fmt.Println(string(bytes))

	err = ioutil.WriteFile(path, bytes, 0)
	if err != nil {
		return err
	}

	return nil
}
