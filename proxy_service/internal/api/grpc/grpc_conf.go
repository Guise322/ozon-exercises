package grpc

import "github.com/Guise322/ozon-exercises/proxy_service/internal/api"

type GRPCConf struct {
	Server struct {
		Host string `yaml:"host"`
		Port int64  `yaml:"port"`
	} `yaml:"grpcServer"`
}

func ReadConfig(path string) (*GRPCConf, error) {
	var conf GRPCConf
	confPtr := &conf
	err := api.ReadConfig(confPtr, path)
	if err != nil {
		return nil, err
	}
	return confPtr, nil
}
