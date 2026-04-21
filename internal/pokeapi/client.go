package pokeapi

import (
	"net/http"
	"time"

	"github.com/potpot1029/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	pokeCache  pokecache.Cache
}

// create a Client
func NewClient(timeout, interval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		pokeCache: pokecache.NewCache(interval),
	}
}
