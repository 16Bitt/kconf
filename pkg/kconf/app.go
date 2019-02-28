package kconf

import "github.com/16bitt/kconf/internal/kubernetes"

const (
	KubernetesAppType = "kubernetes"
	HelmAppType       = "helm"
)

// App is a container within your application
type App struct {
	Name       string
	Build      string
	Configs    []string
	Image      string
	Type       string
	Helm       HelmConfig
	Kubernetes KubernetesParams
	Ports      []string
	Command    []string
	Args       []string
}

type KubernetesParams struct {
	Replicas string
	Limits   kubernetes.ResourceConstraint
	Requests kubernetes.ResourceConstraint
	Labels   map[string]string
	Ports    []string
}

// AllVars returns a list of all variable names for a given app
func (kc KConfig) AllVars(app App) ([]string, error) {
	vars := []string{}

	for _, name := range app.Configs {
		cfgs, err := kc.GetConfig(name)
		if err != nil {
			return vars, err
		}

		vars = append(vars, cfgs.Variables()...)
	}

	return vars, nil
}

func (app App) KubernetesDeployment(kc *KConfig) kubernetes.Deployment {
	return kubernetes.Deployment{
		ApiVersion: kubernetes.DeploymentAPIVersion,
		Metadata: kubernetes.Metadata{
			Name:      app.Name,
			Namespace: kc.Project.Kubernetes.Namespace,
			Labels:    app.Kubernetes.Labels,
		},
		Spec: kubernetes.DeploymentSpec{
			Replicas: app.Kubernetes.Replicas,
			Selector: kubernetes.MatchSelector{
				MatchLabels: app.Kubernetes.Labels,
			},
			Template: kubernetes.DeploymentTemplate{
				Metadata: kubernetes.Metadata{
					Name:      app.Name,
					Namespace: kc.Project.Kubernetes.Namespace,
					Labels:    app.Kubernetes.Labels,
				},
				Spec: []kubernetes.ContainerSpec{
					kubernetes.ContainerSpec{
						Name:    app.Name,
						Image:   app.Image,
						Ports:   app.kubernetesPortSpec(),
						Command: app.Command,
						Args:    app.Args,
						Resources: kubernetes.ResourceConstraints{
							Limits:   app.Kubernetes.Limits,
							Requests: app.Kubernetes.Requests,
						},
						Env: app.kubernetesEnv(kc),
					},
				},
			},
		},
	}
}

func (app App) kubernetesPortSpec() []kubernetes.PortSpec {
	ports := []kubernetes.PortSpec{}
	for _, port := range app.Kubernetes.Ports {
		ports = append(ports, kubernetes.PortSpec{ContainerPort: port})
	}

	return ports
}

func (app App) kubernetesEnv(kc *KConfig) []kubernetes.EnvironmentConfig {
	envs := []kubernetes.EnvironmentConfig{}
	for _, section := range app.Configs {
		configs := kc.Configs[section]
		for _, val := range configs {
			env := kubernetes.EnvironmentConfig{Name: val.Name}
			ref := kubernetes.KeyRef{
				Name: section,
				Key:  val.Name,
			}

			if val.Secret {
				env.ValueFrom.SecretKeyRef = ref
			} else {
				env.ValueFrom.ConfigMapKeyRef = ref
			}

			envs = append(envs, env)
		}
	}

	return envs
}

func (kc KConfig) AppType(env string) string {
	for _, proj := range kc.Project.Environments {
		if proj.Name == env {
			return proj.Type
		}
	}
	return "docker-compose"
}
