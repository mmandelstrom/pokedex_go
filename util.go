package main

import (
	"strings"
	"time"

	api "github.com/mmandelstrom/pokedex_go/internal/pokecache"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text)) //Converts input string to slice of lowercase words
}

type config struct {
	Next     string
	Previous *string
}

var cfg = config{Next: "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"}
var c = api.NewCache(7 * time.Second)

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit(&cfg),
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp(&cfg),
		},
		"map": {
			name:        "map",
			description: "Display a list of 20 locations",
			callback:    commandMap(&cfg, c),
		},
		"mapb": {
			name:        "mapb",
			description: "Display a list of the 20 previous locations",
			callback:    commandMapB(&cfg, c),
		},
		"explore": {
			name:        "explore",
			description: "Displays pokemon in area",
			callback:    commandExplore(c),
		},
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(string) error
}
