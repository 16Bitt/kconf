package kconf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func Load(path string) (*KConfig, error) {
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
