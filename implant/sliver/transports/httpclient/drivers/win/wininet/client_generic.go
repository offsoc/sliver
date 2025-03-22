
package http

import (
	"errors"
	stdhttp "net/http"
)

// Client is an empty struct, if not Windows.
type Client struct{}

// NewClient is only supported on Windows.
func NewClient() (*Client, error) {
	return nil, errors.New("unsupported OS")
}

// Do is only supported on Windows.
func (c *Client) Do(r *stdhttp.Request) (*stdhttp.Response, error) {
	return nil, errors.New("unsupported OS")
}

// Get is only supported on Windows.
func (c *Client) Get(url string) (*stdhttp.Response, error) {
	return nil, errors.New("unsupported OS")
}

// Head is only supported on Windows.
func (c *Client) Head(url string) (*stdhttp.Response, error) {
	return nil, errors.New("unsupported OS")
}
// Post is only supported on Windows.
func (c *Client) Post(
	url string,
	contentType string,
	body []byte,
) (*stdhttp.Response, error) {
	return nil, errors.New("unsupported OS")
}
