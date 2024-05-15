package main

import (
	"github.com/garrettkucinski/pokedex-cli/internal/pokeapi"
	"github.com/garrettkucinski/pokedex-cli/internal/pokecache"
)

func main() {
	cache := pokecache.NewCache(1000 * 60 * 10)
	client := new(pokeapi.MapClient)
	client.Cache = cache

	cfg := &config{
		mapClient: *client,
	}

	startRepl(cfg)
}
