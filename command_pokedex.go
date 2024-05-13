package main

import (
	"fmt"
)

func commandPokedex(cfg *config, pokemon string) error {
	if len(cfg.pokedex) < 1 {
		fmt.Println("You haven't caught any pokemon yet")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, v := range cfg.pokedex {
		fmt.Println("- ", v.Name)
	}

	return nil
}