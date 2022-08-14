package notif_client

import (
	"net/http"

	"github.com/Guise322/ozon-exercises/proxy_service/internal"
)

type notifClient struct {
	client http.Client
	url    string
}

func NewNotifClient(confPath string) (*notifClient, error) {
	var conf NotifClientConf
	err := internal.ReadConfig(&conf, confPath)
	if err != nil {
		return nil, err
	}
	return &notifClient{client: http.Client{}, url: conf.NotifClient.URL}, nil
}
