package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(url *string) (LocationArea, error) {
	req, err := http.NewRequest("GET", *url, nil)
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

	var locs LocationArea
	if err = json.Unmarshal(data, &locs); err != nil {
		return LocationArea{}, fmt.Errorf("error: %w", err)
	}

	return locs, nil
}
