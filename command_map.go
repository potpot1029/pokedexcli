package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, _ string) error {
	locationArea, err := cfg.pokeapiClient.GetLocationAreas(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationArea.Next
	cfg.prevLocationsURL = locationArea.Previous

	for _, location := range locationArea.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapBack(cfg *config, _ string) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationArea, err := cfg.pokeapiClient.GetLocationAreas(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationArea.Next
	cfg.prevLocationsURL = locationArea.Previous

	for _, location := range locationArea.Results {
		fmt.Println(location.Name)
	}
	return nil
}
