package main

import "fmt"

func commandPokedex(cfg *config, name string) error {
	fmt.Printf("Your Pokedex:\n")
	for _, pokemon := range cfg.pokedex {
		fmt.Printf(" -%s\n", pokemon.Name)
	}
	return nil
}
