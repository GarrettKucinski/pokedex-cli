package main

import (
	"fmt"
	"os"

	"github.com/garrettkucinski/pokedex-cli/internal/pokeapi"
)

func exitCommand(cfg *pokeapi.Config, args ...string) error {
	fmt.Println("Goodbye! :)")
	os.Exit(0)
	return nil
}
