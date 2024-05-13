package main

import "fmt"

func commandExplore(cfg *config, area string) error {
	locationsResp, err := cfg.pokeapiClient.ListPokemon(area)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, encounters := range locationsResp.PokemonEncounters {
		fmt.Println("-", encounters.Pokemon.Name)
	}

	return nil
}