package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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

func commandHelp(cfg *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	commands := GetCommands()
	for name, command := range commands {
		fmt.Printf("%s: %s\n", name, command.description)
	}
	fmt.Println("")
	return nil
}

func commandMap(cfg *Config) error {

	data, err := fetchAPI(cfg.Next)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(data), cfg)
	if err != nil {
		return err
	}
	for location := range cfg.Results {
		fmt.Printf("%s\n", cfg.Results[location].Name)
	}
	return nil
}

func commandMapB(cfg *Config) error {
	data, err := fetchAPI(cfg.Previous)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(data), cfg)
	if err != nil {
		return err
	}
	for location := range cfg.Results {
		fmt.Printf("%s\n", cfg.Results[location].Name)
	}
	return nil
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

func commandExit(cfg *Config) error {
	os.Exit(0)
	return nil
}
