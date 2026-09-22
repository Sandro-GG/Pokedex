package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	fmt.Printf("Your Pokedex:\n")

	for name := range cfg.pokedex {
		fmt.Printf(" - %s\n", name)
	}

	return nil
}
