package main

import (
	"fmt"
	"os"

	"github.com/mmandelstrom/pokedex_go/internal/api"
)

func commandHelp(cfg *config) func() error {
	return func() error {
		fmt.Print()
		fmt.Println("Welcome to the Pokedex!")
		fmt.Print("Usage:\n\n")
		for _, cmd := range getCommands() {
			fmt.Printf("%s: %s\n", cmd.name, cmd.description)
		}
		return nil
	}

}

func commandExit(cfg *config) func() error {
	return func() error {
		fmt.Println("Closing the Pokedex... Goodbye!")
		os.Exit(0)
		return nil
	}
}

func commandMap(cfg *config) func() error {
	return func() error {
		location := api.GetLocation(cfg.Next)
		api.PrintPokeLocation(&location)
		cfg.Next = location.Next
		cfg.Previous = location.Previous
		return nil
	}
}

func commandMapB(cfg *config) func() error {
	return func() error {
		if cfg.Previous == nil {
			fmt.Println("you're on the first page")
		} else {
			location := api.GetLocation(*cfg.Previous)
			api.PrintPokeLocation(&location)
			cfg.Next = location.Next
			cfg.Previous = location.Previous
		}
		return nil
	}
}
