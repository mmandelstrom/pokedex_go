package main

import (
	"strings"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text)) //Converts input string to slice of lowercase words
}
