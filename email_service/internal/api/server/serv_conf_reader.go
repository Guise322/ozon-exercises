package server

import "github.com/Guise322/ozon-exercises/email_service/internal/api"

func readConfig(path string) (*ServConf, error) {
	var conf ServConf
	confPtr := &conf
	err := api.ReadConfig(confPtr, path)
	if err != nil {
		return nil, err
	}
	return confPtr, nil
}
