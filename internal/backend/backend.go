package backend

import (
	"net/url"
	"time"
)

type Backend struct {
	URL          *url.URL
	Alive        bool
	Enabled      bool
	Requests     int
	TotalLatency time.Duration
}
