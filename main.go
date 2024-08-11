package main

import (
	"time"

	"github.com/barturba/pokedexcli/internal/pokeapi"
	"github.com/barturba/pokedexcli/internal/pokecache"
)

type Config struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

const (
	baseURL = "https://pokeapi.co/api/v2/location/?limit=20"
)

func main() {
	pokeCache := pokecache.NewCache(15 * time.Second)
	pokeClient := pokeapi.NewClient(5*time.Second, pokeCache)
	pokeCache.Add("abc.com", []byte("content"))
	cfg := &config{
		pokeapiClient: pokeClient,
		pokeCache:     pokeCache,
	}
	url := baseURL
	cfg.nextLocationsURL = &url
	cfg.prevLocationsURL = cfg.nextLocationsURL
	startRepl(cfg)

}
