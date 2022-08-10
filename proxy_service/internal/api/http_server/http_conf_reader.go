package http_server

import "github.com/Guise322/ozon-exercises/proxy_service/internal/api"

func readConfig(path string) (*HTTPConf, error) {
	var conf HTTPConf
	confPtr := &conf
	err := api.ReadConfig(confPtr, path)
	if err != nil {
		return nil, err
	}
	return confPtr, nil
}
