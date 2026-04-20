package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
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
		command, ok := getCommands()[commandName]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			callback := command.callback
			err := callback()
			if err != nil {
				fmt.Errorf("error running command: %v", err)
			}
		}
	}
}

func processCommand(userInput []string) {
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
	callback    func() error
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
	}

}
