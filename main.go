package main

import (
	"fmt"
	"time"

	"github.com/barturba/pokedexcli/internal/pokeapi"
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
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfgNew := &config{
		pokeapiClient: pokeClient,
	}
	fmt.Sprintf("%v", cfgNew)
	url := baseURL
	cfgNew.nextLocationsURL = &url
	cfgNew.prevLocationsURL = cfgNew.nextLocationsURL
	startRepl(cfgNew)

}
