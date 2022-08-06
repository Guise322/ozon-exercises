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

const (
	debugPath string = "../config/config.yml"
	prodPath  string = "config/config.yml"
)

func ReadConfig() (*Config, error) {
	var file *os.File
	var err error
	var p string
	if IsDebugging() {
		p = debugPath
	} else {
		p = prodPath
	}
	file, err = os.Open(p)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var conf Config
	confPtr := &conf
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(confPtr)
	if err != nil {
		return nil, err
	}
	return confPtr, nil
}
