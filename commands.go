package main

import (
	"fmt"
	"os"

	api "github.com/mmandelstrom/pokedex_go/internal/pokecache"
)

func commandHelp(cfg *config) func(param string) error {
	return func(param string) error {
		fmt.Print()
		fmt.Println("Welcome to the Pokedex!")
		fmt.Print("Usage:\n\n")
		for _, cmd := range getCommands() {
			fmt.Printf("%s: %s\n", cmd.name, cmd.description)
		}
		return nil
	}

}

func commandExit(cfg *config) func(param string) error {
	return func(param string) error {
		fmt.Println("Closing the Pokedex... Goodbye!")
		os.Exit(0)
		return nil
	}
}

func commandMap(cfg *config, c *api.Cache) func(area string) error {
	return func(area string) error {
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

func commandMapB(cfg *config, c *api.Cache) func(area string) error {
	return func(area string) error {
		if cfg.Previous == nil {
			fmt.Println("you're on the first page")
		} else {
			data, err := api.MakeRequest(*cfg.Previous, c)
			if err != nil {
				return fmt.Errorf("unable to make request")
			}
			location := api.GetLocation(data)
			api.PrintPokeLocation(&location)
			cfg.Next = location.Next
			cfg.Previous = location.Previous
		}
		return nil
	}
}

func commandExplore(c *api.Cache) func(areaName string) error {
	return func(areaName string) error {
		baseUrl := "https://pokeapi.co/api/v2/location-area/"
		fullUrl := baseUrl + areaName
		data, err := api.MakeRequest(fullUrl, c)
		if err != nil {
			return fmt.Errorf("Unable to make request")
		}
		areaInfo := api.GetAreaDetails(data)
		api.PrintPokemonInArea(&areaInfo)
		return nil
	}
}
