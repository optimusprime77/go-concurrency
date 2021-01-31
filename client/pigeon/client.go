package pigeon

import (
	"time"

	"github.com/optimusprime77/go-concurrency/client"
	"github.com/optimusprime77/go-concurrency/config"
	"golang.org/x/sync/singleflight"
)

// Client contains request and concurrency attributes.
type Client struct {
	Endpoint     string
	HTTPClient   client.HTTPClient
	RequestGroup singleflight.Group
	semaphore    chan struct{}
}

// Init initializes request and concurrency attributes.
func (c *Client) Init(config *config.Config) error {
	timeout := 5 * time.Second
	c.Endpoint = config.Endpoint
	c.HTTPClient = client.NewHTTPClient(client.Parameters{Timeout: &timeout})
	c.semaphore = make(chan struct{}, config.MaximumConcurrency)

	return nil
}
