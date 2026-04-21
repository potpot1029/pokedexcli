package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("no pokemon name provided")
	}

	pokemonName := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	// get pokemon info....
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	// generate catch chance base on base experience
	catchChance := rand.Intn(pokemon.BaseExperience)
	threshold := 30

	if catchChance <= threshold {
		fmt.Printf("%s was caught!\n", pokemonName)

		cfg.caughtPokemon[pokemonName] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}
	fmt.Println("you can now inspect the pokemon with the command!")

	return nil
}
