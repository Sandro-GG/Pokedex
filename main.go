package main

type config struct {
	commands map[string]cliCommand
	nextURL  *string
	prevURL  *string
}

func main() {
	startingURL := "https://pokeapi.co/api/v2/location-area/"

	cfg := config{
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
		},
		nextURL: &startingURL,
		prevURL: nil,
	}

	startRepl(&cfg)
}
