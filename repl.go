package main

import (
	"bufio"
	"errors"
	"fmt"
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
