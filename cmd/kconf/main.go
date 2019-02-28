package main

import (
	"fmt"
	"flag"
	"os"
	"io/ioutil"
	"github.com/16bitt/kconf/pkg/kconf"
)

func main() {
	confPath := flag.String("config", "./kconf.yaml", "Path to your kconf file")
	environment := flag.String("env", "development", "Environment to build configs for")
	init := flag.Bool("init", false, "Generate a default config")
	flag.Parse()

	if *init {
		fmt.Println("Generating default kconf config")
		err := ioutil.WriteFile(*confPath, []byte(kconf.DefaultConfig), 0777)
		if err != nil {
			panic(err)
		}
	}

	config, err := kconf.Load(*confPath)
	if err != nil {
		panic(err)
	}

	envType := config.AppType(*environment)

	// Fast-path for docker-compose
	if envType == "docker-compose" {
		fmt.Println("Writing docker-compose.yml to .")
		dc, err := config.GenerateDockerCompose(*environment)
		if err != nil {
			panic(err)
		}


		err = ioutil.WriteFile("./docker-compose.yml", []byte(dc), 0777)
		if err != nil {
			panic(err)
		}

		return
	}

	targetPath := fmt.Sprintf("%s/%s", config.Project.ConfigDir, *environment)
	if _, err := os.Stat(targetPath); os.IsNotExist(err) {
		fmt.Printf("ConfigDir `%s` does not exist, creating\n", targetPath)
		err = os.MkdirAll(targetPath, 0777)
		if err != nil {
			panic(err)
		}
	}

	configs, err := config.GenerateConfigmaps(*environment)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(resourcePath(config, "configmaps.yml", *environment), []byte(configs), 0777)
	if err != nil {
		panic(err)
	}

	secrets, err := config.GenerateSecrets(*environment)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(resourcePath(config, "secrets.yml", *environment), []byte(secrets), 0777)
	if err != nil {
		panic(err)
	}

	deps, err := config.GenerateDeployments(*environment)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(resourcePath(config, "deployments.yml", *environment), []byte(deps), 0777)
	if err != nil {
		panic(err)
	}
}

func resourcePath(kc *kconf.KConfig, name, env string) string {
	return fmt.Sprintf("%s/%s/%s", kc.Project.ConfigDir, env, name)
}
