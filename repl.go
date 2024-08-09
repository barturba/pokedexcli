package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/barturba/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	cmds := GetCommands()

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
			err := cmds[scanner.Text()].callback(cfg)
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Printf("Pokedex > ")
	}
}
