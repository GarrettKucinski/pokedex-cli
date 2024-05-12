package main

import (
	"fmt"
	"os"
)

func exitCommand() error {
	fmt.Println("Goodbye! :)")
	os.Exit(0)
	return nil
}
