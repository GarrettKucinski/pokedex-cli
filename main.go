package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()

	for {
		fmt.Print("Pokedex> ")

		scanner.Scan()

		input := scanner.Text()

		fmt.Print("\n")

		if command, ok := commands[input]; ok {
			if err := command.callback(); err != nil {
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
