package main

import (
	"fmt"

	"github.com/potpot1029/pokedexcli/internal/pokeapi"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("no pokemon name provided")
	}

	pokemonName := args[0]
	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		fmt.Printf("you haven't caught %s\n", pokemonName)
		return nil
	}

	printPokemonStat(pokemon)

	return nil
}

func printPokemonStat(pokemon pokeapi.Pokemon) {
	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, s := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Println("Types: ")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}
}
