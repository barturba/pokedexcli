package main

import (
	"fmt"
	"io"
	"net/http"
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

func fetchAPI(url string) (string, error) {

	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return "", fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", body), nil
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 location areas in the Pokemon world.",
			callback:    commandMapB,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
