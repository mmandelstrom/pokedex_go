package main

import (
	"fmt"
	"math/rand"
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
			return fmt.Errorf("unable to make request")
		}
		areaInfo := api.GetAreaDetails(data)
		api.PrintPokemonInArea(&areaInfo)
		return nil
	}
}

func commandCatch(c *api.Cache, dex *api.Pokedex) func(pokemonName string) error {
	return func(pokemonName string) error {
		baseUrl := "https://pokeapi.co/api/v2/pokemon/"
		fullUrl := baseUrl + pokemonName

		data, err := api.MakeRequest(fullUrl, c)
		if err != nil {
			return fmt.Errorf("unable to find pokemon")
		}
		pokemon := api.GetPokemon(data)
		fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

		throwRes := CatchPokemon(pokemon)
		if throwRes == false {
			fmt.Printf("%s escaped!\n", pokemon.Name)
			return nil
		}
		fmt.Printf("%s was caught!\n", pokemon.Name)
		dex.PokemonMap[pokemon.Name] = pokemon

		return nil
	}

}

func CatchPokemon(pokemon api.Pokemon) bool {
	chance := 100 - (pokemon.BaseExperience / 2)
	if chance < 5 {
		chance = 5
	}

	roll := rand.Intn(100) + 1

	return roll <= chance
}

func commandInspect(dex *api.Pokedex) func(pokemonName string) error {
	return func(pokemonName string) error {
		pokemon, err := dex.PokemonMap[pokemonName]
		if !err {
			return fmt.Errorf("%s was not found in pokedex", pokemonName)
		}

		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		for _, stat := range pokemon.Stats {
			fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, pokeType := range pokemon.Types {
			fmt.Printf("  - %s\n", pokeType.Type.Name)
		}

		return nil
	}
}

func commandPokedex() func(param string) error {
	if pokeDex.PokemonMap == nil {
		println("Your pokedex is empty")
		return nil
	}

	return func(param string) error {
		if len(pokeDex.PokemonMap) == 0 {
			fmt.Println("Your pokedex is empty")
			return nil
		}
		fmt.Println("Your pokedex:")
		for _, pokemon := range pokeDex.PokemonMap {
			fmt.Printf("- %s\n", pokemon.Name)
		}
		return nil
	}
}
