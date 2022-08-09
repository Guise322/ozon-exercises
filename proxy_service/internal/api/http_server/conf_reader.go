package http_server

import "github.com/Guise322/ozon-exercises/proxy_service/internal/api"

func ReadConfig(path string) (*httpConf, error) {
	var conf httpConf
	confPtr := &conf
	err := api.ReadConfig(confPtr, path)
	if err != nil {
		return nil, err
	}
	return confPtr, nil
}
