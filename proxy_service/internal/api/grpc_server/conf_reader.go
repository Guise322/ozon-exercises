package grpc_server

import "github.com/Guise322/ozon-exercises/proxy_service/internal/api"

func ReadConfig(path string) (*GRPCConf, error) {
	var conf GRPCConf
	confPtr := &conf
	err := api.ReadConfig(confPtr, path)
	if err != nil {
		return nil, err
	}
	return confPtr, nil
}
