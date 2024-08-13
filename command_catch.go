package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, name string) error {
	fmt.Printf("Throwing a ball at %s...\n", name)
	pokemonResp, err := cfg.pokeapiClient.ListPokemon(&name)
	if err != nil {
		return err
	}
	caught := 1000 - rand.Intn(pokemonResp.BaseExperience)
	if caught > 500 {
		fmt.Printf("%v was caught!\n", name)
		cfg.pokedex[name] = pokemonResp
	} else {
		fmt.Printf("%v escaped!\n", name)
	}
	return nil
}
