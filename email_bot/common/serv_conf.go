package common

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Serv_conf struct {
	Server struct {
		Host          string `yaml:"host"`
		Port          int64  `yaml:"port"`
		TimeoutInMils int64  `yaml:"timeoutInMils"`
	} `yaml:"server"`
}

const (
	debugPath string = "../config/config.yml"
	prodPath  string = "config/config.yml"
)

func ReadConfig() (*Serv_conf, error) {
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
	var conf Serv_conf
	confPtr := &conf
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(confPtr)
	if err != nil {
		return nil, err
	}
	return confPtr, nil
}
