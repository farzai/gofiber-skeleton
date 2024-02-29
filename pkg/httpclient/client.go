package httpclient

import (
	"net/http"
	"time"
)

type Client interface {
	Do(req *http.Request) (*http.Response, error)
}

func NewClient() Client {
	return &http.Client{
		Timeout: time.Duration(5 * time.Second),
	}
}
