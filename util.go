package main

import (
	"strings"
	"time"

	api "github.com/mmandelstrom/pokedex_go/internal/pokecache"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

type config struct {
	Next     string
	Previous *string
}

var cfg = config{Next: "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"}
var c = api.NewCache(7 * time.Second)
var pokeDex = api.Pokedex{
	PokemonMap: make(map[string]api.Pokemon),
}

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
		"catch": {
			name:        "catch",
			description: "attempt to catch a pokemon",
			callback:    commandCatch(c, &pokeDex),
		},
		"inspect": {
			name:        "inspect",
			description: "inspect a pokemon in pokedex",
			callback:    commandInspect(&pokeDex),
		},
		"pokedex": {
			name:        "pokedex",
			description: "displays all pokemon in dex",
			callback:    commandPokedex(),
		},
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(string) error
}
