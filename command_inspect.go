package main

import (
	"fmt"
)

func commandInspect(cfg *config, pokemon string) error {
	if p, ok := cfg.pokedex[pokemon]; !ok {
		fmt.Printf("Pokemon %s has not been caught\n", p.Species.Name)
		return nil
	} else {
		fmt.Println()
		fmt.Printf(
			"Name: %s \nWeight: %d\nStats:\n  -hp: %d\n  -attack: %d\n  -defense: %d\n  -special-attack: %d\n  -special-defense: %d\n  -speed: %d\n", 
			p.Name, 
			p.Height, 
			p.Stats[0].BaseStat,
			p.Stats[1].BaseStat,
			p.Stats[2].BaseStat,
			p.Stats[3].BaseStat,
			p.Stats[4].BaseStat,
			p.Stats[5].BaseStat,
		)

		fmt.Println("Types:")
		for _, v := range p.Types {
			fmt.Printf("  - %s\n", v.Type.Name)
		}
	
		return nil
	}

}