package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		ok := scanner.Scan()
		if !ok {

		}
		userInput := scanner.Text()
		cleanUserInput := cleanInput(userInput)
		fmt.Printf("Your command was: %s\n", cleanUserInput[0])
	}
}
