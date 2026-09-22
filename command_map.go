package main

import "fmt"

func commandMap(cfg *config, args ...string) error {
	if cfg.nextURL == nil {
		fmt.Println("You have reached the end of data")
		return nil
	}

	// send get request
	locs, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextURL)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	// print 20 locations
	for _, loc := range locs.Results {
		fmt.Println(loc.Name)
	}

	// move to the next 20
	cfg.nextURL = locs.Next
	cfg.prevURL = locs.Previous

	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.prevURL == nil {
		fmt.Println("You are on the first page. Cannot go back!")
		return nil
	}

	locs, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevURL)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	// print 20 locations
	for _, loc := range locs.Results {
		fmt.Println(loc.Name)
	}

	// move to the prev 20
	cfg.nextURL = locs.Next
	cfg.prevURL = locs.Previous

	return nil
}
