package newrelic

import (
	"errors"
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
	res, err := c.do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("Failed to update dashboard with status '%d'", res.StatusCode))
	}

	return nil
}
