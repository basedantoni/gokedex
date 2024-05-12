package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/basedantoni/gokedex/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(c *config) error
}

type config struct {
	Next 	 string
	Previous string
}

type LocationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand {
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:		 "map",
			description: "Get next location areas",
			callback:	 commandMap,
		},
		"mapb": {
			name:		 "map back",
			description: "Get previous location areas",
			callback:	 commandMapb,
		},
	}
}

func printPrompt() {
	fmt.Print("\npokedex > ")
}

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	c := config{Next: "", Previous: ""}
	cache := pokecache.NewCache(time.Second * 5)
	
	printPrompt()
	for reader.Scan() {
		printPrompt()

		input := reader.Text()

		if _, ok := commands[input]; !ok {
			fmt.Fprintln(os.Stderr, "invalid command")
			continue
		}

		err :=commands[input].callback(&c)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	}

	if err := reader.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}