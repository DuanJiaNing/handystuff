package statsd

import (
	"fmt"
	"handystuff/config"
	"time"

	"gopkg.in/alexcesaro/statsd.v2"
)

type Cli interface {
	Timing(bucket string, value interface{})
}

func Client() Cli {
	if !config.Conf.Tick.Enable {
		return dummyClient{}
	}

	return statsdClient
}

type dummyClient struct{}

func (d dummyClient) Timing(string, interface{}) {
	// do nothing.
}

var statsdClient *statsd.Client

func Init() (err error) {
	cfg := config.Conf.Tick
	if !cfg.Enable {
		return nil
	}

	statsdClient, err = statsd.New(statsd.Address(cfg.Addr))

	return
}

func ReportHTTPCost(method, uri string, dur time.Duration) {
	bucket := fmt.Sprintf("api.%s.%s.cost", method, uri)
	Client().Timing(bucket, dur.Milliseconds())
}
