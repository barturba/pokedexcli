package main

import (
	"fmt"

	"github.com/barturba/pokedexcli/internal/pokeapi"
)

func commandInspect(cfg *config, name string) error {
	pokemon, ok := cfg.pokedex[name]
	if ok {
		printPokemonStats(pokemon)
	} else {
		fmt.Printf("you have not caught that pokemon\n")
	}

	return nil
}

func printPokemonStats(pokemon pokeapi.RespPokemon) {
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, t := range pokemon.Types {
		fmt.Printf(" - %s\n", t.Type.Name)
	}
}
