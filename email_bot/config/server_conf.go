package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
}

func ReadConfig() (*Config, error) {
	var file *os.File
	var err error
	if IsDebugging() {
		file, err = os.Open("../config/config.yml")
	} else {
		file, err = os.Open("email_bot/config/config.yml")
	}
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var conf Config
	decoder := yaml.NewDecoder(file)
	confPtr := &conf
	err = decoder.Decode(confPtr)
	if err != nil {
		return nil, err
	}
	return confPtr, nil
}
