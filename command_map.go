package main

import (
	"encoding/json"
	"fmt"

	"github.com/barturba/pokedexcli/internal/pokeapi"
)

func commandMap(cfg *config, param string) error {
	// check cache before calling pokeapi
	locationsResp := pokeapi.RespShallowLocations{}
	var err error

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

		if locationsResp.Previous == nil {
			url := baseURL
			locationsResp.Previous = &url
		}

		dat, err := json.Marshal(&locationsResp)
		if err != nil {
			return err
		}
		cfg.pokeCache.Add(*cfg.nextLocationsURL, dat)

	}
	cfg.prevLocationsURL = locationsResp.Previous
	cfg.nextLocationsURL = locationsResp.Next

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapB(cfg *config, param string) error {
	// check cache before calling pokeapi
	locationsResp := pokeapi.RespShallowLocations{}
	var err error
	// map -> 0 -> 20
	// map -> 20 -> 40
	// mapb -> 40 -> 20

	val, ok := cfg.pokeCache.Get(*cfg.prevLocationsURL)
	if ok {
		err = json.Unmarshal(val, &locationsResp)
		if err != nil {
			return err
		}
	} else {
		locationsResp, err = cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
		if err != nil {
			return err
		}

		dat, err := json.Marshal(&locationsResp)
		if err != nil {
			return err
		}
		cfg.pokeCache.Add(*cfg.prevLocationsURL, dat)

	}
	cfg.prevLocationsURL = locationsResp.Previous
	cfg.nextLocationsURL = locationsResp.Next

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
