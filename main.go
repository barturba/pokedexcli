package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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

type CliCommand struct {
	name        string
	description string
	Callback    func(cfg *Config) error
}

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfgNew := &config{
		pokeapiClient: pokeClient,
	}
	fmt.Sprintf("%v", cfgNew)

	scanner := bufio.NewScanner(os.Stdin)
	cmds := GetCommands()
	cfg := Config{}
	cfg.Next = "https://pokeapi.co/api/v2/location/?limit=20"
	cfg.Previous = cfg.Next

	fmt.Printf("Pokedex > ")
	for {
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}
		_, ok := cmds[scanner.Text()]
		fmt.Println("")
		if ok {
			err := cmds[scanner.Text()].Callback(&cfg)
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Printf("Pokedex > ")
	}
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

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			Callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world.",
			Callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 location areas in the Pokemon world.",
			Callback:    commandMapB,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			Callback:    commandExit,
		},
	}
}
