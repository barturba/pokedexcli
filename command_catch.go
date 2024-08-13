package main

import "fmt"

func commandCatch(cfg *config, name string) error {
	fmt.Printf("catching Pokemon: %s\n", name)
	// https://pokeapi.co/api/v2/pokemon/{id or name}
	pokemonResp, err := cfg.pokeapiClient.ListPokemon(&name)
	if err != nil {
		return err
	}
	fmt.Printf("caught: %s\n", pokemonResp.Name)
	return nil
}
