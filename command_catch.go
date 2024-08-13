package main

import "fmt"

func commandCatch(cfg *config, name string) error {
	fmt.Printf("catching Pokemon: %s\n", name)
	pokemonResp, err := cfg.pokeapiClient.ListPokemon(&name)
	if err != nil {
		return err
	}
	cfg.pokedex[name] = pokemonResp
	// change the chance of catching the pokemon from 100%
	// derive and use the chance of catching the pokemon
	fmt.Printf("caught: %s\n", pokemonResp.Name)
	return nil
}
