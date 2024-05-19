package main

import (
	"github.com/garrettkucinski/pokedex-cli/internal/pokeapi"
)

func main() {
	cfg := pokeapi.NewConfig()

	startRepl(cfg)
}
