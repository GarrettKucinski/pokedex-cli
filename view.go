package main

import "github.com/garrettkucinski/pokedex-cli/internal/pokeapi"

func viewCommand(cfg *pokeapi.Config, args ...string) (viewError error) {
	cfg.Pokedex.View()

	return
}
