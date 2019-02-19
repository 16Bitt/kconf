package main

import (
	"fmt"
	"github.com/16bitt/kconf/pkg/kconf"
)

func main() {
	config, err := kconf.Load("./kconf.yaml")
	if err != nil {
		panic(err)
	}

	configs, err := config.GenerateConfigmaps("staging")
	if err != nil {
		panic(err)
	}
	fmt.Println(configs)

	secrets, err := config.GenerateSecrets("staging")
	if err != nil {
		panic(err)
	}
	fmt.Println(secrets)

	deps, err := config.GenerateDeployments("staging")
	if err != nil {
		panic(err)
	}

	fmt.Println(deps)
}
