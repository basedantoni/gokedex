package main

import (
	"fmt"
)

func commandHelp(c *config) error {
	commands := getCommands()
	
	fmt.Println("\nWelcome to the Pokedex CLI")
	fmt.Println("Usage:")
	fmt.Println()

	for k, v := range commands {
		fmt.Printf("%s: %s\n", k, v.description)
	}
	printPrompt()
	
	return nil
}