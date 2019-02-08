package kconf

type KConfig struct {
	Version int
	Project Project
	Apps    []App
	Configs map[string]Configs
}

type Project struct {
	ConfigDir    string
	Environments []Environment
	Kubernetes   KubernetesProject
}

type Environment struct {
	Name string
	Type string
}

type KubernetesProject struct {
	Namespace string
}

type HelmConfig struct {
	Chart     string
	Name      string
	Namespace string
}
