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
		return nil, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
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

func PrintPokeLocation(location *pokeLocation) error {
	for _, area := range location.Results {
		fmt.Println(area.Name)
	}

	return nil
}
