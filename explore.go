package main

import (
	"errors"

	"github.com/garrettkucinski/pokedex-cli/internal/pokeapi"
)

func exploreCommand(cfg *pokeapi.Config, args ...string) (exploreError error) {
	var locationName string

	if len(args) > 0 {
		locationName = args[0]
	} else {
		exploreError = errors.New("You need to specify a location")
		return
	}

	cfg.Client.ExploreLocation(locationName)

	return
}
