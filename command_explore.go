package main

import (
	"encoding/json"
	"fmt"

	"github.com/barturba/pokedexcli/internal/pokeapi"
)

func commandExplore(cfg *config, areaName string) error {

	fmt.Printf("Exploring %s...\n", areaName)

	locationResp := pokeapi.RespShallowLocation{}
	var err error

	val, found := cfg.pokeCache.Get(areaName)
	if found {
		err = json.Unmarshal(val, &locationResp)
		if err != nil {
			return err
		}
	} else {
		locationResp, err = cfg.pokeapiClient.ListLocation(&areaName)
		if err != nil {
			return err
		}
		dat, err := json.Marshal(&locationResp)
		if err != nil {
			return err
		}
		cfg.pokeCache.Add(areaName, dat)
	}

	fmt.Println("Found Pokemon:")
	for _, pokemon := range locationResp.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}
	return nil

}
