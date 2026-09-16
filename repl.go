package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func cleanInput(text string) []string {
	clean := strings.ToLower(text)

	return strings.Fields(clean)
}

func commandExit(cfg *config) error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)

	return nil
}

func commandHelp(cfg *config) error {
	fmt.Printf("Welcome to the Pokedex!\n")

	for _, cmd := range cfg.commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	return nil
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		if scanner.Scan() {
			input := scanner.Text()
			clean := cleanInput(input)
			if len(clean) == 0 {
				continue
			}

			if command, ok := cfg.commands[clean[0]]; !ok {
				fmt.Printf("Unknown command\n")
			} else {
				if err := command.callback(cfg); err != nil {
					fmt.Println(err)
				}
			}

		}
	}
}
