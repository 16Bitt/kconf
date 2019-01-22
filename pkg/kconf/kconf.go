package kconf

type KConfig struct {
  Version int
  Project Project
  Apps    []App
  Configs map[string]Config
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

type App struct {
  Name    string
  Build   string
  Configs []string
  Image   string
  Helm    HelmConfig
  Ports   string
}

type HelmConfig struct {
  Chart     string
  Name      string
  Namespace string
}

type Config struct {
  Name     string
  Secret   bool
  Defaults ConfigDefault
}

type ConfigDefault struct {
  Type  string
  Value string
}