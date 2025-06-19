package pokecache

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Pokedex struct {
	PokemonMap map[string]Pokemon
}

type Pokemon struct {
	ID             int           `json:"id"`
	Name           string        `json:"name"`
	BaseExperience int           `json:"base_experience"`
	Height         int           `json:"height"`
	IsDefault      bool          `json:"is_default"`
	Order          int           `json:"order"`
	Weight         int           `json:"weight"`
	Stats          []PokemonStat `json:"stats"`
	Types          []PokemonType `json:"types"`
}

type pokeLocation struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type PokemonStat struct {
	BaseStat int `json:"base_stat"`
	Effort   int `json:"effort"`
	Stat     struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"stat"`
}

type PokemonType struct {
	Slot int `json:"slot"`
	Type struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"type"`
}
type pokeArea struct {
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func MakeRequest(url string, c *Cache) ([]byte, error) {
	if entry, ok := c.Get(url); ok {
		return entry, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if res.StatusCode > 299 {
		return nil, fmt.Errorf("response failed with status code: %d and\nbody: %s", res.StatusCode, body)
	}

	if err != nil {
		return nil, err
	}

	c.Add(url, body)

	return body, nil
}

func GetLocation(data []byte) pokeLocation {
	location := pokeLocation{}
	if err := json.Unmarshal(data, &location); err != nil {
		fmt.Println(err)
	}

	return location
}

func GetAreaDetails(data []byte) pokeArea {
	area := pokeArea{}
	if err := json.Unmarshal(data, &area); err != nil {
		fmt.Println(err)
	}
	return area
}

func PrintPokeLocation(location *pokeLocation) error {
	for _, area := range location.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func PrintPokemonInArea(area *pokeArea) error {
	fmt.Printf("Exploring %s...\n", area.Name)
	fmt.Println("Found Pokemon:")
	for _, pokemon := range area.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}
	return nil
}

func GetPokemon(data []byte) Pokemon {
	pokemon := Pokemon{}
	if err := json.Unmarshal(data, &pokemon); err != nil {
		fmt.Println(err)
	}
	return pokemon
}
