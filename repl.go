package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/potpot1029/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	caughtPokemon    map[string]pokeapi.Pokemon
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		// get and clean user input
		userInput := scanner.Text()
		cleanUserInput := cleanInput(userInput)
		if len(cleanUserInput) == 0 {
			continue
		}

		// process command
		commandName := cleanUserInput[0]
		args := []string{}
		if len(cleanUserInput) > 1 {
			args = cleanUserInput[1:]
		}

		command, ok := getCommands()[commandName]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			callback := command.callback

			err := callback(cfg, args...)
			if err != nil {
				fmt.Printf("error running command: %v\n", err)
			}
		}
	}
}

// split user's input based on whitespace
func cleanInput(text string) []string {
	text = strings.ToLower(text)
	words := strings.Fields(text)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the names of next 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of previous 20 location areas in the Pokemon world",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore <location-area>",
			description: "See a list of all Pokemon in <location-area>",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon-name>",
			description: "Catch <pokemon-name> with a Pokeball",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon-name>",
			description: "Inspect information about caught <pokemon-name>",
			callback:    commandInspect,
		},
	}
}
