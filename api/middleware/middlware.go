package middleware

import "handystuff/api"

type Middleward struct {
	statsd api.StatsdClient
}

func NewMiddleward(
	statsd api.StatsdClient,
) Middleward {
	return Middleward{statsd: statsd}
}
