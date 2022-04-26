package statsd

import (
	"fmt"
	"handystuff/config"
	"time"

	"gopkg.in/alexcesaro/statsd.v2"
)

type Client struct {
	statsdClient *statsd.Client
}

func NewClient(conf *config.TickConfig) (Client, error) {
	statsdClient, err := statsd.New(statsd.Address(conf.Addr))
	if err != nil {
		return Client{}, err
	}

	return Client{statsdClient}, nil
}

func (c Client) ReportHTTPCost(method, uri string, dur time.Duration) {
	bucket := fmt.Sprintf("api.%s.%s.cost", method, uri)
	c.statsdClient.Timing(bucket, dur.Milliseconds())
}

type DummyClient struct{}

func (d DummyClient) ReportHTTPCost(method, uri string, dur time.Duration) {
	// do nothing.
}

func NewDummyClient() DummyClient {
	return DummyClient{}
}
