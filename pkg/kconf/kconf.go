package kconf

const DefaultConfig = `
version: 1
project:
  configDir: config/kubernetes
  environments:
  - name: development
    type: docker-compose
  - name: production
    type: kubernetes
  kubernetes:
    namespace: project

apps:
- name: project
  image: user/project
  type: kubernetes
  build: .
  configs:
  - project
  ports:
  - "8000:8000"
  kubernetes:
    replicas: "3"
    labels:
      tier: backend
      app: project
    limits:
      memory: 1Gi
      cpu: 100m
    requests:
      memory: 1Gi
      cpu: 100m

configs:
  project:
  - name: HELLO_WORLD
    secret: false
    defaults:
    - type: kubernetes
      environment: production
      value: Hello, production!
`

type KConfig struct {
	Version int
	Project Project
	Apps    []App
	Configs map[string]Configs
}

type Project struct {
	ConfigDir    string `yaml:"configDir"`
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
