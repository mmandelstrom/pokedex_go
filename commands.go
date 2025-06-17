package main

import (
	"fmt"
	"os"

	api "github.com/mmandelstrom/pokedex_go/internal/pokecache"
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

func commandMap(cfg *config, c *api.Cache) func() error {
	return func() error {
		data, err := api.MakeRequest(cfg.Next, c)
		if err != nil {
			return fmt.Errorf("Unable to make request")
		}
		location := api.GetLocation(data)
		api.PrintPokeLocation(&location)
		cfg.Next = location.Next
		cfg.Previous = location.Previous
		return nil
	}
}

func commandMapB(cfg *config, c *api.Cache) func() error {
	return func() error {
		if cfg.Previous == nil {
			fmt.Println("you're on the first page")
		} else {
			data, err := api.MakeRequest(*cfg.Previous, c)
			if err != nil {
				return fmt.Errorf("Unable to make request")
			}
			location := api.GetLocation(data)
			api.PrintPokeLocation(&location)
			cfg.Next = location.Next
			cfg.Previous = location.Previous
		}
		return nil
	}
}
