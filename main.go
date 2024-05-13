package main

import (
	"time"

	"github.com/basedantoni/gokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex: map[string]pokeapi.Pokemon{},
	}

	startRepl(cfg)
}