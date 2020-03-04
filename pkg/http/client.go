package http

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	ContentTypeYaml    = "application/yaml"
	ContentTypeJson    = "application/json"
	ContentTypeDefault = ContentTypeYaml
)

type Client struct {
	*http.Client
	token string
}

func (c *Client) Do(method string, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, errors.New("http: Can not create request")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))
	req.Header.Add("Content-Type", ContentTypeDefault)

	return c.Client.Do(req.WithContext(ctx))
}

func NewClient(token string) *Client {
	return &Client{
		Client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		},
		token: token,
	}
}
