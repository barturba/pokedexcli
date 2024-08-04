package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/barturba/pokedexcli/internal/commands"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cmds := commands.GetCommands()
	cfg := commands.Config{}
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
