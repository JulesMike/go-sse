package sse

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

// Options holds server configurations.
type Options struct {
	// RetryInterval change EventSource default retry interval (milliseconds).
	RetryInterval int
	// Headers allow to set custom headers (useful for CORS support).
	Headers map[string]string
	// OnClientConnect allows to execute custom logic when client connects for
	// first time
	OnClientConnect func(c *Client)
	// ChannelNameFunc allow to create custom channel names.
	// Default channel name is the request path.
	ChannelNameFunc func(*http.Request) string
	// All usage logs end up in Logger
	Logger *logrus.Entry
}

func (opt *Options) hasHeaders() bool {
	return opt.Headers != nil && len(opt.Headers) > 0
}
