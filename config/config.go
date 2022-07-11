package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type global struct {
	Addr string
	Auth auth
}

type auth struct {
	Enabled  bool
	Username string
	Password string
	IPList   []string // this will use for allowed ip addresses
}

type envfiles struct {
	Default      string
	Environments map[string]environment
}

type environment struct {
	Path     string
	Excludes []string
}

type Config struct {
	Global   global
	EnvFiles map[string]envfiles
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
