package kconf

import "fmt"

const (
	ConfTypeKubernetes = "kubernetes"
  ConfTypeDockerCompose = "docker-compose"
)

// Config is a configuration section for an application
type Config struct {
	Name     string
	Secret   bool
	Defaults []ConfigDefault
}

// ConfigDefault allows setting default values for configurations by type
type ConfigDefault struct {
	Type        string
	Environment string
	Value       string
}

// Configs - List of configs
type Configs []Config

// Variables returns a list of variable names from a config section
func (cfgs Configs) Variables() []string {
	vars := []string{}
	for _, cfg := range cfgs {
		vars = append(vars, cfg.Name)
	}

	return vars
}

// GetConfig or return an error if it cannot be found
func (kc KConfig) GetConfig(name string) (Configs, error) {
	cfgs := kc.Configs[name]
	if len(cfgs) == 0 {
		return cfgs, fmt.Errorf("%s is not a valid configuration section", name)
	}

	return cfgs, nil
}

// GetDefaultForEnvironment tries to find the best default for a given config section
func (cfg Config) GetDefaultForEnvironment(environment, confType string) string {
	fallback := ""
	for _, def := range cfg.Defaults {
		if def.Environment == environment && confType == def.Type {
			return def.Value
		}

		if def.Type == confType && def.Environment == "default" {
			fallback = def.Value
		}
	}

	return fallback
}
