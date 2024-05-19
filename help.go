package main

import (
	"fmt"

	"github.com/garrettkucinski/pokedex-cli/internal/pokeapi"
)

func helpCommand(cfg *pokeapi.Config, args ...string) (helpError error) {
	fmt.Println("Commands:")

	commands := getCommands()

	for _, command := range commands {
		fmt.Printf("  %s: %s\n", command.name, command.description)
	}

	return
}
