package newrelic

import (
	"fmt"
	"net/http"
)

func (c *Client) UpdateDashboard(id string, dashboard []byte) error {
	req, err := c.newRequest(
		http.MethodPut,
		fmt.Sprintf("/v2/dashboards/%s", id),
		dashboard,
	)
	if err != nil {
		return err
	}
	_, err = c.do(req)
	return err
}
