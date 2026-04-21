package main

import "fmt"

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("no location name provided")
	}

	locationName := args[0]
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

	fmt.Println("Found Pokemon: ")
	for _, pokemon := range pokemons {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
