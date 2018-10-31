package newrelic

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
)

const insightsBaseUrl = "https://api.eu.newrelic.com/" // TODO: Configurable
const userAgent = "insights-client/0.1.0"

type Client struct {
	BaseURL     *url.URL
	UserAgent   string
	AdminApiKey string

	httpClient *http.Client
}

func CreateClient(adminApiKey string) (*Client, error) {
	baseUrl, err := url.Parse(insightsBaseUrl)
	if err != nil {
		return nil, err
	}

	return &Client{
		BaseURL:     baseUrl,
		UserAgent:   userAgent,
		AdminApiKey: adminApiKey,
		httpClient:  http.DefaultClient,
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
	debug(httputil.DumpRequestOut(req, true))
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	debug(httputil.DumpResponse(resp, true))
	return resp, nil
}

// TODO: Verbose logging
func debug(data []byte, err error) {
	if err == nil {
		fmt.Printf("\n%s\n", data)
	} else {
		fmt.Printf("\n%s\n", err)
	}
}
