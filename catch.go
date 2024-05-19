package main

import (
	"errors"

	"github.com/garrettkucinski/pokedex-cli/internal/pokeapi"
)

func catchCommand(cfg *pokeapi.Config, args ...string) (exploreError error) {
	var pokemonName string

	if len(args) > 0 {
		pokemonName = args[0]
	} else {
		exploreError = errors.New("you need to specify a pokemon to catch")
		return
	}

	cfg.Client.CatchPokemon(cfg, pokemonName)

	return
}
