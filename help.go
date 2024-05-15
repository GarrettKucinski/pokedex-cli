package main

import "fmt"

func helpCommand(cfg *config) (helpError error) {
	fmt.Println("Commands:")

	commands := getCommands()

	for _, command := range commands {
		fmt.Printf("  %s: %s\n", command.name, command.description)
	}

	return
}
