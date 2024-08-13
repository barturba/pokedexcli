package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/barturba/pokedexcli/internal/pokeapi"
	"github.com/barturba/pokedexcli/internal/pokecache"
)

type config struct {
	pokeapiClient    pokeapi.Client
	pokeCache        *pokecache.Cache
	nextLocationsURL *string
	prevLocationsURL *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	cmds := getCommands()

	fmt.Printf("Pokedex > ")
	for {
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}
		input := strings.Fields(scanner.Text())
		firstWord := ""
		secondWord := ""
		if len(input) > 0 {
			firstWord = input[0]
		}
		if len(input) > 1 {
			secondWord = input[1]
		}
		_, ok := cmds[firstWord]
		fmt.Println("")
		if ok {
			err := cmds[firstWord].callback(cfg, secondWord)
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Printf("Pokedex > ")
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"explore": {
			name:        "explore",
			description: "Displays the informaion about a location area in the Pokemon world.",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch the given Pokemon.",
			callback:    commandCatch,
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
