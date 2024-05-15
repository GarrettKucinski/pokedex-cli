package main

import "github.com/garrettkucinski/pokedex-cli/internal/pokeapi"

func main() {
	client := new(pokeapi.MapClient)

	cfg := &config{
		mapClient: *client,
	}

	startRepl(cfg)
}
