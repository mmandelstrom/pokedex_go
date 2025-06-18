package pokecache

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type pokeLocation struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type pokeArea struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
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
