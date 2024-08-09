package pokeapi

import (
	"net/http"
	"time"

	"github.com/barturba/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	pokecache  *pokecache.Cache
}

func NewClient(timeout time.Duration, c *pokecache.Cache) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		pokecache: c,
	}
}
