package pokeapi

import (
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
}

// create a Client
func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
