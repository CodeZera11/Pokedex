package main

import (
	"errors"
	"fmt"
)

func commandInspect(c *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("please enter a pokemon name")
	}

	name := args[0]

	pokemon, has := c.pokedex[name]

	if !has {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Println()
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")

	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")

	for _, pokeType := range pokemon.Types {
		fmt.Printf("  -%s\n", pokeType.Type.Name)
	}

	fmt.Println()

	return nil
}