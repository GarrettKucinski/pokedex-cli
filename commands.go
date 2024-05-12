package main

import (
	"github.com/garrettkucinski/pokedex-cli/pkg/pokemonApi"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type Commands struct{}

// Menu map[string]cliCommand
// func (c *Commands) AddCommand(command cliCommand) {
// 	commandName := command.name
// 	c[commandName] = command
// }

func getCommands() map[string]cliCommand {
	client := new(pokemonApi.MapClient)

	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    exitCommand,
		},
		"help": {
			name:        "help",
			description: "List all commands",
			callback:    helpCommand,
		},
		"map": {
			name:        "map",
			description: "Show the current map, additional calls return paginated results",
			callback:    client.DisplayNextLocationList,
		},
		"mapb": {
			name:        "mapb",
			description: "Go back the the previous map",
			callback:    client.DisplayPrevLocationList,
		},
	}
}
