package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("please specify the name of the Pokemon you wish to inspect")
	}

	pok, ok := cfg.pokedex[args[0]]
	if !ok {
		return errors.New("you have not caught that Pokemon")
	}

	fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\n", pok.Name, pok.Height, pok.Weight)

	fmt.Printf("Stats:\n")
	for _, stat := range pok.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Printf("Types:\n")
	for _, tp := range pok.Types {
		fmt.Printf("  - %s\n", tp.Type.Name)
	}

	return nil
}
