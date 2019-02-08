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

	fmt.Println(config)
	err = config.GenerateConfigmaps("./test")
	if err != nil {
		panic(err)
	}
}
