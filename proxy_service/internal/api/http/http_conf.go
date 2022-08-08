package http

import "github.com/Guise322/ozon-exercises/proxy_service/internal/api"

type httpConf struct {
	Server struct {
		Host string `yaml:"host"`
		Port int64  `yaml:"port"`
	} `yaml:"httpServer"`
}

func ReadConfig(path string) (*httpConf, error) {
	var conf httpConf
	confPtr := &conf
	err := api.ReadConfig(confPtr, path)
	if err != nil {
		return nil, err
	}
	return confPtr, nil
}
