package main

import (
	"errors"
	"fmt"

	"github.com/garrettkucinski/pokedex-cli/internal/pokeapi"
)

func inspectCommand(cfg *pokeapi.Config, args ...string) (inspectError error) {
	var pokemonName string

	if len(args) > 0 {
		pokemonName = args[0]
		pokemon, found := cfg.Pokedex.Get(pokemonName)

		if found {
			fmt.Printf("Name: %s\n", pokemon.Name)
			fmt.Printf("Height: %d\n", pokemon.Height)
			fmt.Printf("Weight: %d\n", pokemon.Weight)

			fmt.Println("Stats:")
			for _, statInfo := range pokemon.Stats {
				fmt.Printf("\t-%s: %d\n", statInfo.Stat.Name, statInfo.BaseStat)
			}

			fmt.Println("Types:")
			for _, pokeType := range pokemon.Types {
				fmt.Printf("\t-%s\n", pokeType.Type.Name)
			}
		}
	} else {
		inspectError = errors.New("you need to specify a pokemon to catch")
	}

	return
}
