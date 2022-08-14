package notif_client

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/contract"
)

func (c *notifClient) Notify(cmd *contract.NotifCmd) (interface{}, error) {
	res, err := c.client.Post(c.url, "text/json", nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusOK {
		return nil, nil
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return nil, fmt.Errorf("the server returns: %v\n%v", res.Status, body)
}
