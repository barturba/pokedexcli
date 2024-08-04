package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type config struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

func getCommands() map[string]cliCommand {
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	cfg := config{}
	cfg.Next = "https://pokeapi.co/api/v2/location/?limit=20"
	cfg.Previous = cfg.Next

	fmt.Printf("Pokedex > ")
	for {
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}
		_, ok := commands[scanner.Text()]
		fmt.Println("")
		if ok {
			err := commands[scanner.Text()].callback(&cfg)
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Printf("Pokedex > ")
	}
}

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	commands := getCommands()
	for name, command := range commands {
		fmt.Printf("%s: %s\n", name, command.description)
	}
	fmt.Println("")
	return nil
}

func commandMap(cfg *config) error {

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

func commandMapB(cfg *config) error {
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

func commandExit(cfg *config) error {
	os.Exit(0)
	return nil
}
