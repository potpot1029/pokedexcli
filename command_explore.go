package main

import "fmt"

func commandExplore(cfg *config, locationName string) error {
	if len(locationName) == 0 {
		return fmt.Errorf("no location name provided")
	}

	fmt.Printf("Exploring %s...\n", locationName)

	locationPokemons, err := cfg.pokeapiClient.GetLocationPokemon(locationName)
	if err != nil {
		return err
	}

	pokemons := locationPokemons.PokemonEncounters

	if len(pokemons) == 0 {
		fmt.Println("no pokemon found...")
		return nil
	}

	for _, pokemon := range pokemons {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
