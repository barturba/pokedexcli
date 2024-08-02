package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
			err = commands[scanner.Text()].callback()
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Printf("Pokedex > ")
	}
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	fmt.Println("")
	return nil
}
func commandMap() error {
	fmt.Println("Future map function")
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}
