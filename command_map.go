package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/barturba/pokedexcli/internal/pokeapi"
)

func commandMap(cfg *config) error {
	// check cache before calling pokeapi
	locationsResp := pokeapi.RespShallowLocations{}
	var err error
	fmt.Printf("commandMap()\n")

	val, ok := cfg.pokeCache.Get(*cfg.nextLocationsURL)
	if ok {
		err = json.Unmarshal(val, &locationsResp)
		if err != nil {
			return err
		}
	} else {
		locationsResp, err = cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
		if err != nil {
			return err
		}
		cfg.nextLocationsURL = locationsResp.Next
		cfg.prevLocationsURL = locationsResp.Previous

	}

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapB(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}
	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
