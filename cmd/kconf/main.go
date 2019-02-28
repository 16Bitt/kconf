package main

import (
	"fmt"
	"flag"
	"github.com/16bitt/kconf/pkg/kconf"
)

func main() {
	confPath := flag.String("config", "./kconf.yaml", "Path to your kconf file")
	environment := flag.String("env", "development", "Environment to build configs for")
	flag.Parse()
	// init := flag.Bool("init", false, "Generate a default config")

	config, err := kconf.Load(*confPath)
	if err != nil {
		panic(err)
	}

	envType := config.AppType(*environment)

	// Fast-path for docker-compose
	if envType == "docker-compose" {
		dc, err := config.GenerateDockerCompose(*environment)
		if err != nil {
			panic(err)
		}

		fmt.Println(dc)
		return
	}

	configs, err := config.GenerateConfigmaps(*environment)
	if err != nil {
		panic(err)
	}
	fmt.Println(configs)

	secrets, err := config.GenerateSecrets(*environment)
	if err != nil {
		panic(err)
	}
	fmt.Println(secrets)

	deps, err := config.GenerateDeployments(*environment)
	if err != nil {
		panic(err)
	}

	fmt.Println(deps)
}
