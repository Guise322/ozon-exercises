package notif

import (
	"context"
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

func (c *notifClient) Notify(ctx context.Context, cmd *contract.NotifCmd) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.url, nil)
	if err != nil {
		return err
	}
	res, err := c.client.Do(req)
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
