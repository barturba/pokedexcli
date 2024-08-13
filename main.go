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
	baseURL      = "https://pokeapi.co/api/v2/location/?offset=0&limit=20"
	cacheTimeout = 30 * time.Minute
)

func main() {
	pokeCache := pokecache.NewCache(cacheTimeout)
	pokeClient := pokeapi.NewClient(5*time.Second, pokeCache)
	pokedex := make(map[string]pokeapi.RespPokemon)
	pokeCache.Add("abc.com", []byte("content"))
	cfg := &config{
		pokeapiClient: pokeClient,
		pokeCache:     pokeCache,
		pokedex:       pokedex,
	}
	url := baseURL
	cfg.nextLocationsURL = &url
	cfg.prevLocationsURL = cfg.nextLocationsURL
	startRepl(cfg)

}
