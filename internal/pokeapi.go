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

func CommandMap(cfg *config) error {
	if cfg.nextURL == nil {
		fmt.Println("You have reached the end of the map data!")
		return nil
	}

	req, err := http.NewRequest("GET", *cfg.nextURL, nil)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	var locs LocationArea
	if err = json.Unmarshal(data, &locs); err != nil {
		return fmt.Errorf("error: %w", err)
	}

	for _, loc := range locs.Results {
		fmt.Println(loc.Name)
	}

	cfg.prevURL = cfg.nextURL
	cfg.nextURL = &locs.Next

	return nil
}
