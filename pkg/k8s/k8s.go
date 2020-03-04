package k8s

import (
	"context"
)

type Client interface {
	Apply() Apply
}

type Apply interface {
	Do(ctx context.Context, directory string, filePath string) error
}

type client struct {
	url       string
	namespace string
	token     string
	applysvc  Apply
}

func New(url string, namespace string, token string, applySvc Apply) Client {
	return &client{
		url:       url,
		namespace: namespace,
		token:     token,
		applysvc:  applySvc,
	}
}

func (c *client) Apply() Apply {
	return c.applysvc
}
