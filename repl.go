package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/garrettkucinski/pokedex-cli/internal/pokeapi"
)

type config struct {
	mapClient pokeapi.MapClient
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()

	for {
		fmt.Print("Pokedex> ")

		scanner.Scan()

		input := scanner.Text()

		fmt.Print("\n")

		if command, ok := commands[input]; ok {
			if err := command.callback(cfg); err != nil {
				fmt.Fprintln(os.Stderr, "error executing command:", err)
			}
		} else {
			fmt.Println("Command not found")
		}

		fmt.Print("\n")

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	}
}
