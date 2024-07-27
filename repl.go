package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/codezera11/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

type cliCommand struct {
	name string
	description string
	callback func(config *config, args ...string) error
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
 
		words := cleanInput(reader.Text())

		if len(words) == 0 {
			continue
		}

		command, exists := getCommands()[words[0]]

		args := []string{}

		if len(words) > 1 {
			args = words[1:]
		}

		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
				continue
			}
		} else {
			fmt.Println("Unknown command!")
			continue
		}
	}

}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:  commandExit,
		},
		"map": {
			name: "map",
			description: "Displays areas of pokemon world in incremental manner",
			callback: commandMapf,
		},
		"mapb": {
			name: "mapb",
			description: "Goes back to pre page",
			callback: commandMapb,
		},
		"explore": {
			name: "explore",
			description: "Displays details about a specific area",
			callback: commandExplore,
		},
	}
}