package main

import (
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, pokemon string) error {
	pokemonResp, err := cfg.pokeapiClient.CatchPokemon(pokemon)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a pokeball at %s...\n", pokemon)

	chance := rand.IntN(pokemonResp.BaseExperience)

	if chance < 100 {
		fmt.Printf("%s was caught\n", pokemon)
		
		cfg.pokedex[pokemon] = pokemonResp
	} else {
		fmt.Printf("%s escaped\n", pokemon)
	}

	return nil
}