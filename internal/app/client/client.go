package client

import (
	"crypto/tls"
	"net/http"
	"time"
)

type Client struct {
	Client   *http.Client
	BasePath string
}

func NewHTTPClient() Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c := &http.Client{
		Timeout:   30 * time.Second,
		Transport: tr,
	}

	client := &Client{
		Client:   c,
		BasePath: "http://external.roszdravnadzor.ru",
	}
	return *client
}

func (c *Client) Do(request *http.Request) (*http.Response, error) {
	return c.Client.Do(request)
}
