package main

import (
	"github.com/16bitt/kconf/pkg/kconf"
)

func main() {
	config, err := kconf.Load("./kconf.yaml")
	if err != nil {
		panic(err)
	}

	err = config.GenerateConfigmaps("./test")
	if err != nil {
		panic(err)
	}

	err = config.GenerateSecrets("./test")
	if err != nil {
		panic(err)
	}
}
