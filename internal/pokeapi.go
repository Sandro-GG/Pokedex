package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationArea struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
	} `json:"results"`
}

func FetchLocationAreas(url *string) (LocationArea, error) {
	if url == nil {
		return LocationArea{}, fmt.Errorf("You have reached the end of the map data!")
	}

	req, err := http.NewRequest("GET", *url, nil)
	if err != nil {
		return LocationArea{}, fmt.Errorf("error: %w", err)
	}

	client := &http.Client{}
	res, err := client.Do(req)
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
