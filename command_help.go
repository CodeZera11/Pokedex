package main

import "fmt"

func commandHelp() error {
	fmt.Println()
	fmt.Println("Welcome to Pokedex!")
	fmt.Print("Usage:\n\n")

	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	fmt.Println()

	return nil
}