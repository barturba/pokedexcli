package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type cliCommand struct {
	name string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: "help",
		},
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	for {
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("pokedex > ")
		fmt.Println(commands[scanner.Text()].name)
	}
}
