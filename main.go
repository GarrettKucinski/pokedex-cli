package main

import (
	"github.com/garrettkucinski/pokedex-cli/internal/pokeapi"
	"github.com/garrettkucinski/pokedex-cli/internal/pokecache"
)

func main() {
	client := &pokeapi.MapClient{
		Cache: pokecache.NewCache(1000 * 60 * 10),
	}

	cfg := &config{
		mapClient: *client,
	}

	startRepl(cfg)
}
