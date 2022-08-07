package api

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Serv_conf struct {
	Server struct {
		Host string `yaml:"host"`
		Port int64  `yaml:"port"`
	} `yaml:"server"`
}

func ReadConfig(path string) (*Serv_conf, error) {
	var file *os.File
	var err error
	file, err = os.Open(path)
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
