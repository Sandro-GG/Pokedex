package main

import (
	"time"

	"github.com/Sandro-GG/Pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	commands      map[string]cliCommand
	nextURL       *string
	prevURL       *string
	pokedex       map[string]pokeapi.Pokemon
}

func main() {
	startingUrl := "https://pokeapi.co/api/v2/location-area/"

	cfg := config{
		pokeapiClient: pokeapi.NewClient(30 * time.Second),
		commands: map[string]cliCommand{
			"exit": {
				name:        "exit",
				description: "Exit the Pokedex",
				callback:    commandExit,
			},
			"help": {
				name:        "help",
				description: "Displays a help message",
				callback:    commandHelp,
			},
			"map": {
				name:        "map",
				description: "Displays the names of 20 location areas in the Pokemon world",
				callback:    commandMap,
			},
			"mapb": {
				name:        "mapb",
				description: "Displays names of the previous 20 location areas in the Pokemon world",
				callback:    commandMapb,
			},
			"explore": {
				name:        "explore",
				description: "Displays the names of Pokemon in the area",
				callback:    commandExplore,
			},
			"catch": {
				name:        "catch",
				description: "Attempts to catch a specified Pokemon",
				callback:    commandCatch,
			},
		},
		nextURL: &startingUrl,
		prevURL: nil,
		pokedex: make(map[string]pokeapi.Pokemon),
	}

	startRepl(&cfg)
}
