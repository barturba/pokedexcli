package main

import (
	"fmt"
)

func commandExplore(cfg *config, areaName string) error {
	fmt.Printf("Exploring %s...\n", areaName)

	var err error
	locationResp, err := cfg.pokeapiClient.ListLocation(&areaName)
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, pokemon := range locationResp.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}
	return nil
}
