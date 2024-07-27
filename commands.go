package main

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
		"catch": {
			name: "catch",
			description: "Catches the specified pokemon",
			callback: commandCatch,
		},
		"inspect": {
			name: "inspect",
			description: "Provides info about pokemons you have caught",
			callback: commandInspect,
		},
		"pokedex": {
			name: "pokedex",
			description: "Lists the pokemons you have caught",
			callback: commandPokedex,
		},
	}
}