package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
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

func GetLocation(url string) pokeLocation {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	if err != nil {
		log.Fatal(err)
	}

	location := pokeLocation{}
	if err := json.Unmarshal(body, &location); err != nil {
		fmt.Println(err)
	}

	return location
}

func PrintPokeLocation(location *pokeLocation) error {
	for _, area := range location.Results {
		fmt.Println(area.Name)
	}

	return nil
}
