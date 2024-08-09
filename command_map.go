package main

import (
	"encoding/json"
	"fmt"
)

func commandMap(cfg *Config) error {

	data, err := fetchAPI(cfg.Next)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(data), cfg)
	if err != nil {
		return err
	}
	for location := range cfg.Results {
		fmt.Printf("%s\n", cfg.Results[location].Name)
	}
	return nil
}

func commandMapB(cfg *Config) error {
	data, err := fetchAPI(cfg.Previous)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(data), cfg)
	if err != nil {
		return err
	}
	for location := range cfg.Results {
		fmt.Printf("%s\n", cfg.Results[location].Name)
	}
	return nil
}
