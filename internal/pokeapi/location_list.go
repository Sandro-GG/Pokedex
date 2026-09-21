package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageUrl *string) (LocationAreaList, error) {
	url := baseLocationUrl + "?offset=0&limit=20"
	if pageUrl != nil {
		url = *pageUrl
	}

	if val, ok := c.cache.Get(url); ok {
		var locs LocationAreaList
		if err := json.Unmarshal(val, &locs); err != nil {
			return LocationAreaList{}, fmt.Errorf("error: %w", err)
		}
		return locs, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaList{}, fmt.Errorf("error: %w", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaList{}, fmt.Errorf("error: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaList{}, fmt.Errorf("error: %w", err)
	}

	c.cache.Add(url, data)

	var locs LocationAreaList
	if err = json.Unmarshal(data, &locs); err != nil {
		return LocationAreaList{}, fmt.Errorf("error: %w", err)
	}

	return locs, nil
}
