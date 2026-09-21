package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"math/rand/v2"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, args ...string) error
}

func cleanInput(text string) []string {
	clean := strings.ToLower(text)

	return strings.Fields(clean)
}

func commandExit(cfg *config, args ...string) error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)

	return nil
}

func commandHelp(cfg *config, args ...string) error {
	fmt.Printf("Welcome to the Pokedex!\n")

	for _, cmd := range cfg.commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	return nil
}

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

func commandExplore(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("please specify the name of the area you wish to explore")
	}

	encounters, err := cfg.pokeapiClient.LocationGet(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\nFound Pokemon:\n", args[0])

	for _, encounter := range encounters.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}

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

func commandPokedex(cfg *config, args ...string) error {
	fmt.Printf("Your Pokedex:\n")

	for name := range cfg.pokedex {
		fmt.Printf(" - %s\n", name)
	}

	return nil
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		clean := cleanInput(input)
		if len(clean) == 0 {
			continue
		}

		if command, ok := cfg.commands[clean[0]]; !ok {
			fmt.Printf("Unknown command\n")
		} else {
			if err := command.callback(cfg, clean[1:]...); err != nil {
				fmt.Println(err)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
