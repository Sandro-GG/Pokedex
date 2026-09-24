package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) PokemonGet(name string) (Pokemon, error) {
	url := c.pokemonUrl + name

	if val, ok := c.cache.Get(url); ok {
		var pok Pokemon
		if err := json.Unmarshal(val, &pok); err != nil {
			return Pokemon{}, fmt.Errorf("error: %w", err)
		}
		return pok, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error: %w", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error: %w", err)
	}
	if res.StatusCode > 299 {
		if res.StatusCode == 404 {
			return Pokemon{}, fmt.Errorf("error %d: pokemon name not found", res.StatusCode)
		}
		return Pokemon{}, fmt.Errorf("bad status code: %d", res.StatusCode)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error: %w", err)
	}

	c.cache.Add(url, data)

	var pok Pokemon
	if err = json.Unmarshal(data, &pok); err != nil {
		return Pokemon{}, fmt.Errorf("error: %w", err)
	}

	return pok, nil
}
