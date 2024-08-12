package main

import "fmt"

func commandHelp(cfg *config, param string) error {
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
