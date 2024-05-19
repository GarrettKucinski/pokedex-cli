package main

import "github.com/garrettkucinski/pokedex-cli/internal/pokeapi"

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *pokeapi.Config, args ...string) error
}

// TODO: an add command function would be cool here

func getCommands() map[string]cliCommand {

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
			callback:    mapCommand,
		},
		"mapb": {
			name:        "mapb",
			description: "Go back the the previous map",
			callback:    mapbCommand,
		},
		"explore": {
			name:        "explore",
			description: "Explore areas by name",
			callback:    exploreCommand,
		},
		"catch": {
			name:        "catch",
			description: "Catch a pokemon by throwing a Pokeball at it!",
			callback:    catchCommand,
		},
	}
}
