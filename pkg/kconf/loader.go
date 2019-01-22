package kconf

import (
  "io/ioutil"
  "gopkg.in/yaml.v2"
) 

func Load(path string) *KConfig, err {
  conf := &KConfig{}
  bytes, err := ioutil.ReadFile(path)
  if err != nil {
    return nil, err
  }

  err = yaml.Unmarshal(bytes, conf)
  if err != nil {
    return nil, err
  }

  return conf, nil
}