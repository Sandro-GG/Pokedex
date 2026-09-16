package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	clean := strings.ToLower(text)

	return strings.Fields(clean)
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			input := scanner.Text()
			cleaned := cleanInput(input)
			if len(cleaned) == 0 {
				continue
			}

			fmt.Printf("Your command was: %s\n", cleaned[0])
		}
	}
}
