package newrelic

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *Client) GetDashboard(id string) ([]byte, error) {
	req, err := c.newRequest(
		http.MethodGet,
		fmt.Sprintf("/v2/dashboards/%s", id),
		nil,
	)
	if err != nil {
		return nil, err
	}

	res, err := c.do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Failed to get dashboard with status '%d'", res.StatusCode))
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
