package kconf

// App is a container within your application
type App struct {
	Name    string
	Build   string
	Configs []string
	Image   string
	Helm    HelmConfig
	Ports   []string
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
