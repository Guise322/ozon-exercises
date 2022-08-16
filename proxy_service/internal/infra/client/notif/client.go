package notif

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/contract"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/conf"
)

type notifClient struct {
	client http.Client
	url    string
}

func NewNotifClient(c conf.NotifClientConf) *notifClient {
	return &notifClient{client: http.Client{}, url: c.NotifClient.URL}
}

func (c *notifClient) Notify(cmd *contract.NotifCmd) error {
	res, err := c.client.Post(c.url, "application/json", nil)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusOK {
		return nil
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	return fmt.Errorf("the server returns: %v\n%v", res.Status, body)
}
