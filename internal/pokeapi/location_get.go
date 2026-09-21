package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) LocationGet(name string) (LocationArea, error) {
	url := baseLocationUrl + name

	if val, ok := c.cache.Get(url); ok {
		var locs LocationArea
		if err := json.Unmarshal(val, &locs); err != nil {
			return LocationArea{}, fmt.Errorf("error: %w", err)
		}
		return locs, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, fmt.Errorf("error: %w", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, fmt.Errorf("error: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, fmt.Errorf("error: %w", err)
	}

	c.cache.Add(url, data)

	var locs LocationArea
	if err = json.Unmarshal(data, &locs); err != nil {
		return LocationArea{}, fmt.Errorf("error: %w", err)
	}

	return locs, nil
}
