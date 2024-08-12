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

	val, ok := cfg.pokeCache.Get(areaName)
	if ok {
		err = json.Unmarshal(val, &locationResp)
		if err != nil {
			return err
		}
		fmt.Printf("got explore data from cache")
	} else {

		locationResp, err = cfg.pokeapiClient.ListLocation(&areaName)
		if err != nil {
			return err
		}
		fmt.Printf("got explore data from api")
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
