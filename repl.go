package main

import (
	"strings"
)

// split user's input based on whitespace
func cleanInput(text string) []string {
	text = strings.ToLower(text)
	words := strings.Fields(text)
	return words
}
