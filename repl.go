package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/garrettkucinski/pokedex-cli/internal/pokeapi"
)

func startRepl(cfg *pokeapi.Config) {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()

	for {
		fmt.Print("Pokedex> ")

		scanner.Scan()
		input := scanner.Text()

		splitArgs := strings.Split(input, " ")
		var cmd string
		var arg string

		if len(splitArgs) > 0 {
			cmd = splitArgs[0]
		}

		if len(splitArgs) > 1 {
			arg = splitArgs[1]
		}

		fmt.Print("\n")

		if command, ok := commands[cmd]; ok {
			if err := command.callback(cfg, arg); err != nil {
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
