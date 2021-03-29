package newrelic

import (
	"errors"
	"fmt"
	"net/http"
)

func (c *Client) CreateDashboard(dashboard []byte) error {
	req, err := c.newRequest(
		http.MethodPost,
		"/v2/dashboards",
		dashboard,
	)
	if err != nil {
		return err
	}
	res, err := c.do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("Failed to create dashboard with status '%d'", res.StatusCode))
	}

	return nil
}
