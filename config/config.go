package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type global struct {
	Auth auth
}

type auth struct {
	Enabled  bool
	Username string
	Password string
	IPList   []string // this will use for allowed ip addresses
}

type Config struct {
	Global global
}

func GetConfig() (*Config, error) {
	yfile, err := ioutil.ReadFile("netenv.yml")
	if err != nil {
		return nil, err
	}

	config := &Config{}

	err = yaml.Unmarshal(yfile, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
