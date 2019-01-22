package main

import (
  "fmt"
  "github.com/16bitt/kconf"
)

func main() {
  config, err := kconf.Load("./kconf.yaml")
  if err != nil {
    panic(err)
  }

  fmt.Println(config)
}