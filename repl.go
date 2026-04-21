package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	c := config{
		Next:     baseURL + "/1",
		Previous: "",
	}

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
		command, ok := getCommands()[commandName]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			callback := command.callback
			err := callback(&c)
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
	callback    func(*config) error
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
	}
}

type config struct {
	Next     string
	Previous string
}

const baseURL = "https://pokeapi.co/api/v2/location-area"

func getNextURL(url string, step int) string {
	id, err := strconv.Atoi(path.Base(url))
	if err != nil {
		fmt.Println(err)
	}
	id = id + step
	return fmt.Sprintf("%s/%d", baseURL, id)
}
