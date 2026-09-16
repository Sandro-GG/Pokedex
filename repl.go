package main

import "strings"

func cleanInput(text string) []string {
	clean := strings.ToLower(text)

	return strings.Fields(clean)
}
