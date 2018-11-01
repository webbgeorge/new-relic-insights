package newrelic

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"github.com/sirupsen/logrus"
)

const insightsBaseUrl = "https://api.eu.newrelic.com/" // TODO: Configurable
const userAgent = "insights-client/0.1.0"

type Client struct {
	BaseURL     *url.URL
	UserAgent   string
	AdminApiKey string

	httpClient *http.Client
	logger     *logrus.Logger
}

func CreateClient(adminApiKey string, logger *logrus.Logger) (*Client, error) {
	baseUrl, err := url.Parse(insightsBaseUrl)
	if err != nil {
		return nil, err
	}

	return &Client{
		BaseURL:     baseUrl,
		UserAgent:   userAgent,
		AdminApiKey: adminApiKey,
		httpClient:  http.DefaultClient,
		logger:      logger,
	}, nil
}

func (c *Client) newRequest(method, path string, body []byte) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = bytes.NewBuffer(body)
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("X-Api-Key", c.AdminApiKey)

	return req, nil
}

func (c *Client) do(req *http.Request) (*http.Response, error) {
	c.debug(httputil.DumpRequestOut(req, true))
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	c.debug(httputil.DumpResponse(resp, true))
	return resp, nil
}

func (c *Client) debug(data []byte, err error) {
	if err == nil {
		c.logger.Debugf("\n%s\n", data)
	} else {
		c.logger.Debugf("\n%s\n", err)
	}
}
