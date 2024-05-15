package main

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

// type Commands struct{}

// Menu map[string]cliCommand
// func (c *Commands) AddCommand(command cliCommand) {
// 	commandName := command.name
// 	c[commandName] = command
// }

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
	}
}
