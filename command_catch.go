package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand/v2"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("please specify the name of the Pokemon")
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])

	pokemon, err := cfg.pokeapiClient.PokemonGet(args[0])
	if err != nil {
		return err
	}

	if tryCatch(pokemon.BaseExp) {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		fmt.Printf("You may now inspect it with the inspect command.\n")
		cfg.pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}

func tryCatch(baseExp int) bool {
	// k handles the drop-off speed.
	// At 0.0055, a 64 EXP Bulbasaur has a ~70% catch rate, while a 635 EXP Blissey has a ~3% catch rate.
	const k = 0.0055
	const minChance = 0.02 // Strict 2% floor so no Pokemon is mathematically impossible

	// Calculate exponential drop: P = e^(-k * EXP)
	catchChance := math.Exp(-k * float64(baseExp))

	// Enforce a minimum safety floor
	if catchChance < minChance {
		catchChance = minChance
	}

	// Roll the dice and return true if caught, false if escaped
	return rand.Float64() < catchChance
}
