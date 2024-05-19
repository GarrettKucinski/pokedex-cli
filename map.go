package main

import "github.com/garrettkucinski/pokedex-cli/internal/pokeapi"

func mapCommand(cfg *pokeapi.Config, args ...string) (mapError error) {
	cfg.Client.DisplayNextLocationList(cfg)

	return
}

func mapbCommand(cfg *pokeapi.Config, args ...string) (mapbError error) {
	cfg.Client.DisplayPrevLocationList(cfg)

	return
}
