package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("you haven't catch a pokemon yet...")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for key, _ := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", key)
	}

	return nil
}
